package main

import (
	"fmt"

	"modules-and-tests/hello"
	"modules-and-tests/sum"
)

func main() {
	fmt.Printf("Hello message: %s\n", hello.Hello())
	fmt.Printf("Sum: %d\n", sum.Add(2, 2))
}
