package main

import (
	"fmt"
	"sort"
)

func sortPeople(names []string, heights []int) []string {
	var result []string
	var hh []int
	hh = append(hh, heights...)
	sort.Sort(sort.Reverse(sort.IntSlice(hh)))
	for i := 0; i < len(names); i++ {
		val := hh[i]
		for k := 0; k < len(heights); k++ {
			if val == heights[k] {
				result = append(result, names[k])
			}
		}
	}
	return result
}

func main() {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}
	fmt.Println(sortPeople(names, heights))

}
