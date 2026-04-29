package main

import (
    "bytes"
    "fmt"

    "ci-sha-test/healchain"
)

func main() {
    fmt.Println("=== HealChain Quick Demo ===")

    rs, _ := healchain.New(10, 4)
    original := []byte("Harmony HealChain self-healing test payload")

    encoded, _ := rs.Encode(original)
    encoded[200] ^= 0xFF

    recovered, _ := rs.Decode(encoded)
    if bytes.Equal(recovered, original) {
        fmt.Println("🎉 Self-healing works perfectly!")
    }
}
