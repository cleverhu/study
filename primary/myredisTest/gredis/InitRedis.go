package gredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
)

var redisCli *redis.Client
var redisOnce sync.Once

func Redis() *redis.Client {
	redisOnce.Do(func() {
		redisCli = redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		})

		err := redisCli.Ping(context.Background()).Err()
		if err != nil {
			log.Fatalf("连接redis数据库失败:%v\n", err)
		}
	})
	return redisCli
}
