package tree

// Test the external correctness of the tree implementations.

import (
	"math/rand"
	"testing"
)

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

func testTree(t *testing.T, treeFactory func() Tree) {
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

	for i := 0; i < 100; i++ {
		k := rand.Int()
		_, ok := tree.Get(IntKey(k))
		if ok {
			t.Fatalf("get failed: non-existent key %d was found\n", k)
		}
	}
}

func testSetUnique(t *testing.T, tree Tree) {
	if testing.Verbose() {
		t.Logf("%T: set unique\n", tree)
	}

	for i, v := range rand.Perm(100) {
		_, ok := tree.Set(IntKey(v), v)
		if ok {
			t.Fatalf("set failed: duplicate reported on set %d of %d\n", i, v)
		}
	}

	for i := 0; i < 100; i++ {
		v, ok := tree.Get(IntKey(i))
		if !ok {
			t.Fatalf("set failed: %d was not in tree\n", i)

		}
		if i != v {
			t.Errorf("set failed: got %d, expected %d\n", v, i)
		}
	}
}

func testSetDuplicates(t *testing.T, tree Tree) {
	if testing.Verbose() {
		t.Logf("%T: set duplicates\n", tree)
	}

	for i, v := range rand.Perm(100) {
		_, ok := tree.Set(IntKey(v), v)
		if ok {
			t.Fatalf("set failed: duplicate reported on set %d of %d\n", i, v)
		}
	}

	for i, v := range rand.Perm(100) {
		ov, ok := tree.Set(IntKey(v), -v)
		if !ok {
			t.Errorf("set failed: duplicate missing on set %d of %d\n", i, v)
		}

		if ov != v {
			t.Errorf("set failed: incorrect old value %d, expected %d\n", ov, v)
		}
	}

	for i := 0; i < 100; i++ {
		v, ok := tree.Get(IntKey(i))
		if !ok {
			t.Fatalf("set failed: %d was not in tree\n", i)
			continue
		}
		if v != -i {
			t.Errorf("set failed: got %d, expected %d\n", v, -i)
		}
	}
}

func testDelMissing(t *testing.T, tree Tree) {
	if testing.Verbose() {
		t.Logf("%T: delete missing\n", tree)
	}

	for i := 0; i < 100; i++ {
		k := rand.Int()
		_, ok := tree.Del(IntKey(k))
		if ok {
			t.Fatalf("delete failed: non-existent key %d was found\n", k)
		}
	}
}

func testDel(t *testing.T, tree Tree) {
	if testing.Verbose() {
		t.Logf("%T: delete\n", tree)
	}

	for i, v := range rand.Perm(100) {
		_, ok := tree.Set(IntKey(v), v)
		if ok {
			t.Fatalf("delete failed: duplicate reported on set %d of %d\n", i, v)
		}
	}

	for i := 0; i < 100; i++ {
		v, ok := tree.Del(IntKey(i))
		if !ok {
			t.Errorf("delete failed: %d was not in tree\n", i)
		}
		if v != i {
			t.Errorf("delete failed: got %d, expected %d\n", v, i)
		}
	}

	for i, v := range rand.Perm(100) {
		_, ok := tree.Set(IntKey(v), v)
		if ok {
			t.Fatalf("delete failed: duplicate reported on set %d of %d\n", i, v)
		}
	}

	perm := rand.Perm(100)

	for _, v := range perm[:50] {
		ov, ok := tree.Del(IntKey(v))
		if !ok {
			t.Errorf("delete failed: %d was not in tree\n", v)
		}
		if v != ov {
			t.Errorf("delete failed: got %d, expected %d\n", v, ov)
		}
	}

	for _, v := range perm[:50] {
		_, ok := tree.Get(IntKey(v))
		if ok {
			t.Errorf("delete failed: %d still in tree\n", v)
		}
	}

	for _, v := range perm[50:] {
		ov, ok := tree.Get(IntKey(v))
		if !ok {
			t.Errorf("delete failed: %d was not in tree\n", v)
			continue
		}
		if v != ov {
			t.Errorf("delete failed: got %d, expected %d\n", v, ov)
		}
	}
}
