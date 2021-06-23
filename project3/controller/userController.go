package controller

import (
	"project3/entity"
	"project3/service"

	"github.com/gin-gonic/gin"
	"net/http"
)

func UserCheckCode(r *gin.Engine, url string, userMap map[string]string) {
	r.GET(url, func(c *gin.Context) {
		uid := c.Query("uid")
		code := c.Query("code")
		retCode, data, serviceEro := service.CheckCodeService(uid, code, userMap)
		ret := entity.SetResult(retCode, serviceEro, data)
		c.JSON(http.StatusOK, ret)
	})
}
