package main

import (
	"fmt"
	"sort"
)

func average(salary []int) float64 {
	sort.Ints(salary)
	if len(salary) == 3 {
		return float64(salary[1])
	}
	total := 0
	for i := 1; i < len(salary)-1; i++ {
		total += salary[i]
	}
	return float64(total) / float64(len(salary)-2)
}

func main() {
	salary := []int{4000, 3000, 1000, 2000}
	fmt.Println(average(salary))
}
