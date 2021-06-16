package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "project1/entity"
)

func GetAckById(r *gin.Engine, soldierMap map[string]SoldierDTO){
	// gin.Context，封装了request和response
	r.GET("/soldier/atc", func(c *gin.Context) {
		id := c.Query("id")
		var atc string
		for _,v := range soldierMap{
			if v.Id == id{
				atc = v.Atk
				break
			}
		}
		c.String(http.StatusOK, "士兵的战斗力:" + atc)
	})
}
