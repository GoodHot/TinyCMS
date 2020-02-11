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

func (s *TagService) GetByArticleID(id int) []*model.Tag {
	var rel []*model.RelTagArticle
	s.ORM.DB.Where("article_id = ?", id).Find(&rel)
	if len(rel) == 0 {
		return nil
	}
	var ids []uint
	for _, v := range rel {
		ids = append(ids, v.TagID)
	}
	var tags []*model.Tag
	s.ORM.DB.Where("id in (?)", ids).Find(&tags)
	return tags
}
