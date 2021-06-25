package service

import (
	"fmt"

	. "project1/entity"
	"project1/global"
)

// 需求1:输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
func GetAllByRarity(rarity, unlockArena, cvc string) (int, string, []SoldierDTO) {
	Data := make([]SoldierDTO, 0)
	if global.SoldierMap == nil {
		return 4001, "数据库为空", nil
	} else {
		for _, v := range global.SoldierMap {
			if v.Rarity == rarity && v.UnlockArena == unlockArena && cvc == v.Cvc {
				Data = append(Data, v)
			}
		}
		return 200, "请求成功", Data
	}
}

//需求2：输入士兵id获取稀有度
func GetRarityById(id string) (int, string, string) {
	if global.SoldierMap == nil {
		return 4001, "数据库为空", ""
	} else {
		_, ok := global.SoldierMap[id]
		if ok {
			return 200, "请求成功", global.SoldierMap[id].Rarity
		} else {
			return 4002, "该用户不存在", ""
		}
	}
}

//需求3：输入士兵id获取战力
func GetAckById(id string) (int, string, string) {
	if global.SoldierMap == nil {
		return 4001, "数据库为空", ""
	} else {
		_, ok := global.SoldierMap[id]
		fmt.Println(global.SoldierMap[id])
		if ok {
			return 200, "请求成功", global.SoldierMap[id].Atk
		} else {
			return 4002, "该用户不存在", ""
		}
	}
}

//需求4：输入cvc获取所有合法的士兵
func GetSoldierByCvc(cvc string) (int, string, []SoldierDTO) {
	Data := make([]SoldierDTO, 0)
	if global.SoldierMap == nil {
		return 4001, "数据库为空", nil
	} else {
		for _, v := range global.SoldierMap {
			if cvc == v.Cvc {
				Data = append(Data, v)
			}
		}
		return 200, "请求成功", Data
	}
}

//需求5：获取每个阶段解锁相应士兵的json数据
func GetSoldierByUnlockArena() (int, string, map[string][]SoldierDTO) {
	Data := make(map[string][]SoldierDTO)
	if global.SoldierMap == nil {
		return 4001, "数据库为空", nil
	} else {
		for _, v := range global.SoldierMap {
			unlockArena := v.UnlockArena
			Data[unlockArena] = append(Data[unlockArena], v)
		}
		return 200, "请求成功", Data
	}
}
