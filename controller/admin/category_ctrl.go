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
		return ctx.ResultErr(err)
	}
	if err := s.CategoryService.Save(category); err != nil {
		return ctx.ResultErr(err)
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
		return ctx.ResultErr(err)
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

func (s *AdminCategoryCtrl) GetByName(ctx *ctrl.HTTPContext) error {
	name := ctx.Param("name")
	ctx.Put("categorys", s.CategoryService.GetByName(name))
	return ctx.ResultOK()
}

func (s *AdminCategoryCtrl) GetByPath(ctx *ctrl.HTTPContext) error {
	path := ctx.Param("path")
	ctx.Put("categorys", s.CategoryService.GetByPath(path))
	return ctx.ResultOK()
}

func (s *AdminCategoryCtrl) Sort(ctx *ctrl.HTTPContext) error {
	sort := new(service.CategorySort)
	if err := ctx.Bind(sort); err != nil {
		return ctx.ResultErr(err)
	}
	s.CategoryService.Sort(sort)
	return ctx.ResultOK()
}

type CategoryDeleteFrom struct {
	CategoryIds []int `json:"category_ids"`
}

func (s *AdminCategoryCtrl) Delete(ctx *ctrl.HTTPContext) error {
	ids := new(CategoryDeleteFrom)
	if err := ctx.Bind(ids); err != nil {
		return ctx.ResultErr(err)
	}
	if err := s.CategoryService.Delete(ids.CategoryIds); err != nil {
		return ctx.ResultErr(err)
	}
	return ctx.ResultOK()
}

