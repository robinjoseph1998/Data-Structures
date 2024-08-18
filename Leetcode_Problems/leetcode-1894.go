package main

import "fmt"

func chalkReplacer(chalk []int, k int) int {
	result := -1
	for i := 0; i < len(chalk); i++ {
		k -= chalk[i]
		if k < 0 {
			result = i
			break
		}
		if i == len(chalk)-1 {
			i = -1
		}
	}
	return result
}

func main() {
	chalk := []int{3, 4, 1, 2}
	k := 22
	fmt.Println(chalkReplacer(chalk, k))
}
