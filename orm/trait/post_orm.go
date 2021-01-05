package trait

import (
	"github.com/GoodHot/TinyCMS/core"
	"time"
)

type PostStatus int

const (
	PostStatus_Publish   PostStatus = 0
	PostStatus_Drafts    PostStatus = 1
	PostStatus_Scheduled PostStatus = 2
)

type Post struct {
	BaseORM
	Title           string       `json:"title"`
	Content         string       `json:"content"`
	ContentHTML     string       `json:"content_html"`
	Excerpt         string       `json:"excerpt"`
	Image           string       `json:"image"`
	URL             string       `json:"url"`
	ChannelID       int          `json:"channel_id"`
	TagsID          string       `json:"tags_id"`
	Visible         VisibleLevel `json:"visible"`
	PublishTime     *time.Time   `json:"publish_time"`
	MetaTitle       string       `json:"meta_title"`
	MetaDescription string       `json:"meta_description"`
	Author          int          `json:"author"`
	Status          PostStatus   `json:"status"`
}

type PostORM interface {
	Save(post *Post) *core.Err
	Page(query *Query) (*Page, *core.Err)
}
