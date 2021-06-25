package test

import (
	"testing"

	"project1/global"
	"project1/service"
	"project1/util"
)

func Test_GetAckByIdService(t *testing.T) {
	global.SoldierMap = util.JsonToFile("./config/config.army.model.json")
	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
	for _, v := range []struct {
		id         string
		ret        int
	}{
		{"10101", 200},
		{"10102",  200},
		{"10103",  200},
		{"10104",  200},
	} {
		// 调用排列组合函数，与期望的结果比对，如果不一致输出错误
		if actually,_,_ := service.GetAckById(v.id); actually != v.ret {
			t.Errorf("combination: [%v], actually: [%v]", v, actually)
		}
	}
}

func Test_GetRarityById(t *testing.T) {
	global.SoldierMap = util.JsonToFile("./config/config.army.model.json")
	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
	for _, v := range []struct {
		id         string
		ret        int
	}{
		{"10101", 200},
		{"10102",  200},
		{"10103",  200},
		{"10104",  200},
	} {
		// 调用排列组合函数，与期望的结果比对，如果不一致输出错误
		if actually,_,_  := service.GetRarityById(v.id); actually != v.ret {
			t.Errorf("combination: [%v], actually: [%v]", v, actually)
		}
	}
}

func Test_GetAllByRarityService(t *testing.T) {
	global.SoldierMap = util.JsonToFile("../config/config.army.model.json")
	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
	for _, v := range []struct {
		rarity      string
		unlockArena string
		cvc         string
		ret         int
	}{
		{"1", "1", "1000",  200},
		{"2", "2", "1000", 200},
		{"3", "2", "1000",200},
		{"4", "1","1000",  200},
	} {
		// 调用排列组合函数，与期望的结果比对，如果不一致输出错误
		if actually,_,_ := service.GetAllByRarity(v.rarity, v.unlockArena, v.cvc); actually != v.ret {
			t.Errorf("combination: [%v], actually: [%v]", v, actually)
		}
	}
}

func Test_GetSoldierByUnlockArena(t *testing.T) {
	global.SoldierMap = util.JsonToFile("../config/config.army.model.json")
	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
	for _, v := range []struct {
		ret        int
	}{
		{200},
		{200},
		{200},
		{200},
	} {
		// 调用排列组合函数，与期望的结果比对，如果不一致输出错误
		if actually,_,_ := service.GetSoldierByUnlockArena(); actually != v.ret {
			t.Errorf("combination: [%v], actually: [%v]", v, actually)
		}
	}
}

func Test_GetSoldierByCvc(t *testing.T) {
	global.SoldierMap = util.JsonToFile("../config/config.army.model.json")
	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
	for _, v := range []struct {
		cvc        string
		ret        int
	}{
		{"1000", 200},
		{"1500", 200},
		{"2000", 200},
		{"1100", 200},
	} {
		// 调用排列组合函数，与期望的结果比对，如果不一致输出错误
		if actually,_,_ := service.GetSoldierByCvc(v.cvc); actually != v.ret {
			t.Errorf("combination: [%v], actually: [%v]", v, actually)
		}
	}
}
