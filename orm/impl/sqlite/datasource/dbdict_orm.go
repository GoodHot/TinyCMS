
package datasource

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DBDictORMExample struct {
	showSQL bool
	db      *sqlx.DB
	sql     string
	where   string
	orderBy string
	limit   string
	param   []interface{}
}

func (exp *DBDictORMExample) And() *DBDictORMExample {
	exp.where += " and"
	return exp
}
func (exp *DBDictORMExample) Or() *DBDictORMExample {
	exp.where += " or"
	return exp
}
func (exp *DBDictORMExample) Exec(handler ...func(result sql.Result)) error {
	
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
func (exp *DBDictORMExample) ExecQueryCount() (int, error) {
	sql := "SELECT COUNT(1) FROM t_dict "
	if exp.where != "" {
        sql += " where " + exp.where
    }
    if exp.showSQL {
        fmt.Println("sql: " + sql)
        fmt.Println("param: ", exp.param)
    }
    rows, err := exp.db.Query(sql, exp.param...)
    defer rows.Close()
    if err != nil {
        return 0, err
    }
    if rows.Next() {
	    var count int
	    if err := rows.Scan(&count); err != nil {
            return 0, err
        }
        return count, nil
	}
	return 0, nil
}
func (exp *DBDictORMExample) ExecQuery() ([]*DBDictModel, error) {
	
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
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var result []*DBDictModel
	for rows.Next() {
		var model DBDictModel
        if err := rows.StructScan(&model); err != nil {
            return result, err
        }
        result = append(result, &model)
	}
	return result, nil
}
func (exp *DBDictORMExample) ExecQueryOne() (*DBDictModel, error) {
	
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
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	if rows.Next() {
	    var model DBDictModel
	    if err := rows.StructScan(&model); err != nil {
            return nil, err
        }
        return &model, nil
	}
	return nil, nil
}
func (exp *DBDictORMExample) ExecPage() error {
	sql := exp.sql
	if exp.where != "" {
		sql += " where " + exp.where
	}
	fmt.Println(sql)
	return nil
}
func (exp *DBDictORMExample) AddParam(param interface{}) *DBDictORMExample {
	exp.param = append(exp.param, param)
	return exp
}

func (exp *DBDictORMExample) Limit(limit int, offset int) *DBDictORMExample {
	exp.limit = " limit ? offset ? "
	exp.param = append(exp.param, limit)
	exp.param = append(exp.param, offset)
	return exp
}




func (exp *DBDictORMExample) IDEq(ID int) *DBDictORMExample {
	exp.where += "id = ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBDictORMExample) IDNotEq(ID int) *DBDictORMExample {
	exp.where += "id = ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBDictORMExample) IDGt(ID int) *DBDictORMExample {
	exp.where += "id > ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBDictORMExample) IDLt(ID int) *DBDictORMExample {
	exp.where += "id < ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBDictORMExample) IDEqGt(ID int) *DBDictORMExample {
	exp.where += "id >= ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBDictORMExample) IDEqLt(ID int) *DBDictORMExample {
	exp.where += "id <= ?"
	exp.param = append(exp.param, ID)
	return exp
}








func (exp *DBDictORMExample) NameEq(Name string) *DBDictORMExample {
	exp.where += "name = ?"
	exp.param = append(exp.param, Name)
	return exp
}

func (exp *DBDictORMExample) NameNotEq(Name string) *DBDictORMExample {
	exp.where += "name != ?"
	exp.param = append(exp.param, Name)
	return exp
}

func (exp *DBDictORMExample) NameLike(Name string) *DBDictORMExample {
	exp.where += "name like '%' || ? || '%'"
	exp.param = append(exp.param, Name)
	return exp
}

func (exp *DBDictORMExample) NameLikeAfter(Name string) *DBDictORMExample {
	exp.where += "name like ? || '%'"
	exp.param = append(exp.param, Name)
	return exp
}

func (exp *DBDictORMExample) NameLikeBefore(Name string) *DBDictORMExample {
	exp.where += "name like  '%' || ?"
	exp.param = append(exp.param, Name)
	return exp
}






func (exp *DBDictORMExample) DescriptionEq(Description string) *DBDictORMExample {
	exp.where += "description = ?"
	exp.param = append(exp.param, Description)
	return exp
}

func (exp *DBDictORMExample) DescriptionNotEq(Description string) *DBDictORMExample {
	exp.where += "description != ?"
	exp.param = append(exp.param, Description)
	return exp
}

func (exp *DBDictORMExample) DescriptionLike(Description string) *DBDictORMExample {
	exp.where += "description like '%' || ? || '%'"
	exp.param = append(exp.param, Description)
	return exp
}

func (exp *DBDictORMExample) DescriptionLikeAfter(Description string) *DBDictORMExample {
	exp.where += "description like ? || '%'"
	exp.param = append(exp.param, Description)
	return exp
}

func (exp *DBDictORMExample) DescriptionLikeBefore(Description string) *DBDictORMExample {
	exp.where += "description like  '%' || ?"
	exp.param = append(exp.param, Description)
	return exp
}






func (exp *DBDictORMExample) FormTypeEq(FormType string) *DBDictORMExample {
	exp.where += "form_type = ?"
	exp.param = append(exp.param, FormType)
	return exp
}

func (exp *DBDictORMExample) FormTypeNotEq(FormType string) *DBDictORMExample {
	exp.where += "form_type != ?"
	exp.param = append(exp.param, FormType)
	return exp
}

func (exp *DBDictORMExample) FormTypeLike(FormType string) *DBDictORMExample {
	exp.where += "form_type like '%' || ? || '%'"
	exp.param = append(exp.param, FormType)
	return exp
}

func (exp *DBDictORMExample) FormTypeLikeAfter(FormType string) *DBDictORMExample {
	exp.where += "form_type like ? || '%'"
	exp.param = append(exp.param, FormType)
	return exp
}

func (exp *DBDictORMExample) FormTypeLikeBefore(FormType string) *DBDictORMExample {
	exp.where += "form_type like  '%' || ?"
	exp.param = append(exp.param, FormType)
	return exp
}






func (exp *DBDictORMExample) TypeValueEq(TypeValue string) *DBDictORMExample {
	exp.where += "type_value = ?"
	exp.param = append(exp.param, TypeValue)
	return exp
}

func (exp *DBDictORMExample) TypeValueNotEq(TypeValue string) *DBDictORMExample {
	exp.where += "type_value != ?"
	exp.param = append(exp.param, TypeValue)
	return exp
}

func (exp *DBDictORMExample) TypeValueLike(TypeValue string) *DBDictORMExample {
	exp.where += "type_value like '%' || ? || '%'"
	exp.param = append(exp.param, TypeValue)
	return exp
}

func (exp *DBDictORMExample) TypeValueLikeAfter(TypeValue string) *DBDictORMExample {
	exp.where += "type_value like ? || '%'"
	exp.param = append(exp.param, TypeValue)
	return exp
}

func (exp *DBDictORMExample) TypeValueLikeBefore(TypeValue string) *DBDictORMExample {
	exp.where += "type_value like  '%' || ?"
	exp.param = append(exp.param, TypeValue)
	return exp
}






func (exp *DBDictORMExample) KeyEq(Key string) *DBDictORMExample {
	exp.where += "key = ?"
	exp.param = append(exp.param, Key)
	return exp
}

func (exp *DBDictORMExample) KeyNotEq(Key string) *DBDictORMExample {
	exp.where += "key != ?"
	exp.param = append(exp.param, Key)
	return exp
}

func (exp *DBDictORMExample) KeyLike(Key string) *DBDictORMExample {
	exp.where += "key like '%' || ? || '%'"
	exp.param = append(exp.param, Key)
	return exp
}

func (exp *DBDictORMExample) KeyLikeAfter(Key string) *DBDictORMExample {
	exp.where += "key like ? || '%'"
	exp.param = append(exp.param, Key)
	return exp
}

func (exp *DBDictORMExample) KeyLikeBefore(Key string) *DBDictORMExample {
	exp.where += "key like  '%' || ?"
	exp.param = append(exp.param, Key)
	return exp
}






func (exp *DBDictORMExample) ValueEq(Value string) *DBDictORMExample {
	exp.where += "value = ?"
	exp.param = append(exp.param, Value)
	return exp
}

func (exp *DBDictORMExample) ValueNotEq(Value string) *DBDictORMExample {
	exp.where += "value != ?"
	exp.param = append(exp.param, Value)
	return exp
}

func (exp *DBDictORMExample) ValueLike(Value string) *DBDictORMExample {
	exp.where += "value like '%' || ? || '%'"
	exp.param = append(exp.param, Value)
	return exp
}

func (exp *DBDictORMExample) ValueLikeAfter(Value string) *DBDictORMExample {
	exp.where += "value like ? || '%'"
	exp.param = append(exp.param, Value)
	return exp
}

func (exp *DBDictORMExample) ValueLikeBefore(Value string) *DBDictORMExample {
	exp.where += "value like  '%' || ?"
	exp.param = append(exp.param, Value)
	return exp
}







func (exp *DBDictORMExample) VisibleEq(Visible bool) *DBDictORMExample {
	exp.where += "visible = ?"
	exp.param = append(exp.param, Visible)
	return exp
}
func (exp *DBDictORMExample) VisibleNotEq(Visible bool) *DBDictORMExample {
	exp.where += "visible != ?"
	exp.param = append(exp.param, Visible)
	return exp
}



func (exp *DBDictORMExample) OrderByIDDesc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " id desc"
    return exp
}
func (exp *DBDictORMExample) OrderByIDAsc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " id asc"
    return exp
}

func (exp *DBDictORMExample) OrderByNameDesc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " name desc"
    return exp
}
func (exp *DBDictORMExample) OrderByNameAsc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " name asc"
    return exp
}

func (exp *DBDictORMExample) OrderByDescriptionDesc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " description desc"
    return exp
}
func (exp *DBDictORMExample) OrderByDescriptionAsc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " description asc"
    return exp
}

func (exp *DBDictORMExample) OrderByFormTypeDesc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " form_type desc"
    return exp
}
func (exp *DBDictORMExample) OrderByFormTypeAsc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " form_type asc"
    return exp
}

func (exp *DBDictORMExample) OrderByTypeValueDesc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " type_value desc"
    return exp
}
func (exp *DBDictORMExample) OrderByTypeValueAsc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " type_value asc"
    return exp
}

func (exp *DBDictORMExample) OrderByKeyDesc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " key desc"
    return exp
}
func (exp *DBDictORMExample) OrderByKeyAsc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " key asc"
    return exp
}

func (exp *DBDictORMExample) OrderByValueDesc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " value desc"
    return exp
}
func (exp *DBDictORMExample) OrderByValueAsc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " value asc"
    return exp
}

func (exp *DBDictORMExample) OrderByVisibleDesc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " visible desc"
    return exp
}
func (exp *DBDictORMExample) OrderByVisibleAsc() *DBDictORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " visible asc"
    return exp
}


type DBDictUpdater struct {
    setSQL  string
    sql     string
    param   []interface{}
    showSQL bool
    db      *sqlx.DB
}


func (updater *DBDictUpdater) SetID(ID int) *DBDictUpdater {
	updater.setSQL += ", id = ?"
	updater.param = append(updater.param, ID)
	return updater
}

func (updater *DBDictUpdater) SetName(Name string) *DBDictUpdater {
	updater.setSQL += ", name = ?"
	updater.param = append(updater.param, Name)
	return updater
}

func (updater *DBDictUpdater) SetDescription(Description string) *DBDictUpdater {
	updater.setSQL += ", description = ?"
	updater.param = append(updater.param, Description)
	return updater
}

func (updater *DBDictUpdater) SetFormType(FormType string) *DBDictUpdater {
	updater.setSQL += ", form_type = ?"
	updater.param = append(updater.param, FormType)
	return updater
}

func (updater *DBDictUpdater) SetTypeValue(TypeValue string) *DBDictUpdater {
	updater.setSQL += ", type_value = ?"
	updater.param = append(updater.param, TypeValue)
	return updater
}

func (updater *DBDictUpdater) SetKey(Key string) *DBDictUpdater {
	updater.setSQL += ", key = ?"
	updater.param = append(updater.param, Key)
	return updater
}

func (updater *DBDictUpdater) SetValue(Value string) *DBDictUpdater {
	updater.setSQL += ", value = ?"
	updater.param = append(updater.param, Value)
	return updater
}

func (updater *DBDictUpdater) SetVisible(Visible bool) *DBDictUpdater {
	updater.setSQL += ", visible = ?"
	updater.param = append(updater.param, Visible)
	return updater
}

func (updater *DBDictUpdater) Done() *DBDictORMExample {
    example := &DBDictORMExample{sql: updater.sql + updater.setSQL[1:], db: updater.db, showSQL: updater.showSQL}
    for _, p := range updater.param {
        example.AddParam(p)
    }
	return example
}

type DBDictModel struct {
	ID int
	Name string
	Description string
	FormType string
	TypeValue string
	Key string
	Value string
	Visible bool
	
}

type DBDictORM struct {
	ShowSQL bool
	DB      *sqlx.DB
}

func (orm *DBDictORM) Insert(model *DBDictModel) *DBDictORMExample {
	sql := "INSERT INTO t_dict"
	sql += "(name,description,form_type,type_value,key,value,visible)"
	sql += "VALUES"
	sql += "(?,?,?,?,?,?,?)"
	example := &DBDictORMExample{sql: sql, db: orm.DB, showSQL: orm.ShowSQL}
	
	
	
	
	example.AddParam(model.Name)
	
	
	
	example.AddParam(model.Description)
	
	
	
	example.AddParam(model.FormType)
	
	
	
	example.AddParam(model.TypeValue)
	
	
	
	example.AddParam(model.Key)
	
	
	
	example.AddParam(model.Value)
	
	
	
	example.AddParam(model.Visible)
	
	
	return example
}
func (orm *DBDictORM) Update(model *DBDictModel) *DBDictORMExample {
    sql := "UPDATE t_dict SET "
    sql += "name = ?,description = ?,form_type = ?,type_value = ?,key = ?,value = ?,visible = ?"
    example := &DBDictORMExample{sql: sql, db: orm.DB, showSQL: orm.ShowSQL}
    
    
    
    
    example.AddParam(model.Name)
    
    
    
    example.AddParam(model.Description)
    
    
    
    example.AddParam(model.FormType)
    
    
    
    example.AddParam(model.TypeValue)
    
    
    
    example.AddParam(model.Key)
    
    
    
    example.AddParam(model.Value)
    
    
    
    example.AddParam(model.Visible)
    
    
    return example
}
func (orm *DBDictORM) UpdateField() *DBDictUpdater {
	sql := "UPDATE t_dict SET "
	return &DBDictUpdater{
		sql:     sql,
		param:   nil,
		showSQL: orm.ShowSQL,
		db:      orm.DB,
	}
}
func (orm *DBDictORM) Delete() *DBDictORMExample {
    sql := "DELETE FROM t_dict"
    return &DBDictORMExample{sql: sql, db: orm.DB, showSQL: orm.ShowSQL}
}
func (orm *DBDictORM) Query() *DBDictORMExample {
    sql := "SELECT id as id,name as name,description as description,form_type as formtype,type_value as typevalue,key as key,value as value,visible as visible FROM t_dict"
    return &DBDictORMExample{sql: sql, db: orm.DB, showSQL: orm.ShowSQL}
}



