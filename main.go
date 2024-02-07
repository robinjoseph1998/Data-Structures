package main

type Stack struct {
	Item []interface{}
}

func (s *Stack) push(a interface{}) {
	s.Item = append(s.Item, a)
}

func (s *Stack) pop() interface{} {
	if len(s.Item) == 0 {
		return nil
	}
}

func main() {

}
