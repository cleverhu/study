package main

import (
	"jtthinkStudy/wsTest/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", handlers.Echo)
	http.HandleFunc("/sendAllPod", handlers.SendAllPod)
	//http.HandleFunc("/test", func(w http.ResponseWriter, req *http.Request) {
	//	m:=make(map[string]interface{})
	//	json.Unmarshal(req.Body,m)
	//
	//})
	http.ListenAndServe(":8080", nil)
}
