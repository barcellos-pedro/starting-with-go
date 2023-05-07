package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "Lbh penpxrq gur pbqr"
	message = strings.ToLower(message)

	for i := range message {
		code := message[i] // current char code
		min := byte(97)    // ascii for letter a
		max := byte(122)   // ascii for letter z

		if code <= 109 {
			fmt.Printf("%v", string(code+13))
		} else {
			diff := code + 13
			result := diff - max - 1
			fmt.Printf("%v", string(min+result))
		}
	}
}
