package service

import (
	"fyne.io/fyne/widget"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"

	"project6/global"
	"project6/response"
)

//用户开始连接向服务端发送消息
func LoginService(nameEntry *widget.Entry) {
	global.WS, _, _ = websocket.DefaultDialer.Dial(global.URL, nil)
	global.Connection.Ws = global.WS
	if global.UserName == "" {
		global.UserName = nameEntry.Text
	}
	go global.Connection.Write()
	go global.Connection.Read()

	data1 := response.Data{
		Type:    "login",
		Content: nameEntry.Text,
	}
	data2 := response.Data{
		Type: "user_list",
	}
	data_res1, _ := proto.Marshal(&data1)
	data_res2, _ := proto.Marshal(&data2)
	//发送登录的消息
	global.Connection.WriteChan <- data_res1
	//发送获取用户列表的消息
	global.Connection.WriteChan <- data_res2
}

//用户断开连接向服务端发送消息
func ExitService() {
	data1 := response.Data{
		Type: "exit",
		User: global.UserName,
	}
	data2 := response.Data{
		Type: "user_list",
	}
	data_res1, _ := proto.Marshal(&data1)
	data_res2, _ := proto.Marshal(&data2)
	global.UserName = ""
	//发送退出连接的消息
	global.Connection.WriteChan <- data_res1
	//发送获取用户列表的消息
	global.Connection.WriteChan <- data_res2
}

//发送用户说话的消息
func SendService(multiEntry *widget.Entry) {
	data := response.Data{
		Type:    "talk",
		User:    global.UserName,
		From:    global.UserName,
		Content: multiEntry.Text,
	}
	data_b, _ := proto.Marshal(&data)
	global.Connection.WriteChan <- data_b
}
