package main

import (
	"fmt"
)

func sendToIntChannel(ch chan int) {
	// time.Sleep(2 * time.Second)
	ch <- 11
}

func sendToStringChannel(ch chan string) {
	// here we make the channel unavailable for 2 sec
	// time.Sleep(2 * time.Second)
	ch <- "String channel won"
}

func main() {
	numCh := make(chan int)
	strCh := make(chan string)

	go sendToIntChannel(numCh)
	go sendToStringChannel(strCh)

	// When both multiple channels are ready for execution
	// the select statement executes the channel randomly.
	//
	// However, if only one of the channels is ready for execution, it executes that channel
	//
	// If both multiple channels are not ready for execution,
	// select blocks both the channels for a certain time
	// until one or both are available for execution.
	select {
	case strVal := <-strCh:
		fmt.Println(strVal)
	case intVal := <-numCh:
		fmt.Println("int channel won:", intVal)
	}
}
