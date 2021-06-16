package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	. "project1/entity"
)
func JsonToFile(jsonPath string, target *map[string]Soldier) *map[string]SoldierDTO {
	data,error := ioutil.ReadFile(jsonPath)
	if error != nil {
		panic("文件读取失败")
	}
	if ero := json.Unmarshal(data, target); ero != nil {
		panic("json解析出错了")
	}
	ret := make(map[string]SoldierDTO)
	for k,v := range *target{
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
		panic("json转化失败")
	}
	testWrite(data)
	return &ret
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