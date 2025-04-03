package meta

type OrderBy struct {
	Column int
	Order  string
}

type QueryMeta struct {
	Limit   int
	Offset  int
	OrderBy OrderBy
	SearchBy SearchBy
}

type SearchBy struct {
	Column int
	Term string
	Exact bool
}

func GetDefaultOrderId(s OrderBy) int {
	if s.Column == 0 {
		return 1
	}
	return s.Column
}

func GetDefaultColumnId(s SearchBy) int {
	if s.Column == 0 {
		return 1
	}
	return s.Column
}


func GetDefaultLimit(q QueryMeta) int {
	if q.Limit == 0 {
		return 10
	}
	return q.Limit
}
