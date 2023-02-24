package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"a", "b", "c"}
	fmt.Println(names)
	fmt.Println(strings.Join(names, "-"))

	path := "/some/deep/nested/folder/here"
	fmt.Println(path)
	fmt.Println(strings.Split(path, "/"))

	game := [][]string{
		[]string{"[x]", "[X]", "[O]"},
		[]string{"[O]", "[O]", "[O]"},
		[]string{"[X]", "[X]", "[O]"},
	}

	for i := 0; i < len(game); i++ {
		fmt.Println(game[i])
	}
}
