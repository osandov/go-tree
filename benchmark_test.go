package tree

// Benchmark tree implementations.
// vim: ft=go

import (
	"math/rand"
	"testing"
	"time"
)

var testRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

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

func benchmarkCreateBalancedLarge(b *testing.B, tree Tree) {
	for i := 0; i < b.N; i++ {
		createBalancedTree(tree, 0x80000000, 0x80000000+NUM_NODES)
	}
}

func benchmarkCreateRandom(b *testing.B, tree Tree) {
	for i := 0; i < b.N/NUM_NODES; i++ {
		for j := 0; j < NUM_NODES; j++ {
			x := uint64(testRand.Int())
			tree.Set(Uint64Key(x), x)
		}
	}
}

func benchmarkCreateRandomLarge(b *testing.B, tree Tree) {
	for i := 0; i < b.N/NUM_NODES; i++ {
		for j := 0; j < NUM_NODES; j++ {
			x := 0x80000000 + uint64(testRand.Int())
			tree.Set(Uint64Key(x), x)
		}
	}
}

func benchmarkRandomGet(b *testing.B, tree Tree) {
	createBalancedTree(tree, 0, NUM_NODES)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Get(Uint64Key(testRand.Intn(NUM_NODES)))
	}
}

func benchmarkRandomGetLarge(b *testing.B, tree Tree) {
	createBalancedTree(tree, 0x80000000, 0x80000000+NUM_NODES)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tree.Get(Uint64Key(0x80000000 + testRand.Intn(NUM_NODES)))
	}
}

func benchmarkLocalGet(b *testing.B, tree Tree) {
	createBalancedTree(tree, 0, NUM_NODES)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 5; j < NUM_NODES-5; j++ {
			x := j + testRand.Intn(10) - 5
			tree.Get(Uint64Key(x))
		}
	}
}

func benchmarkLocalGetLarge(b *testing.B, tree Tree) {
	createBalancedTree(tree, 0x80000000, 0x80000000+NUM_NODES)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 5; j < NUM_NODES-5; j++ {
			x := 0x80000000 + j + testRand.Intn(10) - 5
			tree.Get(Uint64Key(x))
		}
	}
}

func benchmarkRandomDel(b *testing.B, tree Tree) {
	createBalancedTree(tree, 0, NUM_NODES)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x := testRand.Intn(NUM_NODES)
		tree.Del(Uint64Key(x))
		tree.Set(Uint64Key(x), x)
	}
}

func benchmarkLocalDel(b *testing.B, tree Tree) {
	createBalancedTree(tree, 0, NUM_NODES)
	b.ResetTimer()
	for i := 0; i < b.N/NUM_NODES; i++ {
		for j := 5; j < NUM_NODES-5; j++ {
			for k := 0; k < 10; k++ {
				x := j + testRand.Intn(10) - 5
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
func BenchmarkBSTCreateBalancedLarge(b *testing.B) {
	benchmarkCreateBalancedLarge(b, NewBST())
}
func BenchmarkBSTLocalGetLarge(b *testing.B) {
	benchmarkLocalGetLarge(b, NewBST())
}
func BenchmarkBSTRandomGetLarge(b *testing.B) {
	benchmarkRandomGetLarge(b, NewBST())
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
func BenchmarkBSTCreateRandomLarge(b *testing.B) {
	benchmarkCreateRandomLarge(b, NewBST())
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
func BenchmarkSplayCreateBalancedLarge(b *testing.B) {
	benchmarkCreateBalancedLarge(b, NewSplay())
}
func BenchmarkSplayLocalGetLarge(b *testing.B) {
	benchmarkLocalGetLarge(b, NewSplay())
}
func BenchmarkSplayRandomGetLarge(b *testing.B) {
	benchmarkRandomGetLarge(b, NewSplay())
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
func BenchmarkSplayCreateRandomLarge(b *testing.B) {
	benchmarkCreateRandomLarge(b, NewSplay())
}
func BenchmarkSplayCreateRandom(b *testing.B) {
	benchmarkCreateRandom(b, NewSplay())
}

// Simple trie.
func BenchmarkSimpleTrieRandomGet(b *testing.B) {
	benchmarkRandomGet(b, NewTreeFromTrie(NewBinaryTrie()))
}
func BenchmarkSimpleTrieRandomDel(b *testing.B) {
	benchmarkRandomDel(b, NewTreeFromTrie(NewBinaryTrie()))
}
func BenchmarkSimpleTrieCreateBalancedLarge(b *testing.B) {
	benchmarkCreateBalancedLarge(b, NewTreeFromTrie(NewBinaryTrie()))
}
func BenchmarkSimpleTrieLocalGetLarge(b *testing.B) {
	benchmarkLocalGetLarge(b, NewTreeFromTrie(NewBinaryTrie()))
}
func BenchmarkSimpleTrieRandomGetLarge(b *testing.B) {
	benchmarkRandomGetLarge(b, NewTreeFromTrie(NewBinaryTrie()))
}
func BenchmarkSimpleTrieCreateBalanced(b *testing.B) {
	benchmarkCreateBalanced(b, NewTreeFromTrie(NewBinaryTrie()))
}
func BenchmarkSimpleTrieLocalGet(b *testing.B) {
	benchmarkLocalGet(b, NewTreeFromTrie(NewBinaryTrie()))
}
func BenchmarkSimpleTrieLocalDel(b *testing.B) {
	benchmarkLocalDel(b, NewTreeFromTrie(NewBinaryTrie()))
}
func BenchmarkSimpleTrieCreateRandomLarge(b *testing.B) {
	benchmarkCreateRandomLarge(b, NewTreeFromTrie(NewBinaryTrie()))
}
func BenchmarkSimpleTrieCreateRandom(b *testing.B) {
	benchmarkCreateRandom(b, NewTreeFromTrie(NewBinaryTrie()))
}

// CLZ trie.
func BenchmarkCLZTrieRandomGet(b *testing.B) {
	benchmarkRandomGet(b, NewTreeFromTrie(NewCLZTrie()))
}
func BenchmarkCLZTrieRandomDel(b *testing.B) {
	benchmarkRandomDel(b, NewTreeFromTrie(NewCLZTrie()))
}
func BenchmarkCLZTrieCreateBalancedLarge(b *testing.B) {
	benchmarkCreateBalancedLarge(b, NewTreeFromTrie(NewCLZTrie()))
}
func BenchmarkCLZTrieLocalGetLarge(b *testing.B) {
	benchmarkLocalGetLarge(b, NewTreeFromTrie(NewCLZTrie()))
}
func BenchmarkCLZTrieRandomGetLarge(b *testing.B) {
	benchmarkRandomGetLarge(b, NewTreeFromTrie(NewCLZTrie()))
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
func BenchmarkCLZTrieCreateRandomLarge(b *testing.B) {
	benchmarkCreateRandomLarge(b, NewTreeFromTrie(NewCLZTrie()))
}
func BenchmarkCLZTrieCreateRandom(b *testing.B) {
	benchmarkCreateRandom(b, NewTreeFromTrie(NewCLZTrie()))
}

// Radix trie.
func BenchmarkRadixTrieRandomGet(b *testing.B) {
	benchmarkRandomGet(b, NewTreeFromTrie(NewRadixTrie()))
}
func BenchmarkRadixTrieRandomDel(b *testing.B) {
	benchmarkRandomDel(b, NewTreeFromTrie(NewRadixTrie()))
}
func BenchmarkRadixTrieCreateBalancedLarge(b *testing.B) {
	benchmarkCreateBalancedLarge(b, NewTreeFromTrie(NewRadixTrie()))
}
func BenchmarkRadixTrieLocalGetLarge(b *testing.B) {
	benchmarkLocalGetLarge(b, NewTreeFromTrie(NewRadixTrie()))
}
func BenchmarkRadixTrieRandomGetLarge(b *testing.B) {
	benchmarkRandomGetLarge(b, NewTreeFromTrie(NewRadixTrie()))
}
func BenchmarkRadixTrieCreateBalanced(b *testing.B) {
	benchmarkCreateBalanced(b, NewTreeFromTrie(NewRadixTrie()))
}
func BenchmarkRadixTrieLocalGet(b *testing.B) {
	benchmarkLocalGet(b, NewTreeFromTrie(NewRadixTrie()))
}
func BenchmarkRadixTrieLocalDel(b *testing.B) {
	benchmarkLocalDel(b, NewTreeFromTrie(NewRadixTrie()))
}
func BenchmarkRadixTrieCreateRandomLarge(b *testing.B) {
	benchmarkCreateRandomLarge(b, NewTreeFromTrie(NewRadixTrie()))
}
func BenchmarkRadixTrieCreateRandom(b *testing.B) {
	benchmarkCreateRandom(b, NewTreeFromTrie(NewRadixTrie()))
}
