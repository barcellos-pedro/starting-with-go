package main

import (
	"fmt"
	"time"
)

func main() {
	targetDate := time.Saturday
	currentDate := time.Now().Weekday()

	fmt.Printf("When's %s?", targetDate)

	switch targetDate {
	case currentDate:
		fmt.Println("\nToday.")
	case currentDate + 1:
		fmt.Println("\nTomorrow.")
	case currentDate + 2:
		fmt.Println("\nIn two days.")
	default:
		fmt.Println("\nToo far away.")
	}
}
