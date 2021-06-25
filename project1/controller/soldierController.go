package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

	. "project1/entity"
	"project1/service"
)

func GetAllByRarity(r *gin.Engine, url string, soldierMap map[string]SoldierDTO) {
	r.GET(url, func(c *gin.Context) {
		rarity := c.Query("rarity")
		unlockArena := c.Query("unlockArena")
		returnData := service.GetAllByRarityService(rarity, unlockArena, soldierMap)
		if returnData == nil {
			ret := SetResult(4001, "没有相关内容", nil)
			c.JSON(http.StatusOK, ret)
		} else {
			ret := SetResult(200, "找到符合信息的士兵", returnData)
			c.JSON(http.StatusOK, ret)
		}
	})
}

func GetAckById(r *gin.Engine, url string, soldierMap map[string]SoldierDTO) {
	// gin.Context，封装了request和response
	r.GET(url, func(c *gin.Context) {
		id := c.Query("id")
		atc := service.GetAckByIdService(id, soldierMap)
		if atc == "" {
			ret := SetResult(4001, "没有相关内容", nil)
			c.JSON(http.StatusOK, ret)
		} else {
			ret := SetResult(200, "找到了该士兵的战斗力", atc)
			c.JSON(http.StatusOK, ret)
		}
	})
}

func GetRarityById(r *gin.Engine, url string, soldierMap map[string]SoldierDTO) {
	// gin.Context，封装了request和response
	r.GET(url, func(c *gin.Context) {
		id := c.Query("id")
		rarity := service.GetRarityById(id, soldierMap)
		if rarity == "" {
			ret := SetResult(4001, "没有相关内容", nil)
			c.JSON(http.StatusOK, ret)
		} else {
			ret := SetResult(200, "找到了该士兵的稀有度", rarity)
			c.JSON(http.StatusOK, ret)
		}
	})
}

func GetSoldierByUnlockArena(r *gin.Engine, url string, soldierMap map[string]SoldierDTO) {
	r.GET(url, func(c *gin.Context) {
		retData := service.GetSoldierByUnlockArena(soldierMap)
		if retData == nil {
			ret := SetResult(4001, "没有相关内容", nil)
			c.JSON(http.StatusOK, ret)
		} else {
			ret := SetResult(200, "找到符合信息的士兵", retData)
			c.JSON(http.StatusOK, ret)
		}
	})
}
