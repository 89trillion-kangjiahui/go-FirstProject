package test

import (
	"testing"

	"project1/entity"
	"project1/service"
	"project1/util"
)

func Test_GetAckByIdService(t *testing.T) {
	target := make(map[string]entity.Soldier)
	newSoldierMap := util.JsonToFile("./config/config.army.model.json", target)
	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
	for _, v := range []struct {
		id         string
		soldierMap map[string]entity.SoldierDTO
		ret        string
	}{
		{"1", newSoldierMap, "-1"},
		{"2", newSoldierMap, "-1"},
		{"3", newSoldierMap, "-1"},
		{"4", newSoldierMap, "-1"},
	} {
		// 调用排列组合函数，与期望的结果比对，如果不一致输出错误
		if actually := service.GetAckByIdService(v.id, v.soldierMap); actually == v.ret {
			t.Errorf("combination: [%v], actually: [%v]", v, actually)
		}
	}
}

func Test_GetRarityById(t *testing.T) {
	target := make(map[string]entity.Soldier)
	newSoldierMap := util.JsonToFile("./config/config.army.model.json", target)
	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
	for _, v := range []struct {
		id         string
		soldierMap map[string]entity.SoldierDTO
		ret        string
	}{
		{"1", newSoldierMap, "-1"},
		{"2", newSoldierMap, "-1"},
		{"3", newSoldierMap, "-1"},
		{"4", newSoldierMap, "-1"},
	} {
		// 调用排列组合函数，与期望的结果比对，如果不一致输出错误
		if actually := service.GetRarityById(v.id, v.soldierMap); actually == v.ret {
			t.Errorf("combination: [%v], actually: [%v]", v, actually)
		}
	}
}

func Test_GetAllByRarityService(t *testing.T) {
	target := make(map[string]entity.Soldier)
	newSoldierMap := util.JsonToFile("./config/config.army.model.json", target)
	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
	for _, v := range []struct {
		rarity      string
		unlockArena string
		soldierMap  map[string]entity.SoldierDTO
		ret         []entity.SoldierDTO
	}{
		{"1", "1", newSoldierMap, nil},
		{"2", "2", newSoldierMap, nil},
		{"3", "2", newSoldierMap, nil},
		{"4", "1", newSoldierMap, nil},
	} {
		// 调用排列组合函数，与期望的结果比对，如果不一致输出错误
		if actually := service.GetAllByRarityService(v.rarity, v.unlockArena, v.soldierMap); actually == nil {
			t.Errorf("combination: [%v], actually: [%v]", v, actually)
		}
	}
}

func Test_GetSoldierByUnlockArena(t *testing.T) {
	target := make(map[string]entity.Soldier)
	newSoldierMap := util.JsonToFile("./config/config.army.model.json", target)
	// 这里定义一个临时的结构体来存储测试case的参数以及期望的返回值
	for _, v := range []struct {
		soldierMap map[string]entity.SoldierDTO
		ret        []entity.SoldierDTO
	}{
		{newSoldierMap, nil},
		{newSoldierMap, nil},
		{newSoldierMap, nil},
		{newSoldierMap, nil},
	} {
		// 调用排列组合函数，与期望的结果比对，如果不一致输出错误
		if actually := service.GetSoldierByUnlockArena(v.soldierMap); actually == nil {
			t.Errorf("combination: [%v], actually: [%v]", v, actually)
		}
	}
}
