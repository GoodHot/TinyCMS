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

const allPluginParamColumn = "pid, key, value, can_edit as canedit, description, value_type as valuetype, visible"

func (orm *PluginORMImpl) SaveParam(param *trait.PluginParam) *core.Err {
	sql := "insert into t_plugin_param(pid, key, value, description, can_edit, value_type, visible) values(?,?,?,?,?,?,?)"
	rst, err := orm.DB.Exec(sql, param.PID, param.Key, param.Value, param.Description, param.CanEdit, param.ValueType, param.Visible)
	if err != nil {
		return core.NewErr(core.Err_Plugin_Save_Param_Fail)
	}
	lastID, err := rst.LastInsertId()
	if err != nil {
		return core.NewErr(core.Err_Plugin_Save_Param_Fail)
	}
	param.ID = int(lastID)
	return nil
}

func (orm *PluginORMImpl) GetParams(pid int) *[]trait.PluginParam {
	sql := "select %v from t_plugin_param where pid = ?"
	var params []trait.PluginParam
	err := orm.DB.Select(&params, fmt.Sprintf(sql, allPluginParamColumn), pid)
	if err != nil {
		return nil
	}
	return &params
}

const allPluginColumn = "id, name, version, description, type, server, internal"

func (orm *PluginORMImpl) GetByType(pluginType string) (*[]trait.Plugin, *core.Err) {
	sql := "select %v from t_plugin where type = ?"
	var plugins []trait.Plugin
	err := orm.DB.Select(&plugins, fmt.Sprintf(sql, allPluginColumn), pluginType)
	if err != nil {
		return nil, core.NewErr(core.Err_Plugin_Type_Not_Exists)
	}
	return &plugins, nil
}

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
	rst, err := orm.DB.Exec(sql, plugin.Name, plugin.Version, plugin.Description, plugin.Type, plugin.Server, plugin.Internal)
	if err != nil {
		return core.NewErr(core.Err_Plugin_Save_Fail)
	}
	lastID, err := rst.LastInsertId()
	if err != nil {
		return core.NewErr(core.Err_Plugin_Save_Fail)
	}
	plugin.ID = int(lastID)
	return nil
}
