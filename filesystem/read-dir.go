package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	entries, entriesErr := os.ReadDir("content")

	handleError(entriesErr)

	for i, file := range entries {
		info, infoErr := file.Info()
		content, contentErr := os.ReadFile("./content/" + file.Name())

		handleError(infoErr)
		handleError(contentErr)

		fmt.Printf("# File %d\n", i)
		fmt.Printf("Name: %s\n", file.Name())
		fmt.Printf("Is dir: %v\n", info.IsDir())
		fmt.Printf("Modification time: %s\n", info.ModTime())
		fmt.Printf("Size (bytes): %d\n", info.Size())
		fmt.Printf("Content %s\n", content)
		fmt.Println(strings.Repeat("-", 15))
	}
}
