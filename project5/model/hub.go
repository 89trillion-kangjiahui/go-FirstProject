package model

import (
	"github.com/golang/protobuf/proto"

	"project5/response"
	"project5/util"
)

type Hub struct {
	//connections 注册了连接器
	Connections map[*Connection]bool
	//从连接器发送的信息
	Broadcast chan []byte
	//从连接器注册请求
	Register chan *Connection
	//销毁请求
	Unregister chan *Connection
}

//将连接器对象初始化
var H = Hub{
	Connections: make(map[*Connection]bool),
	Broadcast:   make(chan []byte),
	Register:    make(chan *Connection),
	Unregister:  make(chan *Connection),
}

//处理ws的逻辑实现
func (h *Hub) Run() {
	//监听数据管道，在后段不断处理管道数据
	for {
		//根据不同的数据管道处理不同的逻辑
		select {
		//注册
		case c := <-h.Register:
			//标识注册
			h.Connections[c] = true
		case c := <-h.Unregister:
			if _, ok := h.Connections[c]; ok {
				delete(h.Connections, c)
				close(c.Send)
			}
		case data := <-h.Broadcast:
			//将数据广播到所有用户
			var log response.Data
			proto.Unmarshal(data, &log)
			//将广播内容输出到日志
			util.PrintLog(log.User, log.Content, "./log/content.log")
			//c是具体的每个连接
			for c := range h.Connections {
				//将数据同步
				select {
				case c.Send <- data:
				default:
					//防止死循环
					delete(h.Connections, c)
					close(c.Send)
				}
			}
		}
	}
}
