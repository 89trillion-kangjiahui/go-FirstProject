package controller

import (
	. "project1/entity"

	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllByRarity(r *gin.Engine, url string, soldierMap map[string]SoldierDTO) {
	r.GET(url, func(c *gin.Context) {
		rarity := c.Query("rarity")
		unlockArena := c.Query("unlockArena")
		returnData := make([]SoldierDTO, 0)
		for _, v := range soldierMap {
			if v.Rarity == rarity && v.UnlockArena == unlockArena {
				returnData = append(returnData, v)
			}
		}

		c.JSON(http.StatusOK, returnData)
	})
}

func GetAckById(r *gin.Engine, url string, soldierMap map[string]SoldierDTO) {
	// gin.Context，封装了request和response
	r.GET(url, func(c *gin.Context) {
		id := c.Query("id")
		var atc string
		for _, v := range soldierMap {
			if v.Id == id {
				atc = v.Atk
				break
			}
		}
		c.String(http.StatusOK, "士兵的战斗力:"+atc)
	})
}

func GetRarityById(r *gin.Engine, url string, soldierMap map[string]SoldierDTO) {
	// gin.Context，封装了request和response
	r.GET(url, func(c *gin.Context) {
		id := c.Query("id")
		var rarity string
		for _, v := range soldierMap {
			if v.Id == id {
				rarity = v.Rarity
				break
			}
		}
		c.String(http.StatusOK, "士兵的稀有度:"+rarity)
	})
}

func GetSoldierByUnlockArena(r *gin.Engine, url string, soldierMap map[string]SoldierDTO) {
	r.GET(url, func(c *gin.Context) {
		ret := make(map[string][]SoldierDTO)
		for _, v := range soldierMap {
			unlockArena := v.UnlockArena
			if i := ret[unlockArena]; i == nil {
				i = make([]SoldierDTO, 0)
			}
			ret[unlockArena] = append(ret[unlockArena], v)
		}
		c.JSON(http.StatusOK, ret)
	})
}
