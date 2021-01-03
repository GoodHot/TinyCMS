package sqlite

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type MemberORMImpl struct {
	DB *sqlx.DB
}

const allMemberColumn = "id, nickname, email, password, role"

func (orm *MemberORMImpl) GetByUsername(username string) (*trait.Member, *core.Err) {
	row := orm.DB.QueryRowx(fmt.Sprintf("select %v from t_member where username = ?", allMemberColumn), username)
	if row.Err() != nil {
		return nil, core.NewErr(core.Err_Auth_Not_Username)
	}
	var member trait.Member
	if err := row.StructScan(&member); err != nil {
		return nil, core.NewErr(core.Err_Auth_Not_Username)
	}
	return &member, nil
}

func (orm *MemberORMImpl) GetByEmail(email string) (*trait.Member, *core.Err) {
	var admin trait.Member
	err := orm.DB.QueryRowx("select * from t_member where email = ?", email).StructScan(&admin)
	if err != nil {
		return nil, core.NewErr(core.Err_Auth_Not_Email)
	}
	return &admin, nil
}
