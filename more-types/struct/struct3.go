package main

import "fmt"

type Person struct {
	name string
}

/* Here the change stays inside the func */
func rename(person Person) {
	person.name = "test"
}

/* Here we mutate the atrr inside the function */
func renameV2(person *Person) {
	// (*person).name = "" // longer way to access struct pointer fields
	person.name = "test"
}

/* One good reason to use pointers instead of pass-by-values */
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
}
