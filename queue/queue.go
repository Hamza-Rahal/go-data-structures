package queue

import "../linkedlist"

type Node struct {
	data interface{}
	next *Node
}

type Queue struct {
	ll *linkedlist.LinkedList
}

func New() *Queue {
	return &Queue{
		ll: linkedlist.New(),
	}
}

func (q *Queue) Enqueue(data interface{}) {
	q.ll.Append(data)
	return
}

func (q *Queue) Dequeue() interface{} {
	return q.ll.RemoveHead()
}

func (q *Queue) IsEmpty() bool {
	return q.ll.Size() == 0
}
