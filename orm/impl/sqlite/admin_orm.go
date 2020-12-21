package sqlite

import (
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type AdminORMImpl struct {
	isInitial bool
	DB        *sqlx.DB
}

func (orm *AdminORMImpl) GetByUsername(username string) (*trait.Admin, error) {
	row := orm.DB.QueryRowx("select * from t_admin where username = ?", username)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var admin trait.Admin
	row.StructScan(&admin)
	return &admin, nil
}

func (orm *AdminORMImpl) GetByEmail(email string) (*trait.Admin, error) {
	var admin trait.Admin
	err := orm.DB.QueryRowx("select * from t_admin where email = ?", email).StructScan(&admin)
	if err != nil {
		return nil, err
	}
	return &admin, err
}
