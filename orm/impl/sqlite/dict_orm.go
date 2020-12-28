package sqlite

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/jmoiron/sqlx"
)

type DictORMImpl struct {
	DB *sqlx.DB
}

func (orm *DictORMImpl) Save(dict *trait.Dict) *core.Err {
	sql := "insert into t_dict(name, description, form_type, type_value, key, value, visible) values(?,?,?,?,?,?,?)"
	rst, err := orm.DB.Exec(sql, dict.Name, dict.Description, dict.FormType, dict.TypeValue, dict.Key, dict.Value, dict.Visible)
	if err != nil {
		return core.NewErr(core.Err_Dict_Save_Fail)
	}
	lastID, err := rst.LastInsertId()
	if err != nil {
		return core.NewErr(core.Err_Dict_Save_Fail)
	}
	dict.ID = int(lastID)
	return nil
}

const allDictColumn = "id, name, description, form_type as formtype, type_value as typevalue, key, value, visible"

func (d *DictORMImpl) GetAll() (*[]trait.Dict, *core.Err) {
	panic("implement me")
}

func (orm *DictORMImpl) GetByKey(key string) (*trait.Dict, *core.Err) {
	sql := "select %v from t_dict where key = ?"
	var dict trait.Dict
	if err := orm.DB.Get(&dict, fmt.Sprintf(sql, allDictColumn), key); err != nil {
		return nil, core.NewErr(core.Err_Dict_Not_Found_Key)
	}
	return &dict, nil
}

func (orm *DictORMImpl) UpdateValue(key string, value string) *core.Err {
	sql := "update t_dict set value = ? where key = ?"
	if _, err := orm.DB.Exec(sql, value, key); err != nil {
		return core.NewErr(core.Err_Dict_Update_Fail)
	}
	return nil

}
