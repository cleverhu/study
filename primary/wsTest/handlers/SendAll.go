package handlers

import (
	"jtthinkStudy/wsTest/core"
	"net/http"
)

func SendAllPod(w http.ResponseWriter, req *http.Request) {
	core.ClientMap.SendAllPod()
	w.Write([]byte("OK"))
}
