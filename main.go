package main

import (
	_ "embed"
	"log"
	"net/http"
	"os"
)

//go:embed index.html
var indexHTML []byte

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func main() {
	version := env("VERSION", "1.0.0")
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok","service":"shop-ui","version":"` + version + `"}`))
	})
	// dilayani di prefix /shop oleh VirtualService gateway
	mux.HandleFunc("GET /shop", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(indexHTML)
	})
	addr := ":" + env("PORT", "8080")
	log.Printf("shop-ui %s listening on %s", version, addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
