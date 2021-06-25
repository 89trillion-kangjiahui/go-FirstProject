package route

import (
	"github.com/gin-gonic/gin"

	"project1/controller"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/soldier/getAll", controller.GetAllByRarity)
	r.GET("/soldier/getRarity", controller.GetRarityById)
	r.GET("/soldier/getAtk", controller.GetAckById)
	r.GET("/soldier/getByCvc", controller.GetSoldierByCvc)
	r.GET("/soldier/getByUnlockArena", controller.GetSoldierByUnlockArena)
	return r
}
