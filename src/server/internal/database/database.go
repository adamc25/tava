package database

import (
	"database/sql"
)

func GetConnectionString() string {
	return "postgres://cragg@localhost:5432/tava-health?sslmode=disable"
}

func GetNullableString(value string) sql.NullString {
	if value == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: value, Valid: true}
}