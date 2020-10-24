package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func newLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
	}
}

func newNode(data interface{}) *Node {
	Node := Node{Data: data}
	return &Node
}

func (ll LinkedList) list() []interface{} {
	currentNode := ll.Head

	list := []interface{}{}

	for currentNode != nil {
		list = append(list, currentNode.Data)
		currentNode = currentNode.Next
	}

	return list
}

func (ll LinkedList) size() int {
	count := 0
	currentNode := ll.Head
	for currentNode != nil {
		count++
		currentNode = currentNode.Next
	}
	return count
}

func (ll *LinkedList) append(data interface{}) {
	newNode := newNode(data)
	current := ll.Head

	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	return
}

func (ll *LinkedList) delete(v interface{}) {

	if ll.Head.Data == v {
		ll.Head = ll.Head.Next
		return
	}

	current := ll.Head
	for current != nil {
		if current.Next.Data == v {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
	return
}

func main() {
	ll := newLinkedList()
	ll.append(1)
	ll.append(2)
	ll.append(3)

	fmt.Println(ll.list())

	ll.delete(1)
	fmt.Println(ll.list())

	ll.delete(3)
	fmt.Println(ll.list())
}

func assert(predicate bool, message string) {
	if !predicate {
		panic(message)
	}
}
