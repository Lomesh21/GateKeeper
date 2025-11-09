package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Lomesh21/GateKeeper/internal/auth"
	"github.com/Lomesh21/GateKeeper/internal/gateway"
	"github.com/Lomesh21/GateKeeper/internal/middleware"
	"github.com/go-chi/chi/v5"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if req.Username != "admin" && req.Password != "password" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	token, err := auth.GenerateJWT(req.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}
	resp := loginResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {

	r := chi.NewRouter()

	proxyHandler, err := gateway.ProxyHandler("http://localhost:8081")
	if err != nil {
		log.Fatalf("Failed to create proxy handler: %v", err)
	}
	r.Post("/login", loginHandler)

	r.With(middleware.JWTMiddleware).HandleFunc("/books/*", proxyHandler)

	log.Println("API Gateway running on :8080")
	http.ListenAndServe(":8080", r)
}
