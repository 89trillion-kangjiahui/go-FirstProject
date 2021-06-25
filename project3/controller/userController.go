package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"project3/entity"
	"project3/global"
	"project3/service"
)

// 用户获取礼品包
func UserCheckCode(c *gin.Context) {
	uid := c.Query("uid")
	code := c.Query("code")
	if uid == "" {
		ret := entity.SetResult(4000, "uid不能为空", "")
		c.JSON(http.StatusOK, ret)
	} else if code == "" {
		ret := entity.SetResult(4001, "code不能为空", "")
		c.JSON(http.StatusOK, ret)
	} else {
		retCode, data, serviceEro := service.CheckCodeService(uid, code, global.UserMap)
		ret := entity.SetResult(retCode, serviceEro, data)
		c.JSON(http.StatusOK, ret)
	}
}
