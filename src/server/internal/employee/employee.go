package employee

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/adamc25/db/internal/database"
	"github.com/adamc25/db/internal/department"
	"github.com/adamc25/db/internal/employment_date"
	"github.com/adamc25/db/internal/employment_status"
)

type Employee struct {
	Id	int	`json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	DepartmentId int `json:"-"`
	DepartmentName string `json:"departmentName"`
	EmploymentStatusId int `json:"-"`
	EmployeeStatus string `json:"employeeStatus"`
	AvatarUrl string `json:"avatarUrl"`
	Quote string `json:"quote"`
	EmploymentDate employmentdate.EmploymentDate `json:"employmentDate"`
}

func (e Employee) Create (db *sql.DB) (Employee, error) {
	if e.Id != 0 {
		return Employee{}, errors.New("Employee id must be 0")
	}
	if e.DepartmentName == "" {
		return e, errors.New("Employee must belong to a department")
	}
	if e.FirstName == "" {
		return e, errors.New("Employee must have a first name set")
	}
	if e.LastName == "" {
		return e, errors.New("Employee must have a last name set")
	}
	if e.EmployeeStatus == "" {
		return e, errors.New("Employee must have a status set")
	}
	avatarUrl := database.GetNullableString(e.AvatarUrl)
	quote := database.GetNullableString(e.Quote)
	tx, err := db.Begin()
	if err != nil {
		return e, errors.New("Failed to begin a transaction to create new employee record")
	}
	employeeDept := department.Department{
		Name: e.DepartmentName,
	}
	employeeDept, err = employeeDept.GetDepartmentByName(db)
	if err != nil {
		return e, errors.New(fmt.Sprintf("Failed to find a matching deparment with the name of %s", e.DepartmentName))
	}
	employeeEmployment := employmentstatus.EmploymentStatus{
		Status: e.EmployeeStatus,
	}
	employeeEmployment, err = employeeEmployment.GetEmploymentStatusByStatus(db)
	if err != nil {
		return e, errors.New(fmt.Sprintf("Failed to find a matching employment status with the status of %s", e.EmployeeStatus))
	}
	defer tx.Rollback()
	s := 
	`
		INSERT INTO EMPLOYEE (first_name, last_name, department_id, employment_status_id, avatar_url, quote)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	err = tx.QueryRow(s, e.FirstName, e.LastName, employeeDept.Id, employeeEmployment.Id, avatarUrl, quote).Scan(&e.Id)
	if err != nil {
		return e, errors.New(fmt.Sprintf("Failed to create an employee with the name of %s %s", e.FirstName, e.LastName))
	}
	employmentDate := employmentdate.EmploymentDate{
		EmployeeId: e.Id,
		StartDate: e.EmploymentDate.StartDate,
		EndDate: e.EmploymentDate.EndDate,
	}
	employmentDate, err = employmentDate.Create(db)
	if err != nil {
		return e, errors.New(fmt.Sprintf("Failed to create a employment date record for new employee %s", err))
	}
	tx.Commit()
	return e, nil
}

func (e Employee) Delete (db *sql.DB) (Employee, error) {
	if e.Id == 0 {
		return Employee{}, errors.New("Employee id must not be 0")
	}
	tx, err := db.Begin()
	if err != nil {
		return e, errors.New("Failed to begin a transaction to delete employee record")
	}
	defer tx.Rollback()
	s := 
	`
		DELETE FROM EMPLOYEE
		WHERE EMPLOYEE.id = $1
	`
	_, err = tx.Exec(s, e.Id)
	if err != nil {
		return e, errors.New(fmt.Sprintf("Failed to delete the employee with id of %d", e.Id))
	}
	employmentDate := employmentdate.EmploymentDate{
		EmployeeId: e.Id,
	}
	employmentDate, err = employmentDate.Delete(db)
	if err != nil {
		return e, errors.New(fmt.Sprintf("Failed to delete employment date record with id of %d", e.Id))
	}
	tx.Commit()
	return e, nil
}

func GetEmployeeById(db *sql.DB, id int) (Employee, error) {
	if id == 0 {
		return Employee{}, errors.New("Employee id = 0 is not a valid value to lookup")
	}
	s := 
	`
		select employee.id, employee.first_name, employee.last_name, department.name, employment_status.status, employee.avatar_url,
		employee.quote, employment_date.start_date, employment_date.end_date
		from employment_date
		join employee on employee.id = employment_date.employee_id
		join department on department.id = employee.department_id
		join employment_status on employment_status.id = employee.employment_status_id
		where employee.id = $1
	`
	var employee Employee
	var quote sql.NullString
	var avatarUrl sql.NullString
	var startDate sql.NullString
	var endDate sql.NullString
	err := db.QueryRow(s, id).Scan(
		&employee.Id,
		&employee.FirstName,
		&employee.LastName,
		&employee.DepartmentName,
		&employee.EmployeeStatus,
		&avatarUrl,
		&quote,
		&startDate,
		&endDate,
	)
	employee.Quote = database.GetNullableString(quote.String).String
	employee.AvatarUrl = database.GetNullableString(avatarUrl.String).String
	employee.EmploymentDate.StartDate = database.GetNullableString(startDate.String).String
	employee.EmploymentDate.EndDate = database.GetNullableString(endDate.String).String
	if err != nil {
		return Employee{}, err
	}
	return employee, nil
}

func (e Employee) Update (db *sql.DB) (Employee, error) {
	if e.Id == 0 {
		return e, errors.New("Employee id = 0 is not a valid value to update")
	}
	if e.DepartmentName == "" {
		return e, errors.New("Employee must belong to a department")
	}
	if e.FirstName == "" {
		return e, errors.New("Employee must have a first name set")
	}
	if e.LastName == "" {
		return e, errors.New("Employee must have a last name set")
	}
	avatarUrl := database.GetNullableString(e.AvatarUrl)
	quote := database.GetNullableString(e.Quote)
	tx, err := db.Begin()
	if err != nil {
		return e, errors.New("Failed to begin a transaction to update the employee record")
	}
	employeeDept := department.Department{
		Name: e.DepartmentName,
	}
	employeeDept, err = employeeDept.GetDepartmentByName(db)
	if err != nil {
		return e, errors.New(fmt.Sprintf("Failed to find a matching deparment with the name of %s", e.DepartmentName))
	}
	employeeStatus := employmentstatus.EmploymentStatus{
		Status: e.EmployeeStatus,
	}
	employeeStatus, err = employeeStatus.GetEmploymentStatusByStatus(db)
	if err != nil {
		return e, errors.New(fmt.Sprintf("Failed to find a employment status with the status of %s", e.EmployeeStatus))
	}
	defer tx.Rollback()
	s := 
	`
		UPDATE EMPLOYEE
		set first_name = $1,
		last_name = $2,
		department_id = $3,
		avatar_url = $4,
		quote = $5,
		employment_status_id = $6
		where id = $7
	`
	result, err := tx.Exec(s, e.FirstName, e.LastName, employeeDept.Id, avatarUrl, quote, employeeStatus.Id, e.Id,)
	if err != nil {
		return e, errors.New(fmt.Sprintf("Failed to update an employee with the name of %s %s", e.FirstName, e.LastName))
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return e, errors.New(fmt.Sprintf("No such employee exists with the name of %s %s", e.FirstName, e.LastName))
	}
	tx.Commit()
	return e, nil
}
