package sqlite

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/impl/sqlite/datasource"
	"github.com/GoodHot/TinyCMS/orm/trait"
)

type CodeORMImpl struct {
	DB *datasource.DBCodeInjectORM
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
