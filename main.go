package main

import (
	"fmt"
	"sort"
)

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	fmt.Println("Channels in golang")

// 	myCh := make(chan int, 2)
// 	wg := &sync.WaitGroup{}
// 	wg.Add(2)
// 	go func(ch <-chan int, wg *sync.WaitGroup) {
// 		for val := range ch {
// 			fmt.Println("val", val)
// 		}
// 		wg.Done()
// 	}(myCh, wg)

// 	go func(ch chan<- int, wg *sync.WaitGroup) {
// 		myCh <- 8
// 		myCh <- 9
// 		close(ch)
// 		wg.Done()
// 	}(myCh, wg)

// 	wg.Wait()
// }

func minOperations(nums []int, k int) int {
	sort.Ints(nums)
	minOp := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < k {
			minOp++
		}
	}
	return minOp
}

func main() {

	nums := []int{2, 11, 10, 1, 3}
	k := 10
	fmt.Println(minOperations(nums, k))
}
