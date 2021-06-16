package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "project1/entity"
)

func GetRarityById(r *gin.Engine, soldierMap map[string]SoldierDTO)  {
	// gin.Context，封装了request和response
	r.GET("/soldier/getRarity", func(c *gin.Context) {
		id := c.Query("id")
		var rarity string
		for _,v := range soldierMap{
			if v.Id == id{
				rarity = v.Rarity
				break
			}
		}
		c.String(http.StatusOK, "士兵的稀有度:" + rarity)
	})
}