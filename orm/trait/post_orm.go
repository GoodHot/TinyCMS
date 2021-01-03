package trait

import "github.com/GoodHot/TinyCMS/core"

type Post struct {
	BaseORM
	Title       string       `json:"title"`
	Content     string       `json:"content"`
	ContentHTML string       `json:"content_html"`
	TagsID      string       `json:"tags_id"`
	ChannelID   int          `json:"channel_id"`
	Visible     VisibleLevel `json:"visible"`
	Author      int          `json:"author"`
}

type PostORM interface {
	Save(post *Post) *core.Err
	Page(query *Query) (*Page, *core.Err)
}
