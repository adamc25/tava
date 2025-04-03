package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"database/sql"
	"strconv"

	"github.com/adamc25/db/internal/meta"
	"github.com/adamc25/db/internal/employee"
)

func CreateEmployee(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	EnableCors(w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Only POST is allowed.", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var e employee.Employee
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	e, err = e.Create(db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating employee: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(e)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response to JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteEmployeeById(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	EnableCors(w, r)

	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method. Only DELETE is allowed.", http.StatusMethodNotAllowed)
		return
	}

	queryParams := r.URL.Query()

	idParam := queryParams.Get("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid id value in queryParams.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	e, err := employee.GetEmployeeById(db, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching employee to delete: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	e, err = e.Delete(db)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error deleting employee: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(e)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response to JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetEmployeeById(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	EnableCors(w, r)

	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method. Only GET is allowed.", http.StatusMethodNotAllowed)
		return
	}

	queryParams := r.URL.Query()

	idParam := queryParams.Get("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid id value in queryParams.", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	e, err := employee.GetEmployeeById(db, id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching employee: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(e)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding response to JSON: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetEmployees(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	EnableCors(w, r)

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Only POST is allowed.", http.StatusMethodNotAllowed)
		return
	}

	var m meta.QueryMeta
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	l, err := employee.GetList(db, m)
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

func UpdateEmployee(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	EnableCors(w, r)

	if r.Method != http.MethodPut {
		http.Error(w, "Invalid request method. Only PUT is allowed.", http.StatusMethodNotAllowed)
		return
	}

	var e employee.Employee
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, "Error decoding JSON: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Print("e is ", e)

	e, err = e.Update(db)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error Update Employee: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
