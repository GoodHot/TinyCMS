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
