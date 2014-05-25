package tree

// Test the external correctness of the tree implementations.
// vim: ft=go

import (
	"testing"
)

const NUM_NODES = 10000

// Test implementations.

func testGetMissing(t *testing.T, tree Tree) {
	for i := 0; i < NUM_NODES; i++ {
		k := testRand.Int()
		_, ok := tree.Get(Uint64Key(k))
		if ok {
			t.Fatalf("get failed: non-existent key %v was found\n", k)
		}
	}
}

func testSetUnique(t *testing.T, tree Tree) {
	for i, v := range testRand.Perm(NUM_NODES) {
		_, ok := tree.Set(Uint64Key(v), v)
		if ok {
			t.Fatalf("set failed: duplicate reported on set %v of %v\n", i, v)
		}
	}

	for i := 0; i < NUM_NODES; i++ {
		v, ok := tree.Get(Uint64Key(i))
		if !ok {
			t.Fatalf("set failed: %v was not in tree\n", i)

		}
		if i != v {
			t.Errorf("set failed: got %v, expected %v\n", v, i)
		}
	}
}

func testSetDuplicates(t *testing.T, tree Tree) {
	for i, v := range testRand.Perm(NUM_NODES) {
		_, ok := tree.Set(Uint64Key(v), v)
		if ok {
			t.Fatalf("set failed: duplicate reported on set %v of %v\n", i, v)
		}
	}

	for i, v := range testRand.Perm(NUM_NODES) {
		ov, ok := tree.Set(Uint64Key(v), -v)
		if !ok {
			t.Errorf("set failed: duplicate missing on set %v of %v\n", i, v)
		}

		if ov != v {
			t.Errorf("set failed: incorrect old value %v, expected %v\n", ov, v)
		}
	}

	for i := 0; i < NUM_NODES; i++ {
		v, ok := tree.Get(Uint64Key(i))
		if !ok {
			t.Fatalf("set failed: %v was not in tree\n", i)
			continue
		}
		if v != -i {
			t.Errorf("set failed: got %v, expected %v\n", v, -i)
		}
	}
}

func testDelMissing(t *testing.T, tree Tree) {
	for i := 0; i < NUM_NODES; i++ {
		k := testRand.Int()
		_, ok := tree.Del(Uint64Key(k))
		if ok {
			t.Fatalf("delete failed: non-existent key %v was found\n", k)
		}
	}
}

func testDel(t *testing.T, tree Tree) {
	for i, v := range testRand.Perm(NUM_NODES) {
		_, ok := tree.Set(Uint64Key(v), v)
		if ok {
			t.Fatalf("delete failed: duplicate reported on set %v of %v\n", i, v)
		}
	}

	for i := 0; i < NUM_NODES; i++ {
		v, ok := tree.Del(Uint64Key(i))
		if !ok {
			t.Errorf("delete failed: %v was not in tree\n", i)
		}
		if v != i {
			t.Errorf("delete failed: got %v, expected %v\n", v, i)
		}
	}

	for i, v := range testRand.Perm(NUM_NODES) {
		_, ok := tree.Set(Uint64Key(v), v)
		if ok {
			t.Fatalf("delete failed: duplicate reported on set %v of %v\n", i, v)
		}
	}

	perm := testRand.Perm(NUM_NODES)

	for _, v := range perm[:NUM_NODES/2] {
		ov, ok := tree.Del(Uint64Key(v))
		if !ok {
			t.Errorf("delete failed: %v was not in tree\n", v)
		}
		if v != ov {
			t.Errorf("delete failed: got %v, expected %v\n", v, ov)
		}
	}

	for _, v := range perm[:NUM_NODES/2] {
		_, ok := tree.Get(Uint64Key(v))
		if ok {
			t.Errorf("delete failed: %v still in tree\n", v)
		}
	}

	for _, v := range perm[NUM_NODES/2:] {
		ov, ok := tree.Get(Uint64Key(v))
		if !ok {
			t.Errorf("delete failed: %v was not in tree\n", v)
			continue
		}
		if v != ov {
			t.Errorf("delete failed: got %v, expected %v\n", v, ov)
		}
	}
}

// Binary search tree.
func TestBSTDelMissing(t *testing.T) {
	testDelMissing(t, NewBST())
}
func TestBSTDel(t *testing.T) {
	testDel(t, NewBST())
}
func TestBSTSetDuplicates(t *testing.T) {
	testSetDuplicates(t, NewBST())
}
func TestBSTGetMissing(t *testing.T) {
	testGetMissing(t, NewBST())
}
func TestBSTSetUnique(t *testing.T) {
	testSetUnique(t, NewBST())
}

// Splay tree.
func TestSplayDelMissing(t *testing.T) {
	testDelMissing(t, NewSplay())
}
func TestSplayDel(t *testing.T) {
	testDel(t, NewSplay())
}
func TestSplaySetDuplicates(t *testing.T) {
	testSetDuplicates(t, NewSplay())
}
func TestSplayGetMissing(t *testing.T) {
	testGetMissing(t, NewSplay())
}
func TestSplaySetUnique(t *testing.T) {
	testSetUnique(t, NewSplay())
}

// Simple trie.
func TestSimpleTrieDelMissing(t *testing.T) {
	testDelMissing(t, NewTreeFromTrie(NewBinaryTrie()))
}
func TestSimpleTrieDel(t *testing.T) {
	testDel(t, NewTreeFromTrie(NewBinaryTrie()))
}
func TestSimpleTrieSetDuplicates(t *testing.T) {
	testSetDuplicates(t, NewTreeFromTrie(NewBinaryTrie()))
}
func TestSimpleTrieGetMissing(t *testing.T) {
	testGetMissing(t, NewTreeFromTrie(NewBinaryTrie()))
}
func TestSimpleTrieSetUnique(t *testing.T) {
	testSetUnique(t, NewTreeFromTrie(NewBinaryTrie()))
}

// CLZ trie.
func TestCLZTrieDelMissing(t *testing.T) {
	testDelMissing(t, NewTreeFromTrie(NewCLZTrie()))
}
func TestCLZTrieDel(t *testing.T) {
	testDel(t, NewTreeFromTrie(NewCLZTrie()))
}
func TestCLZTrieSetDuplicates(t *testing.T) {
	testSetDuplicates(t, NewTreeFromTrie(NewCLZTrie()))
}
func TestCLZTrieGetMissing(t *testing.T) {
	testGetMissing(t, NewTreeFromTrie(NewCLZTrie()))
}
func TestCLZTrieSetUnique(t *testing.T) {
	testSetUnique(t, NewTreeFromTrie(NewCLZTrie()))
}

// Radix-3 trie.
func TestRadixTrie3DelMissing(t *testing.T) {
	testDelMissing(t, NewTreeFromTrie(NewRadixTrie(3)))
}
func TestRadixTrie3Del(t *testing.T) {
	testDel(t, NewTreeFromTrie(NewRadixTrie(3)))
}
func TestRadixTrie3SetDuplicates(t *testing.T) {
	testSetDuplicates(t, NewTreeFromTrie(NewRadixTrie(3)))
}
func TestRadixTrie3GetMissing(t *testing.T) {
	testGetMissing(t, NewTreeFromTrie(NewRadixTrie(3)))
}
func TestRadixTrie3SetUnique(t *testing.T) {
	testSetUnique(t, NewTreeFromTrie(NewRadixTrie(3)))
}

// Radix-4 trie.
func TestRadixTrie4DelMissing(t *testing.T) {
	testDelMissing(t, NewTreeFromTrie(NewRadixTrie(4)))
}
func TestRadixTrie4Del(t *testing.T) {
	testDel(t, NewTreeFromTrie(NewRadixTrie(4)))
}
func TestRadixTrie4SetDuplicates(t *testing.T) {
	testSetDuplicates(t, NewTreeFromTrie(NewRadixTrie(4)))
}
func TestRadixTrie4GetMissing(t *testing.T) {
	testGetMissing(t, NewTreeFromTrie(NewRadixTrie(4)))
}
func TestRadixTrie4SetUnique(t *testing.T) {
	testSetUnique(t, NewTreeFromTrie(NewRadixTrie(4)))
}
