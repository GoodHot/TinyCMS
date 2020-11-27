package admin

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/server/router/http"
)

type Index struct {
}

func (*Index) Index(ctx *http.Context) error {
	fmt.Println("in index ctx")
	return nil
}
