package service

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/GoodHot/TinyCMS/orm/trait"
)

type ChannelService struct {
	ORM *orm.ORMFactory `ioc:"auto"`
}

func (my *ChannelService) Save(channel *trait.Channel) *core.Err {
	return my.ORM.Channel.Save(channel)
}

func (my *ChannelService) GetByPath(path string) (*trait.Channel, *core.Err) {
	return my.ORM.Channel.GetByPath(path)
}

func (my *ChannelService) GetAll() (*[]trait.Channel, *core.Err) {
	// TODO 给channel排序
	return my.ORM.Channel.GetAll()
}
