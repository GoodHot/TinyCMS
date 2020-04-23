package service


import (
	"errors"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/armon/go-radix"
	"strings"
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
	var ids []int
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

func (s *TagService) getArticleIDs(id int) []int {
	var rels []*model.RelTagArticle
	s.ORM.DB.Where("tag_id = ?", id).Find(&rels)
	var ids []int
	for _, rel := range rels {
		ids = append(ids, rel.ArticleID)
	}
	return ids
}

type TagQuery struct {
	Keyword string `json:"keyword"`
}

func (s *TagService) Page(page int, query *TagQuery) *orm.Page {
	where := "1 = 1"
	var param []interface{}
	query.Keyword = strings.TrimSpace(query.Keyword)
	if query.Keyword != "" {
		where += " and name = ?"
		param = append(param, query.Keyword)
	}
	var tags []*model.Tag
	result := s.ORM.Page(orm.ORMPage{
		PageNum:  page,
		PageSize: s.PageSize,
		Result:   &tags,
		Where:    orm.Where(where, param...),
		OrderBy:  "id desc",
	})
	return result
}

func (s *TagService) Get(id int) (*model.Tag, error) {
	model := &model.Tag{}
	err := s.ORM.DB.Find(model, id).Error
	return model, err
}

func (s *TagService) Save(tag *model.Tag) error {
	var tags []*model.Tag
	s.ORM.DB.Where("name = ? and id != ?", tag.Name, tag.ID).Find(&tags)
	if len(tags) > 0 {
		return errors.New("标签名" + tag.Name + "已存在")
	}
	s.ORM.DB.Where("path = ? and id != ?", tag.Path, tag.ID).Find(&tags)
	if len(tags) > 0 {
		return errors.New("路径" + tag.Path + "已存在")
	}
	if tag.ID == 0 {
		return s.ORM.DB.Save(tag).Error
	}
	update := map[string]interface{}{
		"name": tag.Name,
		"path": tag.Path,
		"description": tag.Description,
		"meta_title": tag.MetaTitle,
		"meta_description": tag.MetaDescription,
		"icon": tag.Icon,
	}
	return s.ORM.DB.Model(&model.Tag{}).Where("id = ?", tag.ID).Updates(update).Error
}

func (s *TagService) All() []*model.Tag {
	var tags []*model.Tag
	s.ORM.DB.Find(&tags)
	return tags
}

