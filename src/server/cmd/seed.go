package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/lib/pq"
	"github.com/adamc25/db/internal/database"
	"github.com/adamc25/db/internal/department"
	"github.com/adamc25/db/internal/employee"
	"github.com/adamc25/db/internal/employment_date"
	"github.com/adamc25/db/internal/meta"
)


func createDepartmentTable(tx *sql.Tx) error {
		_, err := tx.Exec(database.DepartmentSchema)
		return err
}

func createEmploymentDate(tx *sql.Tx) error {
		_, err := tx.Exec(database.EmploymentDateSchema)
		return err
}

func createEmployeeTable(tx *sql.Tx) error {
		_, err := tx.Exec(database.EmployeeSchema)
		return err
}

func createEmploymentStatusTable(tx *sql.Tx) error {
		_, err := tx.Exec(database.EmploymentStatusSchema)
		return err
}

func createSearchTable(tx *sql.Tx) error {
		_, err := tx.Exec(database.SearchSchema)
		return err
}

func getDepartmentCount(tx *sql.Tx) (int, error) {
	var count int
	sql :=
	`
		select count(*) from department
	`
	err := tx.QueryRow(sql).Scan(&count)
	return count, err
}

func getEmployeeCount(tx *sql.Tx) (int, error) {
	var count int
	sql :=
	`
		select count(*) from employee
	`
	err := tx.QueryRow(sql).Scan(&count)
	return count, err
}

func getEmploymentDateCount(tx *sql.Tx) (int, error) {
	var count int
	sql :=
	`
		select count(*) from employment_date
	`
	err := tx.QueryRow(sql).Scan(&count)
	return count, err
}

func getEmploymentStatus(tx *sql.Tx) (int, error) {
	var count int
	sql := 
	`
		select count(*) from employment_status
	`
	err := tx.QueryRow(sql).Scan(&count)
	return count, err
}

func getSearchCount(tx *sql.Tx) (int, error) {
	var count int
	sql :=
	`
		select count(*) from search
	`
	err := tx.QueryRow(sql).Scan(&count)
	return count, err
}

func insertDepartments(tx *sql.Tx) error {
	depts := []department.Department{
		{Name: "Engineering"},
		{Name: "Management"},
		{Name: "Operations"},
		{Name: "Food Services"},
	}
	i := 0
	pos := make([]string, 0, len(depts))
	args := make([]any, 0, len(depts))
	for _, dept := range depts {
		pos = append(pos, fmt.Sprintf("($%d)", i+1))
		args = append(args, dept.Name)
		i++
	}
	_, err := tx.Exec(
		fmt.Sprintf("INSERT INTO DEPARTMENT (name) VALUES %s", strings.Join(pos, ",")),
		args...
	)
	return err
}

func insertEmployees(tx *sql.Tx) error {
	employees := []employee.Employee{
		{
			FirstName:          "Peter",
			LastName:           "Gibbons",
			DepartmentId:       1,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "The thing is, Bob, it's not that I'm lazy, it's that I just don't care.",
		},
		{
			FirstName:          "Samir",
			LastName:           "Nagheenanajar",
			DepartmentId:       1,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "No one in this country can ever pronounce my name right.",
		},
		{
			FirstName:          "Michael",
			LastName:           "Bolton",
			DepartmentId:       1,
			EmploymentStatusId: 2,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "That's the worst idea I've ever heard in my life, Tom.",
		},
		{
			FirstName:          "Bill",
			LastName:           "Lumbergh",
			DepartmentId:       2,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Ummm, I'm gonna need you to go ahead come in tomorrow.",
		},
		{
			FirstName:          "Bob",
			LastName:           "Slydell",
			DepartmentId:       2,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "What would you say...you do here?",
		},
		{
			FirstName:          "Bob",
			LastName:           "Porter",
			DepartmentId:       2,
			EmploymentStatusId: 2,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "We need to talk about your TPS reports.",
		},
		{
			FirstName:          "Dom",
			LastName:           "Portwood",
			DepartmentId:       2,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Looks like you've been missing a lot of work lately.",
		},
		{
			FirstName:          "Jane",
			LastName:           "Anderson",
			DepartmentId:       3,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Uh-oh. Sounds like somebody's got a case of the Mondays.",
		},
		{
			FirstName:          "Tom",
			LastName:           "Smykowski",
			DepartmentId:       3,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "I have people skills; I am good at dealing with people.",
		},
		{
			FirstName:          "Nina",
			LastName:           "Schultz",
			DepartmentId:       3,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Corporate accounts payable, Nina speaking. *JUST* a moment.",
		},
		{
			FirstName:          "Milton",
			LastName:           "Waddams",
			DepartmentId:       3,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Excuse me, I believe you have my stapler...",
		},
		{
			FirstName:          "Stan",
			LastName:           "Adams",
			DepartmentId:       4,
			EmploymentStatusId: 2,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "We need to talk about your flair.",
		},
		{
			FirstName:          "Joanna",
			LastName:           "Baker",
			DepartmentId:       4,
			EmploymentStatusId: 2,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "I do want to express myself, okay. And I don't need 37 pieces of flair to do it.",
		},
		{
			FirstName:          "Brian",
			LastName:           "Flores",
			DepartmentId:       4,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Some Pizza Shooters, Shrimp Poppers, or Extreme Fajitas?",
		},
		{
			FirstName:          "John",
			LastName:           "Doe",
			DepartmentId:       1,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "I'll just get it done myself.",
		},
		{
			FirstName:          "Alice",
			LastName:           "Smith",
			DepartmentId:       2,
			EmploymentStatusId: 2,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Maybe I’ll give it a try later.",
		},
		{
			FirstName:          "Jacob",
			LastName:           "Williams",
			DepartmentId:       3,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Just another day at the office.",
		},
		{
			FirstName:          "Emma",
			LastName:           "Johnson",
			DepartmentId:       4,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "I love working with the team.",
		},
		{
			FirstName:          "Mia",
			LastName:           "Davis",
			DepartmentId:       1,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Just another day to make progress.",
		},
		{
			FirstName:          "James",
			LastName:           "Miller",
			DepartmentId:       2,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "I'm here to lead and succeed.",
		},
		{
			FirstName:          "Olivia",
			LastName:           "Taylor",
			DepartmentId:       3,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Teamwork makes the dream work.",
		},
		{
			FirstName:          "Daniel",
			LastName:           "Brown",
			DepartmentId:       1,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Always moving forward.",
		},
		{
			FirstName:          "Sophia",
			LastName:           "Wilson",
			DepartmentId:       4,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Let's keep things running smoothly.",
		},
		{
			FirstName:          "Lucas",
			LastName:           "Moore",
			DepartmentId:       2,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Efficiency is key to success.",
		},
		{
			FirstName:          "Lily",
			LastName:           "Taylor",
			DepartmentId:       3,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "It’s all about the details.",
		},
		{
			FirstName:          "Henry",
			LastName:           "Clark",
			DepartmentId:       2,
			EmploymentStatusId: 2,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Time to take a break and recharge.",
		},
		{
			FirstName:          "Chloe",
			LastName:           "Rodriguez",
			DepartmentId:       1,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Let’s make it happen today.",
		},
		{
			FirstName:          "Ethan",
			LastName:           "Martinez",
			DepartmentId:       4,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Service with a smile!",
		},
		{
			FirstName:          "Mason",
			LastName:           "Hernandez",
			DepartmentId:       1,
			EmploymentStatusId: 2,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Taking it one step at a time.",
		},
		{
			FirstName:          "Avery",
			LastName:           "Young",
			DepartmentId:       3,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Onward and upward!",
		},
		{
			FirstName:          "Mila",
			LastName:           "King",
			DepartmentId:       4,
			EmploymentStatusId: 2,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "I’ve had a good run, but I’m moving on.",
		},
		{
			FirstName:          "Jack",
			LastName:           "Scott",
			DepartmentId:       1,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Challenges are opportunities to grow.",
		},
		{
			FirstName:          "Amelia",
			LastName:           "Green",
			DepartmentId:       2,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Leading by example.",
		},
		{
			FirstName:          "Alexander",
			LastName:           "Adams",
			DepartmentId:       3,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "No task is too big when we work together.",
		},
		{
			FirstName:          "Ella",
			LastName:           "Nelson",
			DepartmentId:       1,
			EmploymentStatusId: 1,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "Consistency is key to success.",
		},
		{
			FirstName:          "Sebastian",
			LastName:           "Carter",
			DepartmentId:       4,
			EmploymentStatusId: 2,
			AvatarUrl:          "https://thispersondoesnotexist.com/image",
			Quote:              "I’ve learned a lot here, but it’s time for a change.",
		},
	}
	i := 0
	pos := make([]string, 0, len(employees))
	args := make([]any, 0, len(employees))
	for _, e := range employees {
		pos = append(pos, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", i*6+1, i*6+2, i*6+3,i*6+4, i*6+5, i*6+6))
		args = append(args, e.FirstName, e.LastName, e.DepartmentId, e.EmploymentStatusId, e.AvatarUrl, e.Quote)
		i++
	}
	_, err := tx.Exec(
		fmt.Sprintf("INSERT INTO EMPLOYEE (first_name, last_name, department_id, employment_status_id, avatar_url, quote) VALUES %s", strings.Join(pos, ",")),
		args...
	)
	return err
}

func insertEmploymentDates(tx *sql.Tx) error {
	dates := []employmentdate.EmploymentDate{
		{
			StartDate:  "1997-03-13T00:00:00.000Z",
			EndDate:    "2001-05-22T00:00:00.000Z",
			EmployeeId: 1,
		},
		{
			StartDate:  "1998-06-25T00:00:00.000Z",
			EndDate:    "2002-09-15T00:00:00.000Z",
			EmployeeId: 2,
		},
		{
			StartDate:  "1998-10-01T00:00:00.000Z",
			EndDate:    "2002-02-10T00:00:00.000Z",
			EmployeeId: 3,
		},
		{
			StartDate:  "1996-08-15T00:00:00.000Z",
			EndDate:    "2001-11-19T00:00:00.000Z",
			EmployeeId: 4,
		},
		{
			StartDate:  "1997-06-13T00:00:00.000Z",
			EndDate:    "2002-04-30T00:00:00.000Z",
			EmployeeId: 5,
		},
		{
			StartDate:  "1997-07-30T00:00:00.000Z",
			EndDate:    "2001-07-01T00:00:00.000Z",
			EmployeeId: 6,
		},
		{
			StartDate:  "1996-05-05T00:00:00.000Z",
			EndDate:    "2000-06-12T00:00:00.000Z",
			EmployeeId: 7,
		},
		{
			StartDate:  "1999-01-22T00:00:00.000Z",
			EndDate:    "2003-03-18T00:00:00.000Z",
			EmployeeId: 8,
		},
		{
			StartDate:  "1997-04-13T00:00:00.000Z",
			EndDate:    "2001-08-10T00:00:00.000Z",
			EmployeeId: 9,
		},
		{
			StartDate:  "1999-02-16T00:00:00.000Z",
			EndDate:    "2003-07-05T00:00:00.000Z",
			EmployeeId: 10,
		},
		{
			StartDate:  "1995-04-18T00:00:00.000Z",
			EndDate:    "2000-02-05T00:00:00.000Z",
			EmployeeId: 11,
		},
		{
			StartDate:  "1997-11-01T00:00:00.000Z",
			EndDate:    "2001-10-14T00:00:00.000Z",
			EmployeeId: 12,
		},
		{
			StartDate:  "1998-09-18T00:00:00.000Z",
			EndDate:    "2002-01-11T00:00:00.000Z",
			EmployeeId: 13,
		},
		{
			StartDate:  "1998-12-25T00:00:00.000Z",
			EndDate:    "2003-04-03T00:00:00.000Z",
			EmployeeId: 14,
		},
		{
			StartDate:  "2001-03-12T00:00:00.000Z",
			EndDate:    "2004-07-10T00:00:00.000Z",
			EmployeeId: 15,
		},
		{
			StartDate:  "2003-06-22T00:00:00.000Z",
			EndDate:    "2007-01-16T00:00:00.000Z",
			EmployeeId: 16,
		},
		{
			StartDate:  "2002-04-19T00:00:00.000Z",
			EndDate:    "2006-09-04T00:00:00.000Z",
			EmployeeId: 17,
		},
		{
			StartDate:  "2005-08-01T00:00:00.000Z",
			EndDate:    "2009-03-30T00:00:00.000Z",
			EmployeeId: 18,
		},
		{
			StartDate:  "2008-10-14T00:00:00.000Z",
			EndDate:    "2011-12-20T00:00:00.000Z",
			EmployeeId: 19,
		},
		{
			StartDate:  "2006-07-29T00:00:00.000Z",
			EndDate:    "2009-11-05T00:00:00.000Z",
			EmployeeId: 20,
		},
		{
			StartDate:  "2004-11-12T00:00:00.000Z",
			EndDate:    "2008-02-15T00:00:00.000Z",
			EmployeeId: 21,
		},
		{
			StartDate:  "2009-05-23T00:00:00.000Z",
			EndDate:    "2013-07-12T00:00:00.000Z",
			EmployeeId: 22,
		},
		{
			StartDate:  "2010-06-11T00:00:00.000Z",
			EndDate:    "2014-02-24T00:00:00.000Z",
			EmployeeId: 23,
		},
		{
			StartDate:  "2011-09-07T00:00:00.000Z",
			EndDate:    "2015-04-15T00:00:00.000Z",
			EmployeeId: 24,
		},
		{
			StartDate:  "2012-08-19T00:00:00.000Z",
			EndDate:    "2016-11-02T00:00:00.000Z",
			EmployeeId: 25,
		},
		{
			StartDate:  "2015-01-14T00:00:00.000Z",
			EndDate:    "2018-06-09T00:00:00.000Z",
			EmployeeId: 26,
		},
		{
			StartDate:  "2016-04-30T00:00:00.000Z",
			EndDate:    "2020-02-21T00:00:00.000Z",
			EmployeeId: 27,
		},
		{
			StartDate:  "2014-02-07T00:00:00.000Z",
			EndDate:    "2017-05-29T00:00:00.000Z",
			EmployeeId: 28,
		},
		{
			StartDate:  "2007-10-23T00:00:00.000Z",
			EndDate:    "2011-01-04T00:00:00.000Z",
			EmployeeId: 29,
		},
		{
			StartDate:  "2013-09-28T00:00:00.000Z",
			EndDate:    "2016-12-14T00:00:00.000Z",
			EmployeeId: 30,
		},
		{
			StartDate:  "2017-07-13T00:00:00.000Z",
			EndDate:    "2020-03-06T00:00:00.000Z",
			EmployeeId: 31,
		},
		{
			StartDate:  "2019-11-05T00:00:00.000Z",
			EndDate:    "2022-08-19T00:00:00.000Z",
			EmployeeId: 32,
		},
		{
			StartDate:  "2020-05-21T00:00:00.000Z",
			EndDate:    "2023-01-30T00:00:00.000Z",
			EmployeeId: 33,
		},
		{
			StartDate:  "2021-02-16T00:00:00.000Z",
			EndDate:    "2022-09-11T00:00:00.000Z",
			EmployeeId: 34,
		},
		{
			StartDate:  "2021-11-09T00:00:00.000Z",
			EndDate:    "2023-06-15T00:00:00.000Z",
			EmployeeId: 35,
		},
		{
			StartDate:  "2022-01-02T00:00:00.000Z",
			EndDate:    "2023-02-25T00:00:00.000Z",
			EmployeeId: 36,
		},
	}
	i := 0
	pos := make([]string, 0, len(dates))
	args := make([]any, 0, len(dates))
	for _, date := range dates {
		pos = append(pos, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
		args = append(args, date.StartDate, date.EndDate, date.EmployeeId)
		i++
	}
	_, err := tx.Exec(
		fmt.Sprintf("INSERT INTO EMPLOYMENT_DATE (start_date, end_date, employee_id) VALUES %s", strings.Join(pos, ",")),
		args...
	)
	return err
}

func insertEmploymentStatus(tx *sql.Tx) error {
	statuses := []string{
		"active",
		"inactive",
	}
	i := 0
	pos := make([]string, 0, len(statuses))
	args := make([]any, 0, len(statuses))
	for _, status := range statuses {
		pos = append(pos, fmt.Sprintf("($%d)", i+1))
		args = append(args, status)
		i++
	}
	_, err := tx.Exec(
		fmt.Sprintf("INSERT INTO EMPLOYMENT_STATUS (status) VALUES %s", strings.Join(pos, ",")),
		args...
	)
	return err
}

func insertSearch(tx *sql.Tx) error {
	searches := []meta.Search{
		{
			HumanReadable: "First Name",
			TableName: "employee",
			ColumnName: "first_name",
		},
		{
			HumanReadable: "Last Name",
			TableName: "employee",
			ColumnName: "last_name",
		},
		{
			HumanReadable: "Employment Status",
			TableName: "employment_status",
			ColumnName: "status",
		},
		{
			HumanReadable: "Department Name",
			TableName: "department",
			ColumnName: "name",
		},
	}
	i := 0
	pos := make([]string, 0, len(searches))
	args := make([]any, 0, len(searches))
	for _, search := range searches {
		pos = append(pos, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
		args = append(args, search.HumanReadable, search.TableName, search.ColumnName)
		i++
	}
	_, err := tx.Exec(
		fmt.Sprintf("INSERT INTO SEARCH (human_readable, table_name, column_name) VALUES %s", strings.Join(pos, ",")),
		args...
	)
	return err
}

func Seed () {
	connStr := database.GetConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to open DB: %v", err)
		return
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback()
	err = createSearchTable(tx)
	if err != nil {
		log.Fatal("Failed to create search table:", err)
	}
	searchCnt, err := getSearchCount(tx)
	if err != nil {
		log.Fatal("Failed to get count of search records:", err)
	}
	if searchCnt == 0 {
		err = insertSearch(tx)
	}
	err = createDepartmentTable(tx)
	if err != nil {
		log.Fatal("Failed to create department table:", err)
	}
	err = createEmploymentStatusTable(tx)
	if err != nil {
		log.Fatal("Failed to create employment status table:", err)
	}
	err = createEmployeeTable(tx)
	if err != nil {
		log.Fatal("Failed to create employee table:", err)
	}
	empCnt, err := getEmployeeCount(tx)
	if err != nil {
		log.Fatal("Failed to get count of employee records:", err)
	}
	if empCnt == 0 {
		err = insertEmployees(tx)
	}
	if err != nil && empCnt == 0 {
		log.Fatal("Failed to insert employee records:", err)
	}
	deptCnt, err := getDepartmentCount(tx)
	if err != nil {
		log.Fatal("Failed to get count of department records:", err)
	}
	if deptCnt == 0 {
		err = insertDepartments(tx)
	}
	if err != nil && deptCnt == 0 {
		log.Fatal("Failed to insert department records:", err)
	}
	employeeStatCnt, err := getEmploymentStatus(tx)
	if err != nil {
		log.Fatal("Failed to get count of employee status records:", err)
	}
	if employeeStatCnt == 0 {
		err = insertEmploymentStatus(tx)
	}
	if err != nil && employeeStatCnt == 0 {
		log.Fatal("Failed to insert employment status records", err)
	}
	err = createEmploymentDate(tx)
	if err != nil {
		log.Fatal("Failed to create employment date table:", err)
	}
	employmentDateCnt, err := getEmploymentDateCount(tx)
	if err != nil {
		log.Fatal("Failed to get count of employment date records:", err)
	}
	if employmentDateCnt == 0 {
		err = insertEmploymentDates(tx)
	}
	if err != nil && employmentDateCnt == 0 {
		log.Fatal("Failed to insert employment date records:", err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}
}


func main() {
	Seed()
}