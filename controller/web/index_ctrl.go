package web

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
)

type IndexCtrl struct {
}

func (s *IndexCtrl) Index(ctx *ctrl.HTTPContext) error {
	return ctx.HTML("index")
}
