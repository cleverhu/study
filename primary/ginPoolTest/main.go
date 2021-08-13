package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"sync"
	"time"
)

var resultPool *sync.Pool

type Result struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func (r *Result) Success(ctx *gin.Context, message interface{}) {
	r.Code = 200
	r.Message = message
	ctx.JSON(200, r)
}

func (r *Result) Error(ctx *gin.Context, message interface{}) {
	r.Code = 400
	r.Message = message
	ctx.JSON(400, r)
}

func main() {
	resultPool = &sync.Pool{New: func() interface{} {
		log.Println("create resultPool")
		return &Result{}
	}}
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		result := resultPool.Get().(*Result)
		defer resultPool.Put(result)
		if time.Now().Unix()%2 == 0 {
			result.Success(ctx, "index")
		} else {
			result.Error(ctx, "error")
		}

	})
	log.Fatal(r.Run(":8899"))
}
