package tree

// Path-compressed radix trie implementation.

// Radix trie.
type radixTrie struct {
	root  radixTrieNode
	radix uint
	count uint64
	mask  uint64
	limit uint
}

type radixTrieNode interface{}

// Node in a bitwise trie.
type radixNode struct {
	key      uint64
	level    uint
	count    uint
	children []radixTrieNode
}

// Leaf containing a value in a bitwise trie.
type radixLeaf struct {
	key   uint64
	value interface{}
}

// NewRadixTrie creates an empty path-compressed radix trie.
func NewRadixTrie(radix uint) Trie {
	if radix < 2 {
		panic("invalid radix")
	}
	rtrie := new(radixTrie)
	rtrie.radix = radix
	rtrie.count = 1 << radix
	rtrie.mask = rtrie.count - 1
	rtrie.limit = (64+(radix+1))/radix - 1
	return rtrie
}

func (rtrie *radixTrie) Get(key uint64) (interface{}, bool) {
	node := rtrie.root

	for node != nil {
		if leaf, ok := node.(*radixLeaf); ok {
			if leaf.key == key {
				return leaf.value, true
			} else {
				break
			}
		} else {
			rnode := node.(*radixNode)
			if rtrie.notDescendant(rnode, key) {
				break
			}
			slot := rtrie.slot(key, rnode.level)
			node = rnode.children[slot]
		}
	}
	return nil, false
}

func (rtrie *radixTrie) Set(key uint64, value interface{}) (interface{}, bool) {
	if rtrie.root == nil {
		rtrie.root = &radixLeaf{key, value}
		return nil, false
	}

	node := rtrie.root
	parent := &rtrie.root

	for {
		if leaf, ok := node.(*radixLeaf); ok {
			if leaf.key == key {
				origValue := leaf.value
				leaf.value = value
				return origValue, true
			}
			level := rtrie.diffLevel(key, leaf.key)
			newKey := rtrie.trimKey(key, level+1)
			node := rtrie.newRadixNode(newKey, level, 2)
			rtrie.setChild(node, key, &radixLeaf{key, value})
			rtrie.setChild(node, leaf.key, leaf)
			*parent = node
			return nil, false
		} else {
			rnode := node.(*radixNode)
			if rtrie.notDescendant(rnode, key) {
				level := rtrie.diffLevel(key, rnode.key)
				newKey := rtrie.trimKey(key, level+1)
				node := rtrie.newRadixNode(newKey, level, 2)
				rtrie.setChild(node, key, &radixLeaf{key, value})
				rtrie.setChild(node, rnode.key, rnode)
				*parent = node
				return nil, false
			}
			slot := rtrie.slot(key, rnode.level)
			if rnode.children[slot] == nil {
				rnode.children[slot] = &radixLeaf{key, value}
				rnode.count++
				return nil, false
			}
			parent = &rnode.children[slot]
			node = rnode.children[slot]
		}
	}
}

func (rtrie *radixTrie) Del(key uint64) (interface{}, bool) {
	if rtrie.root == nil {
		return nil, false
	}

	if leaf, ok := rtrie.root.(*radixLeaf); ok {
		if leaf.key == key {
			rtrie.root = nil
			return leaf.value, true
		} else {
			return nil, false
		}
	}

	var parent *radixNode
	node := rtrie.root

	for node != nil {
		rnode := node.(*radixNode)
		slot := rtrie.slot(key, rnode.level)
		child := rnode.children[slot]
		if leaf, ok := child.(*radixLeaf); ok {
			if leaf.key != key {
				break
			}
			rnode.children[slot] = nil
			rnode.count--
			if rnode.count > 1 {
				return leaf.value, true
			}
			var i uint64
			for ; i < rtrie.count; i++ {
				if rnode.children[i] != nil {
					break
				}
			}
			if i == rtrie.count {
				panic("incorrect count in radix trie")
			}
			if parent == nil {
				rtrie.root = rnode.children[i]
			} else {
				rtrie.setChild(parent, key, rnode.children[i])
			}
			return leaf.value, true
		} else {
			parent = rnode
			node = child
		}
	}

	return nil, false
}

func (rtrie *radixTrie) newRadixNode(key uint64, level uint, count uint) *radixNode {
	children := make([]radixTrieNode, rtrie.count)
	return &radixNode{key, level, count, children}
}

// slot returns the index into the children array of the radix node for the
// given key.
func (rtrie *radixTrie) slot(key uint64, level uint) int {
	return int((key >> (level * rtrie.radix)) & rtrie.mask)
}

// trimKey trims a key after the given level.
func (rtrie *radixTrie) trimKey(key uint64, level uint) uint64 {
	key >>= level * rtrie.radix
	key <<= level * rtrie.radix
	return key
}

// diffLevel finds the highest level at which the keys differ.
func (rtrie *radixTrie) diffLevel(key1, key2 uint64) uint {
	if key1 == key2 {
		panic("equal keys")
	}

	key := key1 ^ key2
	for level := rtrie.limit; ; level-- {
		if rtrie.slot(key, level) != 0 {
			return level
		}
	}
}

// notDescendant returns true if the given key can be determined to not be
// underneath the given node.
func (rtrie *radixTrie) notDescendant(rnode *radixNode, key uint64) bool {
	if rnode.level < rtrie.limit {
		return rtrie.trimKey(key, rnode.level+1) != rnode.key
	} else {
		return false
	}
}

func (rtrie *radixTrie) setChild(rnode *radixNode, key uint64, child radixTrieNode) {
	slot := rtrie.slot(key, rnode.level)
	rnode.children[slot] = child
}
