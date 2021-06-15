package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	. "project1/entity"
	"project1/util"
)

var soldierMap = make(map[string]Soldier)

func main()  {
	//读取ini配置文件
	conf := util.InitConfig("./config/app.ini")
	httpPort := conf.GetValue("server","HttpPort")
	//从命令行中读取json的配置文件
	jsonPath := util.GetJsonPath()
	data,error := ioutil.ReadFile(jsonPath)
	if error != nil {
		fmt.Println("文件读取失败")
		return
	}
	if ero := json.Unmarshal(data, &soldierMap); ero != nil {
		fmt.Println("json解析出错了")
		return
	}
	//将新的数据转化为有用的数据放到新的json文件中
	util.JsonToFile(soldierMap)

	//绑定路由
	r := gin.Default()

	GetAllByRarity(r)
	GetRarityById(r)
	GetAckById(r)
	GetSoldierByUnlockArena(r)

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":" + httpPort)
}

func GetAllByRarity(r *gin.Engine)  {
	r.GET("/soldier/getAll", func(c *gin.Context) {
		rarity := c.Query("rarity")
		unlockArena := c.Query("unlockArena")
		returnData := make([] Soldier, 0)
		for _, v := range soldierMap{
			if v.Rarity == rarity && v.UnlockArena == unlockArena{
				returnData = append(returnData, v)
			}
		}

		c.JSON(http.StatusOK, returnData)
	})
}

func GetRarityById(r *gin.Engine)  {
	// gin.Context，封装了request和response
	r.GET("/soldier/getRarity", func(c *gin.Context) {
		id := c.Query("id")
		var rarity string
		for _,v := range soldierMap{
			if v.Id == id{
				rarity = v.Rarity
				break
			}
		}
		c.String(http.StatusOK, "士兵的稀有度:" + rarity)
	})
}

func GetAckById(r *gin.Engine){
	// gin.Context，封装了request和response
	r.GET("/soldier/atc", func(c *gin.Context) {
		id := c.Query("id")
		var atc string
		for _,v := range soldierMap{
			if v.Id == id{
				atc = v.Atk
				break
			}
		}
		c.String(http.StatusOK, "士兵的战斗力:" + atc)
	})
}


func GetSoldierByUnlockArena(r *gin.Engine)  {
	r.GET("/soldier/getAll/unlockArena", func(c *gin.Context) {
		ret := make(map[string][]Soldier)
		for _, v := range soldierMap{
			unlockArena := v.UnlockArena
			if i := ret[unlockArena]; i == nil {
				i = make([]Soldier, 0)
			}
			ret[unlockArena] = append(ret[unlockArena], v)
		}
		c.JSON(http.StatusOK, ret)
	})
}

