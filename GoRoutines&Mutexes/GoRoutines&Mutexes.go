package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var lc sync.Mutex

type resource struct {
	str   string
	count int
}

func (r *resource) firstMessage(s string) {
	lc.Lock()
	defer wg.Done()
	defer lc.Unlock()
	r.str = s
	r.count += 2
}

func (r *resource) SecondMessage(s string) {
	lc.Lock()
	defer wg.Done()
	defer lc.Unlock()
	r.str = s
	r.count += 3
}

func main() {
	res := resource{}
	count := 6
	wg.Add(6 * 2)
	for i := 0; i < count; i++ {
		go res.firstMessage("Hello")
		go res.SecondMessage("world")
	}
	wg.Wait()
	fmt.Printf("Result: str=%s, count=%d\n", res.str, res.count)
}
