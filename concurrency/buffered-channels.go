package main

import "fmt"

// channels can be buffered
// with the 2nd arg as the length
//
// Blocks on a buffered channel:
// 1- Sends block only when the buffer is full
// 2- Receives block when the buffer is empty
func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
