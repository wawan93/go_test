package main

import "fmt"

func main() {
	fmt.Println(greeting("wawan"))
}

func greeting(name string) string {
	return "Hello, " + name + "!"
}

