package main

import (
	"github.com/gin-gonic/gin"
	. "project1/controller"
	. "project1/entity"
	"project1/util"
)

var SoldierMap = make(map[string]Soldier)

func main()  {
	//读取ini配置文件
	conf := util.InitConfig("./config/app.ini")
	httpPort := conf.GetValue("server","HttpPort")
	//从命令行中读取json的配置文件
	jsonPath := util.GetJsonPath()
	//将新的数据转化为有用的数据放到新的json文件中
	newSoldierMap := *(util.JsonToFile(jsonPath, &SoldierMap))

	//绑定路由
	r := gin.Default()

	GetAllByRarity(r, newSoldierMap)
	GetRarityById(r, newSoldierMap)
	GetAckById(r, newSoldierMap)
	GetSoldierByUnlockArena(r, newSoldierMap)

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":" + httpPort)
}









