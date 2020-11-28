package admin

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/common"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"time"
)

type Index struct {
	Cache *common.Cache `ioc:"auto"`
}

func (my *Index) Index(ctx *http.Context) error {
	time.Sleep(3 * time.Second)
	return nil
}

func (my *Index) List(ctx *http.Context) error {
	my.Cache.Ins().Get("aaa", "bbb")
	fmt.Println("in index ctx")
	return nil
}
