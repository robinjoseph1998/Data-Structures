package main

import (
	"fmt"
	"strings"
)

func restoreString(s string, indices []int) string {
	res := make([]string, len(s))
	for idx, v := range indices {
		res[v] = string(s[idx])
	}
	return strings.Join(res, "")
}

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}
	fmt.Println(restoreString(s, indices))
}
