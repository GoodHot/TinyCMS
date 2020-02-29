package service

import (
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/armon/go-radix"
)

type TagService struct {
	PageSize   int      `val:"${website.page_size}"`
	ORM        *orm.ORM `ioc:"auto"`
	prefixTree *radix.Tree
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

func (s *TagService) PutPrefix(tag *model.Tag) {
	s.initPrefixTree()
	s.prefixTree.Insert(tag.Name, tag)
}

func (s *TagService) initPrefixTree() {
	if s.prefixTree != nil {
		return
	}
	s.prefixTree = radix.New()
	page, size := 1, 20
	row := (page - 1) * size
	var tags []*model.Tag
	for ; ; {
		s.ORM.DB.Limit(size).Offset(row).Find(&tags)
		if len(tags) == 0 {
			break
		}
		for _, tag := range tags {
			s.prefixTree.Insert(tag.Name, tag)
		}
		page++
		row = (page - 1) * size
	}
}

func (s *TagService) PrefixFind(prefix string) []string {
	s.initPrefixTree()
	count := 0
	var result []string
	s.prefixTree.WalkPrefix(prefix, func(s string, v interface{}) bool {
		if count == 8 {
			return true
		}
		tag := v.(*model.Tag)
		result = append(result, tag.Name)
		count++
		return false
	})
	return result
}

func (s *TagService) GetHotTag(count int) []*model.Tag {
	var tags []*model.Tag
	s.ORM.DB.Where("article_count > ?", 0).Order("updated_at desc, article_count desc").Limit(count).Find(&tags)
	return tags
}

func (s *TagService) getArticleIDs(id int) []uint {
	var rels []*model.RelTagArticle
	s.ORM.DB.Where("tag_id = ?", id).Find(&rels)
	var ids []uint
	for _, rel := range rels {
		ids = append(ids, rel.ArticleID)
	}
	return ids
}
