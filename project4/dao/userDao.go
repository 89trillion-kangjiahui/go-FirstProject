package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"project3/entity"
	"project3/util"
)

func SelectByUid(uid string) (*entity.User, error) {
	session, sessionEro, collection := util.GetSession()
	if sessionEro != nil {
		return nil, sessionEro
	}
	defer session.Disconnect(context.TODO())

	result := entity.User{}
	filter := bson.D{{"uid",uid}}
	selectErr := collection.FindOne(context.TODO(), filter).Decode(&result)
	if selectErr != nil {
		return nil, selectErr
	}
	return &result, nil
}

func InsertUser(user *entity.User) error {
	session, sessionEro, collection := util.GetSession()
	if sessionEro != nil {
		return sessionEro
	}
	defer session.Disconnect(context.TODO())

	//一次可以插入多个对象 插入两个Person对象
	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func UpdateBalance(uid string, giftList []entity.Gift) error  {
	session, sessionEro, collection := util.GetSession()
	if sessionEro != nil {
		return sessionEro
	}
	defer session.Disconnect(context.TODO())

	user, selectEro := SelectByUid(uid)
	if selectEro != nil {
		return selectEro
	}
	var GemChance uint64
	var GoldChance uint64
	for _, v := range giftList{
		if v.Gid == 1 {
			GemChance = v.Num
		}
		if v.Name == "gold" {
			GoldChance = v.Num
		}
	}
	user.GemBalance.Num = user.GemBalance.Num + GemChance
	user.GoldBalance.Num = user.GoldBalance.Num + GoldChance
	filter := bson.D{{"uid",uid}}

	update := bson.M{
		"$set": user,
	}
	_, ero := collection.UpdateOne(context.Background(),filter,update)
	return ero
}

