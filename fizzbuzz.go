package main

import "fmt"

/*
FizzBuzz problem

"FizzBuzz" if i is divisible by 3 and 5

"Fizz" if i is divisible by 3

"Buzz" if i is divisible by 5

"i as string" if none of the above
*/
func Fizzbuzz(n int) []string {
	var result []string

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "Fizzbuzz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else {
			result = append(result, fmt.Sprint(i))
		}
	}

	return result
}

func main() {
	fmt.Println(Fizzbuzz(3))
}
