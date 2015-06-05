package main

import (
	"fmt"
	"runtime/debug"
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
	go Recover(disaster, "just wrongness")
	fmt.Printf("Sleeping")
	time.Sleep(5)

	fmt.Printf("All done")
}

func Recover(g func(), message string) {
	defer func() {
		fmt.Printf("[%s] DONE Recoverable processing", message)
		if x := recover(); x != nil {
			fmt.Printf("[%s] PANIC - recovered run time panic: %v. Stacktrace: %s", message, x, string(debug.Stack()))
		}
	}()
	fmt.Printf("[%s] START Recoverable processing", message)
	g()
}
