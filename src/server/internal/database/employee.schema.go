package database

const EmployeeSchema = 
`
-- Create table if it doesn't exist
CREATE TABLE IF NOT EXISTS employee (
		id SERIAL PRIMARY KEY,
		first_name TEXT NOT NULL,
		last_name TEXT NOT NULL,
		department_id INT NOT NULL,
		employment_status_id INT NOT NULL,
		avatar_url TEXT,
		quote TEXT
);

-- Create an index for the first_name column
CREATE INDEX IF NOT EXISTS idx_employee_first_name ON employee (first_name);

-- Create an index for the last_name column
CREATE INDEX IF NOT EXISTS idx_employee_last_name ON employee (last_name);

-- Create an index for the department_id column
CREATE INDEX IF NOT EXISTS idx_employee_department_id ON employee (department_id);

-- Create an index for the employment_status_id column
CREATE INDEX IF NOT EXISTS idx_employment_status_id ON employee (employment_status_id);
`