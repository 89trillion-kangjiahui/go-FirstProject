package service

import (
	"encoding/json"
	"time"

	. "project3/entity"
	"project3/util"
)

func CreateCodeService(codeType, describe, usefulDate string, receiveNum int, jewel, gold, props, hero, batman uint64) (int, string, string) {
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
		return 50001, "json序列化失败", ""
	}
	redisEro := util.SetRedis(key, value)
	if redisEro != nil {
		return 50002, "redis服务错误", ""
	}
	return 200, "请求成功", key
}

func SelectCodeService(code string) (int, string, GiftContent) {
	value, jsonEro := util.GetRedis(code)
	if jsonEro != nil {
		return 50001, "json序列化失败", GiftContent{}
	}
	jsonString := []byte(value)
	var ret GiftContent
	json.Unmarshal(jsonString, &ret)
	return 200, "请求成功", ret
}
