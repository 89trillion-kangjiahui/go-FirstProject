package route

import (
	"github.com/gin-gonic/gin"

	"project2/controller"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/cal/getRet", controller.ComputeRet)
	return r
}
