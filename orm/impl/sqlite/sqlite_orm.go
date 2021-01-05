package sqlite

import (
	"github.com/GoodHot/TinyCMS/orm/trait"
	"strings"
)

func ToSQLiteQuery(query *trait.Query) *SQLiteQuery {
	return &SQLiteQuery{
		Query: query,
	}
}

type SQLiteQuery struct {
	*trait.Query
}

func (query *SQLiteQuery) BuildLimit() (string, []interface{}) {
	//row := query.Page * query.Size
	return " LIMIT ? OFFSET ?", []interface{}{query.Size, (query.Page - 1) * query.Size}
}

func (*SQLiteQuery) BuildWhere() {

}

func (query *SQLiteQuery) BuildOrder() string {
	if query.Order == nil {
		return ""
	}
	var order []string
	for key, val := range query.Order {
		order = append(order, key+" "+val)
	}
	return "ORDER BY " + strings.Join(order, ",")
}

func (query *SQLiteQuery) BuildLastID() (string, []interface{}) {
	//return " AND id > ?", []interface{}{query.LastID}
	return "", nil
}
