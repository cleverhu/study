package main

import "fmt"

func main() {
	a := []int{1, 2, 3,4}
	b := make([]int,cap(a))
	copy(b,a)

	fmt.Printf("%p,%p\n", a, b)
	fmt.Println(a,b)
}
