package main

import (
	"fmt"
	"sync"
)

type user struct {
	ID int
}

var lock sync.Mutex
var m sync.Map
var u *user
var u1 user
var nm map[string]string

func main() {
	lock.Lock()
	lock.Unlock()
	m.Store(1, 2)
	u.ID = 123
	u1.ID = 123
	fmt.Printf("%p\n", nm)
}
