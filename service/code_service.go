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

func (cs *CodeService) Save(t *trait.Code) *core.Err {
	if t.ID == 0 {
		return cs.ORM.Code.Insert(t)
	} else {
		return cs.ORM.Code.Update(t)
	}
}

func (cs *CodeService) DeleteByID(id int) *core.Err {
	return cs.ORM.Code.DeleteByID(id)
}
