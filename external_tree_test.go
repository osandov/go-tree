package tree

// Test the external correctness of the tree implementations.

import (
	"math/rand"
	"testing"
)

const NUM_NODES = 10000

type intKey int

func (n intKey) CompareTo(m Key) int {
	if m, ok := m.(intKey); ok {
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

// Binary search tree.

func TestBSTGetMissing(t *testing.T) {
	testGetMissing(t, NewBST())
}

func TestBSTSetUnique(t *testing.T) {
	testSetUnique(t, NewBST())
}

func TestBSTSetDuplicates(t *testing.T) {
	testSetDuplicates(t, NewBST())
}

func TestBSTDelMissing(t *testing.T) {
	testDelMissing(t, NewBST())
}

func TestBSTDel(t *testing.T) {
	testDel(t, NewBST())
}

// Splay tree.

func TestSplayGetMissing(t *testing.T) {
	testGetMissing(t, NewSplay())
}

func TestSplaySetUnique(t *testing.T) {
	testSetUnique(t, NewSplay())
}

func TestSplaySetDuplicates(t *testing.T) {
	testSetDuplicates(t, NewSplay())
}

func TestSplayDelMissing(t *testing.T) {
	testDelMissing(t, NewSplay())
}

func TestSplayDel(t *testing.T) {
	testDel(t, NewSplay())
}

// Test implementations.

func testGetMissing(t *testing.T, tree Tree) {
	for i := 0; i < NUM_NODES; i++ {
		k := rand.Int()
		_, ok := tree.Get(intKey(k))
		if ok {
			t.Fatalf("get failed: non-existent key %v was found\n", k)
		}
	}
}

func testSetUnique(t *testing.T, tree Tree) {
	for i, v := range rand.Perm(NUM_NODES) {
		_, ok := tree.Set(intKey(v), v)
		if ok {
			t.Fatalf("set failed: duplicate reported on set %v of %v\n", i, v)
		}
	}

	for i := 0; i < NUM_NODES; i++ {
		v, ok := tree.Get(intKey(i))
		if !ok {
			t.Fatalf("set failed: %v was not in tree\n", i)

		}
		if i != v {
			t.Errorf("set failed: got %v, expected %v\n", v, i)
		}
	}
}

func testSetDuplicates(t *testing.T, tree Tree) {
	for i, v := range rand.Perm(NUM_NODES) {
		_, ok := tree.Set(intKey(v), v)
		if ok {
			t.Fatalf("set failed: duplicate reported on set %v of %v\n", i, v)
		}
	}

	for i, v := range rand.Perm(NUM_NODES) {
		ov, ok := tree.Set(intKey(v), -v)
		if !ok {
			t.Errorf("set failed: duplicate missing on set %v of %v\n", i, v)
		}

		if ov != v {
			t.Errorf("set failed: incorrect old value %v, expected %v\n", ov, v)
		}
	}

	for i := 0; i < NUM_NODES; i++ {
		v, ok := tree.Get(intKey(i))
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
		k := rand.Int()
		_, ok := tree.Del(intKey(k))
		if ok {
			t.Fatalf("delete failed: non-existent key %v was found\n", k)
		}
	}
}

func testDel(t *testing.T, tree Tree) {
	for i, v := range rand.Perm(NUM_NODES) {
		_, ok := tree.Set(intKey(v), v)
		if ok {
			t.Fatalf("delete failed: duplicate reported on set %v of %v\n", i, v)
		}
	}

	for i := 0; i < NUM_NODES; i++ {
		v, ok := tree.Del(intKey(i))
		if !ok {
			t.Errorf("delete failed: %v was not in tree\n", i)
		}
		if v != i {
			t.Errorf("delete failed: got %v, expected %v\n", v, i)
		}
	}

	for i, v := range rand.Perm(NUM_NODES) {
		_, ok := tree.Set(intKey(v), v)
		if ok {
			t.Fatalf("delete failed: duplicate reported on set %v of %v\n", i, v)
		}
	}

	perm := rand.Perm(NUM_NODES)

	for _, v := range perm[:NUM_NODES/2] {
		ov, ok := tree.Del(intKey(v))
		if !ok {
			t.Errorf("delete failed: %v was not in tree\n", v)
		}
		if v != ov {
			t.Errorf("delete failed: got %v, expected %v\n", v, ov)
		}
	}

	for _, v := range perm[:NUM_NODES/2] {
		_, ok := tree.Get(intKey(v))
		if ok {
			t.Errorf("delete failed: %v still in tree\n", v)
		}
	}

	for _, v := range perm[NUM_NODES/2:] {
		ov, ok := tree.Get(intKey(v))
		if !ok {
			t.Errorf("delete failed: %v was not in tree\n", v)
			continue
		}
		if v != ov {
			t.Errorf("delete failed: got %v, expected %v\n", v, ov)
		}
	}
}
