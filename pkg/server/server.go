package server

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dcmcand/camp-quansight-2025-go-bof/pkg/even"
)

type NumberRequest struct {
	Number int `json:"number"`
}

type EvenResponse struct {
	IsEven bool `json:"is_even"`
}

type HealthResponse struct {
	Status string `json:"status"`
}

func isEvenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req NumberRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	isEven, err := even.IsEven(req.Number)
	if err != nil {
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		slog.Error("Error in IsEven", "error", err, "number", req.Number)
		return
	}
	response := EvenResponse{IsEven: isEven}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := HealthResponse{Status: "ok"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func ListenaAndServe(port string) error {
	http.HandleFunc("/is_even", isEvenHandler)
	http.HandleFunc("/health", healthHandler)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
