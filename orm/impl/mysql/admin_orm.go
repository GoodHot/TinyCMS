package mysql

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
)

type MemberORMImpl struct {
}

func (a MemberORMImpl) Initial() error {
	panic("implement me")
}

func (a MemberORMImpl) GetByUsername(username string) (*trait.Member, *core.Err) {
	panic("implement me")
}

func (a MemberORMImpl) GetByEmail(email string) (*trait.Member, *core.Err) {
	panic("implement me")
}



