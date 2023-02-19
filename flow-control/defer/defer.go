package main

import "fmt"

func main() {
	// A defer statement defers the execution of a function
	// until the surrounding function returns.
	defer fmt.Println("defered hello")
	fmt.Println("hello")
}
