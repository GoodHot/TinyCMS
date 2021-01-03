package sqlite

import "github.com/GoodHot/TinyCMS/orm/trait"

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
	return " LIMIT ? OFFSET ?", []interface{}{query.Size, 0}
}

func (*SQLiteQuery) BuildWhere() {

}

func (query *SQLiteQuery) BuildLastID() (string, []interface{}) {
	return " AND id > ?", []interface{}{query.LastID}
}
