package web

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/service"
	"strconv"
)

type IndexCtrl struct {
	ChannelService *service.ChannelService `ioc:"auto"`
	ArticleService *service.ArticleService `ioc:"auto"`
}

func (s *IndexCtrl) Index(ctx *ctrl.HTTPContext) error {
	page := s.ArticleService.Page(1, service.ArticlePageQuery{})
	ctx.Put("list", page)
	return ctx.HTML("index")
}

func (s *IndexCtrl) Channel(ctx *ctrl.HTTPContext) error {
	title := ctx.Param("title")
	channel, err := s.ChannelService.GetByTitle(title)
	if err != nil {
		return err
	}
	if channel == nil {
		// TODO redirect 404
		return ctx.HTML("404")
	}
	page := s.ArticleService.Page(1, service.ArticlePageQuery{
		ChannelID: int(channel.ID),
	})
	ctx.Put("parentLink", ctx.Param("parent"))
	ctx.Put("channelId", strconv.Itoa(int(channel.ID)))
	ctx.Put("list", page)
	return ctx.HTML("channel")
}
