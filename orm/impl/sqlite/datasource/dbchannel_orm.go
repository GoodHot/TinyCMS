
package datasource

import (
    "database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type DBChannelORMExample struct {
	showSQL bool
	db      *sqlx.DB
	sql     string
	where   string
	orderBy string
	limit   string
	param   []interface{}
}

func (exp *DBChannelORMExample) And() *DBChannelORMExample {
	exp.where += " and"
	return exp
}
func (exp *DBChannelORMExample) Or() *DBChannelORMExample {
	exp.where += " or"
	return exp
}
func (exp *DBChannelORMExample) Exec(handler ...func(result sql.Result)) error {
	
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
func (exp *DBChannelORMExample) ExecQuery() ([]*DBChannelModel, error) {
	
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
	var result []*DBChannelModel
	for rows.Next() {
		var model DBChannelModel
        if err := rows.Scan( &model.ID, &model.Name, &model.Path, &model.AvatarSVG, &model.Avatar, &model.Visible, &model.MetaTitle, &model.MetaDescription, &model.Sort, &model.ParentID); err != nil {
            return result, err
        }
        result = append(result, &model)
	}
	return result, nil
}
func (exp *DBChannelORMExample) ExecQueryOne() (*DBChannelModel, error) {
	
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
	    var model DBChannelModel
	    if err := rows.Scan( &model.ID, &model.Name, &model.Path, &model.AvatarSVG, &model.Avatar, &model.Visible, &model.MetaTitle, &model.MetaDescription, &model.Sort, &model.ParentID); err != nil {
            return nil, err
        }
        return &model, nil
	}
	return nil, nil
}
func (exp *DBChannelORMExample) ExecPage() error {
	sql := exp.sql
	if exp.where != "" {
		sql += " where " + exp.where
	}
	fmt.Println(sql)
	return nil
}
func (exp *DBChannelORMExample) AddParam(param interface{}) *DBChannelORMExample {
	exp.param = append(exp.param, param)
	return exp
}

func (exp *DBChannelORMExample) Limit(limit int, offset int) *DBChannelORMExample {
	exp.limit = " limit ? offset ? "
	exp.param = append(exp.param, limit)
	exp.param = append(exp.param, offset)
	return exp
}




func (exp *DBChannelORMExample) IDEq(ID int) *DBChannelORMExample {
	exp.where += "id = ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBChannelORMExample) IDGt(ID int) *DBChannelORMExample {
	exp.where += "id > ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBChannelORMExample) IDLt(ID int) *DBChannelORMExample {
	exp.where += "id < ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBChannelORMExample) IDEqGt(ID int) *DBChannelORMExample {
	exp.where += "id >= ?"
	exp.param = append(exp.param, ID)
	return exp
}
func (exp *DBChannelORMExample) IDEqLt(ID int) *DBChannelORMExample {
	exp.where += "id <= ?"
	exp.param = append(exp.param, ID)
	return exp
}







func (exp *DBChannelORMExample) NameEq(Name string) *DBChannelORMExample {
	exp.where += "name = ?"
	exp.param = append(exp.param, Name)
	return exp
}

func (exp *DBChannelORMExample) NameLike(Name string) *DBChannelORMExample {
	exp.where += "name like '%' || ? || '%'"
	exp.param = append(exp.param, Name)
	return exp
}

func (exp *DBChannelORMExample) NameLikeAfter(Name string) *DBChannelORMExample {
	exp.where += "name like ? || '%'"
	exp.param = append(exp.param, Name)
	return exp
}

func (exp *DBChannelORMExample) NameLikeBefore(Name string) *DBChannelORMExample {
	exp.where += "name like  '%' || ?"
	exp.param = append(exp.param, Name)
	return exp
}





func (exp *DBChannelORMExample) PathEq(Path string) *DBChannelORMExample {
	exp.where += "path = ?"
	exp.param = append(exp.param, Path)
	return exp
}

func (exp *DBChannelORMExample) PathLike(Path string) *DBChannelORMExample {
	exp.where += "path like '%' || ? || '%'"
	exp.param = append(exp.param, Path)
	return exp
}

func (exp *DBChannelORMExample) PathLikeAfter(Path string) *DBChannelORMExample {
	exp.where += "path like ? || '%'"
	exp.param = append(exp.param, Path)
	return exp
}

func (exp *DBChannelORMExample) PathLikeBefore(Path string) *DBChannelORMExample {
	exp.where += "path like  '%' || ?"
	exp.param = append(exp.param, Path)
	return exp
}





func (exp *DBChannelORMExample) AvatarSVGEq(AvatarSVG string) *DBChannelORMExample {
	exp.where += "avatar_svg = ?"
	exp.param = append(exp.param, AvatarSVG)
	return exp
}

func (exp *DBChannelORMExample) AvatarSVGLike(AvatarSVG string) *DBChannelORMExample {
	exp.where += "avatar_svg like '%' || ? || '%'"
	exp.param = append(exp.param, AvatarSVG)
	return exp
}

func (exp *DBChannelORMExample) AvatarSVGLikeAfter(AvatarSVG string) *DBChannelORMExample {
	exp.where += "avatar_svg like ? || '%'"
	exp.param = append(exp.param, AvatarSVG)
	return exp
}

func (exp *DBChannelORMExample) AvatarSVGLikeBefore(AvatarSVG string) *DBChannelORMExample {
	exp.where += "avatar_svg like  '%' || ?"
	exp.param = append(exp.param, AvatarSVG)
	return exp
}





func (exp *DBChannelORMExample) AvatarEq(Avatar string) *DBChannelORMExample {
	exp.where += "avatar = ?"
	exp.param = append(exp.param, Avatar)
	return exp
}

func (exp *DBChannelORMExample) AvatarLike(Avatar string) *DBChannelORMExample {
	exp.where += "avatar like '%' || ? || '%'"
	exp.param = append(exp.param, Avatar)
	return exp
}

func (exp *DBChannelORMExample) AvatarLikeAfter(Avatar string) *DBChannelORMExample {
	exp.where += "avatar like ? || '%'"
	exp.param = append(exp.param, Avatar)
	return exp
}

func (exp *DBChannelORMExample) AvatarLikeBefore(Avatar string) *DBChannelORMExample {
	exp.where += "avatar like  '%' || ?"
	exp.param = append(exp.param, Avatar)
	return exp
}



func (exp *DBChannelORMExample) VisibleEq(Visible int) *DBChannelORMExample {
	exp.where += "visible = ?"
	exp.param = append(exp.param, Visible)
	return exp
}
func (exp *DBChannelORMExample) VisibleGt(Visible int) *DBChannelORMExample {
	exp.where += "visible > ?"
	exp.param = append(exp.param, Visible)
	return exp
}
func (exp *DBChannelORMExample) VisibleLt(Visible int) *DBChannelORMExample {
	exp.where += "visible < ?"
	exp.param = append(exp.param, Visible)
	return exp
}
func (exp *DBChannelORMExample) VisibleEqGt(Visible int) *DBChannelORMExample {
	exp.where += "visible >= ?"
	exp.param = append(exp.param, Visible)
	return exp
}
func (exp *DBChannelORMExample) VisibleEqLt(Visible int) *DBChannelORMExample {
	exp.where += "visible <= ?"
	exp.param = append(exp.param, Visible)
	return exp
}







func (exp *DBChannelORMExample) MetaTitleEq(MetaTitle string) *DBChannelORMExample {
	exp.where += "meta_title = ?"
	exp.param = append(exp.param, MetaTitle)
	return exp
}

func (exp *DBChannelORMExample) MetaTitleLike(MetaTitle string) *DBChannelORMExample {
	exp.where += "meta_title like '%' || ? || '%'"
	exp.param = append(exp.param, MetaTitle)
	return exp
}

func (exp *DBChannelORMExample) MetaTitleLikeAfter(MetaTitle string) *DBChannelORMExample {
	exp.where += "meta_title like ? || '%'"
	exp.param = append(exp.param, MetaTitle)
	return exp
}

func (exp *DBChannelORMExample) MetaTitleLikeBefore(MetaTitle string) *DBChannelORMExample {
	exp.where += "meta_title like  '%' || ?"
	exp.param = append(exp.param, MetaTitle)
	return exp
}





func (exp *DBChannelORMExample) MetaDescriptionEq(MetaDescription string) *DBChannelORMExample {
	exp.where += "meta_description = ?"
	exp.param = append(exp.param, MetaDescription)
	return exp
}

func (exp *DBChannelORMExample) MetaDescriptionLike(MetaDescription string) *DBChannelORMExample {
	exp.where += "meta_description like '%' || ? || '%'"
	exp.param = append(exp.param, MetaDescription)
	return exp
}

func (exp *DBChannelORMExample) MetaDescriptionLikeAfter(MetaDescription string) *DBChannelORMExample {
	exp.where += "meta_description like ? || '%'"
	exp.param = append(exp.param, MetaDescription)
	return exp
}

func (exp *DBChannelORMExample) MetaDescriptionLikeBefore(MetaDescription string) *DBChannelORMExample {
	exp.where += "meta_description like  '%' || ?"
	exp.param = append(exp.param, MetaDescription)
	return exp
}



func (exp *DBChannelORMExample) SortEq(Sort int) *DBChannelORMExample {
	exp.where += "sort = ?"
	exp.param = append(exp.param, Sort)
	return exp
}
func (exp *DBChannelORMExample) SortGt(Sort int) *DBChannelORMExample {
	exp.where += "sort > ?"
	exp.param = append(exp.param, Sort)
	return exp
}
func (exp *DBChannelORMExample) SortLt(Sort int) *DBChannelORMExample {
	exp.where += "sort < ?"
	exp.param = append(exp.param, Sort)
	return exp
}
func (exp *DBChannelORMExample) SortEqGt(Sort int) *DBChannelORMExample {
	exp.where += "sort >= ?"
	exp.param = append(exp.param, Sort)
	return exp
}
func (exp *DBChannelORMExample) SortEqLt(Sort int) *DBChannelORMExample {
	exp.where += "sort <= ?"
	exp.param = append(exp.param, Sort)
	return exp
}





func (exp *DBChannelORMExample) ParentIDEq(ParentID int) *DBChannelORMExample {
	exp.where += "parent_id = ?"
	exp.param = append(exp.param, ParentID)
	return exp
}
func (exp *DBChannelORMExample) ParentIDGt(ParentID int) *DBChannelORMExample {
	exp.where += "parent_id > ?"
	exp.param = append(exp.param, ParentID)
	return exp
}
func (exp *DBChannelORMExample) ParentIDLt(ParentID int) *DBChannelORMExample {
	exp.where += "parent_id < ?"
	exp.param = append(exp.param, ParentID)
	return exp
}
func (exp *DBChannelORMExample) ParentIDEqGt(ParentID int) *DBChannelORMExample {
	exp.where += "parent_id >= ?"
	exp.param = append(exp.param, ParentID)
	return exp
}
func (exp *DBChannelORMExample) ParentIDEqLt(ParentID int) *DBChannelORMExample {
	exp.where += "parent_id <= ?"
	exp.param = append(exp.param, ParentID)
	return exp
}





func (exp *DBChannelORMExample) OrderByIDDesc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " id desc"
    return exp
}
func (exp *DBChannelORMExample) OrderByIDAsc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " id asc"
    return exp
}

func (exp *DBChannelORMExample) OrderByNameDesc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " name desc"
    return exp
}
func (exp *DBChannelORMExample) OrderByNameAsc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " name asc"
    return exp
}

func (exp *DBChannelORMExample) OrderByPathDesc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " path desc"
    return exp
}
func (exp *DBChannelORMExample) OrderByPathAsc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " path asc"
    return exp
}

func (exp *DBChannelORMExample) OrderByAvatarSVGDesc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " avatar_svg desc"
    return exp
}
func (exp *DBChannelORMExample) OrderByAvatarSVGAsc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " avatar_svg asc"
    return exp
}

func (exp *DBChannelORMExample) OrderByAvatarDesc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " avatar desc"
    return exp
}
func (exp *DBChannelORMExample) OrderByAvatarAsc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " avatar asc"
    return exp
}

func (exp *DBChannelORMExample) OrderByVisibleDesc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " visible desc"
    return exp
}
func (exp *DBChannelORMExample) OrderByVisibleAsc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " visible asc"
    return exp
}

func (exp *DBChannelORMExample) OrderByMetaTitleDesc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " meta_title desc"
    return exp
}
func (exp *DBChannelORMExample) OrderByMetaTitleAsc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " meta_title asc"
    return exp
}

func (exp *DBChannelORMExample) OrderByMetaDescriptionDesc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " meta_description desc"
    return exp
}
func (exp *DBChannelORMExample) OrderByMetaDescriptionAsc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " meta_description asc"
    return exp
}

func (exp *DBChannelORMExample) OrderBySortDesc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " sort desc"
    return exp
}
func (exp *DBChannelORMExample) OrderBySortAsc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " sort asc"
    return exp
}

func (exp *DBChannelORMExample) OrderByParentIDDesc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " parent_id desc"
    return exp
}
func (exp *DBChannelORMExample) OrderByParentIDAsc() *DBChannelORMExample {
	if exp.orderBy != "" {
        exp.orderBy += " ,"
    }
    exp.orderBy += " parent_id asc"
    return exp
}


type DBChannelUpdater struct {
    setSQL  string
    sql     string
    param   []interface{}
    showSQL bool
    db      *sqlx.DB
}


func (updater *DBChannelUpdater) SetID(ID int) *DBChannelUpdater {
	updater.setSQL += ", id = ?"
	updater.param = append(updater.param, ID)
	return updater
}

func (updater *DBChannelUpdater) SetName(Name string) *DBChannelUpdater {
	updater.setSQL += ", name = ?"
	updater.param = append(updater.param, Name)
	return updater
}

func (updater *DBChannelUpdater) SetPath(Path string) *DBChannelUpdater {
	updater.setSQL += ", path = ?"
	updater.param = append(updater.param, Path)
	return updater
}

func (updater *DBChannelUpdater) SetAvatarSVG(AvatarSVG string) *DBChannelUpdater {
	updater.setSQL += ", avatar_svg = ?"
	updater.param = append(updater.param, AvatarSVG)
	return updater
}

func (updater *DBChannelUpdater) SetAvatar(Avatar string) *DBChannelUpdater {
	updater.setSQL += ", avatar = ?"
	updater.param = append(updater.param, Avatar)
	return updater
}

func (updater *DBChannelUpdater) SetVisible(Visible int) *DBChannelUpdater {
	updater.setSQL += ", visible = ?"
	updater.param = append(updater.param, Visible)
	return updater
}

func (updater *DBChannelUpdater) SetMetaTitle(MetaTitle string) *DBChannelUpdater {
	updater.setSQL += ", meta_title = ?"
	updater.param = append(updater.param, MetaTitle)
	return updater
}

func (updater *DBChannelUpdater) SetMetaDescription(MetaDescription string) *DBChannelUpdater {
	updater.setSQL += ", meta_description = ?"
	updater.param = append(updater.param, MetaDescription)
	return updater
}

func (updater *DBChannelUpdater) SetSort(Sort int) *DBChannelUpdater {
	updater.setSQL += ", sort = ?"
	updater.param = append(updater.param, Sort)
	return updater
}

func (updater *DBChannelUpdater) SetParentID(ParentID int) *DBChannelUpdater {
	updater.setSQL += ", parent_id = ?"
	updater.param = append(updater.param, ParentID)
	return updater
}

func (updater *DBChannelUpdater) Done() *DBChannelORMExample {
    example := &DBChannelORMExample{sql: updater.sql + updater.setSQL[1:], db: updater.db, showSQL: updater.showSQL}
    for _, p := range updater.param {
        example.AddParam(p)
    }
	return example
}

type DBChannelModel struct {
	ID int
	Name string
	Path string
	AvatarSVG string
	Avatar string
	Visible int
	MetaTitle string
	MetaDescription string
	Sort int
	ParentID int
	
}

type DBChannelORM struct {
	ShowSQL bool
	DB      *sqlx.DB
	sql     string
}

func (orm *DBChannelORM) Insert(model *DBChannelModel) *DBChannelORMExample {
	orm.sql = "INSERT INTO t_channel"
	orm.sql += "(name,path,avatar_svg,avatar,visible,meta_title,meta_description,sort,parent_id)"
	orm.sql += "VALUES"
	orm.sql += "(?,?,?,?,?,?,?,?,?)"
	example := &DBChannelORMExample{sql: orm.sql, db: orm.DB, showSQL: orm.ShowSQL}
	
	
	
	
	example.AddParam(model.Name)
	
	
	
	example.AddParam(model.Path)
	
	
	
	example.AddParam(model.AvatarSVG)
	
	
	
	example.AddParam(model.Avatar)
	
	
	
	example.AddParam(model.Visible)
	
	
	
	example.AddParam(model.MetaTitle)
	
	
	
	example.AddParam(model.MetaDescription)
	
	
	
	example.AddParam(model.Sort)
	
	
	
	example.AddParam(model.ParentID)
	
	
	return example
}
func (orm *DBChannelORM) Update(model *DBChannelModel) *DBChannelORMExample {
    orm.sql = "UPDATE t_channel SET "
    orm.sql += "name = ?,path = ?,avatar_svg = ?,avatar = ?,visible = ?,meta_title = ?,meta_description = ?,sort = ?,parent_id = ?"
    example := &DBChannelORMExample{sql: orm.sql, db: orm.DB, showSQL: orm.ShowSQL}
    
    
    
    
    example.AddParam(model.Name)
    
    
    
    example.AddParam(model.Path)
    
    
    
    example.AddParam(model.AvatarSVG)
    
    
    
    example.AddParam(model.Avatar)
    
    
    
    example.AddParam(model.Visible)
    
    
    
    example.AddParam(model.MetaTitle)
    
    
    
    example.AddParam(model.MetaDescription)
    
    
    
    example.AddParam(model.Sort)
    
    
    
    example.AddParam(model.ParentID)
    
    
    return example
}
func (orm *DBChannelORM) UpdateField() *DBChannelUpdater {
	orm.sql = "UPDATE t_channel SET "
	return &DBChannelUpdater{
		sql:     orm.sql,
		param:   nil,
		showSQL: orm.ShowSQL,
		db:      orm.DB,
	}
}
func (orm *DBChannelORM) Delete() *DBChannelORMExample {
    orm.sql = "DELETE FROM t_channel"
    return &DBChannelORMExample{sql: orm.sql, db: orm.DB, showSQL: orm.ShowSQL}
}
func (orm *DBChannelORM) Query() *DBChannelORMExample {
    orm.sql = "SELECT id,name,path,avatar_svg,avatar,visible,meta_title,meta_description,sort,parent_id FROM t_channel"
    return &DBChannelORMExample{sql: orm.sql, db: orm.DB, showSQL: orm.ShowSQL}
}
