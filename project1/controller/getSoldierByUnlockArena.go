package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "project1/entity"
)

func GetSoldierByUnlockArena(r *gin.Engine,soldierMap map[string]SoldierDTO)  {
	r.GET("/soldier/getAll/unlockArena", func(c *gin.Context) {
		ret := make(map[string][]SoldierDTO)
		for _, v := range soldierMap{
			unlockArena := v.UnlockArena
			if i := ret[unlockArena]; i == nil {
				i = make([]SoldierDTO, 0)
			}
			ret[unlockArena] = append(ret[unlockArena], v)
		}
		c.JSON(http.StatusOK, ret)
	})
}
