package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	dir := "content"
	fileName := "hello.txt"
	filePath := path.Join(dir, fileName)

	bytes, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File content in bytes", bytes)
	fmt.Fprint(os.Stdout, string(bytes))
}
