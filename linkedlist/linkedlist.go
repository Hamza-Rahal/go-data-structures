package linkedlist

type Node struct {
	Data interface{}
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func New() *LinkedList {
	return &LinkedList{
		Head: nil,
	}
}

func newNode(data interface{}) *Node {
	Node := Node{Data: data}
	return &Node
}

func (ll LinkedList) ToArray() []interface{} {
	currentNode := ll.Head

	array := []interface{}{}

	for currentNode != nil {
		array = append(array, currentNode.Data)
		currentNode = currentNode.Next
	}

	return array
}

func (ll LinkedList) Size() int {
	count := 0
	currentNode := ll.Head
	for currentNode != nil {
		count++
		currentNode = currentNode.Next
	}
	return count
}

func (ll *LinkedList) Append(data interface{}) {
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

// removes the first occurrence of v
func (ll *LinkedList) Remove(v interface{}) bool {

	if ll.Head.Data == v {
		ll.Head = ll.Head.Next
		return true
	}

	current := ll.Head
	for current != nil {
		if current.Next.Data == v {
			current.Next = current.Next.Next
			return true
		}
		current = current.Next
	}
	return false
}

func (ll *LinkedList) RemoveHead() interface{} {
	if ll.Head != nil {
		data := ll.Head.Data
		ll.Head = ll.Head.Next
		return data
	}
	return nil
}

func (ll *LinkedList) AddHead(v interface{}) {
	newNode := newNode(v)

	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	t := ll.Head
	ll.Head = newNode
	ll.Head.Next = t
	return
}
