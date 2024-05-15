package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func calculate(ch1 <-chan int, ch2 chan<- int) {
	sum := 0
	for val := range ch1 {
		sum += val
	}
	ch2 <- sum
	wg.Done()

}

func Display(ch2 chan int) {
	Sum := <-ch2
	fmt.Println("SUM:", Sum)
	wg.Done()
}

func main() {
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80}
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg.Add(1)
	go calculate(ch1, ch2)
	go func() {
		for _, val := range arr {
			ch1 <- val
		}
		close(ch1)
	}()
	wg.Add(1)
	go Display(ch2)
	wg.Wait()

}
