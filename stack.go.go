
package main

import "fmt"

type Stack struct {
	items []string
}

func (s *Stack) Push(y string) {
	s.items = append(s.items, y)
}

func (s *Stack) Pop() string {
	lastPerson := len(s.items) - 1
	toRemove := s.items[lastPerson]
	s.items = s.items[:lastPerson]
	return toRemove
}

func main() {
	Batch11 := Stack{}
	
	Batch11.Push("kathir")
	Batch11.Push("natraj")
	Batch11.Push("prasanth")
	fmt.Println(Batch11)
	Batch11.Pop()
	fmt.Println(Batch11)
}