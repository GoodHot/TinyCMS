package service

import (
	"github.com/GoodHot/TinyCMS/core"
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
	ds.saveDict(&trait.Dict{
		Name:        "上传文件类型",
		Description: "网站支持的文件上传类型",
		FormType:    trait.FormType_Input,
		TypeValue:   "",
		Key:         "upload:file_type",
		Value:       "*",
		Visible:     false,
	})
	ds.saveDict(&trait.Dict{
		Name:        "上传文件最大大小",
		Description: "网站上传文件支持的最大大小",
		FormType:    trait.FormType_Input,
		TypeValue:   "",
		Key:         "upload:file_max_size",
		Value:       "10",
		Visible:     false,
	})
	ds.saveDict(&trait.Dict{
		Name:        "皮肤模板",
		Description: "当前系统使用的皮肤模板",
		FormType:    trait.FormType_Input,
		TypeValue:   "",
		Key:         "skin:template",
		Value:       "default",
		Visible:     false,
	})
	ds.saveDict(&trait.Dict{
		Name:        "站点名称",
		Description: "网站站点的名称",
		FormType:    trait.FormType_Input,
		TypeValue:   "",
		Key:         "website:name",
		Value:       "TinyCMS",
		Visible:     false,
	})
	ds.saveDict(&trait.Dict{
		Name:        "站点域名",
		Description: "网站站点的域名",
		FormType:    trait.FormType_Input,
		TypeValue:   "",
		Key:         "website:domain",
		Value:       "http://localhost",
		Visible:     false,
	})
	ds.saveDict(&trait.Dict{
		Name:        "网络图片",
		Description: "自动下载网络图片",
		FormType:    trait.FormType_Switch,
		TypeValue:   "",
		Key:         "website:internet_image",
		Value:       "1",
		Visible:     false,
	})
	return nil
}

func (ds *DictService) Update(key string, value string) *core.Err {
	return ds.ORM.Dict.UpdateValue(key, value)
}

func (ds *DictService) Get(key string) (*trait.Dict, *core.Err) {
	return ds.ORM.Dict.GetByKey(key)
}
