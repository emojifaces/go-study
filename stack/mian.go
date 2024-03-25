package main

type Stack struct {
	elements []int
}

func NewStack() *Stack {
	return &Stack{elements: []int{}}
}

func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop(element int) int {
	if len(s.elements) == 0 {
		return -1
	}
	element = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

func (s *Stack) Peek() int {
	if len(s.elements) == 0 {
		return -1
	}
	return s.elements[len(s.elements)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}
