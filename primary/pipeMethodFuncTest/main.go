package main

import (
	"fmt"
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

func M2(in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			out <- num * 2
		}
	}()
	return out
}

func PipeFunc(array []int, c1 Cmd, c2 ...PipeCmd) chan int {
	result := c1(array)
	for _, f := range c2 {
		result = f(result)
	}
	return result
}

func main() {
	array := []int{2, 4, 5, 6, 7, 8}
	resultChan := PipeFunc(array, Even, M10, M10, M10, M10)
	for num := range resultChan {
		fmt.Println(num)
		time.Sleep(1 * time.Second)
	}
}
