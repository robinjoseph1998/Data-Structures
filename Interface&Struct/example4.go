package main

import "fmt"

type Sample interface {
	DataDisplay(request Details) string
}

type Details struct {
	Name  string
	Age   int
	Marks []int
	Avg   int
}

func (d *Details) DataDisplay(request Details) string {
	d.Name = request.Name
	d.Age = request.Age
	d.Marks = request.Marks
	var count int
	for _, m := range request.Marks {
		count = m
	}
	d.Avg = count / len(request.Marks)

	return fmt.Sprintf(" Name: %s \n Age: %d \n Marks: %v \n Avg_Mark: %d", d.Name, d.Age, d.Marks, d.Avg)

}

func main() {

	var request Details

	request.Name = "Robin Joseph"
	request.Age = 25
	request.Marks = []int{100, 30, 40, 50, 70}

	result := request.DataDisplay(request)
	fmt.Println(result)

}
