package main

import "fmt"

/*
We can create slices using make to create dynamically-sized arrays.
The make function allocates a zeroed array and returns a slice that refers to that array
*/

func log(msg string, array []int) {
	fmt.Printf("%s | length=%d | capacity=%d | array=%v", msg, len(array), cap(array), array)
	fmt.Println("")
}

func main() {
	var n = [3]string{"a", "b", "c"}
	fmt.Println("n", n)
	fmt.Println("n slice", n[:])

	a := []int{1, 2, 3}
	log("a", a)

	a = a[1:]
	log("a", a)

	// using make
	arr := make([]int, 3) // len=3, cap=3, [0,0,0]
	log("array 1", arr)

	arr2 := make([]int, 0, 5) // len=0, cap=5, []
	log("array 2", arr2)

	arr3 := arr2[1:cap(arr2)] // len=4, cap=4, [0,0,0,0]
	log("array 3", arr3)
}
