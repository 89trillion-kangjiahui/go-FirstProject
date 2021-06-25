package main

import (
	"github.com/robfig/config"

	"project1/global"
	"project1/route"
	"project1/util"
)

func main() {
	//读取ini配置文件
	c, _ := config.ReadDefault("./config/app.ini")
	httpPort, _ := c.String("server", "HttpPort")
	//从命令行中读取json的配置文件
	jsonPath := util.GetJsonPath()
	//将新的数据转化为有用的数据放到新的json文件中
	global.SoldierMap = util.JsonToFile(jsonPath)

	//绑定路由
	// 3.监听端口，默认在8080
	route.RegisterRoutes().Run(":" + httpPort)
}
