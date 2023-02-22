package main

import "fmt"

/*
A slice has both a length and a capacity
The length of a slice is the number of elements it contains
The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
*/
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length
	// but, mantain its capacity based on the underlying array values
	s = s[:0]
	printSlice(s)

	// You can extend a slice's length by re-slicing it
	// until its maximum capacity!
	s = s[:4]
	printSlice(s)

	// Drop its first two values
	// and decrease slice capacity
	s = s[2:]
	printSlice(s)

	// extend to its full capacity
	s = s[:4]
	printSlice(s)

	// s = s[:7] // error when trying to extend beyond its capacity
	// printSlice(s)
}
