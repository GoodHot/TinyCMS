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

func (query *SQLiteQuery) BuildWhere() (where string, param []interface{}) {
	if query.Param == nil {
		return
	}
	for field, value := range query.Param {
		where += " and " + field + " = ?"
		param = append(param, value)
	}
	return
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
