package main

import "fmt"

func findPeaks(mountain []int) []int {
	var peeks []int
	for i := 1; i < len(mountain); i++ {
		if i+1 < len(mountain) {
			if mountain[i] > mountain[i-1] && mountain[i] > mountain[i+1] {
				peeks = append(peeks, i)
			}
		}
	}
	return peeks
}

func main() {

	mountain := []int{2, 4, 4}

	fmt.Println(findPeaks(mountain))

}
