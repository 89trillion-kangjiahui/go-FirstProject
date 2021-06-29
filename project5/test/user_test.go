package test

import (
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"project5/response"
)

type Client struct {
	Ws        *websocket.Conn
	WriteChan chan []byte
	ReadChan  chan []byte
}

func (this *Client) Write() {
	for message := range this.WriteChan {
		this.Ws.WriteMessage(websocket.TextMessage, message)
	}
	this.Ws.Close()
}

func (this *Client) Read() {
	defer func() {
		this.Ws.Close()
	}()
	for {
		_, data, err := this.Ws.ReadMessage()
		if err != nil {
			break
		}
		this.ReadChan <- data
	}
}

func TestUser(t *testing.T) {
	ws, _, _ := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	client := &Client{
		Ws: ws,
		WriteChan: make(chan []byte),
		ReadChan:  make(chan []byte),
	}
	go client.Write()
	go client.Read()
	data1 := response.Data{
		Type:    "login",
		Content: "jj",
	}
	data2 := response.Data{
		Type:    "user_list",
	}
	data3 := response.Data{
		Type: "talk",
		Content: "nihao",
	}
	data4 := response.Data{
		Type: "exit",
		User: "jj",
	}
	data_res1, _ := proto.Marshal(&data1)
	data_res2, _ := proto.Marshal(&data2)
	data_res3, _ := proto.Marshal(&data3)
	data_res4, _ := proto.Marshal(&data4)
	client.WriteChan <- data_res1
	client.WriteChan <- data_res2
	client.WriteChan <- data_res3
	client.WriteChan <- data_res4
	go func() {
		for {
			select {
			case c := <-client.ReadChan:
				var ret response.Data
				proto.Unmarshal(c, &ret)
				if ret.Type == "login" {
					fmt.Println("登录成功")
					fmt.Println("消息内容:", ret.Content)
				}else if ret.Type == "user_list"{
					fmt.Println("获取用户列表")
					fmt.Println("用户列表:", ret.Userlist)
				}else if ret.Type == "exit"{
					fmt.Println("退出成功")
					fmt.Println("消息内容:", ret.Content)
				}else if ret.Type == "talk"{
					fmt.Println("获取用户说话内容")
					fmt.Println("消息内容:", ret.Content)
				}
			}
		}
	}()
}
