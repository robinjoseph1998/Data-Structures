package main

// import "fmt"

// type Example interface {
// 	DataViewer(Fields *Fields) Fields
// }

// type Fields struct {
// 	Name  string
// 	Age   int
// 	Marks []int
// 	Avg   int
// }

// func (F *Fields) DataViewer(fields Fields) Fields {

// 	F.Name = fields.Name
// 	F.Age = fields.Age
// 	F.Marks = fields.Marks
// 	var count int
// 	for _, m := range fields.Marks {
// 		count += m
// 	}
// 	avg := count / len(F.Marks)
// 	fields.Avg = avg
// 	F.Avg = fields.Avg
// 	return fields
// }

// func main() {

// 	obj := Fields{}

// 	var request Fields
// 	request.Name = "Robin Joseph"
// 	request.Age = 25
// 	request.Marks = []int{100, 200, 300, 50, 90, 45}

// 	fmt.Println(obj.DataViewer(request))

// }

// package main

// import "fmt"

// type Sample interface {
// 	DataDisplay(request Details) string
// }

// type Details struct {
// 	Name  string
// 	Age   int
// 	Marks []int
// 	Avg   int
// }

// func (d *Details) DataDisplay(request Details) string {
// 	d.Name = request.Name
// 	d.Age = request.Age
// 	d.Marks = request.Marks
// 	var count int
// 	for _, m := range request.Marks {
// 		count = m
// 	}
// 	d.Avg = count / len(request.Marks)

// 	return fmt.Sprintf(" Name: %s \n Age: %d \n Marks: %v \n Avg_Mark: %d", d.Name, d.Age, d.Marks, d.Avg)

// }

// func main() {

// 	var request Details

// 	request.Name = "Robin Joseph"
// 	request.Age = 25
// 	request.Marks = []int{100, 30, 40, 50, 70}

// 	result := request.DataDisplay(request)
// 	fmt.Println(result)

// }
