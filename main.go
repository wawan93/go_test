package main

import "fmt"

func main() {
	fmt.Print(greeting("wawan"))
}

func greeting(name string) string {
	return "Hello, " + name + "!\n"
}