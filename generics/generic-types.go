package main

import "fmt"

// simple type declaration for a single-linked list
type List[T any] struct {
	value T
	next  *List[T] // pointer to avoid error: List refers to itself
}

func (*List[T]) newList() *List[T] {
	return &List[T]{}
}

func (list *List[T]) append(val T) {
	for v := list.next; v.next != nil; v = v.next {
		v.value = val
	}
}

func (list *List[T]) getAll() []T {
	result := make([]T, 0)
	result = append(result, list.value)

	for {
		if node := list.next; node.next == nil {
			break
		} else {
			result = append(result, node.value)
		}
	}

	return result
}

func main() {
	var comments List[string]
	comments.value = "1º"
	fmt.Println(comments)

	comments.append("2º")
	fmt.Println(comments)

	// second := comments.newList()
	// second.value = "2º"
	// comments.next = second
	// fmt.Println(comments)

	// third := comments.newList()
	// third.value = "3º"
	// comments.next.next = third
	// fmt.Println(comments)

	// fmt.Println("All values", comments.getAll())
}
