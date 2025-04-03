import "./employee-list.css";

import { useEffect, useState } from 'react';

import { IEmployee, IQueryCriteria, ISearchField } from '../interfaces/';

import UseFetchData from '../hooks/UseFetchData';

import Department from '../department/Department';

const compareFn = (previous, current) => {
	return JSON.stringify(previous) === JSON.stringify(current);
}

const fetchSearchableFields = async () => {
	const res = await fetch('http://localhost:8080/search/list', {
		method: 'GET',
		headers: {
			'Content-Type': 'application/json',
		}
	});
	if (!res.ok) {
		throw new Error(`HTTP error! Status: ${res.status}`);
	}
	return await res.json();
};

const fetchData = async (queryCriteria: IQueryCriteria) => {
	const res = await fetch('http://localhost:8080/employees/list', {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(queryCriteria),
	});
	if (!res.ok) {
	  const errorMessage = await res.text();
    throw new Error(`${errorMessage}`);
	}
	const data = await res.json();
	return await groupData(data);
};

const groupData = async (data: IEmployee[]) => {
	const departments: Map<string, IEmployee[]> = new Map();
	data.map((employee: IEmployee) => {
		if (departments.has(employee.departmentName)) {
			departments.get(employee.departmentName).push(employee);
		} else {
			departments.set(employee.departmentName, [employee]);
		}
	})
	return Array.from(departments);
}

const getInitialValuesFromLocalStorage = () => {
	const initialOrderBy = localStorage.getItem('orderby') || '1'
	const initialField = localStorage.getItem('field') || '1'
	const initialSearchTerm = localStorage.getItem('term') || '';
	const initialHierarchicalOrder = localStorage.getItem('hierarchicalorder') || 'ASC'
	return [initialField, initialSearchTerm, initialOrderBy, initialHierarchicalOrder]
}

export default function Employees() {
	const [cachedField, cachedSearchTerm, cachedOrderBy, cachedHierarchicalOrder] = getInitialValuesFromLocalStorage();
	const hierarchicalOrders = [
		{
			id: "Asc",
			value: "ASC"
		},
		{
			id: "Desc",
			value: "DESC"
		}
	]
	const [deps, setDeps] = useState({
		limit: 30,
		searchBy: {
			column: +cachedField,
			term: cachedSearchTerm
		},
		orderBy: {
			column: +cachedOrderBy,
			order: cachedHierarchicalOrder
		}
	});
	const { data, loading, error, setData } = UseFetchData(fetchData, compareFn, [deps], []);
	const { data: searchFields, error: searchFieldError, loading: searchFieldsLoading } = UseFetchData(fetchSearchableFields, compareFn, [], []);
	const [ searchTerm, setSearchTerm ] = useState(cachedSearchTerm);
	const [ field, setField ] = useState(+cachedField);
	const [ orderField, setOrderField ] = useState(+cachedOrderBy);
	const [ hierarchicalOrder, setHierarchicalOrder] = useState(cachedHierarchicalOrder);

	const findEmployees = () => {
		setDeps(prevState => ({
			...prevState,
			searchBy: {
				column: field,
				term: searchTerm
			},
			orderBy: {
				column: orderField,
				order: hierarchicalOrder
			}
		}));
	}

	const updateField = (event: React.ChangeEvent<HTMLSelectElement>) => {
		const { value } = event.target;
		setField(+value);
		localStorage.setItem('field', +value);
	}

	const updateHierarchicalOrder = (event: React.ChangeEvent<HTMLSelectElement>) => {
		const { value } = event.target;
		setHierarchicalOrder(value);
		localStorage.setItem('hierarchicalorder', value);
	}

	const updateOrderField = (event: React.ChangeEvent<HTMLSelectElement>) => {
		const { value } = event.target;
		setOrderField(+value);
		localStorage.setItem('orderby', +value);
	}

	const updateSearchTerm = (event: React.ChangeEvent<HTMLInputElement>) => {
		const { value } = event.target;
		setSearchTerm(value);
		localStorage.setItem('term', value);
	};

	return (
		<div
			className="employee-list-wrapper"
		>
			<div
				className="flex"
			>
				<label
					className="employee-list-top-margin"
				>
					Search Field
				</label>
				<select
					disabled={searchFieldError || searchFieldsLoading}
					className="employee-list-left-margin"
					value={field}
					onChange={updateField}
				>
					{
						searchFields.map((field: ISearchField) => (
							<option key={field.value} value={field.value}>
								{field.id}
							</option>
						))
					}
				</select>
				<label
					className="employee-list-left-margin employee-list-top-margin"
				>
					Order By
				</label>
				<select
						disabled={searchFieldError || searchFieldsLoading}
						className="employee-list-left-margin"
						value={orderField}
						onChange={updateOrderField}
					>
						{
							searchFields.map((field: ISearchField) => (
								<option key={field.value} value={field.value}>
									{field.id}
								</option>
							))
						}
				</select>
				<label
					className="employee-list-left-margin employee-list-top-margin"
				>
					Asc/Desc
				</label>
				<select
					className="employee-list-left-margin"
					value={hierarchicalOrder}
					onChange={updateHierarchicalOrder}
				>
					{
						hierarchicalOrders.map((order: ISearchField) => (
							<option key={order.value} value={order.value}>
								{order.id}
							</option>
						))
					}
				</select>
				<label
					className="employee-list-left-margin employee-list-top-margin"
				>
					Search By
				</label>
				<input
					className="employee-list-left-margin"
					value={searchTerm}
					onChange={updateSearchTerm}
				/>
				<button
					className="button apply-filters"
					onClick={findEmployees}
				>
					Apply Filters
				</button>
			</div>
				{
					error && 
					<div className="error-message employee-list-top-margin">
		        Failed to fetch list of employees. The following error was encountered: {error.message}
		      </div>
				}
				{
					!error && data.length > 0 && data.map(
						(department: [string, IEmployee[]], index: number) => (
							<Department key={index} name={department[0]} employees={department[1]} />
						)
					)
				}
				{
					!error && data.length === 0 && !loading && 
					<div className="card">
						Your query returned no search results.
					</div>
				}
		</div>
	)
}
