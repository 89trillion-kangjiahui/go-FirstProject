package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"project3/entity"
	"project3/util"
)

func CheckCodeService(uid, code string, userMap map[string]string) ([]entity.Gift, error) {
	username := userMap[uid]
	drawTime := time.Now().Format("2006-01-02 15:04:05")
	value, redisEro := util.GetRedis(code)
	if redisEro != nil {
		return nil, errors.New("礼包码不正确")
	}

	jsonString := []byte(value)
	var ret entity.GiftContent
	json.Unmarshal(jsonString, &ret)

	time1, timeEro1 := time.Parse("2006-01-02 15:04:05", drawTime)
	time2, timeEro2 := time.Parse("2006-01-02 15:04:05", ret.UsefulDate)

	if timeEro1 == nil && timeEro2 == nil && time1.Before(time2) {
		fmt.Println("成功了")
		if ret.CodeType == "1" && ret.DrawId == uid {
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
				return nil, errors.New("可领取为0")
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

		return nil, errors.New("不是指定的用户领取")
	} else {
		return nil, errors.New("领取时间超过限定日期")
	}
}

func giftSetRedis(code string, ret *entity.GiftContent) ([]entity.Gift, error) {
	value, jsonEro := json.Marshal(*ret)
	if jsonEro != nil {
		return nil, jsonEro
	}
	redisEro2 := util.SetRedis(code, value)
	if redisEro2 != nil {
		return nil, redisEro2
	}
	fmt.Println(ret.GiftList)
	return ret.GiftList, nil
}
