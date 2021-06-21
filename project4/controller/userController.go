package controller

import (
	"errors"

	"project3/service"

	"github.com/gin-gonic/gin"
	"net/http"
)

func UserCheckCode(r *gin.Engine, url string) {
	r.GET(url, func(c *gin.Context) {
		uid := c.Query("uid")
		code := c.Query("code")
		data, ero := service.CheckCodeService(uid, code)
		if ero != nil {
			c.JSON(http.StatusInternalServerError, data)
		}
		if ero == nil {
			c.JSON(http.StatusOK, data)
		}
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
			c.JSON(http.StatusOK, *user)
		}
	})
}

func UserRegister(r *gin.Engine, url string) {
	r.GET(url, func(c *gin.Context) {
		username := c.Query("username")
		uid, ero := service.UserRegisterService(username)
		if uid == "" {
			c.String(http.StatusOK, "请输入用户名")
		}
		if ero != nil {
			c.JSON(http.StatusInternalServerError, errors.New("用户注册失败"))
		}
		if ero == nil {
			c.JSON(http.StatusOK, uid)
		}
	})
}
