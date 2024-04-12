package main

import "fmt"

func firstUniqChar(s string) int {
	charCount := make(map[string]int)
	for _, char := range s {
		charCount[string(char)]++

	}

	for i, c := range s {
		if charCount[string(c)] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))

}
