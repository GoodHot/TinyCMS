package ctrl

import (
	"github.com/GoodHot/TinyCMS/model"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"mime/multipart"
	"net/http"
	"strconv"
)

type RESULT_STATUS int

type HTTPResult struct {
	Status int                    `json:"status"` // 状态码
	Msg    string                 `json:"msg"`    // 消息
	Data   map[string]interface{} `json:"data"`   // 返回数据
}

type HTTPContext struct {
	Context echo.Context // echo上下文
	Result  *HTTPResult
	Admin   *model.Admin
}

func (s *HTTPContext) FormFile(name string) (*multipart.FileHeader, error) {
	return s.Context.FormFile(name)
}

func (s *HTTPContext) Bind(i interface{}) error {
	return s.Context.Bind(i)
}

func (s *HTTPContext) Param(name string) string {
	return s.Context.Param(name)
}

func (s *HTTPContext) ParamInt(name string) int {
	param := s.Context.Param(name)
	if param == "" {
		return 0
	}
	rst, err := strconv.Atoi(param)
	if err != nil {
		log.Error(err)
		return 0
	}
	return rst
}

func (s *HTTPContext) FormValue(name string) string {
	return s.Context.FormValue(name)
}

func (s *HTTPContext) FormValueInt(name string) int {
	value := s.FormValue(name)
	if value == "" {
		return 0
	}
	rst, err := strconv.Atoi(value)
	if err != nil {
		log.Error(err)
		return 0
	}
	return rst
}

func (s *HTTPContext) QueryParam(name string) string {
	return s.Context.QueryParam(name)
}

func (s *HTTPContext) QueryParamInt(name string) int {
	param := s.QueryParam(name)
	if param == "" {
		return 0
	}
	rst, err := strconv.Atoi(param)
	if err != nil {
		log.Error(err)
		return 0
	}
	return rst
}

func (s *HTTPContext) JsonErr(code int, i interface{}) error {
	return s.Context.JSON(code, i)
}

func (s *HTTPContext) JsonOk(i interface{}) error {
	return s.Context.JSON(http.StatusOK, i)
}

func (s *HTTPContext) Put(key string, value interface{}) {
	s.Result.Data[key] = value
}

func (s *HTTPContext) ResultOK() error {
	s.Result.Status = http.StatusOK
	s.Result.Msg = "ok"
	return s.JsonOk(s.Result)
}

func (s *HTTPContext) ResultOKMsg(msg string) error {
	s.Result.Status = http.StatusOK
	s.Result.Msg = msg
	return s.JsonOk(s.Result)
}

func (s *HTTPContext) ResultErr(msg string) error {
	s.Result.Status = http.StatusInternalServerError
	s.Result.Msg = msg
	return s.JsonOk(s.Result)
}

func (s *HTTPContext) ResultErrStatus(status int, msg string) error {
	s.Result.Status = status
	s.Result.Msg = msg
	return s.JsonErr(status, s.Result)
}

func (s *HTTPContext) HTML(page string) error {
	return s.Context.Render(http.StatusOK, page, s.Result.Data)
}
