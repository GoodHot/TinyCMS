package sqlite

import (
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/impl/sqlite/datasource"
	"github.com/GoodHot/TinyCMS/orm/trait"
	_ "github.com/mattn/go-sqlite3"
)

type MemberORMImpl struct {
	DB *datasource.DBMemberORM
}

func (orm *MemberORMImpl) GetByUsername(username string) (*trait.Member, *core.Err) {
	model, err := orm.DB.Query().UsernameEq(username).ExecQueryOne()
	if err != nil || model == nil {
		return nil, core.NewErr(core.Err_Auth_Not_Username)
	}
	return orm.convert(model), nil
}

func (orm *MemberORMImpl) GetByEmail(email string) (*trait.Member, *core.Err) {
	model, err := orm.DB.Query().EmailEq(email).ExecQueryOne()
	if err != nil || model == nil {
		return nil, core.NewErr(core.Err_Auth_Not_Username)
	}
	return orm.convert(model), nil
}

func (orm *MemberORMImpl) convert(model *datasource.DBMemberModel) *trait.Member {
	return &trait.Member{
		BaseORM: trait.BaseORM{
			ID: model.ID,
		},
		Nickname: model.Nickname,
		Username: model.Username,
		Email:    model.Email,
		Password: model.Password,
		Role:     trait.RoleType(model.Role),
	}
}
