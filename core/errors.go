package core

import "fmt"

type ErrType struct {
	Code uint16
	Msg  string
}

var (
	Err_Sys_Server        = ErrType{Code: 0x1001, Msg: "Param Bind Error"}
	Err_Auth              = ErrType{Code: 0x2001, Msg: "Not Auth"}
	Err_Auth_Not_Username = ErrType{Code: 0x2002, Msg: "Not Username"}
	Err_Auth_Not_Email    = ErrType{Code: 0x2003, Msg: "Not Email"}
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
