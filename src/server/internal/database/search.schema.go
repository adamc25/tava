package database

const SearchSchema = 
`
-- Create table if it doesn't exist
CREATE TABLE IF NOT EXISTS search (
		id SERIAL PRIMARY KEY,
		human_readable TEXT NOT NULL,
		table_name TEXT NOT NULL,
		column_name TEXT NOT NULL
);
`
