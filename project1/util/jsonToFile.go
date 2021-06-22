package util

import (
	"encoding/json"
	"io/ioutil"
	"os"

	. "project1/entity"
)

func JsonToFile(jsonPath string, target map[string]Soldier) map[string]SoldierDTO {
	data, _ := ioutil.ReadFile(jsonPath)
	json.Unmarshal(data, &target)
	ret := make(map[string]SoldierDTO)

	for k, v := range target {
		ret[k] = SoldierDTO{
			Id:          v.Id,
			Name:        v.Name,
			UnlockArena: v.UnlockArena,
			Rarity:      v.Rarity,
			Atk:         v.Atk,
		}
	}
	data, _ = json.Marshal(ret)
	WriteJsonFile(data)
	return ret
}

func WriteJsonFile(data []byte) {
	fp, _ := os.OpenFile("config/new.soldier.dto.json", os.O_RDWR|os.O_CREATE, 0755)
	defer fp.Close()
	fp.Write(data)
}
