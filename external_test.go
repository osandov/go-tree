package tree

// Test the external correctness of the tree implementations.

import (
	"math/rand"
	"testing"
	"time"
)

const NUM_NODES = 10000

type IntKey int

func (n IntKey) CompareTo(m Key) int {
	if m, ok := m.(IntKey); ok {
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

func TestBST(t *testing.T) {
	testTree(t, NewBST)
}

func TestSplay(t *testing.T) {
	testTree(t, NewSplay)
}

func testTree(t *testing.T, treeFactory func() Tree) {
	rand.Seed(time.Now().UnixNano())
	testGetMissing(t, treeFactory())
	testSetUnique(t, treeFactory())
	testSetDuplicates(t, treeFactory())
	testDelMissing(t, treeFactory())
	testDel(t, treeFactory())
}

func testGetMissing(t *testing.T, tree Tree) {
	if testing.Verbose() {
		t.Logf("%T: get missing\n", tree)
	}

	for i := 0; i < NUM_NODES; i++ {
		k := rand.Int()
		_, ok := tree.Get(IntKey(k))
		if ok {
			t.Fatalf("get failed: non-existent key %v was found\n", k)
		}
	}
}

func testSetUnique(t *testing.T, tree Tree) {
	if testing.Verbose() {
		t.Logf("%T: set unique\n", tree)
	}

	for i, v := range rand.Perm(NUM_NODES) {
		_, ok := tree.Set(IntKey(v), v)
		if ok {
			t.Fatalf("set failed: duplicate reported on set %v of %v\n", i, v)
		}
	}

	for i := 0; i < NUM_NODES; i++ {
		v, ok := tree.Get(IntKey(i))
		if !ok {
			t.Fatalf("set failed: %v was not in tree\n", i)

		}
		if i != v {
			t.Errorf("set failed: got %v, expected %v\n", v, i)
		}
	}
}

func testSetDuplicates(t *testing.T, tree Tree) {
	if testing.Verbose() {
		t.Logf("%T: set duplicates\n", tree)
	}

	for i, v := range rand.Perm(NUM_NODES) {
		_, ok := tree.Set(IntKey(v), v)
		if ok {
			t.Fatalf("set failed: duplicate reported on set %v of %v\n", i, v)
		}
	}

	for i, v := range rand.Perm(NUM_NODES) {
		ov, ok := tree.Set(IntKey(v), -v)
		if !ok {
			t.Errorf("set failed: duplicate missing on set %v of %v\n", i, v)
		}

		if ov != v {
			t.Errorf("set failed: incorrect old value %v, expected %v\n", ov, v)
		}
	}

	for i := 0; i < NUM_NODES; i++ {
		v, ok := tree.Get(IntKey(i))
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
	if testing.Verbose() {
		t.Logf("%T: delete missing\n", tree)
	}

	for i := 0; i < NUM_NODES; i++ {
		k := rand.Int()
		_, ok := tree.Del(IntKey(k))
		if ok {
			t.Fatalf("delete failed: non-existent key %v was found\n", k)
		}
	}
}

func testDel(t *testing.T, tree Tree) {
	if testing.Verbose() {
		t.Logf("%T: delete\n", tree)
	}

	for i, v := range rand.Perm(NUM_NODES) {
		_, ok := tree.Set(IntKey(v), v)
		if ok {
			t.Fatalf("delete failed: duplicate reported on set %v of %v\n", i, v)
		}
	}

	for i := 0; i < NUM_NODES; i++ {
		v, ok := tree.Del(IntKey(i))
		if !ok {
			t.Errorf("delete failed: %v was not in tree\n", i)
		}
		if v != i {
			t.Errorf("delete failed: got %v, expected %v\n", v, i)
		}
	}

	for i, v := range rand.Perm(NUM_NODES) {
		_, ok := tree.Set(IntKey(v), v)
		if ok {
			t.Fatalf("delete failed: duplicate reported on set %v of %v\n", i, v)
		}
	}

	perm := rand.Perm(NUM_NODES)

	for _, v := range perm[:NUM_NODES/2] {
		ov, ok := tree.Del(IntKey(v))
		if !ok {
			t.Errorf("delete failed: %v was not in tree\n", v)
		}
		if v != ov {
			t.Errorf("delete failed: got %v, expected %v\n", v, ov)
		}
	}

	for _, v := range perm[:NUM_NODES/2] {
		_, ok := tree.Get(IntKey(v))
		if ok {
			t.Errorf("delete failed: %v still in tree\n", v)
		}
	}

	for _, v := range perm[NUM_NODES/2:] {
		ov, ok := tree.Get(IntKey(v))
		if !ok {
			t.Errorf("delete failed: %v was not in tree\n", v)
			continue
		}
		if v != ov {
			t.Errorf("delete failed: got %v, expected %v\n", v, ov)
		}
	}
}
