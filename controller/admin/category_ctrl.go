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

func (s *AdminCategoryCtrl) Page(ctx *ctrl.HTTPContext) error {
	page := ctx.ParamInt("page")
	query := &service.CategoryQuery{
		Keyword: ctx.QueryParam("keyword"),
	}
	ctx.Put("page", s.CategoryService.Page(page, query))
	return ctx.ResultOK()
}

func (s *AdminCategoryCtrl) Sort(ctx *ctrl.HTTPContext) error {
	sort := new(service.CategorySort)
	if err := ctx.Bind(sort); err != nil {
		return ctx.ResultErr(err.Error())
	}
	s.CategoryService.Sort(sort)
	return ctx.ResultOK()
}

