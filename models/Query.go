package models

type Query struct {
	QueryString string
	SortBy      string
	Order       string
}

func (this Query) GetQuery(qs string) (q *Query) {
	qu := Query{}
	q.QueryString = qs
	return &qu
}
