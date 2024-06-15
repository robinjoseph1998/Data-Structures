package main

import "fmt"

type SampleInterface interface {
	print(msg string)
}

type SampleStruct struct {
	message string
}

func (s SampleStruct) print(msg string) {
	fmt.Println("Message is", msg)
}

func main() {

	var Si SampleInterface = SampleStruct{}
	var ss SampleStruct
	ss.print("Hello World")

	if dat, ok := Si.(SampleStruct); ok {
		fmt.Println("the interface holds the Sample struct", dat.message)
	} else {
		fmt.Println("type is not defined")
	}

}
