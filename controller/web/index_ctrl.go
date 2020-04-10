package web

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/service"
)

type IndexCtrl struct {
	CategoryService *service.CategoryService `ioc:"auto"`
	ArticleService *service.ArticleService `ioc:"auto"`
}

func (s *IndexCtrl) Index(ctx *ctrl.HTTPContext) error {
	page := s.ArticleService.Page(1, &service.ArticlePageQuery{})
	ctx.Put("parentLink", "")
	ctx.Put("channelId", 1)
	ctx.Put("list", page)
	return ctx.HTML("index")
}

func (s *IndexCtrl) Post(ctx *ctrl.HTTPContext) error {
	id := ctx.ParamInt("id")
	article := s.ArticleService.GetById(id)
	if article == nil || article.Title == "" {
		return ctx.ResultErr("article not exists")
	}
	content := s.ArticleService.GetContent(article.ContentTable, int(article.ContentID))
	ctx.Put("article", article)
	ctx.Put("content", content)
	return ctx.HTML("post")
}

func (s *IndexCtrl) Channel(ctx *ctrl.HTTPContext) error {
	//title := ctx.Param("title")
	//channel, err := s.CategoryService.GetByTitle(title)
	//if err != nil {
	//	return err
	//}
	//if channel == nil {
	//	// TODO redirect 404
	//	return ctx.HTML("404")
	//}
	//page := s.ArticleService.Page(1, service.ArticlePageQuery{
	//	ChannelID: int(channel.ID),
	//})
	//ctx.Put("parentLink", ctx.Param("parent"))
	//ctx.Put("channelId", strconv.Itoa(int(channel.ID)))
	//ctx.Put("list", page)
	return ctx.HTML("channel")
}
