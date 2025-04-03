package handlers

import (
	"net/http"
)

func EnableCors(w http.ResponseWriter, r *http.Request) {
	// Handle preflight request (OPTIONS method)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
	}
}
