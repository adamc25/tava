package meta

import (
	"database/sql"
	"errors"
)


type Search struct {
	Id int `json:"value"`
	HumanReadable string `json:"id"`
	TableName     string `json:"-"`
	ColumnName    string `json:"-"`
}


func GetAllSearchableFields(db *sql.DB) ([]Search, error) {
	s :=
	`
		select search.id, search.human_readable, search.table_name, search.column_name from search
	`
	rows, err := db.Query(s)
	if err != nil {
		return nil, err
	}

	var searches []Search

	for rows.Next() {
		var search Search
		err := rows.Scan(
			&search.Id,
			&search.HumanReadable,
			&search.TableName,
			&search.ColumnName,
		)
		if err != nil {
			return nil, err
		}
		searches = append(searches, search)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return searches, nil
}

func GetSearchableFieldById(db *sql.DB, id int) (Search, error) {
	if id == 0 {
		return Search{}, errors.New("Searchable Field Id = 0 is not a valid value to lookup")
	}
	s :=
	`
		select search.id, search.human_readable, search.table_name, search.column_name from search
		where search.id = $1
	`
	var search Search
	err := db.QueryRow(s, id).Scan(
		&search.Id,
		&search.HumanReadable,
		&search.TableName,
		&search.ColumnName,
	)
	if err != nil {
		return Search{}, err
	}
	return search, nil
}