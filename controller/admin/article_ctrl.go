package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/orm"
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
	}
	ctx.Put("page", s.ArticleService.Page(page, query))
	return ctx.ResultOK()
}

type PublishForm struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	SEOTitle    string   `json:"seo_title"`
	Markdown    string   `json:"markdown"`
	Html        string   `json:"html"`
	Cover       string   `json:"cover"`
	CategoryID  int      `json:"category_id"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Type        string   `json:"type"` // 操作类型 [publish, draft]
}

func (s *AdminArticleCtrl) Publish(ctx *ctrl.HTTPContext) error {
	publish := new(PublishForm)
	if err := ctx.Bind(publish); err != nil {
		return ctx.ResultErr(err.Error())
	}
	publish.Title = strings.TrimSpace(publish.Title)
	if publish.Title == "" {
		return ctx.ResultErr("请输入文章标题")
	}
	var tags []*model.Tag
	if len(publish.Tags) > 0 {
		for _, v := range publish.Tags {
			tags = append(tags, &model.Tag{
				Name: v,
			})
		}
	}
	status := model.ArticleStatusPublish
	if publish.Type != "publish" {
		status = model.ArticleStatusDraft
	}
	article := &model.ArticlePublish{
		Article: &model.Article{
			Model: orm.Model{
				ID: uint(publish.ID),
			},
			Title:       publish.Title,
			SEOTitle:    publish.SEOTitle,
			CategoryID:  uint(publish.CategoryID),
			Description: publish.Description,
			Cover:       publish.Cover,
			Status:      status,
			AuthorID:    ctx.Admin.ID,
		},
		Tags: tags,
		Content: &model.ArticleContent{
			Markdown: publish.Markdown,
			Html:     publish.Html,
		},
	}
	if err := s.ArticleService.Publish(article); err != nil {
		return ctx.ResultErr(err.Error())
	}
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
		return ctx.ResultErr("article not exists")
	}
	tags := s.TagService.GetByArticleID(id)
	tagsName := []string{}
	if tags != nil && len(tags) > 0 {
		for _, v := range tags {
			tagsName = append(tagsName, v.Name)
		}
	}
	content := s.ArticleService.GetContent(article.ContentTable, int(article.ContentID))
	publish := &PublishForm{
		ID:          int(article.ID),
		Title:       article.Title,
		SEOTitle:    article.SEOTitle,
		Markdown:    content.Markdown,
		Html:        content.Html,
		Cover:       article.Cover,
		CategoryID:  int(article.CategoryID),
		Description: article.Description,
		Tags:        tagsName,
	}
	ctx.Put("article", publish)
	return ctx.ResultOK()
}

type DeleteArticle struct {
	IDS []int `json:"ids"`
}

func (s *AdminArticleCtrl) Delete(ctx *ctrl.HTTPContext) error {
	ids := new(DeleteArticle)
	if err := ctx.Bind(ids); err != nil {
		return ctx.ResultErr(err.Error())
	}
	if err := s.ArticleService.Delete(ids.IDS); err != nil {
		return ctx.ResultErr(err.Error())
	}
	return ctx.ResultOK()
}
