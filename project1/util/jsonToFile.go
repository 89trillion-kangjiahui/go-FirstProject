package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	. "project1/entity"
)
func JsonToFile(target map[string]Soldier)  {
	ret := make(map[string]SoldierDTO)
	for k,v := range target{
		ret[k] = SoldierDTO{
			Id: v.Id,
			Name: v.Name,
			UnlockArena: v.UnlockArena,
			Rarity: v.Rarity,
			Atk: v.Atk,
		}
	}
	data, ero := json.Marshal(ret)
	if ero != nil {
		fmt.Println("json转化失败")
		return
	}
	testWrite(data)
}

func testWrite(data []byte) {
	fp, err := os.OpenFile("config/new.soldier.dto.json", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}