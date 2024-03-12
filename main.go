package main

import "fmt"

func isValid(s string) bool {
	var Stack []rune
	brackets := map[rune]rune{'(': ')', '{': '}', '[': ']'}

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			Stack = append(Stack, char)
		} else {
			if len(Stack) == 0 || brackets[Stack[len(Stack)-1]] != char {
				return false
			}
			Stack = Stack[:len(Stack)-1]
		}
	}
	return len(Stack) == 0
}

func main() {

	s := "([)]"
	fmt.Println(isValid(s))

}
