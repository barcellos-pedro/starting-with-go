package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello world")
	// using the exact number of bytes (11)
	// buffer := make([]byte, reader.Len())
	buffer := make([]byte, 4)

	// read the string content 4 bytes at a time
	for {
		bytesRead, err := reader.Read(buffer)

		// Error that indicates end of file
		if err == io.EOF {
			break
		}

		fmt.Printf("Bytes read: %v\n", bytesRead)
		fmt.Printf("Content: %q\n", buffer[:bytesRead])
		fmt.Println("-------")
	}
}
