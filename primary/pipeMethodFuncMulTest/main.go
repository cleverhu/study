package main

import (
	"fmt"
	"sync"
	"time"
)

type Cmd func([]int) chan int

type PipeCmd func(chan int) chan int

func Even(array []int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range array {
			if num%2 == 0 {
				out <- num
			}
		}
	}()
	return out
}

func M10(in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num * 10
		}
	}()
	return out
}

func PipeFunc(array []int, c1 Cmd, c2 PipeCmd, num int) chan int {
	ret := c1(array)
	wg := sync.WaitGroup{}
	out := make(chan int)
	for i := 0; i < num; i++ {
		getChan := c2(ret)
		wg.Add(1)
		go func(in chan int) {
			defer wg.Done()
			for v := range in {
				out <- v
			}

		}(getChan)
	}
	go func() {
		defer close(out)
		wg.Wait()

	}()
	return out
}

func main() {
	array := []int{2, 4, 5, 6, 7, 8}
	resultChan := PipeFunc(array, Even, M10, 5)

	for num := range resultChan {
		fmt.Println(num)
		time.Sleep(1 * time.Second)
	}
}
