package main

import "fmt"

func FindDuplicates(arr []int) []int {

	Dupicates := []int{}
	myMap := make(map[int]int)

	for _, Count := range arr {
		myMap[Count]++
	}

	for num, Count := range myMap {
		if Count > 1 {

			Dupicates = append(Dupicates, num)
		}
	}
	return Dupicates
}

func main() {
	arr := []int{1, 1, 2, 2, 4, 4, 3, 3, 7, 8, 9, 9, 9, 5}

	fmt.Println("Duplicates", FindDuplicates(arr))
}
