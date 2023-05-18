package main

import "fmt"

// Only the sender should close a channel, never the receiver.
func main() {
	ch := make(chan int, 2)
	fmt.Printf("channel len: %v and cap: %v\n", len(ch), cap(ch)) // 0, 2

	ch <- 1
	ch <- 2
	fmt.Printf("channel len: %v and cap: %v\n", len(ch), cap(ch)) // 2, 2

	fmt.Println("channel value:", <-ch) // 1
	fmt.Println("channel value:", <-ch) // 2

	// Sending on a closed channel will cause a panic.
	// close(ch)
	// after reading above, the channel has availabe 'space' to send values according to its len/cap
	fmt.Printf("channel len: %v and cap: %v\n", len(ch), cap(ch)) // 0, 2
	ch <- 33
	ch <- 999

	// fmt.Println(<-ch)... // 3x reads would cause panic

	// Channels aren't like files; you don't usually need to close them.
	// Closing is only necessary when the receiver must be told there are no more values coming,
	// such as to terminate a range loop
	close(ch)
	for {
		if v, ok := <-ch; ok {
			fmt.Println("chan value:", v)
		} else {
			fmt.Println("channel closed")
			break
		}
	}
	// we can use the if statement with channel capacity
	// for i := 0; i < cap(ch); i++ {
	// 	fmt.Println(<-ch)
	// }
}
