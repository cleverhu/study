package main

import (
	"github.com/gin-gonic/gin"
	"jtthinkStudy/casbinTest/middlewares"
	"log"
)

func main() {
	r := gin.New()
	r.Use(middlewares.Middlewares()...)
	r.GET("/depts", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "部门列表"})
	})

	r.GET("/depts/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "部门详情"})
	})

	r.POST("/depts", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "新增部门"})
	})

	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "会员列表"})
	})

	r.GET("/users/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "会员详情"})
	})

	r.POST("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "新增会员"})
	})

	log.Fatal(r.Run(":8888"))
}
