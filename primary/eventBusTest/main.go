package main

import (
	"github.com/gin-gonic/gin"
	"jtthinkStudy/primary/eventBusTest/services"
	"time"
)

func main() {

	r := gin.New()
	r.GET("/prods1", func(c *gin.Context) {
		ch := services.GetProdListCh()
		services.Bus.Pub(services.GetProdList, ch, 1)
		services.Bus.PrintNum(services.GetProdList)
		defer services.Bus.UnSub(services.GetProdList, ch)
		c.JSON(200, ch.Data(time.Second))
	})

	r.Run(":9999")
}
