// Package ret 统一返回结构
package response

import (
	"net/http"
)

const (
	MsgSuccess = "success"
	MsgFail    = "fail"
)

type result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func success1(data interface{}, msg ...string) *result {
	var m = MsgSuccess
	if len(msg) > 0 {
		m = msg[0]
	}
	return &result{
		Code: http.StatusOK,
		Data: data,
		Msg:  m,
	}
}

func fail1(msg ...string) *result {
	var m = MsgFail
	if len(msg) > 0 {
		m = msg[0]
	}
	return &result{
		Code: http.StatusBadRequest,
		Data: "",
		Msg:  m,
	}
}
