package mongodb

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
)

type AdminORMImpl struct {
}

func (a AdminORMImpl) Initial() error {
	panic("implement me")
}

func (a AdminORMImpl) GetByUsername(username string) (*trait.Admin, *core.Err) {
	panic("implement me")
}

func (a AdminORMImpl) GetByEmail(email string) (*trait.Admin, *core.Err) {
	panic("implement me")
}

