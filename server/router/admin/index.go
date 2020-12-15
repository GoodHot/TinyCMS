package admin

import (
	"errors"
	"github.com/GoodHot/TinyCMS/common"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/GoodHot/TinyCMS/service"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Index struct {
	JWTSecret    string                `val:"${props.secret.jwt}"`
	Cache        *common.Cache         `ioc:"auto"`
	AdminService *service.AdminService `ioc:"auto"`
}

type signinForm struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type jwtCustomClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func (my *Index) Signin(ctx *http.Context) error {
	var signin signinForm
	if err := ctx.Bind(&signin); err != nil {
		return err
	}
	admin, err := my.AdminService.GetByUsernameOrEmail(signin.Account)
	if err != nil {
		return err
	}
	if !my.AdminService.CheckPwd(signin.Account, signin.Password) {
		return errors.New("Password fail")
	}
	claims := &jwtCustomClaims{
		admin.ID,
		admin.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(my.JWTSecret))
	if err != nil {
		return err
	}
	ctx.Put("token", t)
	return ctx.JSON()
}
