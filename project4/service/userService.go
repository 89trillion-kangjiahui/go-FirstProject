package service

import (
	"encoding/json"
	"time"

	"github.com/golang/protobuf/proto"

	"project3/dao"
	"project3/entity"
	"project3/response"
	"project3/util"
)

func CheckCodeService(uid, code string) []byte {
	user, daoEro := dao.SelectByUid(uid)
	drawTime := time.Now().Format("2006-01-02 15:04:05")
	//判断用户是否注册
	if daoEro != nil {
		daoEroRes := setResponse(20100, "该用户未注册", nil, entity.Gift{}, entity.Gift{})
		data, _ := proto.Marshal(daoEroRes)
		return data
	}
	//判断礼包吗是否正确
	value, redisEro := util.GetRedis(code)
	if redisEro != nil {
		redisRes := setResponse(30100, "该礼包码不正确", nil, user.GemBalance, user.GoldBalance)
		data, _ := proto.Marshal(redisRes)
		return data
	}

	jsonString := []byte(value)
	var ret entity.GiftContent
	json.Unmarshal(jsonString, &ret)
	time1, timeEro1 := time.Parse("2006-01-02 15:04:05", drawTime)
	time2, timeEro2 := time.Parse("2006-01-02 15:04:05", ret.UsefulDate)

	//判断时间是否过期
	if timeEro1 == nil && timeEro2 == nil && time1.Before(time2) {
		//礼包类型为指定用户，限定兑换次数为1
		if ret.CodeType == "1" && ret.DrawId == uid {
			if ret.ReceiveNum == 0 {
				receRes := setResponse(10100, "该礼包已经被领取了", nil, user.GemBalance, user.GoldBalance)
				data, _ := proto.Marshal(receRes)
				return data
			}
			//指定用户一次性消耗
			ret.ReceiveNum = 0
			ret.AlreadyNum = 1
			oldDrawList := ret.DrawList
			if oldDrawList == nil {
				oldDrawList = make(map[string]string)
			}
			oldDrawList[drawTime] = user.UserName
			ret.DrawList = oldDrawList
			dao.UpdateBalance(uid, ret.GiftList)
			userBalance, _ := dao.SelectByUid(uid)
			return giftSetRedis(code, ret, *userBalance)
		}
		if ret.CodeType == "2" {
			//不指定用户限制兑换次数
			oldReNum := ret.ReceiveNum
			//可领取次数为0
			if oldReNum == 0 {
				receRes := setResponse(10101, "该礼包已经被领完了", nil, user.GemBalance, user.GoldBalance)
				data, _ := proto.Marshal(receRes)
				return data
			}
			//更新可领取次数
			ret.ReceiveNum = oldReNum - 1
			//更新已领取次数
			oldAlNum := ret.AlreadyNum
			ret.AlreadyNum = oldAlNum + 1
			oldDrawList := ret.DrawList
			if oldDrawList == nil {
				oldDrawList = make(map[string]string)
			}
			oldDrawList[drawTime] = user.UserName
			ret.DrawList = oldDrawList
			dao.UpdateBalance(uid, ret.GiftList)
			userBalance, _ := dao.SelectByUid(uid)
			return giftSetRedis(code, ret, *userBalance)
		}
		if ret.CodeType == "3" {
			//不指定用户限制兑换次数
			oldAlNum := ret.AlreadyNum
			ret.AlreadyNum = oldAlNum + 1
			oldDrawList := ret.DrawList
			if oldDrawList == nil {
				oldDrawList = make(map[string]string)
			}
			oldDrawList[drawTime] = user.UserName
			ret.DrawList = oldDrawList
			dao.UpdateBalance(uid, ret.GiftList)
			userBalance, _ := dao.SelectByUid(uid)
			return giftSetRedis(code, ret, *userBalance)
		}
		eroRes := setResponse(10102, "不是指定用户领取", nil, user.GemBalance, user.GoldBalance)
		data, _ := proto.Marshal(eroRes)
		return data
	} else {
		eroRes := setResponse(10103, "超过礼品码限定日期", nil, user.GemBalance, user.GoldBalance)
		data, _ := proto.Marshal(eroRes)
		return data
	}
}

func giftSetRedis(code string, ret entity.GiftContent, user entity.User) []byte {
	value, jsonEro := json.Marshal(&ret)
	if jsonEro != nil {
		jsonRes := setResponse(40100, "json序列化失败", ret.GiftList, user.GemBalance, user.GoldBalance)
		data, _ := proto.Marshal(jsonRes)
		return data
	}
	redisEro2 := util.SetRedis(code, value)
	if redisEro2 != nil {
		redisRes := setResponse(30101, "redis中礼品内容修改失败", ret.GiftList, user.GemBalance, user.GoldBalance)
		data, _ := proto.Marshal(redisRes)
		return data
	}
	res := setResponse(200, "获得奖励", ret.GiftList, user.GemBalance, user.GoldBalance)
	data, _ := proto.Marshal(res)
	return data
}

func UserLoginService(id string) *entity.User {
	user, _ := dao.SelectByUid(id)
	return user
}

func UserRegisterService(userName string) (int, string, string) {
	uid := util.UniqueId()
	if userName == "" {
		return 20001, "用户名为空", ""
	}
	user := entity.User{
		Uid:      uid,
		UserName: userName,
		GemBalance: entity.Gift{
			Gid:  1,
			Name: "jewel",
			Num:  0,
		},
		GoldBalance: entity.Gift{
			Gid:  2,
			Name: "gold",
			Num:  0,
		},
	}
	err := dao.InsertUser(&user)
	if err != nil {
		return 20002, "用户插入失败", ""
	}
	return 200, "请求成功", uid
}

func setResponse(code int32, msg string, giftList []entity.Gift, gemBalance, goldBalance entity.Gift) *response.GeneralReward {
	balance := make(map[uint32]uint64)
	balance[gemBalance.Gid] = gemBalance.Num
	balance[goldBalance.Gid] = gemBalance.Num
	chance := make(map[uint32]uint64)
	if giftList != nil {
		for _, v := range giftList {
			if v.Gid == 1 {
				chance[1] = v.Num
			}
			if v.Gid == 2 {
				chance[2] = v.Num
			}
		}
		response := &response.GeneralReward{
			Code:    code,
			Msg:     msg,
			Changes: chance,
			Balance: balance,
		}
		return response
	} else {
		response := &response.GeneralReward{
			Code:    code,
			Msg:     msg,
			Changes: nil,
			Balance: balance,
		}
		return response
	}
}
