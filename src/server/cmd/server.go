package main


import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/adamc25/db/internal/database"
	"github.com/adamc25/db/internal/handlers"
)

var db *sql.DB

func main() {
	connStr := database.GetConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
  err = db.Ping()
  if err != nil {
      log.Fatal("Failed to connect to the database: ", err)
  }
  http.HandleFunc("/department/list", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllDepartments(db, w, r)
	})
	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
    handlers.EnableCors(w, r)
    switch r.Method {
    case http.MethodDelete:
    	handlers.DeleteEmployeeById(db, w, r)
    case http.MethodGet:
      handlers.GetEmployeeById(db, w, r)
    case http.MethodPut:
      handlers.UpdateEmployee(db, w, r)
    case http.MethodPost:
    	handlers.CreateEmployee(db, w, r)
    default:
      http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
    }
	})
	http.HandleFunc("/employees/list", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetEmployees(db, w, r)
	})
	http.HandleFunc("/employment_status/list", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllEmploymentStatus(db, w, r)
	})
	http.HandleFunc("/search/list", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllSearchableFields(db, w, r)
	})
	http.ListenAndServe(":8080", nil)
}
