package sqlite

import (
	"database/sql"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/impl/sqlite/datasource"
	"github.com/GoodHot/TinyCMS/orm/trait"
)

type CodeORMImpl struct {
	DB *datasource.DBCodeInjectORM
}

func (orm *CodeORMImpl) GetByKey(key string) (*trait.Code, *core.Err) {
	code, err := orm.DB.Query().KeyEq(key).ExecQueryOne()
	if err != nil || code == nil {
		return nil, core.NewErr(core.Err_Code_Key_Not_Exists)
	}
	return &orm.convert(code)[0], nil
}

func (orm *CodeORMImpl) DeleteByID(id int) *core.Err {
	if err := orm.DB.Delete().IDEq(id).Exec(); err != nil {
		return core.NewErr(core.Err_Code_Delete_Fail)
	}
	return nil
}

func (orm *CodeORMImpl) Insert(t *trait.Code) *core.Err {
	err := orm.DB.Insert(&datasource.DBCodeInjectModel{
		Key:         t.Key,
		Description: t.Description,
		Language:    t.Language,
		Code:        t.Code,
	}).Exec(func(result sql.Result) {
		last, _ := result.LastInsertId()
		t.ID = int(last)
	})
	if err != nil {
		return core.NewErr(core.Err_Code_Save_Fail)
	}
	return nil
}

func (orm *CodeORMImpl) Update(t *trait.Code) *core.Err {
	err := orm.DB.Update(&datasource.DBCodeInjectModel{
		Key:         t.Key,
		Description: t.Description,
		Language:    t.Language,
		Code:        t.Code,
	}).IDEq(t.ID).Exec()
	if err != nil {
		return core.NewErr(core.Err_Code_Update_Fail)
	}
	return nil
}

func (orm *CodeORMImpl) GetAll() ([]trait.Code, *core.Err) {
	result, err := orm.DB.Query().ExecQuery()
	if err != nil {
		return nil, core.NewErr(core.Err_Code_Not_Found)
	}
	return orm.convert(result...), nil
}

func (orm *CodeORMImpl) convert(model ...*datasource.DBCodeInjectModel) []trait.Code {
	var rst []trait.Code
	for _, item := range model {
		rst = append(rst, trait.Code{
			BaseORM: trait.BaseORM{
				ID: item.ID,
			},
			Key:         item.Key,
			Description: item.Description,
			Language:    item.Language,
			Code:        item.Code,
		})
	}
	return rst
}
