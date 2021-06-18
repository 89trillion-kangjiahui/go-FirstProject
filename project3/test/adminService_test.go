package test

import (
	"testing"

	"project3/service"
)

// 单元测试
// 测试类成员函数，以TestClass_Function命名
func Test_CreateCodeService(t *testing.T) {
	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
	var usermap = make(map[string]string)
	usermap["1"] = "管理员"
	usermap["2"] = "张三"
	usermap["3"] = "李四"
	usermap["4"] = "王五"
	usermap["5"] = "赵六"
	for _, v := range []struct {
		uid        string
		codeType   string
		describe   string
		receiveNum string
		usefulDate string
		jewel      string
		gold       string
		props      string
		hero       string
		batman     string
		userMap    map[string]string
		ero        error
	}{
		{"1", "1", "很牛逼",
			"1", "2021-06-25 15:04:05", "7", "8", "11",
			"23", "2", usermap, nil},
		{"1", "2", "很牛逼sad",
			"3", "2021-06-25 18:04:05", "3", "5", "1",
			"2", "5", usermap, nil},
		{"1", "2", "很牛逼salad",
			"1", "2021-06-25 17:04:05", "7", "8", "11",
			"23", "1", usermap, nil},
		{"1", "2", "很牛逼我企鹅",
			"11", "2021-06-25 11:04:05", "0", "8", "11",
			"23", "2", usermap, nil},
		{"1", "3", "很牛逼我企鹅去玩",
			"无限", "2021-06-25 14:04:05", "3", "8", "11",
			"23", "2", usermap, nil},
		{"1", "2", "很牛逼深爱的",
			"4", "2021-06-25 13:04:05", "5", "8", "11",
			"23", "1", usermap, nil},
		{"1", "3", "很牛逼深爱的啊",
			"无限", "2021-06-25 16:04:05", "9", "8", "11",
			"23", "2", usermap, nil},
	} {
		// 调用排列组合函数，与期望的结果比对，如果不一致输出错误
		if _, actually := service.CreateCodeService(v.uid, v.codeType, v.describe, v.receiveNum, v.usefulDate,
			v.jewel, v.gold, v.props, v.hero, v.batman, v.userMap); actually != v.ero {
			t.Errorf("combination: [%v], actually: [%v]", v, actually)
		}
	}
}

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
