package main

import (
	"fmt"
	"time"
)

type T struct {
	s *string
}

func disaster() {
	fmt.Println("[warning] Attempting to do something bad")
	t := &T{}
	s := *t.s
	fmt.Printf("s is: %s\n", s)
}

func main() {
	go disaster()
	fmt.Println("Sleeping")
	time.Sleep(5000)

	fmt.Println("All done")
}
