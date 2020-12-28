package service

import (
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/GoodHot/TinyCMS/orm/trait"
)

type DictService struct {
	ORM *orm.ORMFactory `ioc:"auto"`
}

func (ds *DictService) saveDict(dict *trait.Dict) {
	rst, _ := ds.ORM.Dict.GetByKey(dict.Key)
	if rst != nil {
		// 已经存在的字典数据
		return
	}
	ds.ORM.Dict.Save(dict)
}

func (ds *DictService) Startup() error {
	// 初始化配置
	ds.saveDict(&trait.Dict{
		Name:        "上传插件ID",
		Description: "上传插件的ID，查询t_plugin表",
		FormType:    trait.FormType_Input,
		TypeValue:   "",
		Key:         "upload:plugin",
		Value:       "0",
		Visible:     false,
	})
	return nil
}
