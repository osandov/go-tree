package tree

// Simple bitwise trie implementation.

// Bitwise trie.
type trie struct {
	root trieNode
}

// Node in a bitwise trie.
type trieNode struct {
	value    interface{}
	hasValue bool
	children [2]*trieNode
}

func NewTrie() Trie {
	return new(trie)
}

func (tr *trie) Get(key uint64) (interface{}, bool) {
	node := &tr.root

	for i := uint(0); i < 64; i++ {
		if idx := key & (1 << (63 - i)); idx == 0 {
			node = node.children[0]
		} else {
			node = node.children[1]
		}

		if node == nil {
			return nil, false
		}
	}

	return node.value, node.hasValue
}

func (tr *trie) Set(key uint64, value interface{}) (interface{}, bool) {
	node := &tr.root

	for i := uint(0); i < 64; i++ {
		if idx := key & (1 << (63 - i)); idx == 0 {
			if node.children[0] == nil {
				node.children[0] = new(trieNode)
			}
			node = node.children[0]
		} else {
			if node.children[1] == nil {
				node.children[1] = new(trieNode)
			}
			node = node.children[1]
		}
	}

	if node.hasValue {
		orig_value := node.value
		node.value = value
		return orig_value, true
	} else {
		node.value = value
		node.hasValue = true
		return nil, false
	}
}

func (tr *trie) Del(key uint64) (interface{}, bool) {
	var parent *trieNode
	node := &tr.root

	for i := uint(0); i < 64; i++ {
		parent = node

		if idx := key & (1 << (63 - i)); idx == 0 {
			if node.children[0] == nil {
				node.children[0] = new(trieNode)
			}
			node = node.children[0]
		} else {
			if node.children[1] == nil {
				node.children[1] = new(trieNode)
			}
			node = node.children[1]
		}

		if node == nil {
			return nil, false
		}
	}

	if node == parent.children[0] {
		parent.children[0] = nil
	} else {
		parent.children[1] = nil
	}

	return node.value, node.hasValue
}
