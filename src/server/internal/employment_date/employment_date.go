package employmentdate

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/adamc25/db/internal/database"
)

type EmploymentDate struct {
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
	EmployeeId int `json:"-"`
}

func (d EmploymentDate) Create (db *sql.DB) (EmploymentDate, error) {
	if d.EmployeeId == 0 {
		return EmploymentDate{}, errors.New("Employee id must not be 0")
	}
	tx, err := db.Begin()
	if err != nil {
		return d, errors.New("Failed to begin a transaction to create new employee record")
	}
	defer tx.Rollback()
	s := 
	`
		INSERT INTO EMPLOYMENT_DATE (start_date, end_date, employee_id)
		VALUES ($1, $2, $3)
	`
	_, err = tx.Exec(
		s,
		database.GetNullableString(d.StartDate),
		database.GetNullableString(d.EndDate),
		d.EmployeeId,
	)
	if err != nil {
		return d, errors.New(fmt.Sprintf("Failed to create an employment date for employee with id of %d", d.EmployeeId))
	}
	tx.Commit()
	return d, nil
}

func (d EmploymentDate) Delete (db *sql.DB) (EmploymentDate, error) {
	if d.EmployeeId == 0 {
		return EmploymentDate{}, errors.New("employment date id must not be 0")
	}
	tx, err := db.Begin()
	if err != nil {
		return d, errors.New("Failed to begin a transaction to create new employement date record")
	}
	defer tx.Rollback()
	s := 
	`
		DELETE FROM EMPLOYMENT_DATE
		WHERE EMPLOYMENT_DATE.employee_id = $1
	`
	_, err = tx.Exec(s, d.EmployeeId)
	if err != nil {
		return d, errors.New(fmt.Sprintf("Failed to delete the employee with id of %d", d.EmployeeId))
	}
	tx.Commit()
	return d, nil
}