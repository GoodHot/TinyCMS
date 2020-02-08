package service

import (
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
)

type ChannelService struct {
	PageSize int      `val:"${website.page_size}"`
	ORM      *orm.ORM `ioc:"auto"`
}

func (s *ChannelService) Page(page, parentId int) *orm.Page {
	return s.FindChild(page, parentId)
}

func (s *ChannelService) FindChild(page, parentId int) *orm.Page {
	channels := []*model.Channel{}
	result := s.ORM.Page(orm.ORMPage{
		PageNum:  page,
		PageSize: s.PageSize,
		Result:   &channels,
		Where:    orm.Where("parent_id = ?", parentId),
		OrderBy:  "sort desc",
	})
	if len(channels) == 0 {
		return result
	}
	for _, c := range channels {
		if c.HasChild {
			rst := s.FindChild(1, int(c.ID))
			c.Children = rst.List
		}
	}
	return result
}

func (s *ChannelService) Save(channel *model.Channel) error {
	if channel.ID == 0 {
		return s.ORM.DB.Create(channel).Error
	}
	tmp, _ := s.Get(int(channel.ID))
	channel.Icon = tmp.Icon
	channel.Sort = tmp.Sort
	channel.ParentId = tmp.ParentId
	channel.HasChild = tmp.HasChild
	return s.ORM.DB.Save(channel).Error
}

func (s *ChannelService) Delete(id int) error {
	model := &model.Channel{}
	model.ID = uint(id)
	return s.ORM.DB.Delete(model).Error
}

func (s *ChannelService) Get(id int) (*model.Channel, error) {
	model := &model.Channel{}
	err := s.ORM.DB.Find(model, id).Error
	return model, err
}

//func (s *ChannelService) findChild(channels []*model.Channel) {
//	for _, v := range channels {
//		if !v.HasChild {
//			continue
//		}
//		s.ORM.DB.Where("parent_id = ?", v.ID).Find()
//	}
//}
