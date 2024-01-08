package main

import (
	"fmt"
	"strconv"
)

func findTheArrayConcVal(nums []int) int64 {
	i := 0
	var str string
	var sum int
	ed := len(nums) - 1
	for i < len(nums) {
		st := i
		if i != ed {
			str += strconv.Itoa(nums[st])
			fmt.Println("str", str)
			str += strconv.Itoa(nums[ed])
			val, _ := strconv.Atoi(str)
			fmt.Println("val", val)
			sum += val
			str = ""
			fmt.Println("SUM", sum)
		} else {
			str += strconv.Itoa(nums[i])
			val, _ := strconv.Atoi(str)
			sum += val
		}
		i++
		ed--
		if i > ed {
			break
		}
	}
	return int64(sum)
}

func main() {
	nums := []int{5, 14, 13, 8, 12}
	fmt.Println(findTheArrayConcVal(nums))
}
