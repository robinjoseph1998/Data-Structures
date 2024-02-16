package main

import "fmt"

func areSimilar(mat [][]int, k int) bool {
	for _, each := range mat {
		for i := 0; i < len(each); i++ {
			if each[i] != each[(i+k)%len(each)] {
				return false
			}
		}
	}
	return true
}

func main() {
	mat := [][]int{{1, 2, 1, 2}, {5, 5, 5, 5}, {6, 3, 6, 3}}
	k := 2
	fmt.Println(areSimilar(mat, k))

}
