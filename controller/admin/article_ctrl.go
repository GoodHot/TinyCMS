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

type PublishForm struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
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
			CategoryID:  uint(publish.CategoryID),
			Description: publish.Description,
			Cover:       publish.Cover,
			Status:      status,
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
	//if err := ctx.Bind(publish); err != nil {
	//	return ctx.ResultErr(err.Error())
	//}
	//publish.Title = strings.TrimSpace(publish.Title)
	//if publish.Title == "" {
	//	return ctx.ResultErr("Article title not be empty")
	//}
	//if publish.CategoryID == "" {
	//	return ctx.ResultErr("please choose channel")
	//}
	//channelId, err := strconv.Atoi(publish.ChannelID)
	//if err != nil {
	//	return ctx.ResultErr(err.Error())
	//}
	//var publishTime time.Time
	//if publish.PublishTime != "" {
	//	publishTime = times.Parse(publish.PublishTime, "2006-01-02 15:04:05")
	//}
	//var status model.ArticleStatus
	//if publish.Type == "publish" {
	//	status = model.ArticleStatusPublish
	//} else {
	//	status = model.ArticleStatusDraft
	//}
	//var tags []*model.Tag
	//if len(publish.Tags) > 0 {
	//	for _, v := range publish.Tags {
	//		tags = append(tags, &model.Tag{
	//			Name: v,
	//		})
	//	}
	//}
	//articlePublish := &model.ArticlePublish{
	//	Article: &model.Article{
	//		Title:       publish.Title,
	//		ChannelID:   uint(channelId),
	//		Description: publish.Description,
	//		Cover:       publish.Cover,
	//		PublishTime: &publishTime,
	//		Status:      status,
	//	},
	//	Tags: tags,
	//	Content: &model.ArticleContent{
	//		Markdown: publish.Markdown,
	//		Html:     publish.Html,
	//	},
	//}
	//if publish.ID == 0 {
	//	if err = s.ArticleService.Publish(articlePublish); err != nil {
	//		return ctx.ResultErr(err.Error())
	//	}
	//} else {
	//	articlePublish.Article.ID = uint(publish.ID)
	//	s.ArticleService.Edit(articlePublish)
	//}
	//return ctx.ResultOK()
}

func (s *AdminArticleCtrl) Page(ctx *ctrl.HTTPContext) error {
	page := ctx.ParamInt("page")
	ctx.Put("page", s.ArticleService.Page(page, service.ArticlePageQuery{}))
	return ctx.ResultOK()
}

func (s *AdminArticleCtrl) Get(ctx *ctrl.HTTPContext) error {
	//id := ctx.ParamInt("id")
	//article := s.ArticleService.GetById(id)
	//if article == nil || article.Title == "" {
	//	return ctx.ResultErr("article not exists")
	//}
	//tags := s.TagService.GetByArticleID(id)
	//tagsName := []string{}
	//if tags != nil && len(tags) > 0 {
	//	for _, v := range tags {
	//		tagsName = append(tagsName, v.Name)
	//	}
	//}
	//content := s.ArticleService.GetContent(article.ContentTable, int(article.ContentID))
	//publish := &PublishForm{
	//	ID:          int(article.ID),
	//	Title:       article.Title,
	//	Markdown:    content.Markdown,
	//	Html:        content.Html,
	//	Cover:       article.Cover,
	//	ChannelID:   strconv.Itoa(int(article.ChannelID)),
	//	Description: article.Description,
	//	Tags:        tagsName,
	//}
	//ctx.Put("publish", publish)
	return ctx.ResultOK()
}

func (s *AdminArticleCtrl) Delete(ctx *ctrl.HTTPContext) error {
	id := ctx.ParamInt("id")
	if err := s.ArticleService.Delete(id); err != nil {
		return ctx.ResultErr(err.Error())
	}
	return ctx.ResultOK()
}
