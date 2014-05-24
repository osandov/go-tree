package tree

import (
	"math/rand"
	"testing"
)

func TestCLZGoSimple(t *testing.T) {
	testClzSimple(t, clz_g)
}

func TestCLZGoRandom(t *testing.T) {
	testClzRandom(t, clz_g)
}

func TestCLZSimple(t *testing.T) {
	testClzSimple(t, clz)
}

func TestCLZRandom(t *testing.T) {
	testClzRandom(t, clz)
}

func BenchmarkCLZGo(b *testing.B) {
	for i := 0; i < b.N * b.N; i++ {
		clz_g(uint64(rand.Uint32())<<32 | uint64(rand.Uint32()))
	}
}

func BenchmarkCLZ(b *testing.B) {
	for i := 0; i < b.N * b.N; i++ {
		clz(uint64(rand.Uint32())<<32 | uint64(rand.Uint32()))
	}
}

func testClzSimple(t *testing.T, clzImpl func(uint64) int) {
	if y := clzImpl(0); y != 64 {
		t.Errorf("clz failed: expected %d, got %d\n", 64, y)
	}
	x := uint64(1)
	for i := 0; i < 64; i++ {
		if y := clzImpl(x); y != (63 - i) {
			t.Errorf("clz failed: expected %d, got %d\n", (63 - i), y)
		}
		x <<= 1
	}
}

func testClzRandom(t *testing.T, clzImpl func(uint64) int) {
	x := uint64(1)
	for i := 0; i < 63; i++ {
		temp := x | uint64(rand.Intn(int(x)))
		if y := clzImpl(temp); y != (63 - i) {
			t.Errorf("clz failed: expected %d, got %d\n", (63 - i), y)
		}
		x <<= 1
	}
	if y := clzImpl(x | uint64(rand.Int())); y != 0 {
		t.Errorf("clz failed: expected %d, got %d\n", 0, y)
	}
}
