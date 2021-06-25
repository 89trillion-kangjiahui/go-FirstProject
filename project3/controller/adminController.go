package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"project3/entity"
	"project3/global"
	"project3/service"
)

// 需求1：管理员创建礼品包
func AdminCreateCode(c *gin.Context) {
	uid := c.PostForm("uid")
	codeType := c.PostForm("codeType")     //礼品码类型 1:指定用户一次性消耗，2：不指定用户限制兑换次数，3：不限用户，不限兑换次数
	describe := c.PostForm("des")          //礼品描述
	receiveNum := c.PostForm("receiveNum") //可领取次数
	usefulDate := c.PostForm("usefulDate") //有效期
	jewel := c.PostForm("jewel")           //钻石数量
	gold := c.PostForm("gold")             //金币数量
	props := c.PostForm("props")           //道具数量
	hero := c.PostForm("hero")             //英雄数量
	batman := c.PostForm("batman")         //小兵数量
	retCode, data, serviceEro := service.CreateCodeService(uid, codeType, describe, receiveNum, usefulDate, jewel, gold, props, hero, batman, global.UserMap)
	ret := entity.SetResult(retCode, serviceEro, data)
	c.JSON(http.StatusOK, ret)
}

// 需求2: 管理员查询礼品码内容
func AdminSelectCode(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		ret := entity.SetResult(4001, "code不能为空", "")
		c.JSON(http.StatusOK, ret)
	}else{
		retCode, data, serviceEro := service.SelectCodeService(code)
		ret := entity.SetResult(retCode, serviceEro, data)
		c.JSON(http.StatusOK, ret)
	}
}
