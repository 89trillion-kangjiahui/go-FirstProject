package global

import (
	"github.com/gorilla/websocket"

	"project6/model"
)

const URL = "ws://localhost:8080/ws" //服务器地址
var (
	WS         *websocket.Conn
	IP         string
	UserName   string
	Connection = &model.Client{
		WriteChan: make(chan []byte),
		ReadChan:  make(chan []byte),
	}
)
