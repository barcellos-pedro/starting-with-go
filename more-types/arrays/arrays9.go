package main

import "fmt"

/*
Appending to slice

If the backing array of is too small to fit all the given values
a bigger array will be allocated
The returned slice will point to the newly allocated array
*/

func log(msg string, array []string) {
	fmt.Printf("%s | length=%d | capacity=%d | array=%v", msg, len(array), cap(array), array)
	fmt.Println("")
}

func main() {
	var array []string

	fmt.Println("array == nil\n", array == nil)
	log("array", array)

	// it works with nil slices
	array = append(array, "pedro")

	log("array", array)
}
