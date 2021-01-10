package orm

import (
	"errors"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/impl/mongodb"
	"github.com/GoodHot/TinyCMS/orm/impl/mysql"
	"github.com/GoodHot/TinyCMS/orm/impl/sqlite"
	"github.com/GoodHot/TinyCMS/orm/impl/sqlite/datasource"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
)

type ORMFactory struct {
	Config  *core.Config `ioc:"auto"`
	DBType  string       `val:"${props.orm.type}"`
	ShowSQL bool         `val:"${props.orm.showSQL}"`
	Member  trait.MemberORM
	Channel trait.ChannelORM
	Plugin  trait.PluginORM
	Dict    trait.DictORM
	Post    trait.PostORM
}

func (factory *ORMFactory) Startup() error {
	if factory.DBType == "sqlite" {
		return factory.initSqlite()
	} else if factory.DBType == "mysql" {
		factory.Member = &mysql.MemberORMImpl{}
	} else if factory.DBType == "mongodb" {
		factory.Member = &mongodb.AdminORMImpl{}
	} else {
		return errors.New("can not found " + factory.DBType + " orm datasource type")
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
	factory.Member = &sqlite.MemberORMImpl{
		DB: &datasource.DBMemberORM{
			ShowSQL: factory.ShowSQL,
			DB:      db,
		},
	}
	factory.Channel = &sqlite.ChannelORMImpl{DB: db}
	factory.Plugin = &sqlite.PluginORMImpl{DB: db}
	factory.Dict = &sqlite.DictORMImpl{DB: db}
	factory.Post = &sqlite.PostORMImpl{DB: db}
	return nil
}
