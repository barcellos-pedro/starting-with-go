package main

import "fmt"

// TLDR
// Channels are a typed condit where we can
// send and receive values with the operator <-
//
// channels must be created before use
//
// ch <- send v to channel ch
//
// v := <- ch receive from ch and assign
//
// by default sends and receives are blocking
// until the other side is ready
// this allows to sync without
// explicit locks or condition variables
func main() {
	numbers1 := []int{1, 2, 3, 4, 5}
	numbers2 := []int{6, 7, 8, 9, 10}

	ch := make(chan int)

	go sum(numbers1, ch) // execute sum() in another goroutine concurrently
	go sum(numbers2, ch)

	// total1 := <-ch // if we had used only once the channel
	total1, total2 := <-ch, <-ch

	fmt.Println("First sum:", total1)
	fmt.Println("Second sum:", total2)
}

func sum(numbers []int, ch chan int) {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	ch <- sum // send sum to channel
}
