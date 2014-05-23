package tree

// Benchmark tree implementations.

import (
	"math/rand"
	"testing"
)

func createBalancedTree(tree Tree, n, m int) {
	if n >= m {
		return
	}

	mid := (n + m) / 2
	tree.Set(uint64Key(mid), mid)
	createBalancedTree(tree, n, mid)
	createBalancedTree(tree, mid+1, n)
}

func TestBalancedTree(t *testing.T) {
	tree := NewBST()
	createBalancedTree(tree, 0, NUM_NODES)
}

// Binary search tree.

func BenchmarkBSTCreateBalanced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createBalancedTree(NewBST(), 0, b.N)
	}
}

func BenchmarkBSTRandomGet(b *testing.B) {
	benchmarkRandomGet(b, NewBST())
}

func BenchmarkBSTLocalGet(b *testing.B) {
	benchmarkLocalGet(b, NewBST())
}

func BenchmarkBSTRandomDel(b *testing.B) {
	benchmarkRandomDel(b, NewBST())
}

func BenchmarkBSTLocalDel(b *testing.B) {
	benchmarkLocalDel(b, NewBST())
}

// Splay tree.

func BenchmarkSplayCreateBalanced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createBalancedTree(NewSplay(), 0, b.N)
	}
}

func BenchmarkSplayRandomGet(b *testing.B) {
	benchmarkRandomGet(b, NewSplay())
}

func BenchmarkSplayLocalGet(b *testing.B) {
	benchmarkLocalGet(b, NewSplay())
}

func BenchmarkSplayRandomDel(b *testing.B) {
	benchmarkRandomDel(b, NewSplay())
}

func BenchmarkSplayLocalDel(b *testing.B) {
	benchmarkLocalDel(b, NewSplay())
}

// Benchmark implementations.

func benchmarkRandomGet(b *testing.B, tree Tree) {
	createBalancedTree(tree, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Get(uint64Key(rand.Intn(b.N)))
	}
}

func benchmarkLocalGet(b *testing.B, tree Tree) {
	createBalancedTree(tree, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := i + rand.Intn(10) - 5
		tree.Get(uint64Key(x))
	}
}

func benchmarkRandomDel(b *testing.B, tree Tree) {
	createBalancedTree(tree, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := rand.Intn(b.N)
		tree.Del(uint64Key(x))
		tree.Set(uint64Key(x), x)
	}
}

func benchmarkLocalDel(b *testing.B, tree Tree) {
	createBalancedTree(tree, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := i + rand.Intn(10) - 5
		tree.Del(uint64Key(x))
		tree.Set(uint64Key(x), x)
	}
}
