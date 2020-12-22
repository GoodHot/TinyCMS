package trait

import "github.com/GoodHot/TinyCMS/core"

type Admin struct {
	BaseORM
	Username string
	Email    string
	Password string
}

type AdminORM interface {
	GetByUsername(username string) (*Admin, *core.Err)
	GetByEmail(email string) (*Admin, *core.Err)
}
