package admin

import (
	"encoding/json"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/GoodHot/TinyCMS/service"
	"strconv"
	"strings"
)

type Post struct {
	PostService *service.PostService `ioc:"auto"`
}

type saveForm struct {
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

func (my *Post) Save(ctx *http.Context) *core.Err {
	var form saveForm
	if err := ctx.Bind(&form); err != nil {
		return core.NewErr(core.Err_Sys_Server)
	}
	content, err := json.Marshal(form.Content)
	if err != nil {
		content = []byte("")
	}
	post := &trait.Post{
		Title:     form.Title,
		Content:   string(content),
		TagsID:    strings.Join(form.Tags, ","),
		ChannelID: form.Channel,
		Visible:   trait.VisiblePublic.GetByInt(form.Visible),
		Author:    ctx.CurrMember().ID,
	}
	my.PostService.Save(post)
	return ctx.ResultOK("ok")
}

func (my *Post) Page(ctx *http.Context) *core.Err {
	var page, size, lastID int
	if rst, err := strconv.Atoi(ctx.Param("page")); err != nil {
		page = 0
	} else {
		page = rst
	}
	if rst, err := strconv.Atoi(ctx.Param("size")); err != nil {
		size = 20
	} else {
		size = rst
	}
	if rst, err := strconv.Atoi(ctx.QueryParam("lastID")); err != nil {
		lastID = 0
	} else {
		lastID = rst
	}
	query := &trait.Query{
		Page:   page,
		Size:   size,
		LastID: lastID,
		Param:  nil,
		Order: map[string]string{
			"id": "desc",
		},
	}
	result, err := my.PostService.Page(query)
	if err != nil {
		return core.NewErr(core.Err_Post_Get_Page_Fail)
	}
	ctx.Put("page", result)
	return ctx.ResultOK("ok")
}
