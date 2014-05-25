package tree

// Simple bitwise trie implementation.

// Bitwise trie using count-leading-zeroes as a hint into the tree.
type clzTrie struct {
	// Random-access into the nodes starting with zero bits.
	zeroNodes [64]*trieNode
}

// NewCLZTrie creates an empty binary trie. This trie implementation is
// optimized for lexicographically small keys.
func NewCLZTrie() Trie {
	ctr := new(clzTrie)
	ctr.zeroNodes[63] = new(trieNode)
	for i := 62; i >= 0; i-- {
		ctr.zeroNodes[i] = new(trieNode)
		ctr.zeroNodes[i].children[0] = ctr.zeroNodes[i+1]
	}
	return ctr
}

func (ctr *clzTrie) Get(key uint64) (interface{}, bool) {
	if key == 0 {
		node := ctr.zeroNodes[63].children[0]
		if node == nil {
			return nil, false
		} else {
			return node.value, true
		}
	} else {
		lz := clz(key)
		node := ctr.zeroNodes[lz]

		for i := uint(64 - lz); i > 0; i-- {
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

}

func (ctr *clzTrie) Set(key uint64, value interface{}) (interface{}, bool) {
	if key == 0 {
		node := ctr.zeroNodes[63]
		if child := node.children[0]; child == nil {
			node.children[0] = &trieNode{value: value}
			return nil, false
		} else {
			origValue := child.value
			child.value = value
			return origValue, true
		}
	} else {
		lz := clz(key)
		node := ctr.zeroNodes[lz]

		for i := uint(64 - lz); i > 1; i-- {
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
}

func (ctr *clzTrie) Del(key uint64) (interface{}, bool) {
	if key == 0 {
		node := ctr.zeroNodes[63]
		if node.children[0] == nil {
			return nil, false
		} else {
			origValue := node.children[0].value
			node.children[0] = nil
			return origValue, true
		}
	} else {
		lz := clz(key)
		node := ctr.zeroNodes[lz]

		for i := uint(64 - lz); i > 1; i-- {
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
}
