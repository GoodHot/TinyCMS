package service

import "github.com/GoodHot/TinyCMS/orm"

type ChannelService struct {
	PageSize int      `val:"${website.page_size}"`
	ORM      *orm.ORM `ioc:"auto"`
}

func (s *ChannelService) List(page int) {

}
