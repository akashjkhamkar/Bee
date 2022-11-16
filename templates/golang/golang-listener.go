package main

import (
	"log"
	"net/http"
)

func main() {
    http.HandleFunc("/", Entry)
    log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}