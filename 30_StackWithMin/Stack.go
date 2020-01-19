package main

import "errors"

type element struct {
	Value int
	Min   int
}

type Stack struct {
	elements []element
}

func (s *Stack) Top() (int, error) {
	if s.Empty() {
		return 0, errors.New("stack is empty")
	}
	return s.elements[len(s.elements)-1].Value , nil
}

func (s *Stack) Push(value int) {
	var ele element
	if s.Empty() {
		ele = element{value,value}
	}else {
		min := s.elements[len(s.elements)-1].Min
		if min <= value {
			ele = element{value, min}
		}else {
			ele = element{value,value}
		}
	}
	s.elements = append(s.elements, ele)
}

func (s *Stack) Pop() (int, error) {
	if s.Empty() {
		return 0, errors.New("stack is empty")
	}
	top := s.elements[len(s.elements)-1].Value
	s.elements = s.elements[0:len(s.elements)-1]
	return top, nil
}

func (s *Stack) Min() (int, error) {
	if s.Empty() {
		return 0, errors.New("stack is empty")
	}
	min := s.elements[len(s.elements)-1].Min
	return min, nil
}

func (s *Stack) Empty() bool {
	if len(s.elements) <= 0 {
		return true
	}
	return false
}

type IStack interface{
	Push(int)
	Pop()  (int, error)
	Min()  (int, error)
	Top()  (int, error)
	Empty() bool
}

func StackInstance() IStack {
	return &Stack{}
}
func NewStack() *Stack {
	var elements  []element
	return &Stack{elements}
}

