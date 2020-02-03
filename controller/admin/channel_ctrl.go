package admin

import "github.com/GoodHot/TinyCMS/service"

type AdminChannelCtrl struct {
	ChannelService *service.ChannelService `ioc:"auto"`
	Age            string                  `val:"${fff}"`
}
