#include "textflag.h"

// func memclr64SSE2(p []uint64)
TEXT ·memclr64SSE2(SB),NOSPLIT,$0-24
    MOVQ    p_data+0(FP), DI   // point to slice start (DI = &p[0])
    MOVQ    p_len+8(FP), CX    // slice len (CX = len(p))

    // process tail (0-3 items)
    MOVQ    CX, BX
    ANDQ    $3, BX             // BX = len % 4
    JZ      block_processing   // goto main block processing

    XORQ    AX, AX
tail_loop:
    MOVQ    AX, (DI)           // clear single item
    ADDQ    $8, DI
    DECQ    BX
    JNZ     tail_loop

    SUBQ    BX, CX             // reduce length

block_processing:
    // process 4 items per iteration
    SHRQ    $2, CX             // CX = CX >> 2 (count of blocks)
    JZ      done

    PXOR    X0, X0             // fill up XMM0 with 128 zero bits (16 bytes)

loop:
    MOVUPS  X0, (DI)           // write 128 zero bits (16 bytes/2 items) to addr DI
    MOVUPS  X0, 16(DI)         // write another 128 bits to DI with offset 16 (4 items in total)
    ADDQ    $32, DI
    DECQ    CX
    JNZ     loop

done:
    RET

// func memclr64AVX2(p []uint64)
TEXT ·memclr64AVX2(SB),NOSPLIT,$0
    MOVQ    p_data+0(FP), DI   // point to slice start (DI = &p[0])
    MOVQ    p_len+8(FP), CX    // slice len (CX = len(p))

    // process tail (0-7 items)
    MOVQ    CX, BX
    ANDQ    $7, BX             // BX = len % 8
    JZ      block_processing   // goto main block processing

    XORQ    AX, AX
tail_loop:
    MOVQ    AX, (DI)
    ADDQ    $8, DI
    DECQ    BX
    JNZ     tail_loop

    SUBQ    BX, CX             // reduce length

block_processing:
    // process 4 items per iteration
    SHRQ    $3, CX             // CX = CX >> 3 (count of blocks)
    JZ      done

    VZEROUPPER                 // prepare AVX processing (clear hi bits of YMM)
    VPXOR   Y0, Y0, Y0         // fill up YMM0 with 256 zero bits (16 bytes)

block_loop:
    VMOVDQU Y0, (DI)           // write 256 zero bits (32 bytes/4 items) to addr DI
    VMOVDQU Y0, 32(DI)         // write another 256 bits to DI with offset 32 (8 items in total)
    ADDQ    $64, DI
    DECQ    CX
    JNZ     block_loop

    VZEROUPPER

done:
    RET

// func memclr64AVX512(p []uint64)
TEXT ·memclr64AVX512(SB),NOSPLIT,$0-24
    MOVQ    p_data+0(FP), DI   // point to slice start (DI = &p[0])
    MOVQ    p_len+8(FP), CX    // slice len (CX = len(p))

    // process tail (0-7 items)
    MOVQ    CX, BX
    ANDQ    $7, BX             // BX = len % 8
    JZ      block_processing   // goto main block processing

    XORQ    AX, AX
tail_loop:
    MOVQ    AX, (DI)
    ADDQ    $8, DI
    DECQ    BX
    JNZ     tail_loop

    SUBQ    BX, CX             // reduce length

block_processing:
    // process 8 items per iteration
    SHRQ    $3, CX             // CX = CX >> 3 (count of blocks)
    JZ      done

    //VZEROUPPER                 // prepare AVX processing (clear hi bits of ZMM)
    VPXORQ  Z0, Z0, Z0         // fill up ZMM0 with 512 zero bits (16 bytes)

avx512_loop:
    VMOVDQU64 Z0, (DI)         // write 512 bits (64 bytes/8 items) at once to addr DI
    ADDQ    $64, DI
    DECQ    CX
    JNZ     avx512_loop

    //VZEROUPPER

done:
    RET
