package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Status  bool     `json:"status"`
	Hobbies []string `json:"hobbies"`
}

func getUser(w http.ResponseWriter, req *http.Request) {
	person := Person{"Pedro", 26, true, []string{"movies", "restaurants", "travel"}}
	data, _ := json.Marshal(person) // skipped error handling for the sake of simplicity
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintln(w, string(data))
}

func main() {
	http.HandleFunc("/user", getUser)
	http.ListenAndServe(":8080", nil)
}
