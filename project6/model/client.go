package model

import (
	"github.com/gorilla/websocket"
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
