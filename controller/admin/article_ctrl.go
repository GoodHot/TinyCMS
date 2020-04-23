package admin


import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/service"
	"strings"
)

type AdminArticleCtrl struct {
	ArticleService *service.ArticleService `ioc:"auto"`
	TagService     *service.TagService     `ioc:"auto"`
}

func (s *AdminArticleCtrl) Page(ctx *ctrl.HTTPContext) error {
	page := ctx.ParamInt("page")
	query := &service.ArticlePageQuery{
		Type:       ctx.QueryParamInt("type"),
		CategoryID: ctx.QueryParamInt("category"),
		Keyword:    ctx.QueryParam("keyword"),
		TagID:      ctx.QueryParamInt("tag"),
		UserID:     ctx.QueryParamInt("user"),
	}
	ctx.Put("page", s.ArticleService.Page(page, query))
	return ctx.ResultOK()
}

type PublishForm struct {
	Article        *model.Article        `json:"article"`
	Content        *model.ArticleContent `json:"content"`
	Tags           []*model.Tag          `json:"tags"`
	GetCover       bool                  `json:"get_cover"`
	GetNetImage    bool                  `json:"get_net_image"`
	GetDescription bool                  `json:"get_description"`

}

func (s *AdminArticleCtrl) Publish(ctx *ctrl.HTTPContext) error {
	publish := new(model.ArticlePublish)
	if err := ctx.Bind(publish); err != nil {
		return ctx.ResultErr(err)
	}
	publish.Article.Title = strings.TrimSpace(publish.Article.Title)
	if publish.Article.Title == "" {
		publish.Article.Title = "untitled"
	}
	if err := s.ArticleService.Publish(publish, ctx.User.ID); err != nil {
		return ctx.ResultErr(err)
	}
	ctx.Put("article", publish)
	return ctx.ResultOK()
}

func (s *AdminArticleCtrl) Count(ctx *ctrl.HTTPContext) error {
	ctx.Put("category", s.ArticleService.NoCategoryCount())
	ctx.Put("draft", s.ArticleService.DraftCount())
	ctx.Put("total", s.ArticleService.TotalCount())
	return ctx.ResultOK()
}

func (s *AdminArticleCtrl) Get(ctx *ctrl.HTTPContext) error {
	id := ctx.ParamInt("id")
	article := s.ArticleService.GetById(id)
	if article == nil || article.Title == "" {
		return ctx.ResultErrMsg("article not exists")
	}
	content := s.ArticleService.GetContent(article.ContentTable, int(article.ContentID))
	ctx.Put("article", article)
	ctx.Put("content", content)
	return ctx.ResultOK()
}

type DeleteArticle struct {
	IDS []int `json:"ids"`
}

func (s *AdminArticleCtrl) Delete(ctx *ctrl.HTTPContext) error {
	ids := new(DeleteArticle)
	if err := ctx.Bind(ids); err != nil {
		return ctx.ResultErr(err)
	}
	if err := s.ArticleService.Delete(ids.IDS); err != nil {
		return ctx.ResultErr(err)
	}
	return ctx.ResultOK()
}

