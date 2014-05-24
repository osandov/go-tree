package tree

// Benchmark tree implementations.
// vim: ft=go

import (
	"math/rand"
	"testing"
)

func createBalancedTree(tree Tree, n, m int) {
	if n >= m {
		return
	}

	mid := (n + m) / 2
	tree.Set(Uint64Key(mid), mid)
	createBalancedTree(tree, n, mid)
	createBalancedTree(tree, mid+1, n)
}

func TestBalancedTree(t *testing.T) {
	tree := NewBST()
	createBalancedTree(tree, 0, NUM_NODES)
}

// Benchmark implementations.

func benchmarkCreateBalanced(b *testing.B, tree Tree) {
	for i := 0; i < b.N; i++ {
		createBalancedTree(tree, 0, NUM_NODES)
	}
}

func benchmarkCreateRandom(b *testing.B, tree Tree) {
	for i := 0; i < b.N / NUM_NODES; i++ {
		for j := 0; j < NUM_NODES; j++ {
			x := uint64(rand.Int())
			tree.Set(Uint64Key(x), x)
		}
	}
}

func benchmarkRandomGet(b *testing.B, tree Tree) {
	createBalancedTree(tree, 0, NUM_NODES)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Get(Uint64Key(rand.Intn(NUM_NODES)))
	}
}

func benchmarkLocalGet(b *testing.B, tree Tree) {
	createBalancedTree(tree, 0, NUM_NODES)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 5; j < NUM_NODES - 5; j++ {
			x := j + rand.Intn(10) - 5
			tree.Get(Uint64Key(x))
		}
	}
}

func benchmarkRandomDel(b *testing.B, tree Tree) {
	createBalancedTree(tree, 0, NUM_NODES)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := rand.Intn(NUM_NODES)
		tree.Del(Uint64Key(x))
		tree.Set(Uint64Key(x), x)
	}
}

func benchmarkLocalDel(b *testing.B, tree Tree) {
	createBalancedTree(tree, 0, NUM_NODES)
	b.ResetTimer()
	for i := 0; i < b.N / NUM_NODES; i++ {
		for j := 5; j < NUM_NODES - 5; j++ {
			for k := 0; k < 10; k++ {
				x := j + rand.Intn(10) - 5
				tree.Del(Uint64Key(x))
				tree.Set(Uint64Key(x), x)
			}
		}
	}
}

// Binary search tree.
func BenchmarkBSTRandomGet(b *testing.B) {
	benchmarkRandomGet(b, NewBST())
}
func BenchmarkBSTRandomDel(b *testing.B) {
	benchmarkRandomDel(b, NewBST())
}
func BenchmarkBSTCreateBalanced(b *testing.B) {
	benchmarkCreateBalanced(b, NewBST())
}
func BenchmarkBSTLocalGet(b *testing.B) {
	benchmarkLocalGet(b, NewBST())
}
func BenchmarkBSTLocalDel(b *testing.B) {
	benchmarkLocalDel(b, NewBST())
}
func BenchmarkBSTCreateRandom(b *testing.B) {
	benchmarkCreateRandom(b, NewBST())
}

// Splay tree.
func BenchmarkSplayRandomGet(b *testing.B) {
	benchmarkRandomGet(b, NewSplay())
}
func BenchmarkSplayRandomDel(b *testing.B) {
	benchmarkRandomDel(b, NewSplay())
}
func BenchmarkSplayCreateBalanced(b *testing.B) {
	benchmarkCreateBalanced(b, NewSplay())
}
func BenchmarkSplayLocalGet(b *testing.B) {
	benchmarkLocalGet(b, NewSplay())
}
func BenchmarkSplayLocalDel(b *testing.B) {
	benchmarkLocalDel(b, NewSplay())
}
func BenchmarkSplayCreateRandom(b *testing.B) {
	benchmarkCreateRandom(b, NewSplay())
}

// Simple trie.
func BenchmarkSimpleTrieRandomGet(b *testing.B) {
	benchmarkRandomGet(b, NewTreeFromTrie(NewTrie()))
}
func BenchmarkSimpleTrieRandomDel(b *testing.B) {
	benchmarkRandomDel(b, NewTreeFromTrie(NewTrie()))
}
func BenchmarkSimpleTrieCreateBalanced(b *testing.B) {
	benchmarkCreateBalanced(b, NewTreeFromTrie(NewTrie()))
}
func BenchmarkSimpleTrieLocalGet(b *testing.B) {
	benchmarkLocalGet(b, NewTreeFromTrie(NewTrie()))
}
func BenchmarkSimpleTrieLocalDel(b *testing.B) {
	benchmarkLocalDel(b, NewTreeFromTrie(NewTrie()))
}
func BenchmarkSimpleTrieCreateRandom(b *testing.B) {
	benchmarkCreateRandom(b, NewTreeFromTrie(NewTrie()))
}

// CLZ trie.
func BenchmarkCLZTrieRandomGet(b *testing.B) {
	benchmarkRandomGet(b, NewTreeFromTrie(NewCLZTrie()))
}
func BenchmarkCLZTrieRandomDel(b *testing.B) {
	benchmarkRandomDel(b, NewTreeFromTrie(NewCLZTrie()))
}
func BenchmarkCLZTrieCreateBalanced(b *testing.B) {
	benchmarkCreateBalanced(b, NewTreeFromTrie(NewCLZTrie()))
}
func BenchmarkCLZTrieLocalGet(b *testing.B) {
	benchmarkLocalGet(b, NewTreeFromTrie(NewCLZTrie()))
}
func BenchmarkCLZTrieLocalDel(b *testing.B) {
	benchmarkLocalDel(b, NewTreeFromTrie(NewCLZTrie()))
}
func BenchmarkCLZTrieCreateRandom(b *testing.B) {
	benchmarkCreateRandom(b, NewTreeFromTrie(NewCLZTrie()))
}
