package service

import (
	"encoding/json"
	"strconv"
	"time"

	"project3/entity"
	"project3/util"
)

func CheckCodeService(uid, code string, userMap map[string]string) (int, []entity.Gift, string) {
	username := userMap[uid]
	drawTime := time.Now().Format("2006-01-02 15:04:05")
	value, redisEro := util.GetRedis(code)
	if redisEro != nil {
		return 30001, nil, "礼包码不正确"
	}

	jsonString := []byte(value)
	var ret entity.GiftContent
	json.Unmarshal(jsonString, &ret)

	time1, timeEro1 := time.Parse("2006-01-02 15:04:05", drawTime)
	time2, timeEro2 := time.Parse("2006-01-02 15:04:05", ret.UsefulDate)

	if timeEro1 == nil && timeEro2 == nil && time1.Before(time2) {
		if ret.CodeType == "1" && ret.DrawId == uid {
			if ret.ReceiveNum == "0" {
				return 30002, nil, "该礼包已经被领取了"
			}
			//指定用户一次性消耗
			ret.ReceiveNum = "0"
			ret.AlreadyNum = "1"
			oldDrawList := ret.DrawList
			if oldDrawList == nil {
				oldDrawList = make(map[string]string)
			}
			oldDrawList[drawTime] = username
			ret.DrawList = oldDrawList
			return giftSetRedis(code, &ret)
		}
		if ret.CodeType == "2" {
			//不指定用户限制兑换次数
			oldReNum, _ := strconv.Atoi(ret.ReceiveNum)
			//可领取次数为0
			if oldReNum == 0 {
				return 30003, nil, "该礼包已经没有了"
			}
			ret.ReceiveNum = strconv.Itoa(oldReNum - 1)
			if ret.AlreadyNum == "" {
				ret.AlreadyNum = "1"
			} else {
				oldAlNum, _ := strconv.Atoi(ret.AlreadyNum)
				ret.AlreadyNum = strconv.Itoa(oldAlNum + 1)
			}
			oldDrawList := ret.DrawList
			if oldDrawList == nil {
				oldDrawList = make(map[string]string)
			}
			oldDrawList[drawTime] = username
			ret.DrawList = oldDrawList
			return giftSetRedis(code, &ret)
		}
		if ret.CodeType == "3" {
			//不指定用户限制兑换次数
			if ret.AlreadyNum == "" {
				ret.AlreadyNum = "1"
			} else {
				oldAlNum, _ := strconv.Atoi(ret.AlreadyNum)
				ret.AlreadyNum = strconv.Itoa(oldAlNum + 1)
			}
			oldDrawList := ret.DrawList
			if oldDrawList == nil {
				oldDrawList = make(map[string]string)
			}
			oldDrawList[drawTime] = username
			ret.DrawList = oldDrawList
			return giftSetRedis(code, &ret)
		}

		return 30004, nil, "不是指定的用户领取"
	} else {
		return 30005, nil, "领取时间超过限定日期"
	}
}

func giftSetRedis(code string, ret *entity.GiftContent) (int, []entity.Gift, string) {
	value, jsonEro := json.Marshal(*ret)
	if jsonEro != nil {
		return 40001, nil, "json序列化错误"
	}
	redisEro2 := util.SetRedis(code, value)
	if redisEro2 != nil {
		return 50001, nil, "redis服务有问题"
	}
	return 200, ret.GiftList, "请求成功"
}
