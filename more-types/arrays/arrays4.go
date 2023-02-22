package main

import "fmt"

func main() {
	// Slice defaults
	// The default is zero for the low bound
	// and for the higher is the length of the slice

	list := []string{"a", "b", "c", "d"}
	a := list[1:3]
	b := list[:2]
	c := list[2:]
	d := list[:]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
