package entity

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SetResult(code int, msg string, data interface{}) Result {
	return Result{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
