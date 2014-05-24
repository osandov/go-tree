// func clz(x uint64) (n uint)
TEXT ·clz(SB),4,$0-16
	BSRQ x+0(FP), AX
	JZ Z1
	XORQ $63, AX
	MOVQ AX, n+8(FP)
	RET

Z1:	MOVQ $64, n+8(FP)
	RET
