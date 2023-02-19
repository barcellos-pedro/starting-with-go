package main

import (
	"fmt"

	"time"

	"rsc.io/quote"
)

func main() {
	fmt.Println("hello, world")
	fmt.Println(quote.Go())
	fmt.Println(quote.Glass())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
	fmt.Println("time\n", time.Now())
}
