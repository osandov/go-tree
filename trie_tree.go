package tree

// Wrap a Trie as a Tree.

type trieTree struct {
	trie Trie
}

type uint64Key uint64

func (n uint64Key) CompareTo(m Key) int {
	if m, ok := m.(uint64Key); ok {
		if n < m {
			return -1
		} else if n > m {
			return 1
		} else {
			return 0
		}
	} else {
		panic("invalid comparison")
	}
}

// Wrap a Trie in a Tree.
func NewTreeFromTrie(trie Trie) Tree {
	return &trieTree{trie}
}

func (tt *trieTree) Get(key Key) (interface{}, bool) {
	return tt.trie.Get(uint64(key.(uint64Key)))
}

func (tt *trieTree) Set(key Key, value interface{}) (interface{}, bool) {
	return tt.trie.Set(uint64(key.(uint64Key)), value)
}

func (tt *trieTree) Del(key Key) (interface{}, bool) {
	return tt.trie.Del(uint64(key.(uint64Key)))
}
