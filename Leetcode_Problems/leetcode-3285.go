package main

import "fmt"

func stableMountains(height []int, threshold int) []int {
	res := []int{}
	for i := 1; i < len(height); i++ {
		if height[i-1] > threshold {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	height := []int{1, 2, 3, 4, 5}
	threshold := 2
	fmt.Println(stableMountains(height, threshold))
}
