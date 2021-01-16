package http

import (
	"github.com/GoodHot/TinyCMS/config"
	"github.com/GoodHot/TinyCMS/core"
	"github.com/GoodHot/TinyCMS/orm/trait"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"net/http"
	"strconv"
)

const (
	SuffixJSON = ".json"
	SuffixHTML = ".html"
)

type RespResult struct {
	Code int                    `json:"code"`
	Msg  string                 `json:"msg"`
	Data map[string]interface{} `json:"data"`
}

func NewRespResult(code int, msg string, data map[string]interface{}) *RespResult {
	return &RespResult{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

type Context struct {
	Ctx    echo.Context
	Suffix string
	data   map[string]interface{}
}

func (ctx *Context) Put(key string, val interface{}) {
	if ctx.data == nil {
		ctx.data = make(map[string]interface{})
	}
	ctx.data[key] = val
}

// 获取ip地址
func (*Context) IPAddr() string {
	return ""
}

func (ctx *Context) HTML(html string) error {
	ctx.Ctx.String(200, html)
	return nil
}

func (ctx *Context) Bind(i interface{}) error {
	return ctx.Ctx.Bind(i)
}

func (ctx *Context) JSON() *core.Err {
	result := NewRespResult(int(core.ErrTypeOK.Code), core.ErrTypeOK.Msg, ctx.data)
	err := ctx.Ctx.JSON(http.StatusOK, result)
	if err != nil {
		return core.NewErr(core.Err_Sys_Server)
	}
	return nil
}

func (ctx *Context) ResultOK(html string) *core.Err {
	if ctx.Suffix == SuffixJSON {
		result := NewRespResult(int(core.ErrTypeOK.Code), core.ErrTypeOK.Msg, ctx.data)
		err := ctx.Ctx.JSON(http.StatusOK, result)
		if err != nil {
			return core.NewErr(core.Err_Sys_Server)
		}
	} else {
		err := ctx.Ctx.Render(http.StatusOK, html, ctx.data)
		if err != nil {
			return core.NewErr(core.Err_Render_Fail)
		}
	}
	return nil
}

func (ctx *Context) ResultOKLayout(html string, layout string) *core.Err {
	if ctx.Suffix == SuffixJSON {
		result := NewRespResult(int(core.ErrTypeOK.Code), core.ErrTypeOK.Msg, ctx.data)
		err := ctx.Ctx.JSON(http.StatusOK, result)
		if err != nil {
			return core.NewErr(core.Err_Sys_Server)
		}
	} else {
		ctx.data["template::layout"] = layout
		err := ctx.Ctx.Render(http.StatusOK, html, ctx.data)
		if err != nil {
			return core.NewErr(core.Err_Render_Fail)
		}
	}
	return nil
}

func (ctx *Context) CurrMember() *trait.Member {
	user := ctx.Ctx.Get(config.JWTContextKey).(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return &trait.Member{
		BaseORM: trait.BaseORM{
			ID: int(claims["id"].(float64)),
		},
		Username: claims["username"].(string),
		Role:     trait.RoleType(claims["role"].(float64)),
	}
}

func (ctx *Context) Param(name string) string {
	return ctx.Ctx.Param(name)
}

func (ctx *Context) ParamInt(name string) int {
	rst := ctx.Param(name)
	val, _ := strconv.Atoi(rst)
	return val
}

func (ctx *Context) QueryParam(name string) string {
	return ctx.Ctx.QueryParam(name)
}

func (ctx *Context) QueryParamInt(name string) int {
	rst := ctx.Ctx.QueryParam(name)
	val, _ := strconv.Atoi(rst)
	return val
}

func (ctx *Context) MultipartForm() (*multipart.Form, error) {
	return ctx.Ctx.MultipartForm()
}

type HandlerFunc func(ctx *Context) *core.Err
