package main

import "fmt"

type Person struct {
	name string
}

/*
	We can declare methods with pointer receivers.

With a value receiver, the method operates on a copy of the original value.
(This is the same behavior as for any other function argument)
The method must have a pointer receiver to change the value declared in the main function
*/
func (p *Person) changeName(name string) {
	p.name = name
}

// func (*Person)... // it works too
func (Person) hello() string {
	return "Hello"
}

// with value receiver
// func (p Person) ...

func main() {
	p1 := Person{"Pedro"}
	fmt.Println(p1.name)

	// (&p1).changeName("Bob")
	p1.changeName("Bob") // Methods and pointer indirection for convenience
	fmt.Println(p1.name)

	pointer := &p1

	// result := (*pointer).hello()
	result := pointer.hello() // Methods and pointer indirection for convenience

	fmt.Println(result)
}
