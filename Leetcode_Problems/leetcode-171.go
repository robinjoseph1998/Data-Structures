package main

import "fmt"

func titleToNumber(columnTitle string) int {
	res := 0
	for _, ch := range columnTitle {
		res = res*26 + (int(ch) - int('A') + 1)
	}
	return res
}

func main() {
	columnTitle := "AB"
	fmt.Println(titleToNumber(columnTitle))
}
