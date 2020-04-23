package controller

import (
	"github.com/GoodHot/TinyCMS/brain"
	"github.com/GoodHot/TinyCMS/common/ctrl"
	"github.com/GoodHot/TinyCMS/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type RouterFunc func(ctx *ctrl.HTTPContext) error

type RouterRegister struct {
	Log          brain.Logger
	Group        *echo.Group
	Prefix       string
	BuildOption  bool
	AdminService *service.UserService
}

func (s *RouterRegister) OPTIONS(path string) {
	s.Group.OPTIONS(path, func(e echo.Context) error {
		return e.JSON(http.StatusOK, "ok")
	})
}

func (s *RouterRegister) GET(path string, fn RouterFunc) {
	if s.BuildOption {
		s.OPTIONS(path)
	}
	s.Group.GET(path, s.buildHandlerFunc(fn))
	s.Log.Info("Register Router[GET], Path: %v%v", s.Prefix, path)
}

func (s *RouterRegister) POST(path string, fn RouterFunc) {
	if s.BuildOption {
		s.OPTIONS(path)
	}
	s.Group.POST(path, s.buildHandlerFunc(fn))
	s.Log.Info("Register Router[POST], Path: %v%v", s.Prefix, path)
}

func (s *RouterRegister) PUT(path string, fn RouterFunc) {
	if s.BuildOption {
		s.OPTIONS(path)
	}
	s.Group.PUT(path, s.buildHandlerFunc(fn))
	s.Log.Info("Register Router[PUT], Path: %v%v", s.Prefix, path)
}

func (s *RouterRegister) DELETE(path string, fn RouterFunc) {
	if s.BuildOption {
		s.OPTIONS(path)
	}
	s.Group.DELETE(path, s.buildHandlerFunc(fn))
	s.Log.Info("Register Router[DELETE], Path: %v%v", s.Prefix, path)
}

func (s *RouterRegister) Any(path string, fn RouterFunc) {
	s.Group.Any(path, s.buildHandlerFunc(fn))
	s.Log.Info("Register Router[Any], Path: %v%v", s.Prefix, path)
}

func (s *RouterRegister) buildHandlerFunc(fn RouterFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		ctx := new(ctrl.HTTPContext)
		ctx.Context = e
		ctx.Result = &ctrl.HTTPResult{Data: make(map[string]interface{})}
		token := e.Request().Header.Get("ACCESS-TOKEN")
		if token != "" {
			user, _ := s.AdminService.CheckToken(token)
			if user != nil {
				ctx.User = user.User
			}
		}
		return fn(ctx)
	}
}
