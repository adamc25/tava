package department

import (
	"database/sql"
	"errors"
	"fmt"
)

type Department struct {
	Id	int	`json:"id"`
	Name string `json:"name"`
}

func GetDepartmentById (db *sql.DB, id int) (Department, error) {
	if id == 0 {
		return Department{}, errors.New("Department id = 0 is not a valid value to lookup")
	}
	s := 
	`
		select department.id, department.name from department where department.name = $1
	`
	var newDept Department
	err := db.QueryRow(s, id).Scan(
		&newDept.Id,
		&newDept.Name,
	)
	if err == sql.ErrNoRows {
		return Department{}, errors.New(fmt.Sprintf("Wasn't able to find a department with the id of %s", id))
	}
	if err != nil {
		return Department{}, err
	}
	return newDept, nil
}

func (dept Department) GetDepartmentByName (db *sql.DB) (Department, error) {
	if dept.Name == "" {
		return Department{}, errors.New("Cannot look up a department without a name")
	}
	s := 
	`
		select department.id, department.name from department where department.name ILIKE $1
	`
	var newDept Department
	err := db.QueryRow(s, dept.Name).Scan(
		&newDept.Id,
		&newDept.Name,
	)
	if err == sql.ErrNoRows {
		return Department{}, errors.New(fmt.Sprintf("Wasn't able to find a department with the name of %s", dept.Name))
	}
	if err != nil {
		return Department{}, err
	}
	return newDept, nil
}
