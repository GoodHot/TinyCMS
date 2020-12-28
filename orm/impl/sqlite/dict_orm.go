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

func (orm DictORMImpl) Save(dict *trait.Dict) *core.Err {
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

const allDictCloumn = "id, name, description, form_type as formtype, type_value as typevalue, key, value, visible"

func (d DictORMImpl) GetAll() (*[]trait.Dict, *core.Err) {
	panic("implement me")
}

func (orm DictORMImpl) GetByKey(key string) (*trait.Dict, *core.Err) {
	sql := "select %v from t_dict where key = ?"
	var dict trait.Dict
	orm.DB.Get(&dict, fmt.Sprintf(sql, allDictCloumn), key)
}

func (d DictORMImpl) UpdateValue(key string, value string) *core.Err {
	panic("implement me")
}

