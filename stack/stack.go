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

func (s *Stack) Empty() bool {
	return s.ll.Head == nil
}

func (s *Stack) Peek() interface{} {
	if s.Empty() {
		return nil
	}
	return s.ll.Head.Data
}

func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}
	v := s.ll.Head.Data
	s.ll.Delete(v)
	return v
}

func (s *Stack) Add(v interface{}) {
	s.ll.AddToHead(v)
}
