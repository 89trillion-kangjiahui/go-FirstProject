package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"project5/model"
)

func main() {
	//创建路由
	router := mux.NewRouter()
	//ws控制器不断去处理管道数据，进行同步数据
	go model.H.Run()
	//指定ws回调函数
	router.HandleFunc("/ws", model.WsHandler)
	//开启服务端监听
	http.ListenAndServe("127.0.0.1:8080", router)
}
