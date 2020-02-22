package admin

import (
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"strconv"
)

type AdminTagCtrl struct {
}

func (s *AdminTagCtrl) Search(ctx *ctrl.HTTPContext) error {
	prefix := ctx.Param("prefix")
	var result []string
	for i := 0; i < 10; i++ {
		result = append(result, prefix+":"+strconv.Itoa(i))
	}
	ctx.Put("result", result)
	return ctx.ResultOK()
}
