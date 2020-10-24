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

func (ll LinkedList) List() []interface{} {
	currentNode := ll.Head

	list := []interface{}{}

	for currentNode != nil {
		list = append(list, currentNode.Data)
		currentNode = currentNode.Next
	}

	return list
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

func (ll *LinkedList) Delete(v interface{}) {

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

func (ll *LinkedList) AddToHead(v interface{}) {
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
