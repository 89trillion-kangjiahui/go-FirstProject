package service

import (
	"encoding/json"
	"time"

	. "project3/entity"
	"project3/util"
)

func CreateCodeService(codeType, describe, usefulDate string, receiveNum int, jewel, gold, props, hero, batman uint64) (string, error) {
	createTime := time.Now().Format("2006-01-02 15:04:05")
	key := util.RandSeq(8)
	prize := GiftContent{
		GmtCreate:  createTime,
		CreateUser: "管理员",
		Describe:   describe,
		GiftList: []Gift{
			Gift{Gid: 1, Name: "jewel", Num: jewel},
			Gift{Gid: 2, Name: "gold", Num: gold},
			Gift{Gid: 3, Name: "props", Num: props},
			Gift{Gid: 4, Name: "hero", Num: hero},
			Gift{Gid: 5, Name: "batman", Num: batman},
		},
		ReceiveNum: receiveNum,
		UsefulDate: usefulDate,
		CodeType:   codeType,
	}
	if codeType == "1" {
		prize.DrawId = "7cccef0dea75bcc832f924a3bcaaf456"
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
