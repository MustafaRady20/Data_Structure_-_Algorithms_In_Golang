package main

import "fmt"

type Stack struct {
	stack []int
}

func newStack() *Stack {
	return &Stack{make([]int, 0)}
}

func (s *Stack) Push(data int) {
	s.stack = append(s.stack, data)
}

func (s *Stack) Pop() int {
	l := len(s.stack)
	if l == 0 {
		fmt.Println("Stack is empty")
		return 0
	}
	item := s.stack[l-1]
	s.stack = s.stack[:l-1]
	return item

}

func (s *Stack) PrintStack() {
	for len(s.stack) != 0 {
		fmt.Println(s.Pop())
	}
}

func main() {

	s := &Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.PrintStack() // 6=>5=>4=>3=>2=>1
	s.Pop()
	s.PrintStack() // 5=>4=>3=>2=>1

}
