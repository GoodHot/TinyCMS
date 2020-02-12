package service

import (
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
)

type DictService struct {
	PageSize int      `val:"${website.page_size}"`
	ORM      *orm.ORM `ioc:"auto"`
}

func (s *DictService) Init() {
	s.Save(&model.Dict{
		Key:         "website_name",
		Name:        "网站名称",
		Description: "网站名称",
		Type:        model.DictTypeText,
		Options:     "",
		Value:       "网站名称",
		ReadOnly:    false,
	})
	s.Save(&model.Dict{
		Key:         "website_url",
		Name:        "网站地址",
		Description: "网站地址",
		Type:        model.DictTypeText,
		Options:     "",
		Value:       "网站地址",
		ReadOnly:    false,
	})
	s.Save(&model.Dict{
		Key:         "keyword",
		Name:        "关键字",
		Description: "SEO关键字",
		Type:        model.DictTypeText,
		Options:     "",
		Value:       "网站名称",
		ReadOnly:    false,
	})
	s.Save(&model.Dict{
		Key:         "description",
		Name:        "描述",
		Description: "SEO描述",
		Type:        model.DictTypeTextArea,
		Options:     "",
		Value:       "网站名称",
		ReadOnly:    false,
	})
	s.Save(&model.Dict{
		Key:         "select",
		Name:        "select测试",
		Description: "select测试",
		Type:        model.DictTypeSelect,
		Options:     "1,2,3,4,5",
		Value:       "1",
		ReadOnly:    false,
	})
}

func (s *DictService) Save(dict *model.Dict) {
	s.ORM.DB.Save(dict)
}

func (s *DictService) All() []*model.Dict {
	dicts := []*model.Dict{}
	s.ORM.DB.Find(&dicts)
	return dicts
}

func (s *DictService) GetByKey(key string) *model.Dict {
	var dict model.Dict
	s.ORM.DB.Where("key = ?", key).Find(&dict)
	return &dict
}

func (s *DictService) Update(dict *model.Dict) {
	s.ORM.DB.Model(dict).Update("value", dict.Value)
}
