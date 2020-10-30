package stack

import "../linkedlist"

type Stack struct {
	ll *linkedlist.LinkedList
}

func New() *Stack {
	return &Stack{
		ll: linkedlist.New(),
	}
}

func (s *Stack) IsEmpty() bool {
	return s.ll.Head == nil
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.ll.Head.Data
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.ll.RemoveHead()
}

func (s *Stack) Push(v interface{}) {
	s.ll.AddHead(v)
}
