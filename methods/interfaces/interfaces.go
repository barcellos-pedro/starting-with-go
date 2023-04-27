package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Greet interface {
	hello() string
}

/*
Interfaces are implemented implicitly
Calling a method on an interface value executes the method of the same name on its underlying type.
*/
func (p Person) hello() string {
	return "Hello: " + p.name
}

func main() {
	// var greet Greet
	// p := Person{"pedro", 26}
	// greet = p

	var greet Greet = Person{"pedro", 26}
	fmt.Println(greet.hello())
	describe(greet)
}

func describe(i Greet) {
	fmt.Printf("(%v, %T)\n", i, i)
}
