package trait

import "github.com/GoodHot/TinyCMS/core"

type Admin struct {
	BaseORM
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"_"`
}

type AdminORM interface {
	GetByUsername(username string) (*Admin, *core.Err)
	GetByEmail(email string) (*Admin, *core.Err)
}
