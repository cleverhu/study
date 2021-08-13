package main

import (
	"fmt"
	"sync"
)

type user struct {
	ID   int
	Data sync.Map
}

func (u user) change() {
	fmt.Printf("%p,%v\n", &u, u)
	u.ID = 123
	u.Data.Store(1, 2)
}

func (u user) del() {
	fmt.Printf("%p,%v\n", &u, u)
	u.Data.Delete(1)
}

func main() {
	u := user{}
	fmt.Printf("%p\n", &u)
	u.change()
	fmt.Println(u)
	u.del()
	fmt.Println(u)
}
