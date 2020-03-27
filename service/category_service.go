package service

import (
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/importcjj/trie-go"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)

type CategoryService struct {
	ArticleService *ArticleService `ioc:"auto"`
	PageSize       int             `val:"${website.page_size}"`
	ORM            *orm.ORM        `ioc:"auto"`
	dataTree       *trie.Trie
}

func (s *CategoryService) Get(id int) (*model.Category, error) {
	model := &model.Category{}
	err := s.ORM.DB.Find(model, id).Error
	return model, err
}

func (s *CategoryService) Save(category *model.Category) error {
	return s.ORM.DB.Save(category).Error
}

func (s *CategoryService) Tree() []*CategoryTree {
	return s.tree(0)
}

type CategoryTree struct {
	ID       uint            `json:"id"`
	Name     string          `json:"name"`
	Children []*CategoryTree `json:"children"`
	Category *model.Category `json:"category"`
}

func (s *CategoryService) tree(parentId uint) []*CategoryTree {
	var channels []*model.Category
	s.ORM.DB.Where("parent_id = ?", parentId).Order("sort asc, id desc").Find(&channels)
	if len(channels) == 0 {
		return nil
	}
	var trees []*CategoryTree
	for _, v := range channels {
		child := s.tree(v.ID)
		tmp := &CategoryTree{
			ID:       v.ID,
			Name:     v.Name,
			Children: child,
			Category: v,
		}
		trees = append(trees, tmp)
	}
	return trees
}

func (s *CategoryService) TrieGet(id uint) *model.Category {
	if s.dataTree == nil {
		s.dataTree = trie.New()
		page, size := 1, 20
		row := (page - 1) * size
		var categorys []*model.Category
		for ; ; {
			s.ORM.DB.Limit(size).Offset(row).Find(&categorys)
			if len(categorys) == 0 {
				break
			}
			for _, cate := range categorys {
				s.dataTree.Put(strconv.Itoa(int(cate.ID)), cate)
			}
			page++
			row = (page - 1) * size
		}
	}
	exists, result := s.dataTree.Match(strconv.Itoa(int(id)))
	if !exists {
		return nil
	}
	return result.Value.(*model.Category)
}

type CategoryQuery struct {
	Keyword string `json:"keyword"`
}

func (s *CategoryService) Page(page int, query *CategoryQuery) *orm.Page {
	where := "1 = 1"
	var param []interface{}
	query.Keyword = strings.TrimSpace(query.Keyword)
	if query.Keyword != "" {
		where += " and name = ?"
		param = append(param, query.Keyword)
	}
	var category []*model.Category
	result := s.ORM.Page(orm.ORMPage{
		PageNum:  page,
		PageSize: s.PageSize,
		Result:   &category,
		Where:    orm.Where(where, param...),
		OrderBy:  "id desc",
	})
	return result
}

type CategorySort struct {
	ParentID uint `json:"parent_id"`
	Sort     []struct {
		ID    uint `json:"id"`
		Index int  `json:"index"`
	} `json:"sort"`
}

func (s *CategoryService) Sort(sort *CategorySort) error {
	return s.ORM.Tx(func(db *gorm.DB) error {
		data := make(map[string]interface{})
		for _, item := range sort.Sort {
			data["sort"] = item.Index
			data["parent_id"] = sort.ParentID
			err := s.ORM.DB.Model(&model.Category{}).Where("id = ?", item.ID).Updates(data).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
}
