package main

import "fmt"

func minDeletionSize(strs []string) int {
	wordLength := len(strs[0])
	strsLength := len(strs) - 1
	count := 0
	for i := 0; i < wordLength; i++ {
		for j := 0; j < strsLength; j++ {
			if strs[j][i] > strs[j+1][i] {
				count++
				break
			}
		}
	}
	return count
}

func main() {
	strs := []string{"cba", "daf", "ghi"}
	fmt.Println(minDeletionSize(strs))
}
