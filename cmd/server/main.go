package main

import (
	"log"
	"net/http"

	"github.com/Lomesh21/GateKeeper/internal/gateway"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	proxyHandler, err := gateway.ProxyHandler("http://localhost:8081")
	if err != nil {
		log.Fatalf("Failed to create proxy handler: %v", err)
	}

	r.HandleFunc("/books/*", proxyHandler)

	log.Println("API Gateway running on :8080")
	http.ListenAndServe(":8080", r)
}
