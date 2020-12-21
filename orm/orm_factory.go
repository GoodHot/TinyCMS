package orm

import (
	"errors"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/impl/mongodb"
	"github.com/GoodHot/TinyCMS/orm/impl/mysql"
	"github.com/GoodHot/TinyCMS/orm/impl/sqlite"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
)

type ORMFactory struct {
	Config *core.Config `ioc:"auto"`
	DBType string       `val:"${props.orm.type}"`
	Admin  trait.AdminORM
}

func (factory *ORMFactory) Startup() error {
	if factory.DBType == "sqlite" {
		return factory.initSqlite()
	} else if factory.DBType == "mysql" {
		factory.Admin = &mysql.AdminORMImpl{}
	} else if factory.DBType == "mongodb" {
		factory.Admin = &mongodb.AdminORMImpl{}
	} else {
		return errors.New("can not found " + factory.DBType + " orm db type")
	}
	return errors.New("not found database adapter")
}

func (factory *ORMFactory) initSqlite() error {
	dbName, err := factory.Config.GetStr("props.orm.sqlite.db")
	if err != nil {
		return err
	}
	db, err := sqlx.Open("sqlite3", dbName)
	if err != nil {
		return err
	}
	schema, err := ioutil.ReadFile("./config/sqlite/init.sql")
	if err != nil {
		return err
	}
	if _, err := db.Exec(string(schema)); err != nil {
		return err
	}
	factory.Admin = &sqlite.AdminORMImpl{DB: db}
	return nil
}
