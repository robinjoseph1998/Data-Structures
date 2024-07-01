package main

import "fmt"

func duplicateZeros(arr []int) {
	orginal := make([]int, len(arr))
	copy(orginal, arr)
	i, j := 0, 0
	for i < len(orginal) && j < len(arr) {
		if orginal[i] != 0 {
			arr[j] = orginal[i]
			j++
			i++
		} else {
			arr[j] = 0
			if j+1 < len(arr) {
				arr[j+1] = 0
			}
			i++
			j += 2
		}
	}
	fmt.Println("arr", arr)
}

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
}
