package main

import (
	"fmt"
	"sort"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var arr1 []rune
	var arr2 []rune
	for _, val := range s {
		arr1 = append(arr1, val)
	}
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	for _, val := range t {
		arr2 = append(arr2, val)
	}
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	for i := 0; i < len(s); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func main() {

	s := "anagram"

	t := "nagaram"

	fmt.Println(isAnagram(s, t))

}
