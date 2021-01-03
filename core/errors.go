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
	ErrTypeOK = ErrType{Code: 0, Msg: "ok", RespStatus: http.StatusOK}

	Err_Sys_Server = ErrType{Code: 0x1001, Msg: "Param Bind Error", RespStatus: http.StatusInternalServerError}

	Err_Auth_Account_Fail = ErrType{Code: 0x2001, Msg: "Account Fail", RespStatus: http.StatusOK}
	Err_Auth_Not_Username = ErrType{Code: 0x2002, Msg: "Not Username", RespStatus: http.StatusOK}
	Err_Auth_Not_Email    = ErrType{Code: 0x2003, Msg: "Not Email", RespStatus: http.StatusOK}

	Err_Channel_Save_Fail         = ErrType{Code: 0x3001, Msg: "Save Channel Fail", RespStatus: http.StatusOK}
	Err_Channel_Not_Found_By_Path = ErrType{Code: 0x3002, Msg: "channel not found", RespStatus: http.StatusOK}
	Err_Channel_Can_Not_Get_List  = ErrType{Code: 0x3003, Msg: "can not get list", RespStatus: http.StatusOK}
	Err_Channel_Path_Exists       = ErrType{Code: 0x3004, Msg: "channel path exists", RespStatus: http.StatusOK}
	Err_Channel_Change_Sort_Fail  = ErrType{Code: 0x3005, Msg: "channel path exists", RespStatus: http.StatusOK}

	Err_Plugin_Not_Exists      = ErrType{Code: 0x4001, Msg: "plugin not exists", RespStatus: http.StatusOK}
	Err_Plugin_Exists          = ErrType{Code: 0x4002, Msg: "plugin exists", RespStatus: http.StatusOK}
	Err_Plugin_Save_Fail       = ErrType{Code: 0x4003, Msg: "save plugin fail", RespStatus: http.StatusOK}
	Err_Plugin_Save_Param_Fail = ErrType{Code: 0x4004, Msg: "save plugin param fail", RespStatus: http.StatusOK}
	Err_Plugin_Type_Not_Exists = ErrType{Code: 0x4005, Msg: "plugin type not exists", RespStatus: http.StatusOK}

	Err_Dict_Save_Fail     = ErrType{Code: 0x5001, Msg: "dict not exists", RespStatus: http.StatusOK}
	Err_Dict_Not_Found_Key = ErrType{Code: 0x5001, Msg: "dict not found key", RespStatus: http.StatusOK}
	Err_Dict_Update_Fail   = ErrType{Code: 0x5001, Msg: "dict update fail", RespStatus: http.StatusOK}

	Err_Post_Save_Fail            = ErrType{Code: 0x6001, Msg: "post save fail", RespStatus: http.StatusOK}
	Err_Post_Content_Format_Error = ErrType{Code: 0x6002, Msg: "post content format err", RespStatus: http.StatusOK}
	Err_Post_Get_Page_Fail        = ErrType{Code: 0x6003, Msg: "post content format err", RespStatus: http.StatusOK}
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
