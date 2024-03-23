package main

import (
	"fmt"
	"sync"
)

func main() {

	var Score = []int{0}
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}
	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("First routine")
		m.Lock()
		Score = append(Score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Second routine")
		m.Lock()
		Score = append(Score, 2)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Third routine")
		m.Lock()
		Score = append(Score, 3)
		m.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		m.RLock()
		fmt.Println("Score:", Score)
		m.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println("Scores", Score)
}
