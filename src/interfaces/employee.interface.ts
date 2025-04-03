export interface IEmployee {
	avatarUrl: string;
	// this field is not sent by the backend
	departmentId: number;
	departmentName: string;
	employeeStatus: string;
	// this field is not sent by the backend
	employeeStatusId: number;
	id: number;
	firstName: string;
	lastName: string;
	quote: string;
}
