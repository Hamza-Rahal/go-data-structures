package trie

type Node struct {
	children map[byte]*Node
	leaf     bool
}

// Trie prefix tree
type Trie struct {
	root *Node
}

func New() *Trie {
	return &Trie{
		root: &Node{
			children: make(map[byte]*Node),
		},
	}
}

// Insert inserts a word in the prefix tree
// complexity: O(len(word))
func (t *Trie) Insert(word []byte) {
	node := t.root
	for _, b := range word {
		childNode, exists := node.children[b]
		if !exists {
			if node.children == nil {
				node.children = make(map[byte]*Node)
			}
			childNode = new(Node)
			node.children[b] = childNode
		}
		node = childNode
	}
	node.leaf = true
}

// Contains checks if a word is in the prefix tree
// complexity: O(len(word))
func (t *Trie) Contains(word []byte) bool {
	node := t.root
	for _, b := range word {
		childNode, exists := node.children[b]
		if !exists {
			return false
		}
		node = childNode
	}
	return node.leaf
}
