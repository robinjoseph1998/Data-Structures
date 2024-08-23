package main

import "fmt"

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	junk := []int{}
	for i := 0; i < len(arr1); i++ {
		count := 0
		for j := 0; j < len(arr2); j++ {
			if Minus(arr1[i], arr2[j]) > d {
				count++
			}
		}
		fmt.Println("count", count)
		if count == len(arr2) {
			junk = append(junk, arr1[i])
		}
	}
	fmt.Println("junk", junk)
	return len(junk)
}

func Minus(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	arr1 := []int{4, 5, 8}
	arr2 := []int{10, 9, 1, 8}
	d := 2
	fmt.Println(findTheDistanceValue(arr1, arr2, d))
}
