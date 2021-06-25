package main

import (
	"project3/global"
	"project3/route"
)

func main() {
	createUserMap()
	route.RegisterRoutes().Run(":8000")
}

func createUserMap() {
	global.UserMap["1"] = "管理员"
	global.UserMap["2"] = "张三"
	global.UserMap["3"] = "李四"
	global.UserMap["4"] = "王五"
	global.UserMap["5"] = "赵六"
}
