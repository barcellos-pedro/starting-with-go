package main

import (
	"encoding/json"
	"fmt"
)

type Todo struct {
	Text string   `json:"text"`
	Done bool     `json:"done"`
	Tags []string `json:"tags"`
}

func main() {
	todo := Todo{"Do the dishes", false, []string{"daily", "boring"}}
	val, _ := json.Marshal(todo)
	fmt.Println(string(val))

	dataBytes := []byte(`{"message": "hello world", "tags": ["test", "go"], "status": true}`)
	var data map[string]interface{} // to represent a JSON object

	if err := json.Unmarshal(dataBytes, &data); err != nil {
		panic(err)
	}

	fmt.Println(data)
	fmt.Printf("Value: %v | type: %T\n", data["message"], data["message"])
	fmt.Printf("Value: %v | type: %T\n", data["tags"], data["tags"])
	fmt.Printf("Value: %v | type: %T\n", data["status"], data["status"])
}
