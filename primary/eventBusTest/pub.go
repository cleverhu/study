package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"jtthinkStudy/primary/eventBusTest/models"
	"log"
	"strconv"
	"time"
)

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	client, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	r := gin.Default()
	r.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		idInt, _ := strconv.Atoi(id)
		u := models.NewUserModel(models.WithUserID(int32(idInt)), models.WithUserName("xiaohu"))
		encodeUser, err := u.Encode()
		handleError(err)
		msg, err := client.Request("users.get.score", encodeUser, time.Millisecond*500)
		handleError(err)
		err = u.Decode(msg.Data)
		handleError(err)
		c.JSON(200, u)
	})
	r.Run(":8080")

}
