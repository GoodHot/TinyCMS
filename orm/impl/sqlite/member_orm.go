package sqlite

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type MemberORMImpl struct {
	DB *sqlx.DB
}

func (orm *MemberORMImpl) GetByUsername(username string) (*trait.Member, *core.Err) {
	row := orm.DB.QueryRowx("select * from t_member where username = ?", username)
	if row.Err() != nil {
		return nil, core.NewErr(core.Err_Auth_Not_Username)
	}
	var admin trait.Member
	row.StructScan(&admin)
	return &admin, nil
}

func (orm *MemberORMImpl) GetByEmail(email string) (*trait.Member, *core.Err) {
	var admin trait.Member
	err := orm.DB.QueryRowx("select * from t_member where email = ?", email).StructScan(&admin)
	if err != nil {
		return nil, core.NewErr(core.Err_Auth_Not_Email)
	}
	return &admin, nil
}
