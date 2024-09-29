package main

import "fmt"

func distributeCandies(candies int, num_people int) []int {
	run := 0
	i := 1
	res := make([]int, num_people)
	for candies > 0 {
		toGive := Min(i, candies)
		res[run] += toGive
		candies -= i
		run++
		i++
		if run >= num_people && candies != 0 {
			run = 0
		}
	}
	return res
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	candies := 7
	num_people := 4

	fmt.Println(distributeCandies(candies, num_people))

}
