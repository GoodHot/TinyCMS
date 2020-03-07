package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/service"
)

type AdminCategoryCtrl struct {
	CategoryService *service.CategoryService `ioc:"auto"`
}

func (s *AdminCategoryCtrl) Save(ctx *ctrl.HTTPContext) error {
	category := new(model.Category)
	if err := ctx.Bind(category); err != nil {
		return ctx.ResultErr(err.Error())
	}
	if err := s.CategoryService.Save(category); err != nil {
		return ctx.ResultErr(err.Error())
	}
	return ctx.ResultOK()
}

func (s *AdminCategoryCtrl) Tree(ctx *ctrl.HTTPContext) error {
	ctx.Put("tree", s.CategoryService.Tree())
	return ctx.ResultOK()
}

func (s *AdminCategoryCtrl) Get(ctx *ctrl.HTTPContext) error {
	id := ctx.ParamInt("id")
	category, err := s.CategoryService.Get(id)
	if err != nil {
		return ctx.ResultErr(err.Error())
	}
	ctx.Put("category", category)
	return ctx.ResultOK()
}

//func (s *AdminChannelCtrl) Page(ctx *ctrl.HTTPContext) error {
//	parentId := ctx.QueryParamInt("parentId")
//	page := ctx.ParamInt("page")
//	ctx.Put("page", s.ChannelService.Page(page, parentId))
//	return ctx.ResultOK()
//}
//
//func (s *AdminChannelCtrl) Tree(ctx *ctrl.HTTPContext) error {
//	ctx.Put("tree", s.ChannelService.Tree())
//	return ctx.ResultOK()
//}
//
//func (s *AdminChannelCtrl) Save(ctx *ctrl.HTTPContext) error {
//	channel := new(model.Category)
//	if err := ctx.Bind(channel); err != nil {
//		return ctx.ResultErr(err.Error())
//	}
//	if err := s.ChannelService.Save(channel); err != nil {
//		return ctx.ResultErr(err.Error())
//	}
//	return ctx.ResultOK()
//}
//
//func (s *AdminChannelCtrl) Delete(ctx *ctrl.HTTPContext) error {
//	id := ctx.ParamInt("id")
//	if err := s.ChannelService.Delete(id); err != nil {
//		return ctx.ResultErr(err.Error())
//	}
//	return ctx.ResultOK()
//}
//
//func (s *AdminChannelCtrl) Get(ctx *ctrl.HTTPContext) error {
//	id := ctx.ParamInt("id")
//	channel, err := s.ChannelService.Get(id)
//	if err != nil {
//		return ctx.ResultErr(err.Error())
//	}
//	ctx.Put("channel", channel)
//	return ctx.ResultOK()
//}
