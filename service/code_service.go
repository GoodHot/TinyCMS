package service

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm"
	"github.com/GoodHot/TinyCMS/orm/trait"
)

type CodeService struct {
	ORM *orm.ORMFactory `ioc:"auto"`
}

func (cs *CodeService) GetAll() ([]trait.Code, *core.Err) {
	return cs.ORM.Code.GetAll()
}
