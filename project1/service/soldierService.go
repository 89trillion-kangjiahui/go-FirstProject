package service

import . "project1/entity"

func GetAllByRarityService(rarity, unlockArena string, soldierMap map[string]SoldierDTO) []SoldierDTO {
	returnData := make([]SoldierDTO, 0)
	if soldierMap == nil {
		return nil
	} else {
		for _, v := range soldierMap {
			if v.Rarity == rarity && v.UnlockArena == unlockArena {
				returnData = append(returnData, v)
			}
		}
		return returnData
	}
}

func GetAckByIdService(id string, soldierMap map[string]SoldierDTO) string {
	var atc string
	if soldierMap == nil {
		return "-1"
	} else {
		for _, v := range soldierMap {
			if v.Id == id {
				atc = v.Atk
				break
			}
		}
		return atc
	}
}

func GetRarityById(id string, soldierMap map[string]SoldierDTO) string {
	var rarity string
	if soldierMap == nil {
		return "-1"
	} else {
		for _, v := range soldierMap {
			if v.Id == id {
				rarity = v.Rarity
				break
			}
		}
		return rarity
	}
}

func GetSoldierByUnlockArena(soldierMap map[string]SoldierDTO) map[string][]SoldierDTO {
	ret := make(map[string][]SoldierDTO)
	if soldierMap == nil {
		return nil
	} else {
		for _, v := range soldierMap {
			unlockArena := v.UnlockArena
			if i := ret[unlockArena]; i == nil {
				i = make([]SoldierDTO, 0)
			}
			ret[unlockArena] = append(ret[unlockArena], v)
		}
		return ret
	}
}
