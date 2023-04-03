package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(word string) map[string]int {
	count := make(map[string]int)
	wordsList := strings.Fields(word)

	for _, value := range wordsList {
		if _, ok := count[value]; ok {
			count[value] += 1
		} else {
			count[value] = 1
		}
	}

	return count
}

func main() {
	word := "lets go exercise go go"
	result := WordCount(word)
	fmt.Println("Word count result\n", result)

	wc.Test(WordCount)
}
