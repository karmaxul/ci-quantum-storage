package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	binding "ci-sha-test/binding"
	bindingsepolia "ci-sha-test/binding-sepolia"
	"ci-sha-test/healchain"
)

// ── Configuration ─────────────────────────────────────────────────────────────

var (
	contractAddress = getEnv("CONTRACT_ADDRESS", "0x5FbDB2315678afecb367f032d93F642f64180aa3")
	storePrivateKey = getEnv("STORE_PRIVATE_KEY", "b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
	gethURL         = getEnv("GETH_URL", "http://localhost:8545")
	listenAddr      = getEnv("LISTEN_ADDR", ":8080")
	apiKeys         = getEnv("API_KEYS", "") // comma-separated, empty = disabled

	// Set in main() once if Sepolia config is present. Read by /metrics.
	sepoliaOracle    *Oracle
	serviceStartedAt = time.Now()

	// Set in main() before serving. Used by withRateLimit middleware and /metrics.
	serviceLimiter *rateLimiter
)

const serviceVersion = "2.4"

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

// ── Input validation ──────────────────────────────────────────────────────────

const (
	maxDataBytes  = 1024 * 1024 // 1MB max payload
	maxLabelBytes = 64
)

func validateInput(data []byte, label string) error {
	if len(data) == 0 {
		return fmt.Errorf("data cannot be empty")
	}
	if len(data) > maxDataBytes {
		return fmt.Errorf("data too large: %d bytes (max %d)", len(data), maxDataBytes)
	}
	if len(label) > maxLabelBytes {
		return fmt.Errorf("label too long: %d chars (max %d)", len(label), maxLabelBytes)
	}
	return nil
}

// ── Compression helpers ───────────────────────────────────────────────────────

func compressData(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	if _, err := w.Write(data); err != nil {
		return nil, fmt.Errorf("compress: write failed: %w", err)
	}
	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("compress: close failed: %w", err)
	}
	return buf.Bytes(), nil
}

func decompressData(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("decompress: reader failed: %w", err)
	}
	defer r.Close()
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		return nil, fmt.Errorf("decompress: read failed: %w", err)
	}
	return buf.Bytes(), nil
}

// ── JSON helpers ──────────────────────────────────────────────────────────────

func jsonOK(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func jsonErr(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

// ── Geth health check with retry ─────────────────────────────────────────────

func checkGeth() error {
	for attempt := 1; attempt <= 5; attempt++ {
		client, err := ethclient.Dial(gethURL)
		if err == nil {
			_, err = client.ChainID(context.Background())
			client.Close()
			if err == nil {
				return nil
			}
		}
		fmt.Printf("Geth not ready (attempt %d/5): %v\n", attempt, err)
		time.Sleep(time.Duration(attempt) * 500 * time.Millisecond)
	}
	return fmt.Errorf("geth not reachable after 5 attempts")
}

func gethClient() (*ethclient.Client, error) {
	return ethclient.Dial(gethURL)
}

// ── CORS middleware ───────────────────────────────────────────────────────────

func withCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// ── API key middleware ────────────────────────────────────────────────────────

func withAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If no API keys configured, auth is disabled
		if apiKeys == "" {
			h.ServeHTTP(w, r)
			return
		}

		// Health, stats, and metrics are always public
		if r.URL.Path == "/health" || r.URL.Path == "/stats" || r.URL.Path == "/metrics" {
			h.ServeHTTP(w, r)
			return
		}

		key := r.Header.Get("X-API-Key")
		if key == "" {
			// Also accept as query param for convenience
			key = r.URL.Query().Get("api_key")
		}

		if key == "" {
			jsonErr(w, http.StatusUnauthorized, "missing API key — provide X-API-Key header")
			return
		}

		for _, valid := range strings.Split(apiKeys, ",") {
			if strings.TrimSpace(valid) == key {
				h.ServeHTTP(w, r)
				return
			}
		}

		jsonErr(w, http.StatusForbidden, "invalid API key")
	})
}

// ── Main ──────────────────────────────────────────────────────────────────────

func main() {
	mux := http.NewServeMux()

	// Precompile bridge
	mux.HandleFunc("/encode", handleEncode)
	mux.HandleFunc("/decode", handleDecode)

	// Chain operations
	mux.HandleFunc("/storeOnChain", handleStoreOnChain)
	mux.HandleFunc("/store", handleStore)
	mux.HandleFunc("/retrieve", handleRetrieve)
	mux.HandleFunc("/getMetadata", handleGetMetadata)
	mux.HandleFunc("/listRecords", handleListRecords)
	mux.HandleFunc("/delete", handleDelete)
	mux.HandleFunc("/sepoliaStatus", handleSepoliaStatus)

	// Service info
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/stats", handleStats)
	mux.HandleFunc("/metrics", handleMetrics)

	mux.HandleFunc("/nuclear-test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("=== NUCLEAR TEST ROUTE HIT ===")
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Nuclear test successful"))
	})

	// Rate limiter — configurable via env vars, applied between CORS and Auth.
	ipPerMin := 30
	if v := os.Getenv("RATE_LIMIT_IP_PER_MIN"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			ipPerMin = n
		}
	}
	keyPerMin := 60
	if v := os.Getenv("RATE_LIMIT_KEY_PER_MIN"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			keyPerMin = n
		}
	}
	serviceLimiter = newRateLimiter(ipPerMin, keyPerMin, getEnv("RATE_LIMIT_WHITELIST", ""))

	srv := &http.Server{
		Addr:    listenAddr,
		Handler: withCORS(withRateLimit(serviceLimiter, withAuth(mux))),
	}

	// Graceful shutdown on SIGTERM / SIGINT
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
		<-quit
		fmt.Println("\n🛑 Shutting down gracefully...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		srv.Shutdown(ctx)
	}()

	fmt.Println("🚀 HealChain Self-Healing Service v" + serviceVersion)
	fmt.Println("   Endpoints: /encode /decode /storeOnChain /store /retrieve /getMetadata /listRecords /delete /health /stats /metrics")
	fmt.Printf("   Contract:  %s\n", contractAddress)
	fmt.Printf("   Geth:      %s\n", gethURL)
	fmt.Printf("   Listening: %s\n", listenAddr)
	fmt.Printf("   Rate limit: ip=%d/min key=%d/min (exempt: /health /stats /metrics; whitelist=%d)\n",
		ipPerMin, keyPerMin, len(serviceLimiter.whitelist))
	fmt.Println("Listening...")

	// ── Start Sepolia oracle if configured ───────────────────────────────────
	sepoliaRPC := getEnv("SEPOLIA_RPC_URL", "")
	sepoliaContract := getEnv("SEPOLIA_CONTRACT_ADDRESS", "")
	oracleKey := getEnv("ORACLE_PRIVATE_KEY", storePrivateKey)

	if sepoliaRPC != "" && sepoliaContract != "" {
		oracleCfg := OracleConfig{
			SepoliaRPC:      sepoliaRPC,
			ContractAddress: sepoliaContract,
			PrivateKey:      oracleKey,
			ChainID:         11155111,
			PollInterval:    15 * time.Second,
			Confirmations:   2,
			StateFile:       "/app/oracle-state.json",
		}
		oracle, err := NewOracle(oracleCfg)
		if err != nil {
			fmt.Println("Oracle init failed:", err)
		} else {
			sepoliaOracle = oracle
			oracleCtx, oracleCancel := context.WithCancel(context.Background())
			defer oracleCancel()
			go oracle.Start(oracleCtx)
		}
	} else {
		fmt.Println("ℹ️  SEPOLIA_RPC_URL not set — oracle watcher disabled")
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("ListenAndServe error:", err)
	}
}

// ── handleHealth ──────────────────────────────────────────────────────────────

func handleHealth(w http.ResponseWriter, r *http.Request) {
	gethStatus := "ok"
	lastBlock := "unknown"
	chainID := "unknown"

	client, err := gethClient()
	if err != nil {
		gethStatus = "unreachable"
	} else {
		defer client.Close()
		if block, err := client.BlockByNumber(context.Background(), nil); err == nil {
			lastBlock = block.Number().String()
		}
		if id, err := client.ChainID(context.Background()); err == nil {
			chainID = id.String()
		}
	}

	jsonOK(w, map[string]interface{}{
		"status":    "healthy",
		"version":   "2.4",
		"geth":      gethStatus,
		"chainId":   chainID,
		"lastBlock": lastBlock,
		"contract":  contractAddress,
	})
}

// ── handleStats ───────────────────────────────────────────────────────────────

func handleStats(w http.ResponseWriter, r *http.Request) {
	total := "unknown"

	client, err := gethClient()
	if err == nil {
		defer client.Close()
		contractAddr := common.HexToAddress(contractAddress)
		instance, err := binding.NewHealChainStorage(contractAddr, client)
		if err == nil {
			if t, err := instance.TotalRecords(nil); err == nil {
				total = t.String()
			}
		}
	}

	jsonOK(w, map[string]interface{}{
		"status":       "running",
		"version":      "2.4",
		"contract":     contractAddress,
		"geth":         gethURL,
		"totalRecords": total,
	})
}

// ── handleEncode ──────────────────────────────────────────────────────────────

func handleEncode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== ENCODE REQUEST RECEIVED ===")
	var req struct {
		DataStr      string `json:"data"`
		DataShards   int    `json:"dataShards"`
		ParityShards int    `json:"parityShards"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonErr(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	data, err := hex.DecodeString(strings.TrimPrefix(req.DataStr, "0x"))
	if err != nil {
		jsonErr(w, http.StatusBadRequest, "invalid hex data")
		return
	}
	dataShards, parityShards := 10, 4
	if req.DataShards > 0 {
		dataShards = req.DataShards
	}
	if req.ParityShards > 0 {
		parityShards = req.ParityShards
	}
	rs, err := healchain.New(dataShards, parityShards)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "RS init failed: "+err.Error())
		return
	}
	encoded, err := rs.Encode(data)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "encode failed: "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(encoded)
	fmt.Println("✅ Encode completed, length:", len(encoded))
}

// ── handleDecode ──────────────────────────────────────────────────────────────

func handleDecode(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== DECODE REQUEST RECEIVED ===")
	var req struct {
		EncodedStr   string `json:"encoded"`
		DataShards   int    `json:"dataShards"`
		ParityShards int    `json:"parityShards"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonErr(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	encoded, err := hex.DecodeString(strings.TrimPrefix(req.EncodedStr, "0x"))
	if err != nil {
		jsonErr(w, http.StatusBadRequest, "invalid hex data")
		return
	}
	dataShards, parityShards := 10, 4
	if req.DataShards > 0 {
		dataShards = req.DataShards
	}
	if req.ParityShards > 0 {
		parityShards = req.ParityShards
	}
	rs, err := healchain.New(dataShards, parityShards)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "RS init failed: "+err.Error())
		return
	}
	decoded, err := rs.Decode(encoded)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "decode failed: "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(decoded)
	fmt.Println("✅ Decode completed, length:", len(decoded))
}

// ── handleStore (simple) ──────────────────────────────────────────────────────

func handleStore(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== STORE REQUEST RECEIVED ===")
	var req struct {
		DataStr      string `json:"data"`
		Label        string `json:"label"`
		DataShards   uint8  `json:"dataShards"`
		ParityShards uint8  `json:"parityShards"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonErr(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	data, err := hex.DecodeString(strings.TrimPrefix(req.DataStr, "0x"))
	if err != nil {
		jsonErr(w, http.StatusBadRequest, "invalid hex data")
		return
	}
	if req.Label == "" {
		req.Label = "from-go-service"
	}
	if req.DataShards == 0 {
		req.DataShards = 10
	}
	if req.ParityShards == 0 {
		req.ParityShards = 4
	}
	if err := validateInput(data, req.Label); err != nil {
		jsonErr(w, http.StatusBadRequest, err.Error())
		return
	}
	txHash, _, _, _, err := storeOnChain(data, req.DataShards, req.ParityShards, req.Label)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonOK(w, map[string]string{"status": "success", "tx": txHash})
}

// ── handleStoreOnChain (full response) ───────────────────────────────────────

func handleStoreOnChain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== STORE ON CHAIN REQUEST RECEIVED ===")
	var req struct {
		DataStr      string `json:"data"`
		Label        string `json:"label"`
		DataShards   uint8  `json:"dataShards"`
		ParityShards uint8  `json:"parityShards"`
		Network      string `json:"network"` // "devnet" (default) or "sepolia"
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonErr(w, http.StatusBadRequest, "invalid JSON: "+err.Error())
		return
	}
	if req.DataStr == "" {
		jsonErr(w, http.StatusBadRequest, "data field is required")
		return
	}
	data, err := hex.DecodeString(strings.TrimPrefix(req.DataStr, "0x"))
	if err != nil {
		jsonErr(w, http.StatusBadRequest, "invalid hex data")
		return
	}
	if req.Label == "" {
		req.Label = "from-go-service"
	}
	if req.DataShards == 0 {
		req.DataShards = 10
	}
	if req.ParityShards == 0 {
		req.ParityShards = 4
	}
	if err := validateInput(data, req.Label); err != nil {
		jsonErr(w, http.StatusBadRequest, err.Error())
		return
	}
	// ── Route to Sepolia if requested ─────────────────────────────────────
	if strings.ToLower(req.Network) == "sepolia" {
		txHash, requestID, err := sepoliaStoreOnChain(data, req.Label)
		if err != nil {
			jsonErr(w, http.StatusInternalServerError, "sepolia store failed: "+err.Error())
			return
		}
		jsonOK(w, map[string]interface{}{
			"status":      "pending",
			"network":     "sepolia",
			"tx":          txHash,
			"requestId":   requestID,
			"message":     "Store submitted to Sepolia. Oracle will fulfill within ~30s.",
			"statusUrl":   fmt.Sprintf("http://localhost:8080/sepoliaStatus?requestId=%s", requestID),
			"contract":    getEnv("SEPOLIA_CONTRACT_ADDRESS", ""),
			"explorerUrl": fmt.Sprintf("https://sepolia.etherscan.io/tx/%s", txHash),
		})
		return
	}
	fmt.Printf("StoreOnChain: %d bytes, label: %s, shards: %d/%d\n",
		len(data), req.Label, req.DataShards, req.ParityShards)

	// Compress before RS encoding
	compressed, err := compressData(data)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "compression failed: "+err.Error())
		return
	}
	fmt.Printf("Compressed: %d → %d bytes (%.0f%%)\n",
		len(data), len(compressed), float64(len(compressed))/float64(len(data))*100)

	txHash, recordID, originalSize, encodedSize, err := storeOnChain(
		compressed, req.DataShards, req.ParityShards, req.Label)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	jsonOK(w, map[string]interface{}{
		"status":       "success",
		"tx":           txHash,
		"recordId":     recordID,
		"label":        req.Label,
		"bytes":        len(data),
		"originalSize": originalSize,
		"encodedSize":  encodedSize,
		"retrieveUrl":  fmt.Sprintf("http://localhost:8080/retrieve?id=%s", recordID),
	})
}

// ── storeOnChain (core logic) ─────────────────────────────────────────────────

func storeOnChain(data []byte, dataShards, parityShards uint8, label string) (txHash, recordID, originalSize, encodedSize string, err error) {
	if err := checkGeth(); err != nil {
		return "", "", "", "", fmt.Errorf("geth not reachable: %w", err)
	}
	client, err := gethClient()
	if err != nil {
		return "", "", "", "", fmt.Errorf("failed to connect to geth: %w", err)
	}
	defer client.Close()

	contractAddr := common.HexToAddress(contractAddress)
	instance, err := binding.NewHealChainStorage(contractAddr, client)
	if err != nil {
		return "", "", "", "", fmt.Errorf("failed to load contract: %w", err)
	}
	privateKey, err := crypto.HexToECDSA(storePrivateKey)
	if err != nil {
		return "", "", "", "", fmt.Errorf("invalid private key: %w", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		return "", "", "", "", fmt.Errorf("failed to create transactor: %w", err)
	}
	auth.GasLimit = 10_000_000

	var tx *types.Transaction
	for attempt := 1; attempt <= 3; attempt++ {
		tx, err = instance.Store(auth, data, dataShards, parityShards, label)
		if err == nil {
			break
		}
		fmt.Printf("Store attempt %d failed: %v\n", attempt, err)
		time.Sleep(time.Duration(attempt) * 500 * time.Millisecond)
	}
	if err != nil {
		return "", "", "", "", fmt.Errorf("contract store failed after retries: %w", err)
	}

	fmt.Printf("Tx submitted: %s\n", tx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		return "", "", "", "", fmt.Errorf("failed to get receipt: %w", err)
	}
	if receipt.Status == 0 {
		return "", "", "", "", fmt.Errorf("transaction reverted")
	}

	var id *big.Int
	var orig, enc string
	for _, log := range receipt.Logs {
		event, parseErr := instance.ParseStored(*log)
		if parseErr == nil {
			id = event.Id
			orig = event.OriginalSize.String()
			enc = event.EncodedSize.String()
			break
		}
	}
	if id == nil {
		return tx.Hash().Hex(), "", "", "", fmt.Errorf("could not parse record ID")
	}

	fmt.Printf("✅ Stored! Record ID: %s | Original: %s | Encoded: %s | Tx: %s\n",
		id.String(), orig, enc, tx.Hash().Hex())

	return tx.Hash().Hex(), id.String(), orig, enc, nil
}

// ── handleRetrieve ────────────────────────────────────────────────────────────

func handleRetrieve(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		jsonErr(w, http.StatusBadRequest, "missing ?id= parameter")
		return
	}
	id := new(big.Int)
	if _, ok := id.SetString(idStr, 10); !ok {
		jsonErr(w, http.StatusBadRequest, "invalid id: must be an integer")
		return
	}
	fmt.Printf("=== RETRIEVE id=%s ===\n", idStr)

	if err := checkGeth(); err != nil {
		jsonErr(w, http.StatusInternalServerError, "geth not reachable: "+err.Error())
		return
	}
	client, err := gethClient()
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "failed to connect to geth: "+err.Error())
		return
	}
	defer client.Close()

	contractAddr := common.HexToAddress(contractAddress)
	contractABI, err := abi.JSON(strings.NewReader(binding.HealChainStorageABI))
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "failed to parse ABI: "+err.Error())
		return
	}
	callData, err := contractABI.Pack("retrieve", id)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "failed to pack calldata: "+err.Error())
		return
	}
	result, err := client.CallContract(context.Background(), ethereum.CallMsg{
		To:   &contractAddr,
		Data: callData,
	}, nil)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "contract call failed: "+err.Error())
		return
	}
	out, err := contractABI.Unpack("retrieve", result)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "failed to unpack result: "+err.Error())
		return
	}
	decoded, ok := out[0].([]byte)
	if !ok {
		jsonErr(w, http.StatusInternalServerError, "unexpected return type")
		return
	}

	// Decompress after RS decoding
	decompressed, err := decompressData(decoded)
	if err != nil {
		// Fallback: data may not be compressed (backward compat)
		fmt.Printf("⚠ Decompress failed, returning raw: %v\n", err)
		decompressed = decoded
	}

	fmt.Printf("✅ Retrieved record %s: %d bytes (decompressed from %d)\n",
		idStr, len(decompressed), len(decoded))
	jsonOK(w, map[string]interface{}{
		"status":   "success",
		"recordId": idStr,
		"data":     "0x" + hex.EncodeToString(decompressed),
		"text":     string(decompressed),
		"bytes":    len(decompressed),
	})
}

// ── handleGetMetadata ─────────────────────────────────────────────────────────

func handleGetMetadata(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		jsonErr(w, http.StatusBadRequest, "missing ?id= parameter")
		return
	}
	id := new(big.Int)
	if _, ok := id.SetString(idStr, 10); !ok {
		jsonErr(w, http.StatusBadRequest, "invalid id: must be an integer")
		return
	}
	fmt.Printf("=== GET METADATA id=%s ===\n", idStr)

	if err := checkGeth(); err != nil {
		jsonErr(w, http.StatusInternalServerError, "geth not reachable: "+err.Error())
		return
	}
	client, err := gethClient()
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "failed to connect: "+err.Error())
		return
	}
	defer client.Close()

	contractAddr := common.HexToAddress(contractAddress)
	instance, err := binding.NewHealChainStorage(contractAddr, client)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "failed to load contract: "+err.Error())
		return
	}

	meta, err := instance.GetMetadata(nil, id)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "getMetadata failed: "+err.Error())
		return
	}

	jsonOK(w, map[string]interface{}{
		"status":       "success",
		"recordId":     idStr,
		"dataHash":     fmt.Sprintf("0x%x", meta.DataHash),
		"originalSize": meta.OriginalSize.String(),
		"encodedSize":  meta.EncodedSize.String(),
		"dataShards":   meta.DataShards,
		"parityShards": meta.ParityShards,
		"owner":        meta.Owner.Hex(),
		"timestamp":    meta.Timestamp.String(),
		"label":        meta.Label,
	})
}

// ── handleListRecords ─────────────────────────────────────────────────────────

func handleListRecords(w http.ResponseWriter, r *http.Request) {
	fmt.Println("=== LIST RECORDS ===")

	// Parse pagination params
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 0
	limit := 10

	if p, err := strconv.Atoi(pageStr); err == nil && p >= 0 {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 && l <= 50 {
		limit = l
	}

	if err := checkGeth(); err != nil {
		jsonErr(w, http.StatusInternalServerError, "geth not reachable: "+err.Error())
		return
	}
	client, err := gethClient()
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "failed to connect: "+err.Error())
		return
	}
	defer client.Close()

	contractAddr := common.HexToAddress(contractAddress)
	instance, err := binding.NewHealChainStorage(contractAddr, client)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "failed to load contract: "+err.Error())
		return
	}

	total, err := instance.TotalRecords(nil)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "totalRecords failed: "+err.Error())
		return
	}

	totalInt := int(total.Int64())
	start := page * limit
	end := start + limit
	if end > totalInt {
		end = totalInt
	}

	type Record struct {
		ID           string `json:"id"`
		Label        string `json:"label"`
		OriginalSize string `json:"originalSize"`
		EncodedSize  string `json:"encodedSize"`
		Owner        string `json:"owner"`
		Timestamp    string `json:"timestamp"`
		DataShards   uint8  `json:"dataShards"`
		ParityShards uint8  `json:"parityShards"`
	}

	records := []Record{}

	for i := start; i < end; i++ {
		id := big.NewInt(int64(i))
		meta, err := instance.GetMetadata(nil, id)
		if err != nil {
			continue
		}
		records = append(records, Record{
			ID:           fmt.Sprintf("%d", i),
			Label:        meta.Label,
			OriginalSize: meta.OriginalSize.String(),
			EncodedSize:  meta.EncodedSize.String(),
			Owner:        meta.Owner.Hex(),
			Timestamp:    meta.Timestamp.String(),
			DataShards:   meta.DataShards,
			ParityShards: meta.ParityShards,
		})
	}

	jsonOK(w, map[string]interface{}{
		"status":  "success",
		"total":   totalInt,
		"page":    page,
		"limit":   limit,
		"pages":   (totalInt + limit - 1) / limit,
		"records": records,
	})
}

// ── handleDelete ──────────────────────────────────────────────────────────────

func handleDelete(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		jsonErr(w, http.StatusBadRequest, "missing ?id= parameter")
		return
	}
	id := new(big.Int)
	if _, ok := id.SetString(idStr, 10); !ok {
		jsonErr(w, http.StatusBadRequest, "invalid id: must be an integer")
		return
	}
	fmt.Printf("=== DELETE id=%s ===\n", idStr)

	if err := checkGeth(); err != nil {
		jsonErr(w, http.StatusInternalServerError, "geth not reachable: "+err.Error())
		return
	}
	client, err := gethClient()
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "failed to connect: "+err.Error())
		return
	}
	defer client.Close()

	contractAddr := common.HexToAddress(contractAddress)
	instance, err := binding.NewHealChainStorage(contractAddr, client)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "failed to load contract: "+err.Error())
		return
	}

	privateKey, err := crypto.HexToECDSA(storePrivateKey)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "invalid private key: "+err.Error())
		return
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "failed to create transactor: "+err.Error())
		return
	}
	auth.GasLimit = 1_000_000

	tx, err := instance.Remove(auth, id)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "delete failed (you may not be the owner): "+err.Error())
		return
	}

	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "failed to get receipt: "+err.Error())
		return
	}
	if receipt.Status == 0 {
		jsonErr(w, http.StatusForbidden, "transaction reverted — you may not be the owner of this record")
		return
	}

	fmt.Printf("✅ Deleted record %s | Tx: %s\n", idStr, tx.Hash().Hex())
	jsonOK(w, map[string]interface{}{
		"status":   "success",
		"recordId": idStr,
		"tx":       tx.Hash().Hex(),
	})
}

// ── handleMetrics ─────────────────────────────────────────────────────────────

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	uptime := int64(time.Since(serviceStartedAt).Seconds())

	// Geth (devnet) reachability — best-effort, non-blocking on errors
	gethStatus := "ok"
	gethBlock := uint64(0)
	gethChainID := ""
	if client, err := gethClient(); err != nil {
		gethStatus = "unreachable"
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		if block, err := client.BlockByNumber(ctx, nil); err == nil {
			gethBlock = block.NumberU64()
		} else {
			gethStatus = "error"
		}
		if id, err := client.ChainID(ctx); err == nil {
			gethChainID = id.String()
		}
		client.Close()
	}

	// Devnet contract record count — best-effort
	totalRecords := ""
	if client, err := gethClient(); err == nil {
		instance, err2 := binding.NewHealChainStorage(common.HexToAddress(contractAddress), client)
		if err2 == nil {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			if t, err3 := instance.TotalRecords(&bind.CallOpts{Context: ctx}); err3 == nil {
				totalRecords = t.String()
			}
		}
		client.Close()
	}

	resp := map[string]interface{}{
		"service": map[string]interface{}{
			"version":       serviceVersion,
			"uptimeSeconds": uptime,
			"startedAt":     serviceStartedAt.UTC().Format(time.RFC3339),
			"listenAddr":    listenAddr,
		},
		"devnet": map[string]interface{}{
			"gethUrl":      gethURL,
			"gethStatus":   gethStatus,
			"chainId":      gethChainID,
			"latestBlock":  gethBlock,
			"contract":     contractAddress,
			"totalRecords": totalRecords,
		},
	}

	if sepoliaOracle != nil {
		resp["oracle"] = sepoliaOracle.Metrics()
	} else {
		resp["oracle"] = map[string]interface{}{
			"enabled":      false,
			"healthy":      false,
			"healthReason": "SEPOLIA_RPC_URL or SEPOLIA_CONTRACT_ADDRESS not configured",
		}
	}

	if serviceLimiter != nil {
		resp["rateLimit"] = serviceLimiter.Stats()
	}

	jsonOK(w, resp)
}

// ── handleSepoliaStatus ───────────────────────────────────────────────────────

func handleSepoliaStatus(w http.ResponseWriter, r *http.Request) {
	requestIdStr := r.URL.Query().Get("requestId")
	if requestIdStr == "" {
		jsonErr(w, http.StatusBadRequest, "missing ?requestId= parameter")
		return
	}

	sepoliaURL := getEnv("SEPOLIA_RPC_URL", "")
	contractAddr := getEnv("SEPOLIA_CONTRACT_ADDRESS", "")

	if sepoliaURL == "" || contractAddr == "" {
		jsonErr(w, http.StatusInternalServerError, "Sepolia not configured")
		return
	}

	client, err := ethclient.Dial(sepoliaURL)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "failed to connect to Sepolia: "+err.Error())
		return
	}
	defer client.Close()

	addr := common.HexToAddress(contractAddr)
	instance, err := bindingsepolia.NewHealChainStorage(addr, client)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "failed to load contract: "+err.Error())
		return
	}

	requestId := new(big.Int)
	if _, ok := requestId.SetString(requestIdStr, 10); !ok {
		jsonErr(w, http.StatusBadRequest, "invalid requestId")
		return
	}

	pending, err := instance.IsPending(nil, requestId)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "isPending failed: "+err.Error())
		return
	}

	if pending {
		jsonOK(w, map[string]interface{}{
			"status":    "pending",
			"requestId": requestIdStr,
			"message":   "Oracle is processing — check back in a few seconds",
		})
		return
	}

	// Not pending means fulfilled — find the record ID from total records
	total, err := instance.TotalRecords(nil)
	if err != nil {
		jsonErr(w, http.StatusInternalServerError, "totalRecords failed: "+err.Error())
		return
	}

	jsonOK(w, map[string]interface{}{
		"status":       "fulfilled",
		"requestId":    requestIdStr,
		"totalRecords": total.String(),
		"message":      "Store fulfilled by oracle",
		"contract":     contractAddr,
	})
}
