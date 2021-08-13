package core

import (
	"github.com/gorilla/websocket"
	"jtthinkStudy/wsTest/models"
	"log"
	"sync"
	"time"
)

var ClientMap *ClientMapStruct

func init() {
	ClientMap = &ClientMapStruct{}
}

type ClientMapStruct struct {
	data sync.Map //key 是客户端ip value是websocket连接对象

}

func (this *ClientMapStruct) Store(conn *websocket.Conn) {
	wsCli := newWsClient(conn)
	this.data.Store(conn.RemoteAddr().String(), wsCli)
	go wsCli.Ping(1 * time.Second)
	go wsCli.WriteLoop()
	go wsCli.ReadLoop()
	go wsCli.HandleLoop()
}

func (this *ClientMapStruct) SendAll(msg string) {
	this.data.Range(func(key, value interface{}) bool {
		client := value.(*wsClient).conn
		err := client.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			this.Remove(client)
			log.Println(err)
		}
		return true
	})
}

func (this *ClientMapStruct) SendAllPod() {
	this.data.Range(func(key, value interface{}) bool {
		client := value.(*wsClient).conn
		err := client.WriteJSON(models.PodList())
		if err != nil {
			this.Remove(client)
			log.Println(err)
		}
		return true
	})
}

func (this *ClientMapStruct) Remove(client interface{}) {
	if _, ok := client.(*wsClient); ok {
		this.data.Delete(client.(*wsClient).conn.RemoteAddr().String())
	}

	if _, ok := client.(*websocket.Conn); ok {
		this.data.Delete(client.(*websocket.Conn).RemoteAddr().String())
	}
}
