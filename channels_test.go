package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChannels(t *testing.T) {
	mainChannel := make(chan int)

	// fatal error: all goroutines are asleep - deadlock!
	//mainChannel <- 100
	//mainChannel <- 101

	go func(c chan int) {
		for {
			fmt.Println(<-c)
			time.Sleep(time.Second)
		}
	}(mainChannel)

	mainChannel <- 100
	mainChannel <- 101
}

func TestBufferedChannels(t *testing.T) {
	ch := make(chan int, 2)

	ch <- 100
	ch <- 101

	go func(c chan int) {
		for {
			fmt.Println(<-c)
			time.Sleep(time.Second)
		}
	}(ch)
}

