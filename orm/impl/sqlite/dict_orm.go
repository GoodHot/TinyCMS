package sqlite

import (
	"database/sql"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/impl/sqlite/datasource"
	"github.com/GoodHot/TinyCMS/orm/trait"
)

type DictORMImpl struct {
	DB *datasource.DBDictORM
}

func (orm *DictORMImpl) Save(dict *trait.Dict) *core.Err {
	err := orm.DB.Insert(&datasource.DBDictModel{
		Name:        dict.Name,
		Description: dict.Description,
		FormType:    string(dict.FormType),
		TypeValue:   dict.TypeValue,
		Key:         dict.Key,
		Value:       dict.Value,
		Visible:     dict.Visible,
	}).Exec(func(result sql.Result) {
		last, _ := result.LastInsertId()
		dict.ID = int(last)
	})
	if err != nil {
		return core.NewErr(core.Err_Dict_Save_Fail)
	}
	return nil
}

func (orm *DictORMImpl) GetAll() ([]trait.Dict, *core.Err) {
	result, err := orm.DB.Query().ExecQuery()
	if err != nil {
		return nil, core.NewErr(core.Err_Dict_Not_Found)
	}
	return orm.convert(result...), nil
}

func (orm *DictORMImpl) GetByKey(key string) (*trait.Dict, *core.Err) {
	result, err := orm.DB.Query().KeyEq(key).ExecQueryOne()
	if err != nil || result == nil {
		return nil, core.NewErr(core.Err_Dict_Not_Found_Key)
	}
	return &orm.convert(result)[0], nil
}

func (orm *DictORMImpl) UpdateValue(key string, value string) *core.Err {
	err := orm.DB.UpdateField().SetValue(value).Done().KeyEq(key).Exec()
	if err != nil {
		return core.NewErr(core.Err_Dict_Update_Fail)
	}
	return nil
}

func (orm *DictORMImpl) convert(model ...*datasource.DBDictModel) []trait.Dict {
	var rst []trait.Dict
	for _, item := range model {
		rst = append(rst, trait.Dict{
			BaseORM: trait.BaseORM{
				ID: item.ID,
			},
			Name:        item.Name,
			Description: item.Description,
			FormType:    trait.FormType(item.FormType),
			TypeValue:   item.TypeValue,
			Key:         item.Key,
			Value:       item.Value,
			Visible:     item.Visible,
		})
	}
	return rst
}
