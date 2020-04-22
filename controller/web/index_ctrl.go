package web

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/service"
)

type IndexCtrl struct {
	ArticleService *service.ArticleService `ioc:"auto"`
}

func (s *IndexCtrl) Index(ctx *ctrl.HTTPContext) error {
	page := s.ArticleService.Page(1, &service.ArticlePageQuery{Status: int(model.ArticleStatusPublish), Visible: 1})
	ctx.Put("page", page)
	return ctx.HTML("index")
}

func (s *IndexCtrl) Post(ctx *ctrl.HTTPContext) error {
	sid := ctx.Param("shortId")
	article := s.ArticleService.GetByShortId(sid)
	if article == nil || article.Title == "" {
		return ctx.ResultErrMsg("article not exists")
	}
	content := s.ArticleService.GetContent(article.ContentTable, article.ContentID)
	ctx.Put("article", article)
	ctx.Put("content", content)
	return ctx.HTML("post")
}
