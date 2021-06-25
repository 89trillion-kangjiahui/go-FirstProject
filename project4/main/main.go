package main

import (
	"github.com/gin-gonic/gin"

	"project3/controller"
)

const (
	ACreateURL    = "/admin/create"
	ASelectURL    = "/admin/select"
	UCheckCodeURL = "/user/check"
	UserLogin     = "/user/login"
	UserRegister  = "/user/register"
)

func main() {
	r := gin.Default()
	controller.AdminCreateCode(r, ACreateURL)
	controller.AdminSelectCode(r, ASelectURL)
	controller.UserCheckCode(r, UCheckCodeURL)
	controller.UserLogin(r, UserLogin)
	controller.UserRegister(r, UserRegister)
	r.Run(":8000")
}
