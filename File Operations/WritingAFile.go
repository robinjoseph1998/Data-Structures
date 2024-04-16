package main

import (
	"fmt"
	"os"
)

func FileWriter(path string, msg string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("can't open file", err)
		return
	}
	for i := 0; i <= 10; i++ {
		fmt.Fprintf(file, "The Message Is %s\n", msg)
	}
	fmt.Println("File writted Succesfully")
}

func main() {

	path := "text.txt"
	msg := "Welcome To Golang"

	FileWriter(path, msg)

}
