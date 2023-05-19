package main

import "fmt"

// Channel block case 1
//
// When a goroutine sends data into a channel, the operation is blocked until the data is received by another goroutine.
func sendData(ch chan string) {
	fmt.Println("Sending data to the channel...")
	ch <- "Received. Send Operation Successful"
	// Data is sent and the operation blocks from here
	// until the main or another goroutinte receives the data from the channel
	fmt.Println("No receiver! Send Operation Blocked")
}

func main() {
	// create channel
	ch := make(chan string)

	// function call with goroutine
	go sendData(ch)

	// receive channel data
	fmt.Println(<-ch)
}
