package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(item interface{}) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, errors.New("Stack is empty")
	}
	index := len(s.data) - 1
	item := s.data[index]
	s.data = s.data[:index]
	return item, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if len(s.data) == 0 {
		return nil, errors.New("Stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func main() {
	stack := Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	item, err := stack.Pop()
	if err == nil {
		fmt.Println("Popped item:", item)
	}

	item, err = stack.Peek()
	if err == nil {
		fmt.Println("Top item:", item)
	}

	fmt.Println("Is stack empty?", stack.IsEmpty())
}
