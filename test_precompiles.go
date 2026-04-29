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

    fmt.Println("=== HealChain Precompile Test Client ===")
    fmt.Println("✅ Connected successfully to HealChain Devnet!\n")

    payload := []byte("Ci-RS 20-byte test!")

    // CiSHA4096
    fmt.Println("1. Testing CiSHA4096...")
    c17 := common.HexToAddress("0x0000000000000000000000000000000000000c17")
    hash, err := client.CallContract(context.Background(), ethereum.CallMsg{
        To:   &c17,
        Data: payload,
    }, nil)
    if err != nil {
        log.Fatal("CiSHA4096 failed:", err)
    }
    fmt.Printf("Input : %s\n", payload)
    fmt.Printf("Output length: %d bytes\n", len(hash))
    fmt.Printf("First 32 bytes: %x...\n", hash[:32])

    // CiRSRepair
    fmt.Println("\n2. Testing CiRSRepair (Encode + redundancy 8)...")
    c18 := common.HexToAddress("0x0000000000000000000000000000000000000c18")
    encodeInput := append([]byte{0x01, 8}, payload...)

    encoded, err := client.CallContract(context.Background(), ethereum.CallMsg{
        To:   &c18,
        Data: encodeInput,
    }, nil)
    if err != nil {
        log.Fatal("CiRSRepair failed:", err)
    }

    fmt.Printf("Encoded length: %d bytes\n", len(encoded))
    fmt.Printf("First 80 bytes: %x...\n", encoded[:80])

    fmt.Println("\n🎉 Success! Both precompiles are working on your HealChain devnet.")
}
