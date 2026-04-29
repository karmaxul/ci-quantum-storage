package main

import (
    "fmt"
    "log"
    "net/http"

    "ci-sha-test/healchain"
)

func main() {
    fmt.Println("=== HealChain Precompile Service ===")
    rs, _ := healchain.New(10, 4)

    http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
        data := []byte(r.URL.Query().Get("data"))
        encoded, _ := rs.Encode(data)
        w.Write(encoded)
    })

    http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
        encoded := []byte(r.URL.Query().Get("data"))
        decoded, _ := rs.Decode(encoded)
        w.Write(decoded)
    })

    log.Println("HealChain service running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
