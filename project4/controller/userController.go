package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"project3/entity"
	"project3/service"
)

func UserCheckCode(c *gin.Context) {
	uid := c.Query("uid")
	code := c.Query("code")
	if uid == "" {
		ret := entity.SetResult(4001, "uid不能为空", nil)
		c.JSON(http.StatusOK, ret)
	} else if code == "" {
		ret := entity.SetResult(4002, "code不能为空", nil)
		c.JSON(http.StatusOK, ret)
	} else {
		data := service.CheckCodeService(uid, code)
		c.String(http.StatusOK, "%s", data)
	}
}

func UserLogin(c *gin.Context) {
	uid := c.Query("uid")
	if uid == "" {
		ret := entity.SetResult(4001, "uid不能为空", nil)
		c.JSON(http.StatusOK, ret)
	} else {
		user := service.UserLoginService(uid)
		if user == nil {
			//用户未注册,跳到用户注册页面注册
			c.Redirect(http.StatusMovedPermanently, "http://localhost:8000/user/register")
		} else {
			ret := entity.SetResult(200, "请求成功", *user)
			c.JSON(http.StatusOK, ret)
		}
	}
}

func UserRegister(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		ret := entity.SetResult(4003, "username不能为空", nil)
		c.JSON(http.StatusOK, ret)
	} else {
		retCode, msg, data := service.UserRegisterService(username)
		ret := entity.SetResult(retCode, msg, data)
		c.JSON(http.StatusOK, ret)
	}
}
