package admin

import "github.com/GoodHot/TinyCMS/service"

type Dict struct {
	DictService *service.DictService `ioc:"auto"`
}
