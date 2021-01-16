package trait

import "github.com/GoodHot/TinyCMS/core"

type RoleType int

const (
	RoleTypeOwner  RoleType = 1
	RoleTypeAdmin  RoleType = 2
	RoleTypeEditor RoleType = 3
	RoleTypeAuthor RoleType = 4
	RoleTypeMember RoleType = 5
)

type Member struct {
	BaseORM
	Nickname string   `json:"nickname"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"-"`
	Role     RoleType `json:"role"`
	Avatar   string   `json:"avatar"`
}

type MemberORM interface {
	GetByUsername(username string) (*Member, *core.Err)
	GetByEmail(email string) (*Member, *core.Err)
	Page(query *Query) (*Page, *core.Err)
}
