package department

import (
	"database/sql"
)

func GetAllDepartments(db *sql.DB) ([]Department, error) {
	s :=
	`
		select department.id, department.name from department
	`
	rows, err := db.Query(s)
	if err != nil {
		return nil, err
	}

	var depts []Department

	for rows.Next() {
		var dept Department
		err := rows.Scan(
			&dept.Id,
			&dept.Name,
		)
		if err != nil {
			return nil, err
		}
		depts = append(depts, dept)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return depts, nil
}
