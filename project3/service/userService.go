package service

import (
	"encoding/json"
	"strconv"
	"time"

	"project3/entity"
	"project3/util"
)

//用户领取礼包
func CheckCodeService(uid, code string, userMap map[string]string) (int, []entity.Gift, string) {
	//获取领取用户名
	var username string
	_, ok := userMap[uid]
	if ok {
		username = userMap[uid]
	} else {
		return 30000, nil, "不存在该用户"
	}
	//获取领取时间
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

	//判断领取时间是否超过限定天数
	if timeEro1 == nil && timeEro2 == nil && time1.Before(time2) {
		//判断是否属于限制用户，限制领取次数为1次的礼包
		if ret.CodeType == "1" && ret.DrawId == uid {
			//判断礼包是不是不能领了
			if ret.ReceiveNum == "0" {
				return 30002, nil, "该礼包已经被领取了"
			}
			//用户领取以后减少可领取次数，增加已领取次数
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
		//判断礼包类型为不指定用户，限制兑换次数
		if ret.CodeType == "2" {
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
		//判断礼包类型为：不指定用户，不限制兑换次数
		if ret.CodeType == "3" {
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
		return 50001, nil, "json序列化错误"
	}
	redisEro2 := util.SetRedis(code, value)
	if redisEro2 != nil {
		return 50002, nil, "redis服务有问题"
	}
	return 200, ret.GiftList, "请求成功"
}
