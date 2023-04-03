package main

import "fmt"

func helloWorld(msg string, fn func(string) string) string {
	return fn(msg)
}

func main() {
	hello := func(msg string) string {
		return "Hello " + msg
	}

	fmt.Println(hello("Pedro"))
	fmt.Println(helloWorld("Ana", hello))
}
