package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func calculate(arr []int, ch chan int) {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	ch <- sum
	wg.Done()
}

func Display(ch chan int) {
	val := <-ch
	fmt.Println("Sum", val)
	wg.Done()
}

func main() {
	arr := []int{1, 6, 8, 12, 54, 34}
	ch := make(chan int, len(arr))
	wg.Add(2)
	go calculate(arr, ch)
	go Display(ch)
	wg.Wait()

}
