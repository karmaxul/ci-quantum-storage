package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum"
)

func main() {
    client, err := ethclient.Dial("http://localhost:8545")
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()

    fmt.Println("=== HealChain Full Self-Healing Demo ===")
payload := []byte("This is a longer 32-byte test payload for HealChain!")

    c18 := common.HexToAddress("0x0000000000000000000000000000000000000c18")

    // 1. Encode
    fmt.Println("\n1. Encoding with redundancy 8...")
    encodeInput := append([]byte{0x01, 8}, payload...)

    encoded, err := client.CallContract(context.Background(), ethereum.CallMsg{
        To:   &c18,
        Data: encodeInput,
    }, nil)
    if err != nil {
        log.Fatal("Encode failed:", err)
    }

    fmt.Printf("Encoded length: %d bytes\n", len(encoded))

    // 2. Simulate corruption
    if len(encoded) > 40 {
        encoded[40] ^= 0xFF
        fmt.Println("2. Simulated 1-byte corruption")
    }

    // 3. Repair
    fmt.Println("3. Attempting self-healing repair...")
    repaired, err := client.CallContract(context.Background(), ethereum.CallMsg{
        To:   &c18,
        Data: append([]byte{0x02}, encoded...),
    }, nil)
    if err != nil {
        log.Fatal("Repair failed:", err)
    }

    fmt.Printf("Recovered length: %d bytes\n", len(repaired))
    fmt.Printf("Recovered data : %s\n", repaired)

    if len(repaired) >= len(payload) && string(repaired[:len(payload)]) == string(payload) {
        fmt.Println("\n🎉 SUCCESS: Self-healing worked perfectly!")
    } else {
        fmt.Println("\n⚠️  Partial or failed recovery")
    }
}
