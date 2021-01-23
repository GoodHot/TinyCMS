
package datasource

import (
    "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DBCodeInjectORMExample struct {
	showSQL bool
	db      *sqlx.DB
	sql     string
	where   string
	orderBy string
	limit   string
	param   []interface{}
}

func (exp *DBCodeInjectORMExample) And() *DBCodeInjectORMExample {
	exp.where += " and "
	return exp
}
func (exp *DBCodeInjectORMExample) Or() *DBCodeInjectORMExample {
	exp.where += " or "
	return exp
}
func (exp *DBCodeInjectORMExample) Exec(handler ...func(result sql.Result)) error {
	
    sql := exp.sql
    if exp.where != "" {
        sql += " where " + exp.where
    }
    if exp.orderBy != "" {
        sql += " order by " + exp.orderBy
    }
    if exp.limit != "" {
        sql += " limit ? offset ?"
    }
    if exp.showSQL {
        fmt.Println("sql: " + sql)
        fmt.Println("param: ", exp.param)
    }

	result, err := exp.db.Exec(sql, exp.param...)
	if err != nil {
	    return err
	}
    if len(handler) > 0 {
        for _, h := range handler {
            h(result)
        }
    }
	return nil
}
func (exp *DBCodeInjectORMExample) ExecQueryCount() (int, error) {
	sql := "SELECT COUNT(1) FROM t_code_inject "
	if exp.where != "" {
        sql += " where " + exp.where
    }
    if exp.showSQL {
        fmt.Println("sql: " + sql)
        fmt.Println("param: ", exp.param)
    }
    rows, err := exp.db.Query(sql, exp.param...)
    if err != nil {
        return 0, err
    }
    defer rows.Close()
    if rows.Next() {
	    var count int
	    if err := rows.Scan(&count); err != nil {
            return 0, err
        }
        return count, nil
	}
	return 0, nil
}
func (exp *DBCodeInjectORMExample) ExecQuery() ([]*DBCodeInjectModel, error) {
	
    sql := exp.sql
    if exp.where != "" {
        sql += " where " + exp.where
    }
    if exp.orderBy != "" {
        sql += " order by " + exp.orderBy
    }
    if exp.limit != "" {
        sql += " limit ? offset ?"
    }
    if exp.showSQL {
        fmt.Println("sql: " + sql)
        fmt.Println("param: ", exp.param)
    }

	rows, err := exp.db.Queryx(sql, exp.param...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*DBCodeInjectModel
	for rows.Next() {
		var model DBCodeInjectModel
        if err := rows.StructScan(&model); err != nil {
            return result, err
        }
        result = append(result, &model)
	}
	return result, nil
}
func (exp *DBCodeInjectORMExample) ExecQueryOne() (*DBCodeInjectModel, error) {
	
    sql := exp.sql
    if exp.where != "" {
        sql += " where " + exp.where
    }
    if exp.orderBy != "" {
        sql += " order by " + exp.orderBy
    }
    if exp.limit != "" {
        sql += " limit ? offset ?"
    }
    if exp.showSQL {
        fmt.Println("sql: " + sql)
        fmt.Println("param: ", exp.param)
    }

	rows, err := exp.db.Queryx(sql, exp.param...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
	    var model DBCodeInjectModel
	    if err := rows.StructScan(&model); err != nil {
            return nil, err
        }
        return &model, nil
	}
	return nil, nil
}
func (exp *DBCodeInjectORMExample) ExecPage() error {
	sql := exp.sql
	if exp.where != "" {
		sql += " where " + exp.where
	}
	fmt.Println(sql)
	return nil
}
func (exp *DBCodeInjectORMExample) AddParam(param interface{}) *DBCodeInjectORMExample {
	exp.param = append(exp.param, param)
	return exp
}

func (exp *DBCodeInjectORMExample) Limit(limit int, offset int) *DBCodeInjectORMExample {
	exp.limit = " limit ? offset ? "
	exp.param = append(exp.param, limit)
	exp.param = append(exp.param, offset)
	return exp
}




func (exp *DBCodeInjectORMExample) IDEq(ID int) *DBCodeInjectORMExample {
	exp.where += "id = ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBCodeInjectORMExample) IDNotEq(ID int) *DBCodeInjectORMExample {
	exp.where += "id = ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBCodeInjectORMExample) IDGt(ID int) *DBCodeInjectORMExample {
	exp.where += "id > ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBCodeInjectORMExample) IDLt(ID int) *DBCodeInjectORMExample {
	exp.where += "id < ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBCodeInjectORMExample) IDEqGt(ID int) *DBCodeInjectORMExample {
	exp.where += "id >= ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBCodeInjectORMExample) IDEqLt(ID int) *DBCodeInjectORMExample {
	exp.where += "id <= ?"
	exp.param = append(exp.param, ID)
	return exp
}








func (exp *DBCodeInjectORMExample) KeyEq(Key string) *DBCodeInjectORMExample {
	exp.where += "key = ?"
	exp.param = append(exp.param, Key)
	return exp
}

func (exp *DBCodeInjectORMExample) KeyNotEq(Key string) *DBCodeInjectORMExample {
	exp.where += "key != ?"
	exp.param = append(exp.param, Key)
	return exp
}

func (exp *DBCodeInjectORMExample) KeyLike(Key string) *DBCodeInjectORMExample {
	exp.where += "key like '%' || ? || '%'"
	exp.param = append(exp.param, Key)
	return exp
}

func (exp *DBCodeInjectORMExample) KeyLikeAfter(Key string) *DBCodeInjectORMExample {
	exp.where += "key like ? || '%'"
	exp.param = append(exp.param, Key)
	return exp
}

func (exp *DBCodeInjectORMExample) KeyLikeBefore(Key string) *DBCodeInjectORMExample {
	exp.where += "key like  '%' || ?"
	exp.param = append(exp.param, Key)
	return exp
}






func (exp *DBCodeInjectORMExample) DescriptionEq(Description string) *DBCodeInjectORMExample {
	exp.where += "description = ?"
	exp.param = append(exp.param, Description)
	return exp
}

func (exp *DBCodeInjectORMExample) DescriptionNotEq(Description string) *DBCodeInjectORMExample {
	exp.where += "description != ?"
	exp.param = append(exp.param, Description)
	return exp
}

func (exp *DBCodeInjectORMExample) DescriptionLike(Description string) *DBCodeInjectORMExample {
	exp.where += "description like '%' || ? || '%'"
	exp.param = append(exp.param, Description)
	return exp
}

func (exp *DBCodeInjectORMExample) DescriptionLikeAfter(Description string) *DBCodeInjectORMExample {
	exp.where += "description like ? || '%'"
	exp.param = append(exp.param, Description)
	return exp
}

func (exp *DBCodeInjectORMExample) DescriptionLikeBefore(Description string) *DBCodeInjectORMExample {
	exp.where += "description like  '%' || ?"
	exp.param = append(exp.param, Description)
	return exp
}






func (exp *DBCodeInjectORMExample) LanguageEq(Language string) *DBCodeInjectORMExample {
	exp.where += "language = ?"
	exp.param = append(exp.param, Language)
	return exp
}

func (exp *DBCodeInjectORMExample) LanguageNotEq(Language string) *DBCodeInjectORMExample {
	exp.where += "language != ?"
	exp.param = append(exp.param, Language)
	return exp
}

func (exp *DBCodeInjectORMExample) LanguageLike(Language string) *DBCodeInjectORMExample {
	exp.where += "language like '%' || ? || '%'"
	exp.param = append(exp.param, Language)
	return exp
}

func (exp *DBCodeInjectORMExample) LanguageLikeAfter(Language string) *DBCodeInjectORMExample {
	exp.where += "language like ? || '%'"
	exp.param = append(exp.param, Language)
	return exp
}

func (exp *DBCodeInjectORMExample) LanguageLikeBefore(Language string) *DBCodeInjectORMExample {
	exp.where += "language like  '%' || ?"
	exp.param = append(exp.param, Language)
	return exp
}






func (exp *DBCodeInjectORMExample) CodeEq(Code string) *DBCodeInjectORMExample {
	exp.where += "code = ?"
	exp.param = append(exp.param, Code)
	return exp
}

func (exp *DBCodeInjectORMExample) CodeNotEq(Code string) *DBCodeInjectORMExample {
	exp.where += "code != ?"
	exp.param = append(exp.param, Code)
	return exp
}

func (exp *DBCodeInjectORMExample) CodeLike(Code string) *DBCodeInjectORMExample {
	exp.where += "code like '%' || ? || '%'"
	exp.param = append(exp.param, Code)
	return exp
}

func (exp *DBCodeInjectORMExample) CodeLikeAfter(Code string) *DBCodeInjectORMExample {
	exp.where += "code like ? || '%'"
	exp.param = append(exp.param, Code)
	return exp
}

func (exp *DBCodeInjectORMExample) CodeLikeBefore(Code string) *DBCodeInjectORMExample {
	exp.where += "code like  '%' || ?"
	exp.param = append(exp.param, Code)
	return exp
}




func (exp *DBCodeInjectORMExample) OrderByIDDesc() *DBCodeInjectORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " id desc"
    return exp
}
func (exp *DBCodeInjectORMExample) OrderByIDAsc() *DBCodeInjectORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " id asc"
    return exp
}

func (exp *DBCodeInjectORMExample) OrderByKeyDesc() *DBCodeInjectORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " key desc"
    return exp
}
func (exp *DBCodeInjectORMExample) OrderByKeyAsc() *DBCodeInjectORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " key asc"
    return exp
}

func (exp *DBCodeInjectORMExample) OrderByDescriptionDesc() *DBCodeInjectORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " description desc"
    return exp
}
func (exp *DBCodeInjectORMExample) OrderByDescriptionAsc() *DBCodeInjectORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " description asc"
    return exp
}

func (exp *DBCodeInjectORMExample) OrderByLanguageDesc() *DBCodeInjectORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " language desc"
    return exp
}
func (exp *DBCodeInjectORMExample) OrderByLanguageAsc() *DBCodeInjectORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " language asc"
    return exp
}

func (exp *DBCodeInjectORMExample) OrderByCodeDesc() *DBCodeInjectORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " code desc"
    return exp
}
func (exp *DBCodeInjectORMExample) OrderByCodeAsc() *DBCodeInjectORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " code asc"
    return exp
}


type DBCodeInjectUpdater struct {
    setSQL  string
    sql     string
    param   []interface{}
    showSQL bool
    db      *sqlx.DB
}


func (updater *DBCodeInjectUpdater) SetID(ID int) *DBCodeInjectUpdater {
	updater.setSQL += ", id = ?"
	updater.param = append(updater.param, ID)
	return updater
}

func (updater *DBCodeInjectUpdater) SetKey(Key string) *DBCodeInjectUpdater {
	updater.setSQL += ", key = ?"
	updater.param = append(updater.param, Key)
	return updater
}

func (updater *DBCodeInjectUpdater) SetDescription(Description string) *DBCodeInjectUpdater {
	updater.setSQL += ", description = ?"
	updater.param = append(updater.param, Description)
	return updater
}

func (updater *DBCodeInjectUpdater) SetLanguage(Language string) *DBCodeInjectUpdater {
	updater.setSQL += ", language = ?"
	updater.param = append(updater.param, Language)
	return updater
}

func (updater *DBCodeInjectUpdater) SetCode(Code string) *DBCodeInjectUpdater {
	updater.setSQL += ", code = ?"
	updater.param = append(updater.param, Code)
	return updater
}

func (updater *DBCodeInjectUpdater) Done() *DBCodeInjectORMExample {
    example := &DBCodeInjectORMExample{sql: updater.sql + updater.setSQL[1:], db: updater.db, showSQL: updater.showSQL}
    for _, p := range updater.param {
        example.AddParam(p)
    }
	return example
}

type DBCodeInjectModel struct {
	ID int
	Key string
	Description string
	Language string
	Code string
	
}

type DBCodeInjectORM struct {
	ShowSQL bool
	DB      *sqlx.DB
}

func (orm *DBCodeInjectORM) Insert(model *DBCodeInjectModel) *DBCodeInjectORMExample {
	sql := "INSERT INTO t_code_inject"
	sql += "(key,description,language,code)"
	sql += "VALUES"
	sql += "(?,?,?,?)"
	example := &DBCodeInjectORMExample{sql: sql, db: orm.DB, showSQL: orm.ShowSQL}
	
	
	
	
	example.AddParam(model.Key)
	
	
	
	example.AddParam(model.Description)
	
	
	
	example.AddParam(model.Language)
	
	
	
	example.AddParam(model.Code)
	
	
	return example
}
func (orm *DBCodeInjectORM) Update(model *DBCodeInjectModel) *DBCodeInjectORMExample {
    sql := "UPDATE t_code_inject SET "
    sql += "key = ?,description = ?,language = ?,code = ?"
    example := &DBCodeInjectORMExample{sql: sql, db: orm.DB, showSQL: orm.ShowSQL}
    
    
    
    
    example.AddParam(model.Key)
    
    
    
    example.AddParam(model.Description)
    
    
    
    example.AddParam(model.Language)
    
    
    
    example.AddParam(model.Code)
    
    
    return example
}
func (orm *DBCodeInjectORM) UpdateField() *DBCodeInjectUpdater {
	sql := "UPDATE t_code_inject SET "
	return &DBCodeInjectUpdater{
		sql:     sql,
		param:   nil,
		showSQL: orm.ShowSQL,
		db:      orm.DB,
	}
}
func (orm *DBCodeInjectORM) Delete() *DBCodeInjectORMExample {
    sql := "DELETE FROM t_code_inject"
    return &DBCodeInjectORMExample{sql: sql, db: orm.DB, showSQL: orm.ShowSQL}
}
func (orm *DBCodeInjectORM) Query() *DBCodeInjectORMExample {
    sql := "SELECT id as id,key as key,description as description,language as language,code as code FROM t_code_inject"
    return &DBCodeInjectORMExample{sql: sql, db: orm.DB, showSQL: orm.ShowSQL}
}



