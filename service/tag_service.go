package service

import (
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
)

type TagService struct {
	PageSize int      `val:"${website.page_size}"`
	ORM      *orm.ORM `ioc:"auto"`
}

func (s *TagService) GetByName(name string) *model.Tag {
	var tag model.Tag
	s.ORM.DB.Where("name = ?", name).First(&tag)
	return &tag
}
