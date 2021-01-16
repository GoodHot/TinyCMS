
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
	exp.where += " and "
	return exp
}
func (exp *DBMemberORMExample) Or() *DBMemberORMExample {
	exp.where += " or "
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
func (exp *DBMemberORMExample) ExecQueryCount() (int, error) {
	sql := "SELECT COUNT(1) FROM t_member "
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
func (exp *DBMemberORMExample) ExecQuery() ([]*DBMemberModel, error) {
	
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
	var result []*DBMemberModel
	for rows.Next() {
		var model DBMemberModel
        if err := rows.StructScan(&model); err != nil {
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
	    var model DBMemberModel
	    if err := rows.StructScan(&model); err != nil {
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
func (exp *DBMemberORMExample) IDNotEq(ID int) *DBMemberORMExample {
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

func (exp *DBMemberORMExample) NicknameNotEq(Nickname string) *DBMemberORMExample {
	exp.where += "nickname != ?"
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

func (exp *DBMemberORMExample) UsernameNotEq(Username string) *DBMemberORMExample {
	exp.where += "username != ?"
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

func (exp *DBMemberORMExample) PasswordNotEq(Password string) *DBMemberORMExample {
	exp.where += "password != ?"
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

func (exp *DBMemberORMExample) EmailNotEq(Email string) *DBMemberORMExample {
	exp.where += "email != ?"
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
func (exp *DBMemberORMExample) RoleNotEq(Role int) *DBMemberORMExample {
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








func (exp *DBMemberORMExample) AvatarEq(Avatar string) *DBMemberORMExample {
	exp.where += "avatar = ?"
	exp.param = append(exp.param, Avatar)
	return exp
}

func (exp *DBMemberORMExample) AvatarNotEq(Avatar string) *DBMemberORMExample {
	exp.where += "avatar != ?"
	exp.param = append(exp.param, Avatar)
	return exp
}

func (exp *DBMemberORMExample) AvatarLike(Avatar string) *DBMemberORMExample {
	exp.where += "avatar like '%' || ? || '%'"
	exp.param = append(exp.param, Avatar)
	return exp
}

func (exp *DBMemberORMExample) AvatarLikeAfter(Avatar string) *DBMemberORMExample {
	exp.where += "avatar like ? || '%'"
	exp.param = append(exp.param, Avatar)
	return exp
}

func (exp *DBMemberORMExample) AvatarLikeBefore(Avatar string) *DBMemberORMExample {
	exp.where += "avatar like  '%' || ?"
	exp.param = append(exp.param, Avatar)
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

func (exp *DBMemberORMExample) OrderByAvatarDesc() *DBMemberORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " avatar desc"
    return exp
}
func (exp *DBMemberORMExample) OrderByAvatarAsc() *DBMemberORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " avatar asc"
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

func (updater *DBMemberUpdater) SetAvatar(Avatar string) *DBMemberUpdater {
	updater.setSQL += ", avatar = ?"
	updater.param = append(updater.param, Avatar)
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
	Avatar string
	
}

type DBMemberORM struct {
	ShowSQL bool
	DB      *sqlx.DB
}

func (orm *DBMemberORM) Insert(model *DBMemberModel) *DBMemberORMExample {
	sql := "INSERT INTO t_member"
	sql += "(nickname,username,password,email,role,avatar)"
	sql += "VALUES"
	sql += "(?,?,?,?,?,?)"
	example := &DBMemberORMExample{sql: sql, db: orm.DB, showSQL: orm.ShowSQL}
	
	
	
	
	example.AddParam(model.Nickname)
	
	
	
	example.AddParam(model.Username)
	
	
	
	example.AddParam(model.Password)
	
	
	
	example.AddParam(model.Email)
	
	
	
	example.AddParam(model.Role)
	
	
	
	example.AddParam(model.Avatar)
	
	
	return example
}
func (orm *DBMemberORM) Update(model *DBMemberModel) *DBMemberORMExample {
    sql := "UPDATE t_member SET "
    sql += "nickname = ?,username = ?,password = ?,email = ?,role = ?,avatar = ?"
    example := &DBMemberORMExample{sql: sql, db: orm.DB, showSQL: orm.ShowSQL}
    
    
    
    
    example.AddParam(model.Nickname)
    
    
    
    example.AddParam(model.Username)
    
    
    
    example.AddParam(model.Password)
    
    
    
    example.AddParam(model.Email)
    
    
    
    example.AddParam(model.Role)
    
    
    
    example.AddParam(model.Avatar)
    
    
    return example
}
func (orm *DBMemberORM) UpdateField() *DBMemberUpdater {
	sql := "UPDATE t_member SET "
	return &DBMemberUpdater{
		sql:     sql,
		param:   nil,
		showSQL: orm.ShowSQL,
		db:      orm.DB,
	}
}
func (orm *DBMemberORM) Delete() *DBMemberORMExample {
    sql := "DELETE FROM t_member"
    return &DBMemberORMExample{sql: sql, db: orm.DB, showSQL: orm.ShowSQL}
}
func (orm *DBMemberORM) Query() *DBMemberORMExample {
    sql := "SELECT id as id,nickname as nickname,username as username,password as password,email as email,role as role,avatar as avatar FROM t_member"
    return &DBMemberORMExample{sql: sql, db: orm.DB, showSQL: orm.ShowSQL}
}



