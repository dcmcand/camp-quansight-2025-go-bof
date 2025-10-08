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
	slog.Info("Request received", "method", r.Method, "path", r.URL.Path, "remote_addr", r.RemoteAddr)

	if r.Method != http.MethodPost {
		slog.Warn("Method not allowed", "method", r.Method, "path", r.URL.Path)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req NumberRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		slog.Error("Invalid request body", "error", err, "path", r.URL.Path)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	isEven, err := even.IsEven(req.Number)
	if err != nil {
		slog.Error("Error processing IsEven request", "error", err, "number", req.Number)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}
	response := EvenResponse{IsEven: isEven}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	slog.Info("Request completed successfully", "number", req.Number, "is_even", isEven)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	slog.Info("Health check received", "method", r.Method, "path", r.URL.Path, "remote_addr", r.RemoteAddr)

	if r.Method != http.MethodGet {
		slog.Warn("Method not allowed on health endpoint", "method", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := HealthResponse{Status: "ok"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
	slog.Debug("Health check completed successfully")
}

func ListenaAndServe(port string) error {
	http.HandleFunc("/", isEvenHandler)
	http.HandleFunc("/health", healthHandler)

	addr := fmt.Sprintf(":%s", port)
	slog.Info("Starting server", "port", port, "address", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		slog.Error("Server failed to start", "error", err, "port", port)
		return err
	}
	return nil
}
