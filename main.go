package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("content-type", "application/json")
		_, _ = fmt.Fprintln(w, `{"ok":true}`)
	})
	mux.HandleFunc("/version", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "v1")
	})
	return mux
}

func main() {
	mux := newMux()
	addr := os.Getenv("LISTEN_ADDR")
	if addr == "" {
		addr = ":8081"
	}
	fmt.Println("listening on", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
