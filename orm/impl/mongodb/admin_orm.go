package mongodb

import "github.com/GoodHot/TinyCMS/orm/trait"

type AdminORMImpl struct {
}

func (a AdminORMImpl) GetByUsername(username string) (*trait.Admin, error) {
	panic("implement me")
}

func (a AdminORMImpl) GetByEmail(email string) (*trait.Admin, error) {
	panic("implement me")
}

