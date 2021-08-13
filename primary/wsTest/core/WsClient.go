package core

import (
	"fmt"
	"github.com/gorilla/websocket"
	"jtthinkStudy/wsTest/models"
	"time"
)

type wsClient struct {
	conn      *websocket.Conn
	readChan  chan *WsMessage
	writeChan chan *models.WsResponse
	closeChan chan int
}

func newWsClient(conn *websocket.Conn) *wsClient {
	return &wsClient{conn: conn, writeChan: make(chan *models.WsResponse), readChan: make(chan *WsMessage), closeChan: make(chan int)}
}

func (this *wsClient) Ping(delay time.Duration) {
	for {
		time.Sleep(delay)
		err := this.conn.WriteMessage(websocket.TextMessage, []byte("ping"))
		if err != nil {
			ClientMap.Remove(this.conn)
			return
		}
	}
}

func (this *wsClient) WriteLoop() {
	for {
		select {
		case msg := <-this.writeChan:
			err := this.conn.WriteMessage(websocket.TextMessage, msg.ToJson())
			if err != nil {
				ClientMap.Remove(this.conn)
				this.closeChan <- 1
				return
			}
		}
	}
}

func (this *wsClient) ReadLoop() {
	for {
		t, data, err := this.conn.ReadMessage()
		if err != nil {
			ClientMap.Remove(this.conn)
			this.closeChan <- 1
			return
		}
		this.readChan <- NewWsMessage(t, data)
	}
}

func (this *wsClient) HandleLoop() {
	for {
		select {
		case msg := <-this.readChan:
			rsp, err := msg.parseForCmd()
			if err != nil {
				fmt.Println(err)
			}
			if rsp != nil {
				this.writeChan <- rsp
			}
		case <-this.closeChan:
			return
		}
	}
}
