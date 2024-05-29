package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Linkedlist struct {
	head *Node
}

func NewLinkedlist() *Linkedlist {
	return &Linkedlist{head: nil}
}

func (list *Linkedlist) Add(data int) {
	newNode := &Node{data: data, next: nil}
	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *Linkedlist) Display() {
	current := list.head
	for current != nil {
		fmt.Println("data :", current.data)
		current = current.next
	}
}

func main() {
	var list Linkedlist
	list.Add(10)
	list.Add(20)
	list.Add(30)
	list.Add(40)
	list.Add(50)
	list.Display()

}
