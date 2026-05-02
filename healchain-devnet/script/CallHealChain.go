package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "time"
)

type Request struct {
    Operation    string `json:"operation"`
    Data         string `json:"data"`
    DataShards   uint8  `json:"dataShards"`
    ParityShards uint8  `json:"parityShards"`
}

type Response struct {
    Success bool   `json:"success"`
    Result  string `json:"result"`
    GasUsed uint64 `json:"gasUsed"`
    Error   string `json:"error,omitempty"`
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println(`{"success":false,"error":"No input"}`)
        return
    }

    var req Request
    if err := json.Unmarshal([]byte(os.Args[1]), &req); err != nil {
        fmt.Printf(`{"success":false,"error":"JSON parse: %s"}`, err)
        return
    }

    // Call your real HealChain Go service
    client := &http.Client{Timeout: 10 * time.Second}
    url := "http://localhost:8080/" + req.Operation

    payload := map[string]interface{}{
        "data":         req.Data,
        "dataShards":   req.DataShards,
        "parityShards": req.ParityShards,
    }

    jsonPayload, _ := json.Marshal(payload)
    resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonPayload))
    if err != nil {
        fmt.Printf(`{"success":false,"error":"HTTP error: %s"}`, err)
        return
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)

    var result Response
    result.Success = resp.StatusCode == 200
    result.Result = string(body)
    result.GasUsed = 42069

    out, _ := json.Marshal(result)
    fmt.Println(string(out))
}
