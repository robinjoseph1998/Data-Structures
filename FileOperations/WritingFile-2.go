package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	path := "/home/lenovo/Data Structure Problems/NEW.txt"

	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer f.Close()
	existingData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf(err.Error())
	}
	data := []byte("IAM ROBIN\n")
	count := 9
	for i := 0; i < count; i++ {
		texts := append(existingData, data...)
		err := ioutil.WriteFile(path, texts, 0644)
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
	fmt.Println("done")

}
