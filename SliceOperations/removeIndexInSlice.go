package main

import "fmt"

var courses = []string{"golang", "fluutter", "python"}

func main() {
	courses = append(courses[:2], courses[2+1:]...)
	fmt.Println("couuurses", courses)

}
