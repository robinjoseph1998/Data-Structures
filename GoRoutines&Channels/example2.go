package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	//Receive Only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		_, isChannelOpen := <-myCh
		if isChannelOpen {
			for val := range myCh {
				fmt.Println("values", val)
			}
		}
		wg.Done()
	}(myCh, wg)
	//Send Only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myCh <- 1
		myCh <- 8
		myCh <- 9
		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
