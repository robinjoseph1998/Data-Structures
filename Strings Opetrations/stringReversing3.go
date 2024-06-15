package main

import "fmt"

func main() {
	str := "Hello Wold Good Morning"
	strRune := []rune(str)
	j := 0
	for i := len(str) - 1; i >= len(str)/2; i-- {
		strRune[i], strRune[j] = strRune[j], strRune[i]
		j++
	}
	fmt.Println(string(strRune))
}
