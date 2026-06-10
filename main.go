package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handleHealth)

	// TODO: register your P&L endpoint here, e.g. GET /pnl

	log.Println("listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
