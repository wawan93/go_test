package main

import (
	"testing"
	"fmt"
	"time"
)

type test struct {
	data     string
	expected string
	actual   string
}

func TestGreeting(t *testing.T) {
	tests := []test{
		{
			data:     "Ivan",
			expected: "Hello, Ivan!",
		},
		{
			data:     "wawan",
			expected: "Hello, wawan!",
		},
		{
			data:     "товарищ майор",
			expected: "Hello, товарищ майор!",
		},
	}

	for _, test := range tests {
		test.actual = greeting(test.data)

		if test.actual != test.expected {
			t.Error("Неверный результат", test)
		}
	}
}

func ExampleMain() {
	main()
	// Output:
	// Hello, wawan!
}

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

	for i := 100; i < 105; i++ {
		mainChannel <- i
	}
}
