package main

import "fmt"

func countBits(n int) []int {
	arr := []int{}
	count := 0
	for i := 0; i <= n; i++ {
		each := fmt.Sprintf("%b", i)
		for j := 0; j < len(each); j++ {
			if each[j] == '1' {
				count++
			}
		}
		arr = append(arr, count)
		count = 0
	}
	return arr
}

func main() {
	n := 5
	fmt.Println(countBits(n))
}
