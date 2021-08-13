package main

import (
	"jtthinkStudy/rateLimitTest/lib"
	"time"
)

func main() {
	cache := lib.NewGCache(lib.WithMaxSize(5))
	cache.Set("age", "19", 5*time.Second)
	cache.Set("name", "test", 0*time.Second)
	cache.Set("salary", "10000", 5*time.Second)
	//fmt.Println(cache.Get("age"))
	cache.Print()

}
