package main

import "fmt"


type Stack[T comparable] struct { 
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T // we cannot just return nil because it doesn't match nil
		// In go fashion, we create a instance of T using var and automatically get the zero value
		return zero
	}
	v := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items)]
	return v
}

func (s *Stack[T]) Empty () bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Contain(val T) bool {
	if s.Empty() {
		return false
	}
	for _, v := range s.items {
		if v == val {
			return true
		}
	}
	return false
}


func main () {
	stack := Stack[int] {}
	stack.Push(20)
	stack.Push(30)
	stack.Push(43)
	fmt.Println("What does stack have to offer?", stack.Pop())
	fmt.Println("Do we have 30?", stack.Contain(30))
}