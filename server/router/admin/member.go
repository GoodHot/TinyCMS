package admin

import (
	"github.com/GoodHot/TinyCMS/common"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/server/router/http"
	"github.com/GoodHot/TinyCMS/service"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Member struct {
	JWTSecret     string                 `val:"${props.secret.jwt}"`
	Cache         *common.Cache          `ioc:"auto"`
	MemberService *service.MemberService `ioc:"auto"`
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

func (my *Member) Signin(ctx *http.Context) *core.Err {
	var signin signinForm
	if err := ctx.Bind(&signin); err != nil {
		return core.NewErr(core.Err_Sys_Server)
	}
	member, err := my.MemberService.GetByUsernameOrEmail(signin.Account)
	if err != nil {
		return err
	}
	if member.ID == 0 {
		return core.NewErr(core.Err_Auth_Account_Fail)
	}
	if !my.MemberService.CheckPwd(member.Password, signin.Password) {
		return core.NewErr(core.Err_Auth_Account_Fail)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = member.ID
	claims["username"] = member.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, e := token.SignedString([]byte(my.JWTSecret))
	if e != nil {
		return err
	}
	ctx.Put("token", t)
	return ctx.JSON()
}

func (my *Member) Info(ctx *http.Context) *core.Err {
	full := ctx.QueryParam("full")
	if full != "" {
		// TODO 查询Admin完整信息
		ctx.Put("type", "full")
	} else {
		ctx.Put("info", ctx.CurrMember())
		ctx.Put("type", "simple")
	}
	return ctx.ResultOK("ok")
}
