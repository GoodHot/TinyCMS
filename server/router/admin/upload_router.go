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

func (my *Upload) Upload(ctx *http.Context) *core.Err {
	form, err := ctx.MultipartForm()
	if err != nil {
		return core.NewErr(core.Err_Upload_Fail)
	}
	dict, err1 := my.DictService.Get(config.DictKeyUploadFileMaxSize)
	if err1 != nil {
		return err1
	}
	maxSize, err := strconv.Atoi(dict.Value)
	if err != nil {
		return core.NewErr(core.Err_Upload_Fail)
	}
	maxSize = maxSize * 1024 * 1024
	files := form.File["files"]
	for _, file := range files {
		fmt.Println(file.Filename)
		fmt.Println(file.Size)
		fmt.Println(file.Header)
	}
	return ctx.ResultOK("ok")
}

func (my *Upload) EditorUpload(ctx *http.Context) *core.Err {
	ctx.Put("url", "http://wx2.sinaimg.cn/mw600/00893JKXgy1gmwz73ad31j30g40rctgy.jpg")
	return ctx.ResultOK("ok")
}
