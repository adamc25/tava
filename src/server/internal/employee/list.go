package employee

import (
	"database/sql"
	"fmt"

	"github.com/adamc25/db/internal/database"
	"github.com/adamc25/db/internal/meta"
)

func getDefaultDirection(direction string) string {
    if direction == "ASC" || direction == "DESC" {
    	return direction
    }
    return "ASC"
}

func getWhereClause(q meta.QueryMeta, column string, tableName string) (string, []interface{}) {
	var args []interface{}

	if q.SearchBy.Exact == true {
		return fmt.Sprintf("where %s.%s ILIKE $1", tableName, column), append(args, q.SearchBy.Term)
	}

	return fmt.Sprintf("where %s.%s ILIKE $1", tableName, column), append(args, fmt.Sprintf("%%%s%%", q.SearchBy.Term))
}

func getOrderByClause(q meta.QueryMeta, column string, tableName string) string {

	// We need to ensure any custom ordering is done within the context of each department. In case the user selects that column to order
	// by, we just need to consume their selection for ascending / descending
	if tableName == "department" && column == "name" {
		return fmt.Sprintf("order by department.name %s", getDefaultDirection(q.OrderBy.Order))
	}

	return fmt.Sprintf("order by department.name asc, %s.%s %s", tableName, column, getDefaultDirection(q.OrderBy.Order))
}

func GetList(db *sql.DB, q meta.QueryMeta) ([]Employee, error) {
	searchByColumnnId := meta.GetDefaultColumnId(q.SearchBy)
	search, err:= meta.GetSearchableFieldById(db, searchByColumnnId)
	if err != nil {
		return nil, err
	}
	orderByColumnnId := meta.GetDefaultOrderId(q.OrderBy)
	order, err:= meta.GetSearchableFieldById(db, orderByColumnnId)
	if err != nil {
		return nil, err
	}
	if q.SearchBy.Term == "" {
		return getUnfilteredList(db, q, order.ColumnName, order.TableName)
	}
	whereClause, args := getWhereClause(q, search.ColumnName, search.TableName)
	s := fmt.Sprintf(
	`
		select employee.id, employee.first_name, employee.last_name, department.name as department_name, employment_status.status, employee.avatar_url,
		employee.quote, employment_date.start_date, employment_date.end_date
		from employment_date
		join employee on employee.id = employment_date.employee_id
		join department on department.id = employee.department_id
		join employment_status on employment_status.id = employee.employment_status_id
		%s
		%s
		limit %d offset %d
	`, whereClause, getOrderByClause(q, order.ColumnName, order.TableName) , meta.GetDefaultLimit(q), q.Offset)
	rows, err := db.Query(s, args...)
	if err != nil {
		return nil, err
	}

	var employees []Employee

	for rows.Next() {
		var employee Employee
		var quote sql.NullString
		var avatarUrl sql.NullString
		var startDate sql.NullString
		var endDate sql.NullString
		err := rows.Scan(
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
			return nil, err
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if employees == nil {
		return []Employee{}, nil
	}

	return employees, nil
}

func getUnfilteredList(db *sql.DB, q meta.QueryMeta, column string, tableName string) ([]Employee, error) {
	s := fmt.Sprintf(
	`
		select employee.id, employee.first_name, employee.last_name, department.name, employment_status.status, employee.avatar_url,
		employee.quote, employment_date.start_date, employment_date.end_date
		from employment_date
		join employee on employee.id = employment_date.employee_id
		join department on department.id = employee.department_id
		join employment_status on employment_status.id = employee.employment_status_id
		%s
		limit %d offset %d
	`, getOrderByClause(q, column, tableName), meta.GetDefaultLimit(q), q.Offset)
	rows, err := db.Query(s)
	if err != nil {
		return nil, err
	}

	var employees []Employee

	for rows.Next() {
		var employee Employee
		var quote sql.NullString
		var avatarUrl sql.NullString
		var startDate sql.NullString
		var endDate sql.NullString
		err := rows.Scan(
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
			return nil, err
		}
		employees = append(employees, employee)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}
