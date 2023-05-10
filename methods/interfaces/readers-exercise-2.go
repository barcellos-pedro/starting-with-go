// Implement a rot13Reader that implements io.Reader and reads from an io.Reader,
// modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

// About ROT13
// https://en.wikipedia.org/wiki/ROT13
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	reader io.Reader
}

// TODO:
func (rot *rot13Reader) Read(p []byte) (int, error) {
	bytesRead, err := rot.reader.Read(p)

	if err != nil {
		return 0, io.EOF
	}

	message := string(p[:bytesRead])
	message = strings.ToLower(message)

	for i := 0; i < len(message); i++ {
		min := byte(97)  // ascii for letter a
		max := byte(122) // ascii for letter z

		if code := message[i]; code <= 109 {
			p[i] = code + 13
		} else {
			diff := code + 13
			result := diff - max - 1
			p[i] = min + result
		}
	}

	return len(p), nil
}

func main() {
	messageReader := strings.NewReader("Lbh penpxrq gur pbqr!")
	rot13 := rot13Reader{messageReader}
	io.Copy(os.Stdout, &rot13)
}
