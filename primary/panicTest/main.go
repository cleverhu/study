package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		for {
			test()
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(1 * time.Minute)
}

func test() {

	for {
		b := 0
		fmt.Println(1 / b)
		fmt.Println(1 / b)
		fmt.Println(1 / b)

	}

}
