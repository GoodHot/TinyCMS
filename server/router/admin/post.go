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
	Tags        []string               `json:"tags"`
	Channel     string                 `json:"channel"`
	Visible     string                 `json:"visible"`
	PublishDate string                 `json:"publishDate"`
	PublishTime string                 `json:"publishTime"`
}

func (my *Post) Save(ctx *http.Context) *core.Err {
	var form saveForm
	if err := ctx.Bind(&form); err != nil {
		return core.NewErr(core.Err_Sys_Server)
	}
	cid, err := strconv.Atoi(form.Channel)
	if err != nil {
		cid = 0
	}
	visible, err := strconv.Atoi(form.Visible)
	if err != nil {
		visible = 0
	}
	content, err := json.Marshal(form.Content)
	if err != nil {
		content = []byte("")
	}
	post := &trait.Post{
		Title:     form.Title,
		Content:   string(content),
		TagsID:    strings.Join(form.Tags, ","),
		ChannelID: cid,
		Visible: trait.VisiblePublic.GetByInt(visible),
	}
	my.PostService.Save(post)
	return ctx.ResultOK("ok")
}
