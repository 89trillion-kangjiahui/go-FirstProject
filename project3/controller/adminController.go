package controller

import (
	"project3/entity"
	"project3/service"

	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminCreateCode(r *gin.Engine, url string, userMap map[string]string) {
	r.POST(url, func(c *gin.Context) {
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
		retCode, data, serviceEro := service.CreateCodeService(uid, codeType, describe, receiveNum, usefulDate, jewel, gold, props, hero, batman, userMap)
		ret := entity.SetResult(retCode, serviceEro, data)
		c.JSON(http.StatusOK, ret)
	})
}

func AdminSelectCode(r *gin.Engine, url string) {
	r.GET(url, func(c *gin.Context) {
		code := c.Query("code")
		retCode, data, serviceEro := service.SelectCodeService(code)
		ret := entity.SetResult(retCode, serviceEro, data)
		c.JSON(http.StatusOK, ret)
	})
}
