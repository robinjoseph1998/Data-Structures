package main

import (
	"fmt"
	"sort"
)

func findLUSlength(strs []string) int {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})
	for i, v := range strs {
		isUncommon := true
		for j, eac := range strs {
			if i != j && isSubsequence(v, eac) {
				isUncommon = false
				break
			}
		}
		if isUncommon {
			return len(v)
		}
	}
	return -1
}

func isSubsequence(st1, st2 string) bool {
	j := 0
	for i := 0; i < len(st2) && j < len(st1); i++ {
		if st1[j] == st2[i] {
			j++
		}
	}
	return len(st1) == j
}

func main() {
	strs := []string{"aba", "cdc", "eae"}
	fmt.Println(findLUSlength(strs))
}
