package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(200 * time.Millisecond)
	boom := time.After(1000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("hello.")
		case <-boom:
			fmt.Println("boom .")
			return
		default:
			fmt.Println("     .")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
