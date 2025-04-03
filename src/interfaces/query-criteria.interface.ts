import { IOrderBy } from './order-by.interface';
import { ISearchBy } from './search-by.interface';

export interface IQueryCriteria {
	limit: number;
	offset: number;
	orderBy: IOrderBy;
	searchBy: ISearchBy;
}
