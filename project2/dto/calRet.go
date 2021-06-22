package dto

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data int    `json:"data"`
}

func SetResult(code, data int, msg string) Result {
	return Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
