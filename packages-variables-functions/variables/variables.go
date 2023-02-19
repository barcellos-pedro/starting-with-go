package main

import "fmt"

var alive, ok bool
var school, city = "e.e", "seul"
var a = "a"
var b string = "b"

func main() {
	// short assign ':=' only inside functions
	language, studying := "go", true
	t := "t"
	var name, age string
	name = "pedro"
	fmt.Println(alive, ok, name, age, school, city, language, studying, t, a, b)
}
