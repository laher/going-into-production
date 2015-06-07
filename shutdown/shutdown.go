package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	TIMEOUT = 5
)

func stoplistening() {
	fmt.Println("notify service to stop listening for more requests")
}

func cleanup(doneChan chan int) {
	fmt.Println("cleanup (dummy time.Sleep for illustration purposes)")
	time.Sleep(6 * time.Second) //dummy processing
	fmt.Println("cleanup done")
	doneChan <- 0
}

func main() {
	c := make(chan os.Signal, 0)
	doneChan := make(chan int)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	//when done, exit
	go func() {
		exitCode := <-doneChan
		fmt.Println("Shutting down with code", exitCode)
		os.Exit(exitCode)
	}()

	//shutdown handler
	go func() {
		<-c
		stoplistening()
		go cleanup(doneChan)
		time.Sleep(time.Duration(TIMEOUT) * time.Second)
		fmt.Println("WARN: cleanup took too long")
		doneChan <- 1
	}()

	for {
		fmt.Println("Main goroutine is sleeping...")
		time.Sleep(10 * time.Second)
	}
}
