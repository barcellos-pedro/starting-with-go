package main

import "fmt"

// Channel block case 2
//
// When a goroutine receives data from a channel, the operation is blocked until another goroutine sends the data to the channel.
func receiveData(ch chan string) {
	fmt.Println("Receiving data from the channel...")
	// receive data from the channel
	// blocks here until the main or another goroutinte sends data to the channel
	fmt.Println(<-ch)
}

func main() {
	// create channel
	ch := make(chan string)

	// function call with goroutine
	go receiveData(ch)

	fmt.Println("No data sent yet. Receive Operation Blocked")

	// send data to the channel
	// the goroutine of receiveData() can continue its execution
	ch <- "Data Received. Receive Operation Successful"
}
