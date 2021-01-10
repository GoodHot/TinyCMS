
package datasource

import (
    "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DBMemberORMExample struct {
	showSQL bool
	db      *sqlx.DB
	sql     string
	where   string
	orderBy string
	limit   string
	param   []interface{}
}

func (exp *DBMemberORMExample) And() *DBMemberORMExample {
	exp.where += " and"
	return exp
}
func (exp *DBMemberORMExample) Or() *DBMemberORMExample {
	exp.where += " or"
	return exp
}
func (exp *DBMemberORMExample) Exec(handler ...func(result sql.Result)) error {
	
    sql := exp.sql
    if exp.where != "" {
        sql += " where " + exp.where
    }
    if exp.orderBy != "" {
        sql += " order by " + exp.orderBy
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
func (exp *DBMemberORMExample) ExecQuery() ([]*DBMemberModel, error) {
	
    sql := exp.sql
    if exp.where != "" {
        sql += " where " + exp.where
    }
    if exp.orderBy != "" {
        sql += " order by " + exp.orderBy
    }

    if exp.showSQL {
        fmt.Println("sql: " + sql)
        fmt.Println("param: ", exp.param)
    }

	rows, err := exp.db.Query(sql, exp.param...)
	if err != nil {
		return nil, err
	}
	var result []*DBMemberModel
	for rows.Next() {
		var model DBMemberModel
        if err := rows.Scan( &model.ID, &model.Nickname, &model.Username, &model.Password, &model.Email, &model.Role); err != nil {
            return result, err
        }
        result = append(result, &model)
	}
	return result, nil
}
func (exp *DBMemberORMExample) ExecQueryOne() (*DBMemberModel, error) {
	
    sql := exp.sql
    if exp.where != "" {
        sql += " where " + exp.where
    }
    if exp.orderBy != "" {
        sql += " order by " + exp.orderBy
    }

    if exp.showSQL {
        fmt.Println("sql: " + sql)
        fmt.Println("param: ", exp.param)
    }

	rows, err := exp.db.Query(sql, exp.param...)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
	    var model DBMemberModel
	    if err := rows.Scan( &model.ID, &model.Nickname, &model.Username, &model.Password, &model.Email, &model.Role); err != nil {
            return nil, err
        }
        return &model, nil
	}
	return nil, nil
}
func (exp *DBMemberORMExample) ExecPage() error {
	sql := exp.sql
	if exp.where != "" {
		sql += " where " + exp.where
	}
	fmt.Println(sql)
	return nil
}
func (exp *DBMemberORMExample) AddParam(param interface{}) *DBMemberORMExample {
	exp.param = append(exp.param, param)
	return exp
}

func (exp *DBMemberORMExample) Limit(limit int, offset int) *DBMemberORMExample {
	exp.limit = " limit ? offset ? "
	exp.param = append(exp.param, limit)
	exp.param = append(exp.param, offset)
	return exp
}




func (exp *DBMemberORMExample) IDEq(ID int) *DBMemberORMExample {
	exp.where += "id = ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBMemberORMExample) IDGt(ID int) *DBMemberORMExample {
	exp.where += "id > ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBMemberORMExample) IDLt(ID int) *DBMemberORMExample {
	exp.where += "id < ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBMemberORMExample) IDEqGt(ID int) *DBMemberORMExample {
	exp.where += "id >= ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBMemberORMExample) IDEqLt(ID int) *DBMemberORMExample {
	exp.where += "id <= ?"
	exp.param = append(exp.param, ID)
	return exp
}







func (exp *DBMemberORMExample) NicknameEq(Nickname string) *DBMemberORMExample {
	exp.where += "nickname = ?"
	exp.param = append(exp.param, Nickname)
	return exp
}

func (exp *DBMemberORMExample) NicknameLike(Nickname string) *DBMemberORMExample {
	exp.where += "nickname like '%' || ? || '%'"
	exp.param = append(exp.param, Nickname)
	return exp
}

func (exp *DBMemberORMExample) NicknameLikeAfter(Nickname string) *DBMemberORMExample {
	exp.where += "nickname like ? || '%'"
	exp.param = append(exp.param, Nickname)
	return exp
}

func (exp *DBMemberORMExample) NicknameLikeBefore(Nickname string) *DBMemberORMExample {
	exp.where += "nickname like  '%' || ?"
	exp.param = append(exp.param, Nickname)
	return exp
}





func (exp *DBMemberORMExample) UsernameEq(Username string) *DBMemberORMExample {
	exp.where += "username = ?"
	exp.param = append(exp.param, Username)
	return exp
}

func (exp *DBMemberORMExample) UsernameLike(Username string) *DBMemberORMExample {
	exp.where += "username like '%' || ? || '%'"
	exp.param = append(exp.param, Username)
	return exp
}

func (exp *DBMemberORMExample) UsernameLikeAfter(Username string) *DBMemberORMExample {
	exp.where += "username like ? || '%'"
	exp.param = append(exp.param, Username)
	return exp
}

func (exp *DBMemberORMExample) UsernameLikeBefore(Username string) *DBMemberORMExample {
	exp.where += "username like  '%' || ?"
	exp.param = append(exp.param, Username)
	return exp
}





func (exp *DBMemberORMExample) PasswordEq(Password string) *DBMemberORMExample {
	exp.where += "password = ?"
	exp.param = append(exp.param, Password)
	return exp
}

func (exp *DBMemberORMExample) PasswordLike(Password string) *DBMemberORMExample {
	exp.where += "password like '%' || ? || '%'"
	exp.param = append(exp.param, Password)
	return exp
}

func (exp *DBMemberORMExample) PasswordLikeAfter(Password string) *DBMemberORMExample {
	exp.where += "password like ? || '%'"
	exp.param = append(exp.param, Password)
	return exp
}

func (exp *DBMemberORMExample) PasswordLikeBefore(Password string) *DBMemberORMExample {
	exp.where += "password like  '%' || ?"
	exp.param = append(exp.param, Password)
	return exp
}





func (exp *DBMemberORMExample) EmailEq(Email string) *DBMemberORMExample {
	exp.where += "email = ?"
	exp.param = append(exp.param, Email)
	return exp
}

func (exp *DBMemberORMExample) EmailLike(Email string) *DBMemberORMExample {
	exp.where += "email like '%' || ? || '%'"
	exp.param = append(exp.param, Email)
	return exp
}

func (exp *DBMemberORMExample) EmailLikeAfter(Email string) *DBMemberORMExample {
	exp.where += "email like ? || '%'"
	exp.param = append(exp.param, Email)
	return exp
}

func (exp *DBMemberORMExample) EmailLikeBefore(Email string) *DBMemberORMExample {
	exp.where += "email like  '%' || ?"
	exp.param = append(exp.param, Email)
	return exp
}



func (exp *DBMemberORMExample) RoleEq(Role int) *DBMemberORMExample {
	exp.where += "role = ?"
	exp.param = append(exp.param, Role)
	return exp
}
func (exp *DBMemberORMExample) RoleGt(Role int) *DBMemberORMExample {
	exp.where += "role > ?"
	exp.param = append(exp.param, Role)
	return exp
}
func (exp *DBMemberORMExample) RoleLt(Role int) *DBMemberORMExample {
	exp.where += "role < ?"
	exp.param = append(exp.param, Role)
	return exp
}
func (exp *DBMemberORMExample) RoleEqGt(Role int) *DBMemberORMExample {
	exp.where += "role >= ?"
	exp.param = append(exp.param, Role)
	return exp
}
func (exp *DBMemberORMExample) RoleEqLt(Role int) *DBMemberORMExample {
	exp.where += "role <= ?"
	exp.param = append(exp.param, Role)
	return exp
}





func (exp *DBMemberORMExample) OrderByIDDesc() *DBMemberORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " id desc"
    return exp
}
func (exp *DBMemberORMExample) OrderByIDAsc() *DBMemberORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " id asc"
    return exp
}

func (exp *DBMemberORMExample) OrderByNicknameDesc() *DBMemberORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " nickname desc"
    return exp
}
func (exp *DBMemberORMExample) OrderByNicknameAsc() *DBMemberORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " nickname asc"
    return exp
}

func (exp *DBMemberORMExample) OrderByUsernameDesc() *DBMemberORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " username desc"
    return exp
}
func (exp *DBMemberORMExample) OrderByUsernameAsc() *DBMemberORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " username asc"
    return exp
}

func (exp *DBMemberORMExample) OrderByPasswordDesc() *DBMemberORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " password desc"
    return exp
}
func (exp *DBMemberORMExample) OrderByPasswordAsc() *DBMemberORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " password asc"
    return exp
}

func (exp *DBMemberORMExample) OrderByEmailDesc() *DBMemberORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " email desc"
    return exp
}
func (exp *DBMemberORMExample) OrderByEmailAsc() *DBMemberORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " email asc"
    return exp
}

func (exp *DBMemberORMExample) OrderByRoleDesc() *DBMemberORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " role desc"
    return exp
}
func (exp *DBMemberORMExample) OrderByRoleAsc() *DBMemberORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " role asc"
    return exp
}


type DBMemberUpdater struct {
    setSQL  string
    sql     string
    param   []interface{}
    showSQL bool
    db      *sqlx.DB
}


func (updater *DBMemberUpdater) SetID(ID int) *DBMemberUpdater {
	updater.setSQL += ", id = ?"
	updater.param = append(updater.param, ID)
	return updater
}

func (updater *DBMemberUpdater) SetNickname(Nickname string) *DBMemberUpdater {
	updater.setSQL += ", nickname = ?"
	updater.param = append(updater.param, Nickname)
	return updater
}

func (updater *DBMemberUpdater) SetUsername(Username string) *DBMemberUpdater {
	updater.setSQL += ", username = ?"
	updater.param = append(updater.param, Username)
	return updater
}

func (updater *DBMemberUpdater) SetPassword(Password string) *DBMemberUpdater {
	updater.setSQL += ", password = ?"
	updater.param = append(updater.param, Password)
	return updater
}

func (updater *DBMemberUpdater) SetEmail(Email string) *DBMemberUpdater {
	updater.setSQL += ", email = ?"
	updater.param = append(updater.param, Email)
	return updater
}

func (updater *DBMemberUpdater) SetRole(Role int) *DBMemberUpdater {
	updater.setSQL += ", role = ?"
	updater.param = append(updater.param, Role)
	return updater
}

func (updater *DBMemberUpdater) Done() *DBMemberORMExample {
    example := &DBMemberORMExample{sql: updater.sql + updater.setSQL[1:], db: updater.db, showSQL: updater.showSQL}
    for _, p := range updater.param {
        example.AddParam(p)
    }
	return example
}

type DBMemberModel struct {
	ID int
	Nickname string
	Username string
	Password string
	Email string
	Role int
	
}

type DBMemberORM struct {
	ShowSQL bool
	DB      *sqlx.DB
	sql     string
}

func (orm *DBMemberORM) Insert(model *DBMemberModel) *DBMemberORMExample {
	orm.sql = "INSERT INTO t_member"
	orm.sql += "(nickname,username,password,email,role)"
	orm.sql += "VALUES"
	orm.sql += "(?,?,?,?,?)"
	example := &DBMemberORMExample{sql: orm.sql, db: orm.DB, showSQL: orm.ShowSQL}
	
	
	
	
	example.AddParam(model.Nickname)
	
	
	
	example.AddParam(model.Username)
	
	
	
	example.AddParam(model.Password)
	
	
	
	example.AddParam(model.Email)
	
	
	
	example.AddParam(model.Role)
	
	
	return example
}
func (orm *DBMemberORM) Update(model *DBMemberModel) *DBMemberORMExample {
    orm.sql = "UPDATE t_member SET "
    orm.sql += "nickname = ?,username = ?,password = ?,email = ?,role = ?"
    example := &DBMemberORMExample{sql: orm.sql, db: orm.DB, showSQL: orm.ShowSQL}
    
    
    
    
    example.AddParam(model.Nickname)
    
    
    
    example.AddParam(model.Username)
    
    
    
    example.AddParam(model.Password)
    
    
    
    example.AddParam(model.Email)
    
    
    
    example.AddParam(model.Role)
    
    
    return example
}
func (orm *DBMemberORM) UpdateField() *DBMemberUpdater {
	orm.sql = "UPDATE t_member SET "
	return &DBMemberUpdater{
		sql:     orm.sql,
		param:   nil,
		showSQL: orm.ShowSQL,
		db:      orm.DB,
	}
}
func (orm *DBMemberORM) Delete() *DBMemberORMExample {
    orm.sql = "DELETE FROM t_member"
    return &DBMemberORMExample{sql: orm.sql, db: orm.DB, showSQL: orm.ShowSQL}
}
func (orm *DBMemberORM) Query() *DBMemberORMExample {
    orm.sql = "SELECT id,nickname,username,password,email,role FROM t_member"
    return &DBMemberORMExample{sql: orm.sql, db: orm.DB, showSQL: orm.ShowSQL}
}
