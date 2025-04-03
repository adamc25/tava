import "./department.css";

import { IDepartmentProps, IEmployee } from '../interfaces/';

import EmployeeListItem from '../employee/EmployeeListItem'

export default function Department({ name, employees }: IDepartmentProps) {
	return (
		<div
			className="card department-header-card"
		>
			<h2>
				{name}
			</h2>
			<div
				className="department-header flex"
			>
				<div
					className="department-header-field"
				>
					Name
				</div>
				<div
					className="department-header-field"
				>
					Start Date
				</div>
				<div
					className="department-header-field"
				>
					Quote
				</div>
				<div
					className="department-header-field"
				>
					Status
				</div>
			</div>
			{
				employees.map((employee: IEmployee, index: number) => (
					<EmployeeListItem key={index} employee={employee} />
				))
			}
		</div>
	);
}