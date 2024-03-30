package main

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
