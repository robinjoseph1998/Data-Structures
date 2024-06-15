package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	str := "Hello World"
	fmt.Println("str", str)
	data := []byte(str)

	hasher := sha256.New()
	hasher.Write(data)

	hashed := hasher.Sum(nil)

	fmt.Println("Hashed:", hashed)
	fmt.Printf("Hashed: %x\n", hashed)

}
