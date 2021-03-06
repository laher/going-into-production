package main

import (
	"fmt"
	"time"
)

func fibonacci(n int) []int {
	x, y := 0, 1
	res := []int{}
	for i := 0; i < n; i++ {
		time.Sleep(300 * time.Millisecond)
		res = append(res, x)
		x, y = y, x+y
	}
	return res
}

func main() {
	c := fibonacci(10)
	for _, i := range c {
		fmt.Println(i)
	}
}
