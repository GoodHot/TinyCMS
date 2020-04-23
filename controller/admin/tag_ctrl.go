package admin


import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
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

func (s *AdminTagCtrl) Page(ctx *ctrl.HTTPContext) error {
	page := ctx.ParamInt("page")
	query := &service.TagQuery{Keyword: ctx.QueryParam("keyword")}
	ctx.Put("page", s.TagService.Page(page, query))
	return ctx.ResultOK()
}

type TagDeleteFrom struct {
	TagIds []int `json:"tag_ids"`
}

func (s *AdminTagCtrl) Delete(ctx *ctrl.HTTPContext) error {
	ids := new(TagDeleteFrom)
	if err := ctx.Bind(ids); err != nil {
		return ctx.ResultErr(err)
	}
	if err := s.TagService.Delete(ids.TagIds); err != nil {
		return ctx.ResultErr(err)
	}
	return ctx.ResultOK()
}

func (s *AdminTagCtrl) Get(ctx *ctrl.HTTPContext) error {
	id := ctx.ParamInt("id")
	tag, err := s.TagService.Get(id)
	if err != nil {
		return ctx.ResultErr(err)
	}
	ctx.Put("tag", tag)
	return ctx.ResultOK()
}

func (s *AdminTagCtrl) Save(ctx *ctrl.HTTPContext) error {
	tag := new(model.Tag)
	if err := ctx.Bind(tag); err != nil {
		return ctx.ResultErr(err)
	}
	if err := s.TagService.Save(tag); err != nil {
		return ctx.ResultErr(err)
	}
	return ctx.ResultOK()
}
