package tree

// Path-compressed radix trie implementation.

const (
	RADIX_WIDTH = 4
	RADIX_COUNT = 1 << RADIX_WIDTH
	RADIX_MASK  = RADIX_COUNT - 1
	RADIX_LIMIT = (64+(RADIX_WIDTH+1))/RADIX_WIDTH - 1
)

// Radix trie.
type radixTrie struct {
	root radixTrieNode
}

type radixTrieNode interface{}

// Node in a bitwise trie.
type radixNode struct {
	key      uint64
	level    uint
	count    uint
	children [RADIX_COUNT]radixTrieNode
}

// Leaf containing a value in a bitwise trie.
type radixLeaf struct {
	key   uint64
	value interface{}
}

// NewRadixTrie creates an empty path-compressed radix trie.
func NewRadixTrie() Trie {
	return new(radixTrie)
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
			if rnode.notDescendant(key) {
				break
			}
			slot := radixSlot(key, rnode.level)
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
			level := radixDiffLevel(key, leaf.key)
			newKey := radixTrimKey(key, level+1)
			node := newRadixNode(newKey, level, 2)
			node.setChild(key, &radixLeaf{key, value})
			node.setChild(leaf.key, leaf)
			*parent = node
			return nil, false
		} else {
			rnode := node.(*radixNode)
			if rnode.notDescendant(key) {
				level := radixDiffLevel(key, rnode.key)
				newKey := radixTrimKey(key, level+1)
				node := newRadixNode(newKey, level, 2)
				node.setChild(key, &radixLeaf{key, value})
				node.setChild(rnode.key, rnode)
				*parent = node
				return nil, false
			}
			slot := radixSlot(key, rnode.level)
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
		slot := radixSlot(key, rnode.level)
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
			i := 0
			for ; i < RADIX_COUNT; i++ {
				if rnode.children[i] != nil {
					break
				}
			}
			if i == RADIX_COUNT {
				panic("incorrect count in radix trie")
			}
			if parent == nil {
				rtrie.root = rnode.children[i]
			} else {
				parent.setChild(key, rnode.children[i])
			}
			return leaf.value, true
		} else {
			parent = rnode
			node = child
		}
	}

	return nil, false
}

func newRadixNode(key uint64, level uint, count uint) *radixNode {
	return &radixNode{key: key, level: level, count: count}
}

// radixSlot returns the index into the children array of the radix node for
// the given key.
func radixSlot(key uint64, level uint) int {
	return int((key >> (level * RADIX_WIDTH)) & RADIX_MASK)
}

// radixTrimKey trims a key after the given level.
func radixTrimKey(key uint64, level uint) uint64 {
	key >>= level * RADIX_WIDTH
	key <<= level * RADIX_WIDTH
	return key
}

// radixDiffLevel finds the highest level at which the keys differ.
func radixDiffLevel(key1, key2 uint64) uint {
	if key1 == key2 {
		panic("equal keys")
	}

	key := key1 ^ key2
	for level := uint(RADIX_LIMIT); ; level-- {
		if radixSlot(key, level) != 0 {
			return level
		}
	}
}

// notDescendant returns true if the given key can be determined to not be
// underneath the given node.
func (rnode *radixNode) notDescendant(key uint64) bool {
	if rnode.level < RADIX_LIMIT {
		return radixTrimKey(key, rnode.level+1) != rnode.key
	} else {
		return false
	}
}

// setChild adds this child in its slot.
func (rnode *radixNode) setChild(key uint64, child radixTrieNode) {
	slot := radixSlot(key, rnode.level)
	rnode.children[slot] = child
}
