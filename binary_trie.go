package tree

// Simple bitwise trie implementation.

// Bitwise trie.
type trie struct {
	root trieNode
}

// Node in a bitwise trie.
type trieNode struct {
	value    interface{}
	children [2]*trieNode
}

// NewBinaryTrie creates an empty binary trie. Time complexity is
// O(m), where m is the size of the bit string (64). This implementation does
// not do any optimizations for special cases.
func NewBinaryTrie() Trie {
	return new(trie)
}

func (tr *trie) Get(key uint64) (interface{}, bool) {
	node := &tr.root

	for i := uint(64); i > 0; i-- {
		if key&(1<<(i-1)) == 0 {
			node = node.children[0]
		} else {
			node = node.children[1]
		}

		if node == nil {
			return nil, false
		}
	}

	return node.value, true
}

func (tr *trie) Set(key uint64, value interface{}) (interface{}, bool) {
	node := &tr.root

	for i := uint(64); i > 1; i-- {
		if key&(1<<(i-1)) == 0 {
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

	idx := key & 1
	if node.children[idx] == nil {
		node.children[idx] = &trieNode{value: value}
		return nil, false
	} else {
		origValue := node.children[idx].value
		node.children[idx].value = value
		return origValue, true
	}
}

func (tr *trie) Del(key uint64) (interface{}, bool) {
	node := &tr.root

	for i := uint(64); i > 1; i-- {
		if key&(1<<(i-1)) == 0 {
			node = node.children[0]
		} else {
			node = node.children[1]
		}

		if node == nil {
			return nil, false
		}
	}

	idx := key & 1
	if node.children[idx] == nil {
		return nil, false
	} else {
		origValue := node.children[idx].value
		node.children[idx] = nil
		return origValue, true
	}
}
