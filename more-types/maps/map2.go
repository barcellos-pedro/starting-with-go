package main

import "fmt"

// mutating maps
func main() {
	m := make(map[string]int)
	m["a"] = 300
	m["b"] = 100
	fmt.Println(m)

	delete(m, "b")
	fmt.Println(m)      // map[a:value]
	fmt.Println(m["b"]) // = 0

	elem, ok := m["a"] // value, boolean
	fmt.Println(elem, ok)

	elemB, okB := m["b"]
	fmt.Println(elemB, okB)

	if valueC, existsC := m["c"]; existsC {
		fmt.Println("C key exists: value =", valueC)
	} else {
		fmt.Println("C key do not exists: value =", valueC)
	}

	m["a"] = 1
	fmt.Println(m["a"])

	item := m["a"]
	fmt.Println(item)

	item = 2 // modifies only the copy
	fmt.Println(item)
	fmt.Println(m["a"])
}
