package main

import "fmt"

// simple type declaration for a single-linked list
type List[T any] struct {
	value T
	next  *List[T] // pointer to avoid error: List refers to itself
}

// interface declaration to specify methods for the struct
type SingleLinkedList[T any] interface {
	newList() *List[T]
	peek() T
	values() []T
	push(val T)
}

func (*List[T]) newList() *List[T] {
	return &List[T]{}
}

func (list *List[T]) peek() T {
	return list.value
}

func (list *List[T]) values() []T {
	if list.next == nil {
		return []T{list.value}
	}

	values := make([]T, 0)
	values = append(values, list.value)

	node := list.next

	// while node is not nil
	for node != nil {
		values = append(values, node.value)
		node = node.next
	}

	return values
}

func (list *List[T]) push(val T) {
	newElement := list.newList()
	newElement.value = val

	if list.next == nil {
		list.next = newElement
		return
	}

	node := list.next

	for {
		if node.next == nil {
			node.next = newElement
			break
		}
		node = node.next
	}
}

func main() {
	var comments List[string]
	comments.value = "nice"
	comments.push("cool")
	comments.push("dope")

	// Testing methods
	fmt.Println("Peek:", comments.peek())
	fmt.Println("All values:", comments.values())
	fmt.Println("Range")
	for _, comment := range comments.values() {
		fmt.Println(comment)
	}

	// Using an interface
	// pointer to comments because of methods that have pointer receivers
	var i SingleLinkedList[string] = &comments
	fmt.Println("Using Interface")
	fmt.Println(i.peek())
	fmt.Println(i.values())
}
