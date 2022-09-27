package main

import (
    "net/http"
    "log"
)

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(
        http.ListenAndServe("localhost:8000", nil),
    )
}
