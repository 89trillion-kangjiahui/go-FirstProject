package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

	. "project1/entity"
	"project1/service"
)

//需求1：输入稀有度，当前解锁阶段和cvc，获取该稀有度cvc合法且已解锁的所有士兵
func GetAllByRarity(c *gin.Context) {
	rarity := c.Query("rarity")
	unlockArena := c.Query("unlockArena")
	cvc := c.Query("cvc")
	if rarity == "" {
		ret := SetResult(3000, "rarity不能为空", nil)
		c.JSON(http.StatusOK, ret)
	} else if unlockArena == "" {
		ret := SetResult(3001, "unlockArena不能为空", nil)
		c.JSON(http.StatusOK, ret)
	} else if cvc == "" {
		ret := SetResult(3002, "cvc不能为空", nil)
		c.JSON(http.StatusOK, ret)
	} else {
		code, msg, data := service.GetAllByRarity(rarity, unlockArena, cvc)
		ret := SetResult(code, msg, data)
		c.JSON(http.StatusOK, ret)
	}
}

//输入士兵id获取战力
func GetAckById(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		ret := SetResult(3003, "id不能为空", nil)
		c.JSON(http.StatusOK, ret)
	} else {
		code, msg, data := service.GetAckById(id)
		ret := SetResult(code, msg, data)
		c.JSON(http.StatusOK, ret)
	}
}

//输入士兵id获取稀有度
func GetRarityById(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		ret := SetResult(3003, "id不能为空", nil)
		c.JSON(http.StatusOK, ret)
	} else {
		code, msg, data := service.GetRarityById(id)
		ret := SetResult(code, msg, data)
		c.JSON(http.StatusOK, ret)
	}
}

//输入cvc获取所有合法的士兵
func GetSoldierByCvc(c *gin.Context) {
	cvc := c.Query("cvc")
	if cvc == "" {
		ret := SetResult(3002, "cvc不能为空", nil)
		c.JSON(http.StatusOK, ret)
	} else {
		code, msg, data := service.GetSoldierByCvc(cvc)
		ret := SetResult(code, msg, data)
		c.JSON(http.StatusOK, ret)
	}
}

//获取每个阶段解锁相应士兵的json数据
func GetSoldierByUnlockArena(c *gin.Context) {
	code, msg, data := service.GetSoldierByUnlockArena()
	ret := SetResult(code, msg, data)
	c.JSON(http.StatusOK, ret)
}
