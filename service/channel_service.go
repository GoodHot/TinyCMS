package service

import (
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
)

type ChannelService struct {
	PageSize int      `val:"${website.page_size}"`
	ORM      *orm.ORM `ioc:"auto"`
}

func (s *ChannelService) Page(page int) *orm.Page {
	result := s.ORM.Page(orm.ORMPage{
		PageNum:  page,
		PageSize: s.PageSize,
		Result:   &[]model.Channel{},
	})
	return result
}
