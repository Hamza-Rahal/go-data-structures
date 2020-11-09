package binarytree

import (
	"../queue"
)

// https://en.wikipedia.org/wiki/Binary_tree#:~:text=In%20computer%20science%2C%20a%20binary,child%20and%20the%20right%20child.
type Node struct {
	data  interface{}
	left  *Node
	right *Node
}

type BinaryTree struct {
	root *Node
}

func New() *BinaryTree {
	return &BinaryTree{
		root: nil,
	}
}

// O(n) time complexity
func (bn *BinaryTree) Insert(data interface{}) {
	newNode := &Node{
		data: data,
	}
	queue := queue.New()

	if bn.root == nil {
		bn.root = newNode
		return
	}

	queue.Enqueue(bn.root)

	for {
		current := queue.Dequeue().(*Node)

		if current.left == nil {
			current.left = newNode
			return
		}
		if current.right == nil {
			current.right = newNode
			return
		}

		queue.Enqueue(current.left)
		queue.Enqueue(current.right)
	}
}

func (bn *BinaryTree) Size() int {
	return size(bn.root)
}

func size(root *Node) int {
	if root == nil {
		return 0
	}
	return size(root.left) + size(root.right) + 1
}

func (bn *BinaryTree) Traverse() []interface{} {
	a := []interface{}{}
	queue := queue.New()

	if bn.root == nil {
		return a
	}
	queue.Enqueue(bn.root.left)
	queue.Enqueue(bn.root.right)

	for {
		node := queue.Dequeue()
		if node != nil {
			a = append(a, node.(*Node).data)
		}

		queue.Enqueue(node.(*Node).left)
		queue.Enqueue(node.(*Node).right)
	}

	return a
}

func (bn *BinaryTree) Height() int {
	return height(bn.root)
}

func height(root *Node) int {
	if root == nil || isLeaf(*root) {
		return 0
	}
	return max(height(root.left), height(root.right)) + 1
}

func isLeaf(node Node) bool {
	return (node.left == nil && node.right == nil)
}

func max(a int, b int) int {
	if a <= b {
		return b
	}
	return a
}
