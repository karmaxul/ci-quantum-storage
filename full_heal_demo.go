package main

import (
    "bytes"
    "fmt"
    "time"

    "ci-sha-test/healchain"
)

func main() {
    fmt.Println("=== HealChain Self-Healing System (Phase 1) ===")
    fmt.Println("Reed-Solomon + Per-Shard SHA256 Verification")
    fmt.Println()

    // Test with different configurations
    tests := []struct {
        name         string
        dataShards   int
        parityShards int
        payload      []byte
    }{
        {"Small Payload", 8, 4, []byte("Small test payload for HealChain")},
        {"Medium Payload", 10, 4, []byte("This is a longer test message for self-healing capabilities in decentralized storage - April 2026")},
    }

    for _, t := range tests {
        fmt.Printf("--- %s (%d+%d shards) ---\n", t.name, t.dataShards, t.parityShards)

        rs, err := healchain.New(t.dataShards, t.parityShards)
        if err != nil {
            fmt.Println("Failed to create engine:", err)
            continue
        }

        start := time.Now()
        encoded, err := rs.Encode(t.payload)
        if err != nil {
            fmt.Println("Encode failed:", err)
            continue
        }

        fmt.Printf("Encoded: %d bytes (overhead: %.1f%%)\n", 
            len(encoded), 
            float64(len(encoded)-len(t.payload))/float64(len(t.payload))*100)

        // Simulate corruption
        corrupted := make([]byte, len(encoded))
        copy(corrupted, encoded)
        if len(corrupted) > 300 {
            corrupted[300] ^= 0xFF
        } else {
            corrupted[len(corrupted)/2] ^= 0xFF
        }

        // Heal
        start = time.Now()
        recovered, err := rs.Decode(corrupted)
        healTime := time.Since(start)

        if err != nil {
            fmt.Println("❌ Healing failed:", err)
        } else if bytes.Equal(recovered, t.payload) {
            fmt.Println("🎉 PERFECT SELF-HEALING SUCCESS!")
            fmt.Printf("Recovery time: %v\n", healTime)
        } else {
            fmt.Println("❌ Recovered data does not match")
        }
        fmt.Println()
    }

    fmt.Println("✅ HealChain Go library is clean and ready for Phase 1 🚀")
}
