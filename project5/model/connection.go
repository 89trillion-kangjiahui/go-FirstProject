package model

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"project5/response"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

//抽象出需要的数据结构
//ws连接器
type Connection struct {
	//ws 连接器
	Ws *websocket.Conn
	//管道
	Send chan []byte
	//数据
	Data *response.Data
}

//先实现的读和写
func (c *Connection) writer() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Ws.Close()
	}()

	//从管道遍历数据
	for {
		//数据写出
		select {
		case message, ok := <-c.Send:
			c.Ws.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.Ws.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := c.Ws.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			n := len(c.Send)
			for i := 0; i < n; i++ {
				w.Write([]byte{'\n'})
				w.Write(<-c.Send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.Ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Ws.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

var user_list = []string{}

//读数据
func (c *Connection) reader() {
	defer func() {
		c.Ws.Close()
	}()
	c.Ws.SetReadLimit(maxMessageSize)
	c.Ws.SetReadDeadline(time.Now().Add(pongWait))
	c.Ws.SetPongHandler(func(string) error {
		c.Ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	//不断的读数据
	for {
		_, message, err := c.Ws.ReadMessage()
		if err != nil {
			// 读不进数据，将用户移除
			H.Unregister <- c
			break
		}
		//读取数据
		proto.Unmarshal(message, c.Data)
		//根据data的type判断该做什么
		switch c.Data.Type {
		case "login":
			//弹出窗口，输入用户名
			c.Data.Ip = c.Ws.RemoteAddr().String()
			c.Data.User = c.Data.Content
			c.Data.From = c.Data.User
			//登录后，将用户加入到用户列表
			user_list = append(user_list, c.Data.User)
			//每个用户都加载所有的登录的列表中
			c.Data.Userlist = user_list
			c.Data.Content = "system message:" + c.Data.User + ",up!"
			//将数据序列化
			data_b, _ := proto.Marshal(c.Data)
			H.Broadcast <- data_b
		case "talk":
			//用户聊天
			c.Data.Type = "talk"
			c.Data.Content = c.Data.User + ":" + c.Data.Content
			data_b, _ := proto.Marshal(c.Data)
			H.Broadcast <- data_b
		case "exit":
			c.Data.Type = "exit"
			//用户列表删除
			user_list = remove(user_list, c.Data.User)
			c.Data.Userlist = user_list
			c.Data.Content = c.Data.User
			c.Data.Content = "system message:" + c.Data.User + ",down!"
			data_b, _ := proto.Marshal(c.Data)
			H.Broadcast <- data_b
			H.Unregister <- c
		case "user_list":
			//获取用户列表
			c.Data.Type = "user_list"
			c.Data.Userlist = user_list
			data_b, _ := proto.Marshal(c.Data)
			H.Broadcast <- data_b
		default:
			fmt.Println("其他")
		}
	}
}

//删除用户切片数据
func remove(slice []string, user string) []string {
	//严谨判断
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}
	//定义新的返回切片
	var my_slice = []string{}
	//删除传入切片中的指定用户，其他用户放到新的切片
	for i := range slice {
		//利用索引删除用户
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			my_slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return my_slice
}

//定义一个升级器
var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	}}

// ws的回调函数
func WsHandler(w http.ResponseWriter, r *http.Request) {
	//1.获取ws对象
	ws, ero := upgrader.Upgrade(w, r, nil)
	if ero != nil {
		return
	}
	//创建链接对象做事
	//初始化连接对象
	c := &Connection{
		Ws:   ws,
		Send: make(chan []byte, 128),
		Data: &response.Data{},
	}
	H.Register <- c
	// ws将数据读写跑起来
	go c.writer()
	go c.reader()
}
