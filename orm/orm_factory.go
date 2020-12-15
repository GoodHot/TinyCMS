package orm

import (
	"errors"
	"github.com/GoodHot/TinyCMS/orm/impl/mysql"
	"github.com/GoodHot/TinyCMS/orm/impl/sqlite"
	"github.com/GoodHot/TinyCMS/orm/trait"
)

type AdminORMFactory struct {
	DBType string `val:"${props.orm.type}"`
	impl   trait.AdminORM
}

func (factory *AdminORMFactory) Startup() error {
	if factory.DBType == "sqlite" {
		factory.impl = &sqlite.AdminORMImpl{}
	} else if factory.DBType == "mysql" {
		factory.impl = &mysql.AdminORMImpl{}
	} else {
		return errors.New("can not found " + factory.DBType + " orm db type")
	}
	return nil
}

func (factory *AdminORMFactory) Get() trait.AdminORM {
	return factory.impl
}
