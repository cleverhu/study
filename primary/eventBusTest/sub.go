package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"jtthinkStudy/primary/eventBusTest/models"
	"log"
)

func main() {
	client, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	client.Subscribe("users.get.score", func(msg *nats.Msg) {
		go func(msg *nats.Msg) {
			u := models.NewUserModel()
			_ = u.Decode(msg.Data)
			u.Score = 999
			b, _ := u.Encode()
			client.Publish(msg.Reply, b)
		}(msg)
	})

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.Run(":8081")
}
