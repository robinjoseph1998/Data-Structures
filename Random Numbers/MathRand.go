package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//random number generation using math/rand
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(5))

}
