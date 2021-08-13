package handlers

import (
	"jtthinkStudy/wsTest/core"
	"log"
	"net/http"
)

func Echo(w http.ResponseWriter, req *http.Request) {
	conn, err := core.Upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println(err)
		return
	}
	core.ClientMap.Store(conn)
	//for {
	//	err = conn.WriteMessage(websocket.TextMessage, []byte("hello"))
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	} else {
	//
	//		time.Sleep(2 * time.Second)
	//	}
	//
	//}

}
