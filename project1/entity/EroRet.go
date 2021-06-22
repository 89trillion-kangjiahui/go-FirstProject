package entity

type EroRet struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SetEroRet(code int, msg string, data interface{}) *EroRet {
	return &EroRet{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
