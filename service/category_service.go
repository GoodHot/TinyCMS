package service

import (
	"github.com/GoodHot/TinyCMS/orm"
)

type ChannelCategory struct {
	ArticleService *ArticleService `ioc:"auto"`
	PageSize       int             `val:"${website.page_size}"`
	ORM            *orm.ORM        `ioc:"auto"`
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
