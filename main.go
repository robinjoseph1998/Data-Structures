package main

import "fmt"

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {

	var Bag []int
	if numOnes != 0 {
		for i := 0; i < numOnes; i++ {
			Bag = append(Bag, 1)
		}
	}
	if numZeros != 0 {
		for i := 0; i < numZeros; i++ {
			Bag = append(Bag, 0)
		}
	}
	if numNegOnes == 0 {
		Bag = append(Bag)
	} else {
		for i := 0; i < numNegOnes; i++ {
			Bag = append(Bag, -1)
		}
	}
	fmt.Println("Bag", Bag)
	Sum := 0
	for i := 0; i < k; i++ {
		fmt.Println("k", k)
		fmt.Println("i", i)
		Sum += Bag[i]
	}
	return Sum
}

func main() {

	numOnes := 1
	numZeros := 4
	numNegOnes := 1
	k := 3

	fmt.Println(kItemsWithMaximumSum(numOnes, numZeros, numNegOnes, k))

}
