package main

import "fmt"

func longestValidParentheses(s string) int {
	stack := []int{-1}
	MaxLength := 0
	for cuIdx, pare := range s {
		if pare == '(' {
			stack = append(stack, cuIdx)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, cuIdx)
			} else {
				MaxLength = Max(MaxLength, cuIdx-stack[len(stack)-1])
			}
		}
	}
	return MaxLength
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	s := "()(())"
	fmt.Println(longestValidParentheses(s))
}
