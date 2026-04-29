package main

import (
    "bytes"
    "fmt"

    "ci-sha-test/healchain"
)

func main() {
    fmt.Println("=== Full HealChain Test (Go + Precompile) ===")

    pre, _ := healchain.NewPrecompile(10, 4)
    original := []byte("Harmony HealChain full system test payload - should recover perfectly")

    // Encode
    encoded, _ := pre.Encode(original)
    fmt.Printf("Encoded: %d bytes\n", len(encoded))

    // Corrupt
    encoded[250] ^= 0xFF
    fmt.Println("1-byte corruption applied")

    // Heal
    recovered, err := pre.Decode(encoded)
    if err != nil {
        fmt.Println("Heal failed:", err)
    } else if bytes.Equal(recovered, original) {
        fmt.Println("🎉 FULL SYSTEM SUCCESS - Self-healing works end-to-end!")
    } else {
        fmt.Println("Partial recovery only")
    }
}
