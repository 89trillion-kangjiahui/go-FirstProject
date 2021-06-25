package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	. "project1/entity"
)

func JsonToFile(jsonPath string) map[string]SoldierDTO {
	var target = make(map[string]Soldier)
	data, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(data, &target)
	ret := make(map[string]SoldierDTO)
	for k, v := range target {
		ret[k] = SoldierDTO{
			Id:          v.Id,
			Name:        v.Name,
			UnlockArena: v.UnlockArena,
			Rarity:      v.Rarity,
			Atk:         v.Atk,
			Cvc:         v.Cvc,
		}
	}
	data, _ = json.Marshal(ret)
	WriteJsonFile(data)
	return ret
}

func WriteJsonFile(data []byte) {
	fp, _ := os.OpenFile("../config/new.soldier.dto.json", os.O_RDWR|os.O_CREATE, 0755)
	defer fp.Close()
	fp.Write(data)
}
