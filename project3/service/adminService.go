package service

import (
	"encoding/json"
	"time"

	. "project3/entity"
	"project3/util"
)

func CreateCodeService(uid, codeType, describe, receiveNum, usefulDate, jewel, gold, props, hero, batman string, userMap map[string]string) (string, error) {
	username := userMap[uid]
	createTime := time.Now().Format("2006-01-02 15:04:05")
	key := util.RandSeq(8)
	prize := GiftContent{
		GmtCreate:  createTime,
		CreateUser: username,
		Describe:   describe,
		GiftList: []Gift{
			Gift{Name: "jewel", Num: jewel},
			Gift{Name: "gold", Num: gold},
			Gift{Name: "props", Num: props},
			Gift{Name: "hero", Num: hero},
			Gift{Name: "batman", Num: batman},
		},
		ReceiveNum: receiveNum,
		UsefulDate: usefulDate,
		CodeType:   codeType,
	}
	for i := 0; i < 5; i++ {

	}
	if codeType == "1" {
		prize.DrawId = "2"
	}
	value, jsonEro := json.Marshal(prize)
	if jsonEro != nil {
		return "", jsonEro
	}
	redisEro := util.SetRedis(key, value)
	if redisEro != nil {
		return "", redisEro
	}
	return key, nil
}

func SelectCodeService(code string) (GiftContent, error) {
	value, jsonEro := util.GetRedis(code)
	if jsonEro != nil {
		return GiftContent{}, jsonEro
	}
	jsonString := []byte(value)
	var ret GiftContent
	json.Unmarshal(jsonString, &ret)
	return ret, nil
}
