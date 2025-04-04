package database

const EmploymentDateSchema =
`
-- Create table if it doesn't exist
CREATE TABLE IF NOT EXISTS employment_date (
		start_date TIMESTAMPTZ,
		end_date TIMESTAMPTZ,
		employee_id INT NOT NULL
);

-- Create a unique index on the 'start_date' column if it doesn't exist
CREATE INDEX IF NOT EXISTS idx_start_date ON employment_date (start_date);

-- Create a unique index on the 'end_date' column if it doesn't exist
CREATE INDEX IF NOT EXISTS idx_end_date ON employment_date (end_date);

-- Create a unique index on the 'employee_id' column if it doesn't exist
CREATE UNIQUE INDEX IF NOT EXISTS idx_employee_id ON employment_date (employee_id);
`
