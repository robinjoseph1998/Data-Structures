package main

import (
	"fmt"
	"strconv"
)

func separateDigits(nums []int) []int {
	var result []int
	for _, val := range nums {
		if val > 9 {
			str := strconv.Itoa(val)
			for _, char := range str {
				intval := int(char)
				c, _ := strconv.Atoi(string(intval))
				result = append(result, c)
			}
		} else {
			result = append(result, val)
			fmt.Println("Trees", result)
		}
	}
	fmt.Println("result", result)
	return result
}

func main() {
	nums := []int{9, 81, 33, 17, 58, 42, 8, 75}
	fmt.Println(separateDigits(nums))
}
