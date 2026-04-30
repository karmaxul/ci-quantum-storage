package main

import (
    "encoding/base64"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "net/http"
    "sync"
    "time"

    "ci-sha-test/healchain"
)

type EncodeRequest struct {
    Data         string `json:"data"`
    DataShards   int    `json:"data_shards,omitempty"`
    ParityShards int    `json:"parity_shards,omitempty"`
    Compressed   bool   `json:"compressed,omitempty"`
    UseBase64    bool   `json:"use_base64,omitempty"`
}

type Response struct {
    Success     bool   `json:"success"`
    Encoded     string `json:"encoded,omitempty"` // base64 or hex
    OriginalLen int    `json:"original_len,omitempty"`
    EncodedLen  int    `json:"encoded_len,omitempty"`
    TimeTaken   string `json:"time_taken,omitempty"`
    Error       string `json:"error,omitempty"`
}

type StatsResponse struct {
    Status          string  `json:"status"`
    Version         string  `json:"version"`
    Uptime          string  `json:"uptime"`
    TotalRequests   int64   `json:"total_requests"`
    AvgEncodeTimeMs float64 `json:"avg_encode_time_ms"`
    ActiveEncodes   int     `json:"active_encodes"`
}

var (
    startTime      = time.Now()
    totalRequests  int64
    totalEncodeMs  float64
    mu             sync.Mutex
    activeEncodes  int
)

func main() {
    http.HandleFunc("/encode", handleEncode)
    http.HandleFunc("/decode", handleDecode)
    http.HandleFunc("/health", handleHealth)
    http.HandleFunc("/stats", handleStats)

    fmt.Println("🚀 HealChain Self-Healing Service v2.2 (Enhanced)")
    fmt.Println("   Endpoints: /encode, /decode, /health, /stats")
    fmt.Println("   Ready on http://localhost:8080")

    http.ListenAndServe(":8080", nil)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "status":  "healthy",
        "version": "2.2",
    })
}

func handleStats(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    avg := 0.0
    if totalRequests > 0 {
        avg = totalEncodeMs / float64(totalRequests)
    }
    mu.Unlock()

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(StatsResponse{
        Status:          "running",
        Version:         "2.2",
        Uptime:          time.Since(startTime).String(),
        TotalRequests:   totalRequests,
        AvgEncodeTimeMs: avg,
        ActiveEncodes:   activeEncodes,
    })
}

func handleEncode(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    mu.Lock()
    totalRequests++
    activeEncodes++
    mu.Unlock()

    defer func() {
        mu.Lock()
        activeEncodes--
        mu.Unlock()
    }()

    var req EncodeRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    dataShards := 10
    parityShards := 4
    if req.DataShards > 0 {
        dataShards = req.DataShards
    }
    if req.ParityShards > 0 {
        parityShards = req.ParityShards
    }

    // Decode input
    var rawData []byte
    if req.Compressed {
        rawData, _ = hex.DecodeString(req.Data)
    } else {
        rawData = []byte(req.Data)
    }

    rs, err := healchain.New(dataShards, parityShards)
    if err != nil {
        sendError(w, err.Error())
        return
    }

    encoded, err := rs.Encode(rawData)
    if err != nil {
        sendError(w, err.Error())
        return
    }

    // Return as base64 by default for easier Flask handling
    encodedStr := base64.StdEncoding.EncodeToString(encoded)
    if !req.UseBase64 {
        encodedStr = hex.EncodeToString(encoded)
    }

    duration := time.Since(start).Seconds() * 1000

    mu.Lock()
    totalEncodeMs += duration
    mu.Unlock()

    resp := Response{
        Success:     true,
        Encoded:     encodedStr,
        OriginalLen: len(rawData),
        EncodedLen:  len(encoded),
        TimeTaken:   fmt.Sprintf("%.2fms", duration),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func handleDecode(w http.ResponseWriter, r *http.Request) {
    // Similar improvements can be added here later
    // For now, keep basic decode
}

func sendError(w http.ResponseWriter, msg string) {
    w.WriteHeader(http.StatusInternalServerError)
    json.NewEncoder(w).Encode(Response{Success: false, Error: msg})
}
