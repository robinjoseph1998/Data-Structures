package main

import "fmt"

func percentageLetter(s string, letter byte) int {
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] == letter {
			count++
		}
	}
	if count == 0 {
		return 0
	}
	return (count * 100) / 6
}

func main() {

	s := "kue"
	letter := byte('e')
	fmt.Println(percentageLetter(s, letter))

}
