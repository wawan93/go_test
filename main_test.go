package main

import "testing"

type test struct {
	data     string
	expected string
	actual   string
}

func Test(t *testing.T) {
	tests := []test{
		{
			data:     "Ivan",
			expected: "Hello, Ivan!\n",
		},
		{
			data:     "wawan",
			expected: "Hello, wawan!\n",
		},
	}

	for _, test := range tests {
		test.actual = greeting(test.data)

		if test.actual != test.expected {
			t.Error("Неверный результат", test.expected, test.actual)
		}
	}
}