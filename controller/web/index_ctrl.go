package web

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/service"
)

type IndexCtrl struct {
	ArticleService  *service.ArticleService  `ioc:"auto"`
	TagService      *service.TagService      `ioc:"auto"`
	UserService     *service.UserService     `ioc:"auto"`
	CategoryService *service.CategoryService `ioc:"auto"`
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

func (s *IndexCtrl) Tags(ctx *ctrl.HTTPContext) error {
	tags := s.TagService.All()
	ctx.Put("tags", tags)
	return ctx.HTML("tags")
}

func (s *IndexCtrl) Tag(ctx *ctrl.HTTPContext) error {
	name := ctx.Param("tag")
	tag := s.TagService.GetByName(name)
	page := s.ArticleService.Page(1, &service.ArticlePageQuery{
		TagID: tag.ID,
	})
	ctx.Put("page", page)
	ctx.Put("tag", tag)
	return ctx.HTML("tag")
}

func (s *IndexCtrl) Authors(ctx *ctrl.HTTPContext) error {
	ctx.Put("users", s.UserService.VisibleUser())
	return ctx.HTML("authors")
}

func (s *IndexCtrl) Author(ctx *ctrl.HTTPContext) error {
	username := ctx.Param("username")
	user := s.UserService.GetByUsername(username)
	page := s.ArticleService.Page(1, &service.ArticlePageQuery{
		UserID: user.ID,
	})
	ctx.Put("user", user)
	ctx.Put("page", page)
	return ctx.HTML("author")
}

func (s *IndexCtrl) Categorys(ctx *ctrl.HTTPContext) error {
	ctx.Put("categorys", s.CategoryService.Tree())
	return ctx.HTML("categorys")
}

func (s *IndexCtrl) Category(ctx *ctrl.HTTPContext) error {
	path := ctx.Param("path")
	category := s.CategoryService.GetByPath(path)
	page := s.ArticleService.Page(1, &service.ArticlePageQuery{
		CategoryID: category.ID,
	})
	ctx.Put("category", category)
	ctx.Put("page", page)
	return ctx.HTML("category")
}
