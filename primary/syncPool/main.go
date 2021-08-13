package main

import (
	"fmt"
	"log"
	"sync"
)

type user struct {
	name string
}

var userPool *sync.Pool

func main() {
	userPool = &sync.Pool{New: func() interface{} {
		log.Println("create user")
		return &user{name: "张三"}
	}}
	u1 := userPool.Get()
	fmt.Println(u1)
	userPool.Put(u1)
	u2 := userPool.Get()
	fmt.Println(u2)
}
