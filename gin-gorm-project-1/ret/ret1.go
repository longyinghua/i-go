package ret

import "net/http"

const (
	MsgSuccess1 = "success"
	MsgFail1    = "fail"
)

type Result1 struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func Success1(msg ...string) *Result1 {
	var m = MsgSuccess1
	if len(msg) > 0 {
		m = msg[0]
	}
	return &Result1{
		Code: http.StatusOK,
		Msg:  m,
	}
}

func Fail1(msg ...string) *Result1 {
	var m = MsgFail1
	if len(msg) > 0 {
		m = msg[0]
	}
	return &Result1{
		Code: http.StatusBadRequest,
		Msg:  m,
	}
}
