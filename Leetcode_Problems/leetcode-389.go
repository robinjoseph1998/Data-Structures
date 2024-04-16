package main

import "fmt"

func findTheDifference(s string, t string) byte {
	CountMapS := make(map[byte]int)
	CountMapT := make(map[byte]int)
	for _, char := range s {
		CountMapS[byte(char)]++
	}
	for _, char := range t {
		CountMapT[byte(char)]++
	}
	for char, count := range CountMapT {
		if CountMapS[char] != count {
			return char
		}
	}
	return 0
}

func main() {
	s := "abcd"
	t := "abcde"
	fmt.Println(findTheDifference(s, t))

}
