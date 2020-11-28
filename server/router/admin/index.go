package admin

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/common"
	"github.com/GoodHot/TinyCMS/server/router/http"
)

type Index struct {
	Cache *common.Cache `ioc:"auto"`
}

func (my *Index) Index(ctx *http.Context) error {
	my.Cache.Get().Get("aaa", "bbb")
	fmt.Println("in index ctx")
	return nil
}