package employmentstatus

import (
	"database/sql"
	"errors"
	"fmt"
)

type EmploymentStatus struct {
	Id int `json:"id"`
	Status string `json:"status"`
}

func (e EmploymentStatus) GetEmploymentStatusByStatus (db *sql.DB) (EmploymentStatus, error) {
	if e.Status == "" {
		return EmploymentStatus{}, errors.New("Cannot look up a employment status without a status")
	}
	s := 
	`
		select employment_status.id, employment_status.status from employment_status where employment_status.status ILIKE $1
	`
	var newE EmploymentStatus
	err := db.QueryRow(s, e.Status).Scan(
		&newE.Id,
		&newE.Status,
	)
	if err == sql.ErrNoRows {
		return EmploymentStatus{}, errors.New(fmt.Sprintf("Wasn't able to find a employment status with the status of %s", e.Status))
	}
	if err != nil {
		return EmploymentStatus{}, err
	}
	return newE, nil
}