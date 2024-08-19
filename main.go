package main

import "fmt"

func pivotInteger(n int) int {
	totalSum := 0
	for i := 1; i <= n; i++ {
		totalSum += i
	}
	leftSum := 0
	for x := 1; x <= n; x++ {
		leftSum += x
		rightSum := totalSum - leftSum + x
		if leftSum == rightSum {
			return x
		}
	}
	return -1
}

func main() {
	n := 8
	fmt.Println(pivotInteger(n))
}
