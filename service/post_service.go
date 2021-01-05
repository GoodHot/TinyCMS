package service

import (
	"encoding/json"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/davidscottmills/goeditorjs"
	"strings"
	"time"
)

type PostService struct {
	ORM *orm.ORMFactory `ioc:"auto"`
}

type SavePostForm struct {
	Title       string                 `json:"title"`
	Content     map[string]interface{} `json:"content"`
	AutoImage   bool                   `json:"autoImage"`
	Image       string                 `json:"image"`
	AutoExcerpt bool                   `json:"auto_excerpt"`
	Excerpt     string                 `json:"excerpt"`
	URL         string                 `json:"url"`
	Channel     int                    `json:"channel"`
	Tags        []string               `json:"tags"`
	Visible     int                    `json:"visible"`
	PublishDate string                 `json:"publishDate"`
	PublishTime string                 `json:"publishTime"`
	Meta        struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"meta"`
}

func (ps *PostService) Save(form *SavePostForm, author int) *core.Err {
	content, err := json.Marshal(form.Content)
	if err != nil {
		content = []byte("")
	}
	htmlEngine := goeditorjs.NewHTMLEngine()
	htmlEngine.RegisterBlockHandlers(
		&goeditorjs.HeaderHandler{},
		&goeditorjs.ParagraphHandler{},
		&goeditorjs.ListHandler{},
		&goeditorjs.CodeBoxHandler{},
	)
	html, err := htmlEngine.GenerateHTML(string(content))
	if err != nil {
		return core.NewErr(core.Err_Post_Content_Format_Error)
	}
	if strings.TrimSpace(form.Title) == "" {
		form.Title = "Untitle"
	}
	now := time.Now()
	post := &trait.Post{
		BaseORM: trait.BaseORM{
			Created:  &now,
			Modified: &now,
		},
		Title:           form.Title,
		Content:         string(content),
		ContentHTML:     html,
		Excerpt:         form.Excerpt,
		Image:           form.Image,
		URL:             form.URL,
		ChannelID:       form.Channel,
		TagsID:          strings.Join(form.Tags, ","),
		Visible:         trait.VisiblePublic.GetByInt(form.Visible),
		PublishTime:     nil,
		MetaTitle:       form.Meta.Title,
		MetaDescription: form.Meta.Description,
		Author:          author,
		Status:          trait.PostStatus_Publish,
	}
	return ps.ORM.Post.Save(post)
}

func (ps *PostService) Page(query *trait.Query) (*trait.Page, *core.Err) {
	return ps.ORM.Post.Page(query)
}
