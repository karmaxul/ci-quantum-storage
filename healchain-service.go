package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "encoding/hex" 
    "strings"

    "ci-sha-test/healchain"
)

func main() {
    http.HandleFunc("/encode", handleEncode)
    http.HandleFunc("/decode", handleDecode)
    http.HandleFunc("/health", handleHealth)
    http.HandleFunc("/stats", handleStats)

    fmt.Println("🚀 HealChain Self-Healing Service v2.2 (Enhanced)")
    fmt.Println("   Endpoints: /encode, /decode, /health, /stats")
    fmt.Println("   Ready on http://localhost:8080")

    http.HandleFunc("/nuclear-test", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("=== NUCLEAR TEST ROUTE HIT ===")
        w.Header().Set("Content-Type", "text/plain")
        w.Write([]byte("Nuclear test successful"))
    })

    fmt.Println("Listening on :8080...")

    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("ListenAndServe error:", err)
    }
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "healthy", "version": "2.2"})
}

func handleStats(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "running"})
}

func handleEncode(w http.ResponseWriter, r *http.Request) {
    fmt.Println("=== ENCODE REQUEST RECEIVED FROM PRECOMPILE ===")

    var req struct {
        DataStr      string `json:"data"`
        DataShards   int    `json:"dataShards"`
        ParityShards int    `json:"parityShards"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        fmt.Println("JSON decode error:", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    data, err := hex.DecodeString(strings.TrimPrefix(req.DataStr, "0x"))
    if err != nil {
        fmt.Println("Hex decode error:", err)
        http.Error(w, "invalid hex data", http.StatusBadRequest)
        return
    }

    fmt.Println("Data length:", len(data))
    fmt.Println("Shards:", req.DataShards, req.ParityShards)

    dataShards := 10
    parityShards := 4
    if req.DataShards > 0 {
        dataShards = req.DataShards
    }
    if req.ParityShards > 0 {
        parityShards = req.ParityShards
    }

    rs, err := healchain.New(dataShards, parityShards)
    if err != nil {
        fmt.Println("New RS error:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    encoded, err := rs.Encode(data)
    if err != nil {
        fmt.Println("Encode error:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(map[string]interface{}{
    "success": true,
    "data": hex.EncodeToString(encoded),  // or decoded
})
}

func handleDecode(w http.ResponseWriter, r *http.Request) {
    fmt.Println("=== DECODE REQUEST RECEIVED FROM PRECOMPILE ===")

    var req struct {
        EncodedStr   string `json:"encoded"`
        DataShards   int    `json:"dataShards"`
        ParityShards int    `json:"parityShards"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        fmt.Println("JSON decode error:", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    encoded, err := hex.DecodeString(strings.TrimPrefix(req.EncodedStr, "0x"))
    if err != nil {
        fmt.Println("Hex decode error:", err)
        http.Error(w, "invalid hex data", http.StatusBadRequest)
        return
    }

    fmt.Println("Encoded length:", len(encoded))
    fmt.Println("Shards:", req.DataShards, req.ParityShards)

    dataShards := 10
    parityShards := 4
    if req.DataShards > 0 {
        dataShards = req.DataShards
    }
    if req.ParityShards > 0 {
        parityShards = req.ParityShards
    }

    rs, err := healchain.New(dataShards, parityShards)
    if err != nil {
        fmt.Println("New RS error:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    decoded, err := rs.Decode(encoded)
    if err != nil {
        fmt.Println("Decode error:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/octet-stream")
    w.Write(decoded)
    fmt.Println("✅ Decode completed, length:", len(decoded))
}
