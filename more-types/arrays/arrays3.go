package main

import "fmt"

// Slice literal is like Array literal without the length
func main() {
	array := [3]bool{true, false, false}
	// this creates the same array as above
	// then builds a slice that references it
	slice := []bool{true, false, false}
	fmt.Println(array)
	fmt.Println(slice)
}
