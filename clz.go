package tree

func clz(x uint64) (n int)

func clz_g(x uint64) int {
	x |= (x >> 1)
	x |= (x >> 2)
	x |= (x >> 4)
	x |= (x >> 8)
	x |= (x >> 16)
	x |= (x >> 32)

	x -= (x >> 1) & 0x5555555555555555
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x += x >> 8
	x += x >> 16
	x += x >> 32
	x &= 0x7f

	return 64 - int(x)
}
