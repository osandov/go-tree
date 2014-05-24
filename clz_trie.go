package tree

// Simple bitwise trie implementation.

// Bitwise trie using count-leading-zeroes as a hint into the tree.
type clzTrie struct {
	// Random-access into the nodes starting with zero bits.
	zeroNodes [65]*trieNode
}

// NewCLZTrie creates an empty bitwise trie. This trie implementation is
// optimized for lexicographically small keys.
func NewCLZTrie() Trie {
	ctr := new(clzTrie)
	ctr.zeroNodes[0] = new(trieNode)
	return ctr
}

func (ctr *clzTrie) Get(key uint64) (interface{}, bool) {
	lz := clz(key)
	node := ctr.zeroNodes[lz]
	if node == nil {
		return nil, false
	}

	for i := uint(lz); i < 64; i++ {
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

func (ctr *clzTrie) Set(key uint64, value interface{}) (interface{}, bool) {
	lz := clz(key)
	node := ctr.zeroNodes[lz]
	for node == nil {
		lz--
		node = ctr.zeroNodes[lz];
	}

	for i := uint(lz); i < 64; i++ {
		if idx := key & (1 << (63 - i)); idx == 0 {
			if node.children[0] == nil {
				node.children[0] = new(trieNode)
				if lz >= 0 {
					lz++
					ctr.zeroNodes[lz] = node.children[0]
				}
			}
			node = node.children[0]
		} else {
			if node.children[1] == nil {
				node.children[1] = new(trieNode)
			}
			node = node.children[1]
			lz = -1
		}
	}

	if node.hasValue {
		origValue := node.value
		node.value = value
		return origValue, true
	} else {
		node.value = value
		node.hasValue = true
		return nil, false
	}
}

func (ctr *clzTrie) Del(key uint64) (interface{}, bool) {
	lz := clz(key)
	node := ctr.zeroNodes[lz]
	if node == nil {
		return nil, false
	}

	for i := uint(lz); i < 64; i++ {
		if idx := key & (1 << (63 - i)); idx == 0 {
			node = node.children[0]
		} else {
			node = node.children[1]
		}

		if node == nil {
			return nil, false
		}
	}

	hadValue := node.hasValue
	node.hasValue = false
	return node.value, hadValue
}
