package employmentstatus

import (
	"database/sql"
)

func GetAllEmploymentStatus(db *sql.DB) ([]EmploymentStatus, error) {
	s :=
	`
		select employment_status.id, employment_status.status from employment_status
	`
	rows, err := db.Query(s)
	if err != nil {
		return nil, err
	}

	var employmentStatus []EmploymentStatus

	for rows.Next() {
		var e EmploymentStatus
		err := rows.Scan(
			&e.Id,
			&e.Status,
		)
		if err != nil {
			return nil, err
		}
		employmentStatus = append(employmentStatus, e)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employmentStatus, nil
}