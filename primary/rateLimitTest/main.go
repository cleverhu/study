package main

import (
	"github.com/gin-gonic/gin"
	"jtthinkStudy/rateLimitTest/lib"
	"log"
)

func test(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}

func main() {

	r := gin.New()
	r.Use(func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				c.AbortWithStatusJSON(400, gin.H{"message": e})
			}
		}()
		c.Next()
	})
	r.GET("/", lib.IpLimiter(5, 1)(test))

	log.Fatal(r.Run(":8080"))

}
