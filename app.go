package main

import (
    "bytes"
    "fmt"

    "ci-sha-test/healchain"
)

func main() {
    fmt.Println("=== HealChain Self-Healing System ===")

    pre, err := healchain.NewPrecompile(10, 4)
    if err != nil {
        panic(err)
    }

    fmt.Printf("✅ Precompile registered at %s\n\n", healchain.HealChainPrecompileAddress)

    original := []byte("Harmony HealChain self-healing payload test - 89 byte message that should recover cleanly now")

    encoded, _ := pre.Encode(original)
    fmt.Printf("Encoded: %d bytes\n", len(encoded))

    encoded[300] ^= 0xFF // simulate corruption
    fmt.Println("Corruption applied...")

    recovered, err := pre.Decode(encoded)
    if err != nil {
        fmt.Println("❌ Heal failed:", err)
    } else if bytes.Equal(recovered, original) {
        fmt.Println("🎉 PERFECT SELF-HEALING SUCCESS!")
    }
}
