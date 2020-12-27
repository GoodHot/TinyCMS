package sqlite

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/jmoiron/sqlx"
)

type PluginORMImpl struct {
	DB *sqlx.DB
}

const allPluginColumn = "id, name, version, description, type, server, internal"

func (orm *PluginORMImpl) GetByInfo(name string, version string, pluginType trait.PluginType) (*trait.Plugin, *core.Err) {
	sql := "select %v from t_plugin where name = ? and version = ? and type = ?"
	var plugin trait.Plugin
	err := orm.DB.Get(&plugin, fmt.Sprintf(sql, allPluginColumn), name, version, pluginType)
	if err != nil {
		return nil, core.NewErr(core.Err_Plugin_Not_Exists)
	}
	return &plugin, nil
}

func (orm *PluginORMImpl) Save(plugin *trait.Plugin) *core.Err {
	sql := "insert into t_plugin(name, version, description, type, server, internal) values(?,?,?,?,?,?)"
	_, err := orm.DB.Exec(sql, plugin.Name, plugin.Version, plugin.Description, plugin.Type, plugin.Server, plugin.Internal)
	if err != nil {
		return core.NewErr(core.Err_Plugin_Save_Fail)
	}
	return nil
}