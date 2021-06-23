package service

import (
	"encoding/json"
	"time"

	. "project3/entity"
	"project3/util"
)

func CreateCodeService(uid, codeType, describe, receiveNum, usefulDate, jewel, gold, props, hero, batman string, userMap map[string]string) (int, string, string) {
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
	if codeType == "1" {
		prize.DrawId = "2"
	}
	value, jsonEro := json.Marshal(prize)
	if jsonEro != nil {
		return 40001, "", "json序列化错误"
	}
	redisEro := util.SetRedis(key, value)
	if redisEro != nil {
		return 50001, "", "redis服务有问题"
	}
	return 200, key, "请求成功"
}

func SelectCodeService(code string) (int, GiftContent, string) {
	value, jsonEro := util.GetRedis(code)
	if jsonEro != nil {
		return 40001, GiftContent{}, "json序列化错误"
	}
	jsonString := []byte(value)
	var ret GiftContent
	json.Unmarshal(jsonString, &ret)
	return 200, ret, "请求成功"
}
