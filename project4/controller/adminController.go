package controller

import (
	"fmt"
	"strconv"

	"project3/service"

	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminCreateCode(r *gin.Engine, url string) {
	r.POST(url, func(c *gin.Context) {
		codeType := c.PostForm("codeType")     //礼品码类型 1:指定用户一次性消耗，2：不指定用户限制兑换次数，3：不限用户，不限兑换次数
		describe := c.PostForm("des")          //礼品描述
		receiveNum,_ := strconv.Atoi(c.PostForm("receiveNum")) //可领取次数
		usefulDate := c.PostForm("usefulDate") //有效期
		jewel,_ := strconv.Atoi(c.PostForm("jewel"))           //钻石数量
		gold,_ := strconv.Atoi(c.PostForm("gold"))             //金币数量
		props,_ := strconv.Atoi(c.PostForm("props"))         //道具数量
		hero,_ := strconv.Atoi(c.PostForm("hero"))             //英雄数量
		batman,_ := strconv.Atoi(c.PostForm("batman"))         //小兵数量
		ret, serviceEro := service.CreateCodeService(codeType, describe, usefulDate, receiveNum, uint64(jewel), uint64(gold), uint64(props), uint64(hero), uint64(batman))
		if serviceEro != nil {
			fmt.Println(serviceEro)
			c.JSON(http.StatusInternalServerError, serviceEro.Error())
		}
		if serviceEro == nil {
			c.JSON(http.StatusOK, ret)
		}
	})
}

func AdminSelectCode(r *gin.Engine, url string) {
	r.GET(url, func(c *gin.Context) {
		code := c.Query("code")
		ret, serviceEro := service.SelectCodeService(code)
		if serviceEro != nil {
			fmt.Println(serviceEro)
			c.JSON(http.StatusInternalServerError, serviceEro.Error())
		}
		if serviceEro == nil {
			c.JSON(http.StatusOK, ret)
		}
	})
}
