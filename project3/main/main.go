package main

import (
	"github.com/gin-gonic/gin"

	"project3/controller"
)

var UserMap = make(map[string]string)

const (
	ACreateURL    = "/admin/create"
	ASelectURL    = "/admin/select"
	UCheckCodeURL = "/user/check"
)

func main() {
	createUserMap()
	r := gin.Default()
	controller.AdminCreateCode(r, ACreateURL, UserMap)
	controller.AdminSelectCode(r, ASelectURL)
	controller.UserCheckCode(r, UCheckCodeURL, UserMap)
	r.Run(":8000")
}

func createUserMap() {
	UserMap["1"] = "管理员"
	UserMap["2"] = "张三"
	UserMap["3"] = "李四"
	UserMap["4"] = "王五"
	UserMap["5"] = "赵六"
}
