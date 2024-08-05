package main

import "fmt"

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	c := NumArray{}
	c.nums = nums
	return c
}

func (this *NumArray) SumRange(left int, right int) int {
	res := 0
	for i := left; i <= right; i++ {
		res += this.nums[i]
	}
	return res
}

func main() {
	var v NumArray
	v = Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(v.SumRange(0, 2))
}
