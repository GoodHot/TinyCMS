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
}

type PostORM interface {
	Save(post *Post) *core.Err
}
