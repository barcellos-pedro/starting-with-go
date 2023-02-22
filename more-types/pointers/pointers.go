package main

import "fmt"

func log(text string) {
	// var aux *string = &text
	aux := &text
	fmt.Println(aux)  // memory address
	fmt.Println(*aux) // value through the pointer
}

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	log("pedro")
}
