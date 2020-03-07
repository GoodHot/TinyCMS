package service

import (
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/importcjj/trie-go"
	"strconv"
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
	Channel  *model.Category `json:"channel"`
}

func (s *CategoryService) tree(parentId uint) []*CategoryTree {
	var channels []*model.Category
	s.ORM.DB.Where("parent_id = ?", parentId).Order("sort desc, id desc").Find(&channels)
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
			Channel:  v,
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

//func (s *ChannelService) Page(page, parentId int) *orm.Page {
//	return s.FindChild(page, parentId)
//}
//
//func (s *ChannelService) FindChild(page, parentId int) *orm.Page {
//	channels := []*model.Category{}
//	result := s.ORM.Page(orm.ORMPage{
//		PageNum:  page,
//		PageSize: s.PageSize,
//		Result:   &channels,
//		Where:    orm.Where("parent_id = ?", parentId),
//		OrderBy:  "sort desc",
//	})
//	if len(channels) == 0 {
//		result.List = nil
//		return result
//	}
//	for _, c := range channels {
//		rst := s.FindChild(1, int(c.ID))
//		c.Children = rst.List
//	}
//	return result
//}
//
//func (s *ChannelService) Save(channel *model.Category) error {
//	if channel.ID == 0 {
//		return s.ORM.DB.Create(channel).Error
//	}
//	tmp, _ := s.Get(int(channel.ID))
//	channel.Sort = tmp.Sort
//	channel.ParentId = tmp.ParentId
//	return s.ORM.DB.Save(channel).Error
//}
//
//func (s *ChannelService) Delete(id int) error {
//	channel, err := s.Get(id)
//	if err != nil {
//		return err
//	}
//	if channel == nil {
//		return errors.New("channel is not exists")
//	}
//	childs := s.FindChild(1, int(channel.ID))
//	if childs.TotalCount > 0 {
//		return errors.New("this channel has child channel, please delete child channel first")
//	}
//	articles := s.ArticleService.Page(1, ArticlePageQuery{
//		ChannelID: id,
//	})
//	if articles.TotalCount > 0 {
//		return errors.New("find " + strconv.Itoa(articles.TotalCount) + " article in this channel, please delete article first")
//	}
//	model := &model.Channel{}
//	model.ID = uint(id)
//	return s.ORM.DB.Delete(model).Error
//}
//
//func (s *ChannelService) Get(id int) (*model.Channel, error) {
//	model := &model.Channel{}
//	err := s.ORM.DB.Find(model, id).Error
//	return model, err
//}
//
//type ChannelTree struct {
//	Title    string         `json:"title"`
//	Value    string         `json:"value"`
//	Key      string         `json:"key"`
//	Icon     string         `json:"icon"`
//	Children []*ChannelTree `json:"children"`
//	Channel  *model.Channel `json:"channel"`
//}
//
//func (s *ChannelService) Tree() []*ChannelTree {
//	return s.tree(0)
//}
//
//func (s *ChannelService) tree(parentId uint) []*ChannelTree {
//	var channels []*model.Channel
//	s.ORM.DB.Where("parent_id = ?", parentId).Find(&channels)
//	if len(channels) == 0 {
//		return nil
//	}
//	var trees []*ChannelTree
//	for _, v := range channels {
//		child := s.tree(v.ID)
//		tmp := &ChannelTree{
//			Title:    v.Name,
//			Value:    strconv.Itoa(int(v.ID)),
//			Key:      strconv.Itoa(int(v.ID)),
//			Icon:     v.Icon,
//			Children: child,
//			Channel:  v,
//		}
//		trees = append(trees, tmp)
//	}
//	return trees
//}
//
//func (s *ChannelService) GetByTitle(title string) (*model.Channel, error) {
//	model := &model.Channel{}
//	err := s.ORM.DB.Where("title = ?", title).Find(model).Error
//	return model, err
//}
