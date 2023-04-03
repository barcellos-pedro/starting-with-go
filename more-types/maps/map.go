package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	// var m map[string]string = ...
	// zero value o map is nil
	var m = map[string]string{"A": "first"}
	m["B"] = "second"

	fmt.Println(m)
	fmt.Println(m["A"])
	fmt.Println(m["B"])

	for i, v := range m {
		fmt.Println(i, v)
	}

	// - - - - - - - - -

	m1 := make(map[string]int)
	m1["A1"] = 1
	fmt.Println(m1)
	fmt.Println(m1["A1"])

	// - - - - - - - - -

	m2 := map[int]Person{0: {"joe", 20}}
	m2[1] = Person{"lala", 23}
	fmt.Println(m2)
	fmt.Println(m2[0])
	fmt.Println(m2[1])

	for i, v := range m2 {
		fmt.Println(i, v)
	}
}
