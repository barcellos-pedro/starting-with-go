// Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
package main

import (
	"fmt"
	"strings"
)

type MyReader struct {
	char string
}

// Populates the given slice in the first position with the ASCII value for 'A'
func (reader *MyReader) Read(p []byte) (n int, err error) {
	asciiNum := 65
	reader.char = fmt.Sprintf("%c", asciiNum)
	p[0] = byte(asciiNum)
	return len(p), nil
}

func main() {
	buff := make([]byte, 1)
	reader := &MyReader{}

	for {
		bytesRead, err := reader.Read(buff)

		if err != nil {
			fmt.Println("Error on Reader")
			break
		}

		fmt.Printf("ASCII char 65 = %c\n", buff[0])
		fmt.Println("Reader char", reader.char)
		fmt.Println("Bytes read", bytesRead)
		fmt.Println(strings.Repeat("-", 20))
	}

}
