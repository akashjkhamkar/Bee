package main

import (
	"fmt"
	"html"
	"net/http"
)

func Entry(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there, %q", html.EscapeString(r.URL.Path))
}