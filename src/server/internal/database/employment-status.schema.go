package database

const EmploymentStatusSchema = `
-- Create table if it doesn't exist
CREATE TABLE IF NOT EXISTS employment_status (
		id SERIAL PRIMARY KEY,
		status TEXT NOT NULL
);

-- Create a unique index on the 'status' column if it doesn't exist
CREATE UNIQUE INDEX IF NOT EXISTS idx_status ON employment_status (status);
`