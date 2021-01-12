package admin

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/config"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/GoodHot/TinyCMS/service"
	"strconv"
)

type Upload struct {
	DictService *service.DictService `ioc:"auto"`
}

func (my *Upload) Update(ctx *http.Context) *core.Err {
	form, err := ctx.MultipartForm()
	if err != nil {
		return core.NewErr(core.Err_Upload_Fail)
	}
	dict, err1 := my.DictService.Get(config.DictKeyUploadFileMaxSize)
	if err != nil {
		return err1
	}
	maxSize, err := strconv.Atoi(dict.Value)
	if err != nil {
		return core.NewErr(core.Err_Upload_Fail)
	}
	maxSize = maxSize * 1024 * 1024
	fmt.Println(maxSize)
	files := form.File["files"]
	for _, file := range files {
		fmt.Println(file.Filename)
		fmt.Println(file.Size)
		fmt.Println(file.Header)
	}
	return ctx.ResultOK("ok")
}
