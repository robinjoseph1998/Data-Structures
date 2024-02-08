package main

import "fmt"

type Stack struct {
	Item []interface{}
}

func (s *Stack) push(a interface{}) {
	s.Item = append(s.Item, a)
}

func (s *Stack) pop() interface{} {
	if len(s.Item) == 0 {
		return nil
	}
	item := s.Item[len(s.Item)-1]
	s.Item = s.Item[:len(s.Item)-1]
	return item
}

func (s *Stack) peek() interface{} {
	if len(s.Item) == 0 {
		return nil
	}
	return s.Item[len(s.Item)-1]
}

func (s *Stack) isEmpty() bool {
	return len(s.Item) == 0
}

func (s *Stack) Size() int {
	return len(s.Item)
}

func main() {
	stack := &Stack{}
	stack.push(900)
	stack.push(600)
	stack.push(700)

	fmt.Println("POP", stack.pop())
	fmt.Println("peek After POP", stack.peek())

}
