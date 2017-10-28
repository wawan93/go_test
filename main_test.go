package main

import "testing"

type test struct {
	data     string
	expected string
	actual   string
}

func TestGreeting(t *testing.T) {
	tests := []test{
		{
			data:     "Ivan",
			expected: "Hello, Ivan!\n",
		},
		{
			data:     "wawan",
			expected: "Hello, wawan!\n",
		},
		{
			data:     "товарищ майор",
			expected: "Hello, товарищ майор!\n",
		},
	}

	for _, test := range tests {
		test.actual = greeting(test.data)

		if test.actual != test.expected {
			t.Error("Неверный результат", test.expected, test.actual)
		}
	}
}

func ExampleMain() {
	main()
	// Output:
	// Hello, wawan!
	//
}
