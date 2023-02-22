package main

import "fmt"

// The zero value of a slice is nil
// A nil slice has a length and capacity of 0 and has no underlying array.
func main() {
	var slice []int //  !== var slice = []int{}
	fmt.Println(slice, len(slice), cap(slice))

	if slice == nil {
		fmt.Println("slice is nil!")
	}
}
