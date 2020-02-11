package web

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/service"
)

type IndexCtrl struct {
	ChannelService *service.ChannelService `ioc:"auto"`
	ArticleService *service.ArticleService `ioc:"auto"`
}

func (s *IndexCtrl) Index(ctx *ctrl.HTTPContext) error {
	channels := s.ChannelService.Tree()
	page := s.ArticleService.Page(1, service.ArticlePageQuery{})
	ctx.Put("channels", channels)
	ctx.Put("list", page)
	ctx.Put("showSection", false)
	return ctx.HTML("index")
}
