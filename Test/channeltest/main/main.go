package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

var channel chan int

func pl() {
	for {
		select {
		case i := <-channel:
			fmt.Printf("i: %v\n", i)
		}
	}
}

func main() {
	go pl()

	channel = make(chan int, 0)

	for i := 0; i < 10; i++ {
		channel <- i
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	s := <-c

	fmt.Printf("server exit because get signal: %v", s)
}
