package main

import (
	"fmt"
	"net/http"
	"strings"
)

func logHeaders(req *http.Request) {
	for key, header := range req.Header {
		fmt.Printf("Key: %v | Value: %v\n", key, header)
		for _, val := range header {
			fmt.Printf("Value: %v\n", val)
			fmt.Println(strings.Repeat("-", 20))
		}
	}
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		logHeaders(req)
		fmt.Fprintln(w, "Hello world")
	})

	http.ListenAndServe(":8080", nil)
}
