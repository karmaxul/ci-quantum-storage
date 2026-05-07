package main

import (
	"encoding/json"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// ── Rate limiter ──────────────────────────────────────────────────────────────
//
// Per-key sliding window over a fixed time window (default 60s).
// Buckets are stored in a sync.Map; each bucket has its own mutex protecting
// the timestamp slice. A background goroutine sweeps stale buckets every
// minute so the map does not grow unbounded.
//
// Limits are applied differently depending on the key type:
//   - "ip:<addr>"   — unauthenticated requests, capped at ipLimit per window
//   - "key:<value>" — authenticated requests, capped at keyLimit per window
//
// Setting either limit to 0 disables limiting for that class.

const rlWindow = time.Minute

type rlBucket struct {
	mu         sync.Mutex
	timestamps []time.Time
}

type rateLimiter struct {
	ipLimit   int
	keyLimit  int
	window    time.Duration
	whitelist map[string]bool

	buckets      sync.Map // string -> *rlBucket
	totalAllowed atomic.Uint64
	totalDenied  atomic.Uint64
}

func newRateLimiter(ipPerMin, keyPerMin int, whitelistCSV string) *rateLimiter {
	rl := &rateLimiter{
		ipLimit:   ipPerMin,
		keyLimit:  keyPerMin,
		window:    rlWindow,
		whitelist: map[string]bool{},
	}
	for _, raw := range strings.Split(whitelistCSV, ",") {
		ip := strings.TrimSpace(raw)
		if ip != "" {
			rl.whitelist[ip] = true
		}
	}
	go rl.gcLoop()
	return rl
}

// gcLoop sweeps expired buckets every minute.
func (rl *rateLimiter) gcLoop() {
	t := time.NewTicker(time.Minute)
	defer t.Stop()
	for range t.C {
		rl.cleanup()
	}
}

func (rl *rateLimiter) cleanup() {
	cutoff := time.Now().Add(-rl.window)
	rl.buckets.Range(func(key, val any) bool {
		b := val.(*rlBucket)
		b.mu.Lock()
		i := 0
		for i < len(b.timestamps) && b.timestamps[i].Before(cutoff) {
			i++
		}
		b.timestamps = b.timestamps[i:]
		empty := len(b.timestamps) == 0
		b.mu.Unlock()
		if empty {
			rl.buckets.Delete(key)
		}
		return true
	})
}

func (rl *rateLimiter) getBucket(k string) *rlBucket {
	if v, ok := rl.buckets.Load(k); ok {
		return v.(*rlBucket)
	}
	actual, _ := rl.buckets.LoadOrStore(k, &rlBucket{})
	return actual.(*rlBucket)
}

// allow returns (allowed, remaining, retryAfterSeconds).
func (rl *rateLimiter) allow(key string, limit int) (bool, int, int) {
	if limit <= 0 {
		rl.totalAllowed.Add(1)
		return true, -1, 0
	}
	b := rl.getBucket(key)
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-rl.window)
	i := 0
	for i < len(b.timestamps) && b.timestamps[i].Before(cutoff) {
		i++
	}
	b.timestamps = b.timestamps[i:]

	if len(b.timestamps) >= limit {
		oldest := b.timestamps[0]
		retry := int(oldest.Add(rl.window).Sub(now).Seconds())
		if retry < 1 {
			retry = 1
		}
		rl.totalDenied.Add(1)
		return false, 0, retry
	}
	b.timestamps = append(b.timestamps, now)
	rl.totalAllowed.Add(1)
	return true, limit - len(b.timestamps), 0
}

// clientIP extracts the client's IP, honoring proxy headers.
func clientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		return strings.TrimSpace(strings.Split(xff, ",")[0])
	}
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return strings.TrimSpace(xri)
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}

// keyFor returns (bucketKey, limit, whitelisted). Authenticated requests
// (those that present a non-empty API key when API_KEYS is configured) are
// keyed by the API key and use keyLimit. Everything else is keyed by IP
// and uses ipLimit. Whitelist applies only to IP keys.
func (rl *rateLimiter) keyFor(r *http.Request) (string, int, bool) {
	if apiKeys != "" {
		k := r.Header.Get("X-API-Key")
		if k == "" {
			k = r.URL.Query().Get("api_key")
		}
		if k != "" {
			return "key:" + k, rl.keyLimit, false
		}
	}
	ip := clientIP(r)
	return "ip:" + ip, rl.ipLimit, rl.whitelist[ip]
}

// withRateLimit wraps a handler. /health, /stats, /metrics, and CORS preflights
// bypass the limiter entirely (they're polled by the UI status bar and have no
// side effects). Everything else is subject to the appropriate cap.
func withRateLimit(rl *rateLimiter, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions ||
			r.URL.Path == "/health" ||
			r.URL.Path == "/stats" ||
			r.URL.Path == "/metrics" {
			h.ServeHTTP(w, r)
			return
		}

		key, limit, whitelisted := rl.keyFor(r)
		if whitelisted || limit <= 0 {
			h.ServeHTTP(w, r)
			return
		}

		allowed, remaining, retryAfter := rl.allow(key, limit)

		w.Header().Set("X-RateLimit-Limit", strconv.Itoa(limit))
		w.Header().Set("X-RateLimit-Window", strconv.Itoa(int(rl.window.Seconds())))
		if remaining >= 0 {
			w.Header().Set("X-RateLimit-Remaining", strconv.Itoa(remaining))
		}

		if !allowed {
			w.Header().Set("Retry-After", strconv.Itoa(retryAfter))
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			_ = json.NewEncoder(w).Encode(map[string]interface{}{
				"error":         "rate limit exceeded",
				"limit":         limit,
				"windowSeconds": int(rl.window.Seconds()),
				"retryAfter":    retryAfter,
			})
			return
		}
		h.ServeHTTP(w, r)
	})
}

// rateLimitStats is the JSON-friendly snapshot exposed via /metrics.
type rateLimitStats struct {
	Enabled           bool     `json:"enabled"`
	IPLimitPerMin     int      `json:"ipLimitPerMin"`
	APIKeyLimitPerMin int      `json:"apiKeyLimitPerMin"`
	WindowSeconds     int      `json:"windowSeconds"`
	ActiveBuckets     int      `json:"activeBuckets"`
	TotalAllowed      uint64   `json:"totalAllowed"`
	TotalDenied       uint64   `json:"totalDenied"`
	Whitelist         []string `json:"whitelist,omitempty"`
}

func (rl *rateLimiter) Stats() rateLimitStats {
	n := 0
	rl.buckets.Range(func(_, _ any) bool {
		n++
		return true
	})

	wl := make([]string, 0, len(rl.whitelist))
	for k := range rl.whitelist {
		wl = append(wl, k)
	}
	return rateLimitStats{
		Enabled:           rl.ipLimit > 0 || rl.keyLimit > 0,
		IPLimitPerMin:     rl.ipLimit,
		APIKeyLimitPerMin: rl.keyLimit,
		WindowSeconds:     int(rl.window.Seconds()),
		ActiveBuckets:     n,
		TotalAllowed:      rl.totalAllowed.Load(),
		TotalDenied:       rl.totalDenied.Load(),
		Whitelist:         wl,
	}
}
