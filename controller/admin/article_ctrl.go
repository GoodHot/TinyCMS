package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/common/times"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/service"
	"strconv"
	"strings"
	"time"
)

type AdminArticleCtrl struct {
	ArticleService *service.ArticleService `ioc:"auto"`
}

type PublishForm struct {
	Title       string   `json:"title"`
	Markdown    string   `json:"markdown"`
	Html        string   `json:"html"`
	Cover       string   `json:"cover"`
	ChannelID   string   `json:"channel_id"`
	Description string   `json:"description"`
	PublishTime string   `json:"publish_time"`
	Tags        []string `json:"tags"`
	Type        string   `json:"type"`
}

func (s *AdminArticleCtrl) Publish(ctx *ctrl.HTTPContext) error {
	publish := new(PublishForm)
	if err := ctx.Bind(publish); err != nil {
		return ctx.ResultErr(err.Error())
	}
	publish.Title = strings.TrimSpace(publish.Title)
	if publish.Title == "" {
		return ctx.ResultErr("Article title not be empty")
	}
	if publish.ChannelID == "" {
		return ctx.ResultErr("please choose channel")
	}
	channelId, err := strconv.Atoi(publish.ChannelID)
	if err != nil {
		return ctx.ResultErr(err.Error())
	}
	var publishTime time.Time
	if publish.PublishTime != "" {
		publishTime = times.Parse(publish.PublishTime, "2006-01-02 15:04:05")
	}
	var status model.ArticleStatus
	if publish.Type == "publish" {
		status = model.ArticleStatusPublish
	} else {
		status = model.ArticleStatusDraft
	}
	var tags []*model.Tag
	if len(publish.Tags) > 0 {
		for _, v := range publish.Tags {
			tags = append(tags, &model.Tag{
				Name: v,
			})
		}
	}
	articlePublish := &model.ArticlePublish{
		Article: &model.Article{
			Title:       publish.Title,
			ChannelID:   uint(channelId),
			Description: publish.Description,
			Cover:       publish.Cover,
			PublishTime: &publishTime,
			Status:      status,
		},
		Tags: tags,
		Content: &model.ArticleContent{
			Markdown: publish.Markdown,
			Html:     publish.Html,
		},
	}
	err = s.ArticleService.Publish(articlePublish)
	if err != nil {
		return ctx.ResultErr(err.Error())
	}
	return ctx.ResultOK()
}

func (s *AdminArticleCtrl) Page(ctx *ctrl.HTTPContext) error {
	return nil
}

func (s *AdminArticleCtrl) Delete(ctx *ctrl.HTTPContext) error {
	return nil
}
