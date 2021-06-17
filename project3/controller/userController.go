package controller

import (
	"project3/service"

	"github.com/gin-gonic/gin"
	"net/http"
)

func UserCheckCode(r *gin.Engine, url string, userMap map[string]string) {
	r.GET(url, func(c *gin.Context) {
		uid := c.Query("uid")
		code := c.Query("code")
		ret, serviceEro := service.CheckCodeService(uid, code, userMap)
		if serviceEro != nil {
			c.JSON(http.StatusInternalServerError, serviceEro.Error())
		}
		c.JSON(http.StatusOK, ret)
	})
}
