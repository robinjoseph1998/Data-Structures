package main

import (
	"fmt"
	"slices"
)

func checkPerfectNumber(num int) bool {
	perfects := []int{6, 28, 496, 8128, 33550336, 8589869056}

	return slices.Contains(perfects, num)

}

// This is a Valid Solution But Time Limit Exceeds :(
//
//	func checkPerfectNumber(num int) bool {
//		sum := 0
//		for i := 1; i < num; i++ {
//			if num%i == 0 {
//				fmt.Println("i", i)
//				sum += i
//			}
//		}
//		return sum == num
//	}
func main() {
	num := 28
	fmt.Println(checkPerfectNumber(num))
}
