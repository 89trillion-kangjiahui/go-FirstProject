package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"project3/entity"
	"project3/service"
)

func AdminCreateCode(c *gin.Context) {
	codeType := c.PostForm("codeType")                      //礼品码类型 1:指定用户一次性消耗，2：不指定用户限制兑换次数，3：不限用户，不限兑换次数
	describe := c.PostForm("des")                           //礼品描述
	receiveNum, _ := strconv.Atoi(c.PostForm("receiveNum")) //可领取次数
	usefulDate := c.PostForm("usefulDate")                  //有效期
	jewel, _ := strconv.Atoi(c.PostForm("jewel"))           //钻石数量
	gold, _ := strconv.Atoi(c.PostForm("gold"))             //金币数量
	props, _ := strconv.Atoi(c.PostForm("props"))           //道具数量
	hero, _ := strconv.Atoi(c.PostForm("hero"))             //英雄数量
	batman, _ := strconv.Atoi(c.PostForm("batman"))         //小兵数量
	retCode, msg, data := service.CreateCodeService(codeType, describe, usefulDate, receiveNum, uint64(jewel), uint64(gold), uint64(props), uint64(hero), uint64(batman))
	ret := entity.SetResult(retCode, msg, data)
	c.JSON(http.StatusOK, ret)
}

func AdminSelectCode(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		ret := entity.SetResult(4002, "code不能为空", nil)
		c.JSON(http.StatusOK, ret)
	} else {
		retCode, msg, data := service.SelectCodeService(code)
		ret := entity.SetResult(retCode, msg, data)
		c.JSON(http.StatusOK, ret)
	}
}
