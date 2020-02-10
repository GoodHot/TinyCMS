package admin

import "github.com/GoodHot/TinyCMS/common/ctrl"

type AdminArticleCtrl struct {
}

type PublishForm struct {
	Title       string
	Markdown    string
	Html        string
	Cover       string
	ChannelID   uint
	Description string
	PublishTime string
	Tags        []string
	Status      uint
}

func (s *AdminArticleCtrl) Publish(ctx *ctrl.HTTPContext) error {
	return ctx.ResultOK()
}
