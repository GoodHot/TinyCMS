package core

import (
	"fmt"
	"net/http"
)

type ErrType struct {
	Code       uint16
	Msg        string
	RespStatus int
}

var (
	ErrTypeOK             = ErrType{Code: 0, Msg: "ok", RespStatus: http.StatusOK}
	Err_Sys_Server        = ErrType{Code: 0x1001, Msg: "Param Bind Error", RespStatus: http.StatusInternalServerError}
	Err_Auth_Account_Fail = ErrType{Code: 0x2001, Msg: "Account Fail", RespStatus: http.StatusOK}
	Err_Auth_Not_Username = ErrType{Code: 0x2002, Msg: "Not Username", RespStatus: http.StatusOK}
	Err_Auth_Not_Email    = ErrType{Code: 0x2003, Msg: "Not Email", RespStatus: http.StatusOK}
	Err_Channel_Save_fail = ErrType{Code: 0x3001, Msg: "Save Channel Fail", RespStatus: http.StatusOK}
)

type Err struct {
	ErrType ErrType
}

func (e *Err) Error() string {
	return fmt.Sprintf("error message: %v, error code: %v", e.ErrType.Msg, e.ErrType.Code)
}

func NewErr(et ErrType) *Err {
	return &Err{
		ErrType: et,
	}
}
