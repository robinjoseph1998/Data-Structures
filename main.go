package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type linkedlist struct {
	head *Node
}

func (ll *linkedlist) Append(data int) {
	newNode := &Node{data: data, next: nil}
	if ll.head == nil {
		ll.head = newNode
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode

}

func (ll *linkedlist) Display() {
	current := ll.head
	for current != nil {
		fmt.Println("data", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := linkedlist{}
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)
	ll.Append(40)

	ll.Display()
}
