package main

import "fmt"

func main() {

	str := "Hello World"
	runeStr := []rune(str)
	j := len(str) - 1
	fmt.Println("Before Reversing: ", str)
	for i := 0; i < len(str)/2; i++ {
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
		j--
	}
	fmt.Println("After Reversing ", string(runeStr))
}
