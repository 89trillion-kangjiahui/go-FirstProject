package test

import (
	"testing"

	"project3/service"
)

func Test_CheckCodeService(t *testing.T) {
	var usermap = make(map[string]string)
	usermap["1"] = "管理员"
	usermap["2"] = "张三"
	usermap["3"] = "李四"
	usermap["4"] = "王五"
	usermap["5"] = "赵六"
	usermap["6"] = "赵七"
	usermap["7"] = "赵八"
	usermap["8"] = "赵九"

	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
	for _, v := range []struct {
		uid     string
		code    string
		userMap map[string]string
		ero     int
	}{
		{"2", "APi6eK30", usermap, 200},
		{"6", "b2rWopwr", usermap, 200},
		{"2", "Lhr31KWk", usermap, 200},
		{"4", "QMvFVKAI", usermap, 200},
		{"5", "It7dRPnm", usermap, 200},
		{"7", "sUy1FSTu", usermap, 200},
		{"8", "tzg1f1nd", usermap, 200},
	} {
		// 调用排列组合函数，与期望的结果比对，如果不一致输出错误
		if actually, _, _ := service.CheckCodeService(v.uid, v.code, v.userMap); actually != v.ero {
			t.Errorf("combination: [%v], actually: [%v]", v, actually)
		}
	}
}
