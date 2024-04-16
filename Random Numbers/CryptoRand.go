package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	//random number generation using crypto/rand
	num, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(num)

}
