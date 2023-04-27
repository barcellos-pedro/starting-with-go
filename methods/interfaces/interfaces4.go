package main

import "fmt"

/*
The interface type that specifies zero methods is known as the empty interface
Empty interfaces are used by code that handles values of unknown type.
*/
func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
