package test

import (
	"testing"

	. "project2/util"
)

func Test_SelectCodeService(t *testing.T) {
	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
	for _, v := range []struct {
		exp string
		ret int
	}{
		{"1 + 2 + 3 + 5", 11},
		{"1 + 2 - 3 + 5", 5},
		{"1 + 2 * 3 - 5", 2},
		{"11 + 2 * ( 4 + 5 )", 29},
		{"1 - 2 * ( 4 + 5 )", -17},
		{"62 + 2 + 3 + 5", 72},
		{"35 + 21 * 33 - 55", 673},
	} {
		// 调用排列组合函数，与期望的结果比对，如果不一致输出错误
		if actually := GetResult(MixToPost(v.exp)); actually != v.ret {
			t.Errorf("combination: [%v], actually: [%v]", v, actually)
		}
	}
}
