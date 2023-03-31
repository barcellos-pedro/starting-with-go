package main

import "fmt"

type Person struct {
	name string
}

// Here the change stays inside the func
// and ends up mutating the copy
func rename(person Person) {
	person.name = "test"
}

// Here we mutate the atrr inside the function
// throug the struct pointer
func renameV2(person *Person) {
	// (*person).name = "" // longer way to access struct pointer fields
	person.name = "test2"
}

// in a simple way, similar to how append works
// we can mutate the copy and then return it
func renameV3(person Person) Person {
	person.name = "test3"
	return person
}

// One example of pointers x pass-by-values
func main() {
	p := Person{"Lala"}
	fmt.Println(p.name)
	rename(p)
	fmt.Println(p.name)

	// var p2 *Person = &Person{"peter"}
	// fmt.Println(p2.name)
	// renameV2(p2)

	p2 := Person{"peter"}
	fmt.Println(p2.name)
	renameV2(&p2)
	fmt.Println(p2.name)

	p3 := Person{"james"}
	fmt.Println(p3.name)
	p3 = renameV3(p3)
	fmt.Println(p3.name)
}
