package test

import (
	"fmt"
	"testing"

	"github.com/gorilla/websocket"

	"project6/global"
)

func TestConnection(t *testing.T) {
	global.WS, _, _ = websocket.DefaultDialer.Dial(global.URL, nil)
	if global.WS == nil {
		fmt.Println("连接失败")
	} else {
		fmt.Println("连接成功")
	}
}
