package main

import "testing"

func TestOverride(t *testing.T) {
	st := new(SomeNewType)

	cases := []struct{
		fn func() string
		expected string
	}{
		{st.SomeFunc, "hello, world!"},
		{st.SomeType.SomeFunc, "hello"},
	}

	for _, tc := range cases {
		got := tc.fn()
		if tc.expected != got {
			t.Errorf("expected \"%s\" got \"%s\"", tc.expected, got)
		}
	}
}