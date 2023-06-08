package main

import "fmt"

type MyErr struct{}

func (MyErr) Error() string {
	return "Something went wrong"
}

func handleError() {
	if err := recover(); err != nil {
		fmt.Println("Recovered from error =>", err)
	}
}

func main() {
	defer handleError()
	panic(MyErr{}) // panic calls Error()
}
