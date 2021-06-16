package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "project1/entity"
)

func GetAllByRarity(r *gin.Engine, soldierMap map[string]SoldierDTO)  {
	r.GET("/soldier/getAll", func(c *gin.Context) {
		rarity := c.Query("rarity")
		unlockArena := c.Query("unlockArena")
		returnData := make([] SoldierDTO, 0)
		for _, v := range soldierMap{
			if v.Rarity == rarity && v.UnlockArena == unlockArena{
				returnData = append(returnData, v)
			}
		}

		c.JSON(http.StatusOK, returnData)
	})
}
