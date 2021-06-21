package test

import (
	"testing"

	"project3/service"
)

func Test_SelectCodeService(t *testing.T) {
	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
	for _, v := range []struct {
		code string
		ero  error
	}{
		{"APi6eK30", nil},
		{"b2rWopwr", nil},
		{"Lhr31KWk", nil},
		{"QMvFVKAI", nil},
		{"It7dRPnm", nil},
		{"sUy1FSTu", nil},
		{"tzg1f1nd", nil},
	} {
		// 调用排列组合函数，与期望的结果比对，如果不一致输出错误
		if _, actually := service.SelectCodeService(v.code); actually != v.ero {
			t.Errorf("combination: [%v], actually: [%v]", v, actually)
		}
	}
}
