package trait

import "github.com/GoodHot/TinyCMS/core"

type RoleType int

const (
	RoleTypeOwner  RoleType = 1
	RoleTypeAdmin  RoleType = 2
	RoleTypeEditor RoleType = 3
	RoleTypeAuthor RoleType = 4
)

type Admin struct {
	BaseORM
	Nickname string   `json:"nickname"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"-"`
	Role     RoleType `json:"role"`
}

type AdminORM interface {
	GetByUsername(username string) (*Admin, *core.Err)
	GetByEmail(email string) (*Admin, *core.Err)
}
