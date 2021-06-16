package dto

type CalRet struct {
	Expression string `json:"expression"`
	Result int `json:"result"`
}

func GetCalRet(exp string, newRet int) *CalRet {
	return &CalRet{
		Expression: exp,
		Result: newRet,
	}
}
