package main

import (
	"fmt"
)

type MyReader struct{}

func (MyReader) Read(p []byte) (n int, err error) {
	p[0] = 65
	return 1, nil
}

func main() {
	buff := make([]byte, 1)

	for {
		if bytesRead, err := MyReader.Read(MyReader{}, buff); err != nil {
			fmt.Println("Error on Reader")
			break
		} else {
			fmt.Printf("ASCII char 65 = %c\n", buff[:bytesRead])
		}
	}

}
