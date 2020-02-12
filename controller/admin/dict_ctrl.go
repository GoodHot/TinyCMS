package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/model"
	"github.com/GoodHot/TinyCMS/service"
)

type AdminDictCtrl struct {
	DictService *service.DictService `ioc:"auto"`
}

func (s *AdminDictCtrl) Init(ctx *ctrl.HTTPContext) error {
	s.DictService.Init()
	return ctx.ResultOK()
}

func (s *AdminDictCtrl) All(ctx *ctrl.HTTPContext) error {
	dicts := s.DictService.All()
	ctx.Put("dicts", dicts)
	return ctx.ResultOK()
}

type EditData struct {
	Dicts []*model.Dict `json:"dicts"`
}

func (s *AdminDictCtrl) Edit(ctx *ctrl.HTTPContext) error {
	data := new(EditData)
	if err := ctx.Bind(data); err != nil {
		return ctx.ResultErr(err.Error())
	}
	for _, v := range data.Dicts {
		s.DictService.Update(v)
	}
	return ctx.ResultOK()
}
