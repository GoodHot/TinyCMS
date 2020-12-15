package admin

import (
	"fmt"
	"github.com/GoodHot/TinyCMS/common"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/GoodHot/TinyCMS/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
	"time"
)

type Index struct {
	Cache        *common.Cache         `ioc:"auto"`
	AdminService *service.AdminService `ioc:"auto"`
}

func (my *Index) Index(ctx *http.Context) error {
	go time.Sleep(3 * time.Second)
	ctx.Ctx.JSON(200, gin.H{"hello": "world"})
	return nil
}

func (my *Index) List(ctx *http.Context) error {
	my.Cache.Ins().Get("aaa", "bbb")
	fmt.Println("in index ctx")
	return nil
}

type signinForm struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func (my *Index) Signin(ctx *http.Context) error {
	var signin signinForm
	if err := ctx.Bind(&signin); err != nil {
		return err
	}
	if signin.Account != "x" && signin.Password != "x" {
		return echo.ErrUnauthorized
	}
	claims := &jwtCustomClaims{
		"Jon Snow",
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	ctx.Put("token", t)
	return ctx.JSON()
}
