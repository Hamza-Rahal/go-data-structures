package linkedlist

type Node struct {
	data interface{}
	next *Node
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
	Node := Node{data: data}
	return &Node
}

func (ll LinkedList) ToArray() []interface{} {
	currentNode := ll.Head

	array := []interface{}{}

	for currentNode != nil {
		array = append(array, currentNode.data)
		currentNode = currentNode.next
	}

	return array
}

func (ll LinkedList) Size() int {
	count := 0
	currentNode := ll.Head
	for currentNode != nil {
		count++
		currentNode = currentNode.next
	}
	return count
}

// Append adds a node at the tail of the list
func (ll *LinkedList) Append(data interface{}) {
	newNode := newNode(data)
	current := ll.Head

	if ll.Head == nil {
		ll.Head = newNode
		return
	}
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	return
}

// Remove removes the first occurrence of v
func (ll *LinkedList) Remove(v interface{}) bool {

	if ll.Head.data == v {
		ll.Head = ll.Head.next
		return true
	}

	current := ll.Head
	for current != nil {
		if current.next.data == v {
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func (ll *LinkedList) RemoveHead() interface{} {
	if ll.Head != nil {
		data := ll.Head.data
		ll.Head = ll.Head.next
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
	ll.Head.next = t
	return
}
