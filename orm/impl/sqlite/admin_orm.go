package sqlite

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
)

type AdminORMImpl struct {
	isInitial bool
	db        *sqlx.DB
	Cfg       *core.Config
}

func (orm *AdminORMImpl) Initial() error {
	if orm.isInitial {
		return nil
	}
	dbName, err := orm.Cfg.GetStr("props.orm.sqlite.db")
	if err != nil {
		return err
	}
	orm.db, err = sqlx.Open("sqlite3", dbName)
	if err != nil {
		return err
	}
	schema, err := ioutil.ReadFile("./config/sqlite/init.sql")
	if err != nil {
		return err
	}
	if _, err := orm.db.Exec(string(schema)); err != nil {
		return err
	}
	return nil
}

func (orm *AdminORMImpl) GetByUsername(username string) (*trait.Admin, error) {
	row := orm.db.QueryRowx("select * from t_admin where username = ?", username)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var admin trait.Admin
	row.StructScan(&admin)
	return &admin, nil
}

func (orm *AdminORMImpl) GetByEmail(email string) (*trait.Admin, error) {
	var admin trait.Admin
	err := orm.db.QueryRowx("select * from t_admin where email = ?", email).StructScan(&admin)
	if err != nil {
		return nil, err
	}
	return &admin, err
}
