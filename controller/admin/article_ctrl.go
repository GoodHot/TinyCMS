package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/service"
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
	//ID             int      `json:"id"`
	//Title          string   `json:"title"`           // 标题
	//SEOTitle       string   `json:"seo_title"`       // SEO标题
	//SEODescription string   `json:"seo_description"` // SEO描述
	//SEOKeyword     string   `json:"seo_keyword"`     // SEO关键字
	//Markdown       string   `json:"markdown"`        // markdown代码
	//Html           string   `json:"html"`            // HTML代码
	//Cover          string   `json:"cover"`           // 封面图
	//CategoryID     int      `json:"category_id"`     // 分类
	//Description    string   `json:"description"`     // 内容描述
	//Tags           []string `json:"tags"`            // 标签
	//Type           string   `json:"type"`            // 操作类型 [publish, draft]
	//Visibility     string   `json:"visibility"`      // 可见性 [public, private]
	//Template       string   `json:"template"`        // 页面模板
	//Author         uint     `json:"author"`          // 作者ID

}

func (s *AdminArticleCtrl) Publish(ctx *ctrl.HTTPContext) error {
	publish := new(PublishForm)
	if err := ctx.Bind(publish); err != nil {
		return ctx.ResultErr(err.Error())
	}
	ctx.Put("publish", publish)
	return ctx.ResultOK()
	//publish.Title = strings.TrimSpace(publish.Title)
	//if publish.Title == "" {
	//	return ctx.ResultErr("请输入文章标题")
	//}
	//var tags []*model.Tag
	//if len(publish.Tags) > 0 {
	//	for _, v := range publish.Tags {
	//		tags = append(tags, &model.Tag{
	//			Name: v,
	//		})
	//	}
	//}
	//status := model.ArticleStatusPublish
	//if publish.Type != "publish" {
	//	status = model.ArticleStatusDraft
	//}
	//article := &model.ArticlePublish{
	//	Article: &model.Article{
	//		Model: orm.Model{
	//			ID: uint(publish.ID),
	//		},
	//		Title:       publish.Title,
	//		SEOTitle:    publish.SEOTitle,
	//		CategoryID:  uint(publish.CategoryID),
	//		Description: publish.Description,
	//		Cover:       publish.Cover,
	//		Status:      status,
	//		AuthorID:    ctx.Admin.ID,
	//	},
	//	Tags: tags,
	//	Content: &model.ArticleContent{
	//		Markdown: publish.Markdown,
	//		Html:     publish.Html,
	//	},
	//}
	//if err := s.ArticleService.Publish(article); err != nil {
	//	return ctx.ResultErr(err.Error())
	//}
	//return ctx.ResultOK()
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
	content := s.ArticleService.GetContent(article.ContentTable, int(article.ContentID))
	//tags := s.TagService.GetByArticleID(id)
	//tagsName := []string{}
	//if tags != nil && len(tags) > 0 {
	//	for _, v := range tags {
	//		tagsName = append(tagsName, v.Name)
	//	}
	//}

	//publish := &PublishForm{
	//	ID:          int(article.ID),
	//	Title:       article.Title,
	//	SEOTitle:    article.SEOTitle,
	//	Markdown:    content.Markdown,
	//	Html:        content.Html,
	//	Cover:       article.Cover,
	//	CategoryID:  int(article.CategoryID),
	//	Description: article.Description,
	//	Tags:        tagsName,
	//}
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
		return ctx.ResultErr(err.Error())
	}
	if err := s.ArticleService.Delete(ids.IDS); err != nil {
		return ctx.ResultErr(err.Error())
	}
	return ctx.ResultOK()
}
