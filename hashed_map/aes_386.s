TEXT ·strhash(SB),NOSPLIT,$0-12
	MOVL	p+0(FP), AX	// ptr to string object
	MOVL	4(AX), BX	// length of string
	MOVL	(AX), AX	// string data
	LEAL	ret+8(FP), DX
	JMP	aeshashbody<>(SB)

TEXT aeshashbody<>(SB),NOSPLIT,$0-0
	MOVL	h+4(FP), X0	            // 32 bits of per-table hash seed
	PINSRW	$4, BX, X0	            // 16 bits of length
	PSHUFHW	$0, X0, X0	            // replace size with its low 2 bytes repeated 4 times
	MOVO	X0, X1                      // save unscrambled seed
	PXOR	runtime·aeskeysched(SB), X0 // xor in per-process seed
	AESENC	X0, X0                      // scramble seed

	CMPL	BX, $16
	JB	aes0to15
	JE	aes16
	CMPL	BX, $32
	JBE	aes17to32
	CMPL	BX, $64
	JBE	aes33to64
	JMP	aes65plus

aes0to15:
	TESTL	BX, BX
	JE	aes0

	ADDL	$16, AX
	TESTW	$0xff0, AX
	JE	endofpage

	// 16 bytes loaded at this address won't cross
	// a page boundary, so we can load it directly.
	MOVOU	-16(AX), X1
	ADDL	BX, BX
	PAND	masks<>(SB)(BX*8), X1

final1:
	AESENC	X0, X1  // scramble input, xor in seed
	AESENC	X1, X1  // scramble combo 2 times
	AESENC	X1, X1
	MOVL	X1, (DX)
	RET

endofpage:
	// address ends in 1111xxxx. Might be up against
	// a page boundary, so load ending at last byte.
	// Then shift bytes down using pshufb.
	MOVOU	-32(AX)(BX*1), X1
	ADDL	BX, BX
	PSHUFB	shifts<>(SB)(BX*8), X1
	JMP	final1

aes0:
	// Return scrambled input seed
	AESENC	X0, X0
	MOVL	X0, (DX)
	RET

aes16:
	MOVOU	(AX), X1
	JMP	final1

aes17to32:
	// make second starting seed
	PXOR	runtime·aeskeysched+16(SB), X1
	AESENC	X1, X1

	// load data to be hashed
	MOVOU	(AX), X2
	MOVOU	-16(AX)(BX*1), X3

	// scramble 3 times
	AESENC	X0, X2
	AESENC	X1, X3
	AESENC	X2, X2
	AESENC	X3, X3
	AESENC	X2, X2
	AESENC	X3, X3

	// combine results
	PXOR	X3, X2
	MOVL	X2, (DX)
	RET

aes33to64:
	// make 3 more starting seeds
	MOVO	X1, X2
	MOVO	X1, X3
	PXOR	runtime·aeskeysched+16(SB), X1
	PXOR	runtime·aeskeysched+32(SB), X2
	PXOR	runtime·aeskeysched+48(SB), X3
	AESENC	X1, X1
	AESENC	X2, X2
	AESENC	X3, X3

	MOVOU	(AX), X4
	MOVOU	16(AX), X5
	MOVOU	-32(AX)(BX*1), X6
	MOVOU	-16(AX)(BX*1), X7

	AESENC	X0, X4
	AESENC	X1, X5
	AESENC	X2, X6
	AESENC	X3, X7

	AESENC	X4, X4
	AESENC	X5, X5
	AESENC	X6, X6
	AESENC	X7, X7

	AESENC	X4, X4
	AESENC	X5, X5
	AESENC	X6, X6
	AESENC	X7, X7

	PXOR	X6, X4
	PXOR	X7, X5
	PXOR	X5, X4
	MOVL	X4, (DX)
	RET

aes65plus:
	// make 3 more starting seeds
	MOVO	X1, X2
	MOVO	X1, X3
	PXOR	runtime·aeskeysched+16(SB), X1
	PXOR	runtime·aeskeysched+32(SB), X2
	PXOR	runtime·aeskeysched+48(SB), X3
	AESENC	X1, X1
	AESENC	X2, X2
	AESENC	X3, X3

	// start with last (possibly overlapping) block
	MOVOU	-64(AX)(BX*1), X4
	MOVOU	-48(AX)(BX*1), X5
	MOVOU	-32(AX)(BX*1), X6
	MOVOU	-16(AX)(BX*1), X7

	// scramble state once
	AESENC	X0, X4
	AESENC	X1, X5
	AESENC	X2, X6
	AESENC	X3, X7

	// compute number of remaining 64-byte blocks
	DECL	BX
	SHRL	$6, BX

aesloop:
	// scramble state, xor in a block
	MOVOU	(AX), X0
	MOVOU	16(AX), X1
	MOVOU	32(AX), X2
	MOVOU	48(AX), X3
	AESENC	X0, X4
	AESENC	X1, X5
	AESENC	X2, X6
	AESENC	X3, X7

	// scramble state
	AESENC	X4, X4
	AESENC	X5, X5
	AESENC	X6, X6
	AESENC	X7, X7

	ADDL	$64, AX
	DECL	BX
	JNE	aesloop

	// 2 more scrambles to finish
	AESENC	X4, X4
	AESENC	X5, X5
	AESENC	X6, X6
	AESENC	X7, X7

	AESENC	X4, X4
	AESENC	X5, X5
	AESENC	X6, X6
	AESENC	X7, X7

	PXOR	X6, X4
	PXOR	X7, X5
	PXOR	X5, X4
	MOVL	X4, (DX)
	RET
