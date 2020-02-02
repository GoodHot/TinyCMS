package admin

import "github.com/GoodHot/TinyCMS/service"

type AdminChannelCtrl struct {
	channelService *service.ChannelService `ioc:"auto"`
}
