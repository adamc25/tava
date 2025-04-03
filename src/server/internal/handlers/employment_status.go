package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"

	"github.com/adamc25/db/internal/employment_status"
)


func GetAllEmploymentStatus(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	EnableCors(w, r)

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method. Only GET is allowed.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	l, err := employmentstatus.GetAllEmploymentStatus(db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching list: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(l)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response to JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}