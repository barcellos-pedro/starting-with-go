package main

import "fmt"

/*
Slices, are like references to arrays
A slice does not store any data, it just describes a section of an underlying array.
Changing the elements of a slice modifies the corresponding elements of its underlying array.
Other slices that share the same underlying array will see those changes.
*/
func main() {
	// var primes [4]int = [4]int{2, 3, 5, 7}
	primes := [4]int{2, 3, 5, 7} // creating with array literal

	// var a []int = primes[1:3]
	a := primes[0:2]
	b := primes[1:3]

	fmt.Println("primes:", primes)
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	b[0] = 99 // 3 will change on every slice and array

	fmt.Println("primes:", primes)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
