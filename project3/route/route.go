package route

import (
	"github.com/gin-gonic/gin"

	"project3/controller"
)

func RegisterRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/admin/create", controller.AdminCreateCode)
	r.GET("/admin/select", controller.AdminSelectCode)
	r.GET("/user/check", controller.UserCheckCode)
	return r
}
