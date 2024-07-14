package main

import "fmt"

func validMountainArray(arr []int) bool {
	if len(arr) < 3 {
		return false
	}
	mid := peak(arr)
	return mid != 0 && mid != len(arr)-1 && isDecreasing(arr, mid) && isIncreasing(arr, mid)

}

func isIncreasing(arr []int, mid int) bool {
	for i := mid + 1; i < len(arr); i++ {
		if arr[i-1] <= arr[i] {
			return false
		}
	}
	return true
}

func isDecreasing(arr []int, mid int) bool {
	for i := 1; i <= mid; i++ {
		if arr[i-1] >= arr[i] {
			return false
		}
	}
	return true
}

func peak(arr []int) int {
	mid := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[mid] {
			mid = i
		}
	}
	return mid
}

func main() {
	arr := []int{1, 3, 2}
	fmt.Println(validMountainArray(arr))
}
