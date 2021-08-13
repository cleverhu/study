package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"jtthinkStudy/myredis/lib"
	"jtthinkStudy/myredis/models"
	"jtthinkStudy/myredis/redisLocker"
	"log"
	"time"
)

func main() {

	a := 0
	r := gin.New()
	r.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(200, gin.H{"message": err})
				//c.AbortWithStatusJSON(400, gin.H{"message": err})
			}
		}()
		c.Next()
	})
	r.GET("/item/:id", func(c *gin.Context) {
		id := c.Param("id")
		cache := lib.NewSimpleCache()
		defer lib.ReleaseSimpleCache(cache)
		cache.DBGetter = lib.DBGetter(id)
		ret := cache.GetCache(id).(string)
		item := models.NewItemModel()
		_ = json.Unmarshal([]byte(ret), item)
		c.JSON(200, gin.H{"code": "success", "result": item})
	})
	r.GET("/test", func(c *gin.Context) {
		a++
		lock := redisLocker.NewLocker("lock1").Lock()
		defer lock.UnLock()
		if c.Query("t") != "" {
			time.Sleep(5 * time.Second)
		}
		c.JSON(200, gin.H{"result": a})
	})
	log.Fatal(r.Run(":8899"))
}
