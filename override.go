package main

import "github.com/wawan93/go_test/somepkg"

type SomeNewType struct {
	somepkg.SomeType
}

func (SomeNewType) SomeFunc() string {
	return "hello, world!"
}
