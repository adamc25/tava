package database

const DepartmentSchema =
`
	-- Create table if it doesn't exist
	CREATE TABLE IF NOT EXISTS department (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL
	);

	-- Create a unique index on the 'name' column if it doesn't exist
	CREATE UNIQUE INDEX IF NOT EXISTS idx_department ON department (name);
`
