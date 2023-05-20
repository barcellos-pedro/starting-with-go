package main

import (
	"fmt"
	"time"
)

func sendToIntChannel(ch chan int) {
	time.Sleep(2 * time.Second)
	ch <- 11
}

func sendToStringChannel(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "String channel won"
}

func main() {
	numCh := make(chan int)
	strCh := make(chan string)

	go sendToIntChannel(numCh)
	go sendToStringChannel(strCh)

	// When none of the channels are ready
	// the select blocks the program
	// However, it's better to display some messages
	// instead of blocking until the channels are ready.

	// here we keep showing the default case until
	// one of the channels runs
	for {
		select {
		case strVal := <-strCh:
			fmt.Println(strVal)
			return
		case intVal := <-numCh:
			fmt.Println("int channel won:", intVal)
			return
		default:
			fmt.Println("None of the channels are ready for execution")
		}
	}
}
