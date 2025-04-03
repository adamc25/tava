import './edit-employee.css';

import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';

import { IDepartment } from '../interfaces/';

import UseFetchData from '../hooks/UseFetchData'

const fetchDepartments = async () => {
	const res = await fetch('http://localhost:8080/department/list', {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
		},
	});
	if (!res.ok) {
		const errorMessage = await res.text();
		throw new Error(`${errorMessage}`);
	}
	return await res.json();
};

const fetchEmploymentStatus = async () => {
	const res = await fetch('http://localhost:8080/employment_status/list', {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
		},
	});
	if (!res.ok) {
		const errorMessage = await res.text();
		throw new Error(`${errorMessage}`);
	}
	return await res.json();
};

const fetchEmployee = async (id: number) => {
	const queryParams = new URLSearchParams({
		id
	});
	const res = await fetch(`http://localhost:8080/employee?${queryParams.toString()}`, {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
		},
	});
	if (!res.ok) {
		const errorMessage = await res.text();
		throw new Error(`${errorMessage}`);
	}
	return await res.json();
};

const compareFn = (previous, current) => {
	return previous !== undefined;
}

export default function EditEmployee() {
	const { id } = useParams();
	const navigate = useNavigate();
	const { data: departments, loading: departmentLoading, error: departmentLoadError } = UseFetchData(fetchDepartments, compareFn, [], []);
	const { data: employee, loading: employeeLoading, error: employeeLoadError, setData: setEmployeeData } = UseFetchData(fetchEmployee, compareFn, [id], {});
	const { data: employmentStatus, loading: employmentStatusLoading, error: employmentStatusLoadError } = UseFetchData(fetchEmploymentStatus, compareFn, [], []);
	const [selectedValue, setSelectedValue] = useState('');
	const [saveError, setSaveError] = useState('');

	// Set initial value for the department select option
	useEffect(() => {
		if (typeof employee.id !== 'number' || departments.length === 0) {
			return;
		}
		if (employeeLoadError || departmentLoadError) {
			return;
		}
		const department = departments.find((department: IDepartment) => department.name === employee.departmentName);
		setEmployeeData(prevState => ({
			...prevState,
			departmentId: department.id
		}));
	}, [employeeLoading, departmentLoading]);

	// Set initial value for the employment status select option
	useEffect(() => {
		if (typeof employee.id !== 'number' || employmentStatus.length === 0) {
			return;
		}
		if (employeeLoadError || employmentStatusLoadError) {
			return;
		}
		const status = employmentStatus.find((status: IEmploymentStatus) => status.status === employee.employeeStatus);
		setEmployeeData(prevState => ({
			...prevState,
			employeeStatusId: status.id
		}));
	}, [employeeLoading, employmentStatusLoading]);


	const handleChange = (e) => {
		setSelectedValue(e.target.value);
	};

	const handleInputChange = (event: React.ChangeEvent<HTMLInputElement>) => {
		const { name, value } = event.target;
		setEmployeeData(prevState => ({
			...prevState,
			[name]: value
		}));
	};

	const setDepartmentName = (event: React.ChangeEvent<HTMLSelectElement>) => {
		const { value } = event.target;
		const department = departments.find((department: IDepartment) => department.id === +value);
		setEmployeeData(prevState => ({
			...prevState,
			departmentId: +value,
			departmentName: department.name
		}));
	}

	const setEmploymentStatusName = (event: React.ChangeEvent<HTMLSelectElement>) => {
		const { value } = event.target;
		const status = employmentStatus.find((status: IEmploymentStatus) => status.status === value);
		setEmployeeData(prevState => ({
			...prevState,
			employeeStatusId: status.id,
			employeeStatus: status.status
		}));
	}

	const saveEmployee = async () => {
		setSaveError('');
		try {
			const res = await fetch(`http://localhost:8080/employee`, {
				method: 'PUT',
				headers: {
					'Content-Type': 'application/json',
				},
				body: JSON.stringify(employee)
			});
			if (!res.ok) {
				const errorMessage = await res.text();
				throw new Error(`${errorMessage}`);
			}
			navigate('/employees')
		} catch (error) {
			setSaveError(error.message);
		}
	};

	if (employeeLoadError) {
		 return (
			<div className="edit-employee card">
				<div>
					<div className="error-message">
						Failed to load employee with id of {id}. The following error was encountered: {employeeLoadError.message}
					</div>
				</div>
			</div>
		)
	}

	if (departmentLoadError) {
		 return (
			<div className="edit-employee card">
				<div>
					<div className="error-message">
						Failed to fetch list of departments. The following error was encountered: {departmentLoadError.message}
					</div>
				</div>
			</div>
		)
	}

	if (employmentStatusLoadError) {
		 return (
			<div className="edit-employee card">
				<div>
					<div className="error-message">
						Failed to load the list of employment status. The following error was encountered: {employmentStatusLoadError.message}
					</div>
				</div>
			</div>
		)
	}

	if (!employeeLoading && !departmentLoading && !employmentStatusLoading) {
		return (
			<div
				className="edit-employee card">
					<h1>
						Edit Employee
					</h1>
					<div>
						{saveError && (
							<div className="error-message">
								Failed to save the employee for the following reason: {saveError}
							</div>
						)}
					</div>
					<label>
						First Name
					</label>
					<input 
						name="firstName"
						value={employee.firstName}
						onChange={handleInputChange}
					/>
					<label>
						Last Name
					</label>
					<input 
						name="lastName"
						value={employee.lastName}
						onChange={handleInputChange}
					/>
					<label>
						Start Date
					</label>
					<input
						disabled
						value={employee.employmentDate.startDate}
					/>
					<label>
						End Date
					</label>
					<input
						disabled
						value={employee.employmentDate.endDate}
					/>
					<label>
						Department
					</label>
					<select
						value={employee.departmentId}
						onChange={setDepartmentName}
					>
						{departments.map((department: IDepartment) => (
							<option key={department.id} value={department.id}>
								{department.name}
							</option>
						))}
					</select>
					<label>
						Employment Status
					</label>
					<select
						value={employee.employeeStatus}
						onChange={setEmploymentStatusName}
					>
						{employmentStatus.map((status: IEmploymentStatus) => (
							<option key={status.id} value={status.status}>
								{status.status}
							</option>
						))}
					</select>
					<label>
						Avatar URL
					</label>
					<input 
						name="avatarUrl"
						value={employee.avatarUrl}
						onChange={handleInputChange}
					/>
					<label>
						Quote
					</label>
					<textarea
						name="quote"
						value={employee.quote}
						onChange={handleInputChange}
						style={{
							"minWidth": "100%"
						}}>
					</textarea>
					<div>
						<button
							className="button save-button"
							onClick={saveEmployee}
						>
							Save
						</button>
					</div>
			</div>
		)
	}
	return (
		<div
			className="edit-employee card">
		</div>
	)
}
