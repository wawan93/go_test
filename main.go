package main

func main() {
	print(greeting("wawan"))
}

func greeting(name string) string {
	return "Hello, " + name + "!\n"
}