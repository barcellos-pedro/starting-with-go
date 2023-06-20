package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	url := "https://jsonplaceholder.typicode.com/todos/1"

	response, responseErr := http.Get(url)
	handleError(responseErr)

	respBytes, bytesErr := io.ReadAll(response.Body)
	handleError(bytesErr)

	var todo Todo

	jsonErr := json.Unmarshal(respBytes, &todo)
	handleError(jsonErr)

	fmt.Println(todo)
	fmt.Printf("ID: %v\n", todo.Id)
	fmt.Printf("User ID: %v\n", todo.UserId)
	fmt.Printf("Title: %v\n", todo.Title)
	fmt.Printf("Completed: %v\n", todo.Completed)
}
