package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/service"
)

type AdminTagCtrl struct {
	TagService *service.TagService `ioc:"auto"`
}

func (s *AdminTagCtrl) Search(ctx *ctrl.HTTPContext) error {
	prefix := ctx.Param("prefix")
	ctx.Put("result", s.TagService.PrefixFind(prefix))
	return ctx.ResultOK()
}

func (s *AdminTagCtrl) HotTag(ctx *ctrl.HTTPContext) error {
	tags := s.TagService.GetHotTag(10)
	ctx.Put("tags", tags)
	return ctx.ResultOK()
}
