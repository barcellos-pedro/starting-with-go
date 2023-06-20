package main

import "fmt"

func main() {
	// data from api would come like this
	person := map[string]interface{}{
		"name":    "pedro",
		"age":     26,
		"hobbies": []string{"samba", "movies"},
	}

	fmt.Println(person)
	fmt.Println()

	// property access
	fmt.Printf("Name: %v\n", person["name"])
	fmt.Printf("Age: %v\n", person["age"])
	fmt.Printf("hobbies: %v\n", person["hobbies"])

	fmt.Println()

	// range property access
	for key, val := range person {
		fmt.Printf("key %s | value %v\n", key, val)
	}
}
