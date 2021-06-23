package test

import (
	"testing"

	"project3/service"
)

func Test_SelectCodeService(t *testing.T) {
	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
	for _, v := range []struct {
		code    string
		retCode int
	}{
		{"APi6eK30", 200},
		{"b2rWopwr", 200},
		{"Lhr31KWk", 200},
		{"QMvFVKAI", 200},
		{"It7dRPnm", 200},
		{"sUy1FSTu", 200},
		{"tzg1f1nd", 200},
	} {
		// 调用排列组合函数，与期望的结果比对，如果不一致输出错误
		if actually, _, _ := service.SelectCodeService(v.code); actually != v.retCode {
			t.Errorf("combination: [%v], actually: [%v]", v, actually)
		}
	}
}
