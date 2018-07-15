package main

import (
	"testing"
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
