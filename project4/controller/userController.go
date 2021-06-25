package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"project3/entity"
	"project3/service"
)

func UserCheckCode(r *gin.Engine, url string) {
	r.GET(url, func(c *gin.Context) {
		uid := c.Query("uid")
		code := c.Query("code")
		data := service.CheckCodeService(uid, code)
		c.String(http.StatusOK, "%s", data)
	})
}

func UserLogin(r *gin.Engine, url string) {
	r.GET(url, func(c *gin.Context) {
		uid := c.Query("uid")
		user := service.UserLoginService(uid)
		if user == nil {
			//用户未注册,跳到用户注册页面注册
			c.Redirect(http.StatusMovedPermanently, "http://localhost:8000/user/register")
		}
		if user != nil {
			ret := entity.SetResult(200, "请求成功", *user)
			c.JSON(http.StatusOK, ret)
		}
	})
}

func UserRegister(r *gin.Engine, url string) {
	r.GET(url, func(c *gin.Context) {
		username := c.Query("username")
		retCode, msg, data := service.UserRegisterService(username)
		ret := entity.SetResult(retCode, msg, data)
		c.JSON(http.StatusOK, ret)
	})
}
