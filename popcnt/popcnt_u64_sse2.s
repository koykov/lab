#include "textflag.h"

TEXT ·popcntU64SSE2s(SB), NOSPLIT, $0-32
    MOVQ data+0(FP), SI   // point to slice start (SI = &data[0])
    MOVQ len+8(FP), CX    // slice len (CX = len(data))
    XORQ AX, AX           // reset acc (AX = 0)

    // check if slice len is less than 2
    CMPQ CX, $2
    JL   remainder        // go to remainder label

    // prepare SSE2 regs
    XORPS X0, X0          // clean reg X0 (acc)
    MOVQ $0x5555555555555555, DX
    MOVQ DX, X1           // apply mask 0x5555555555555555 to X1
    MOVQ $0x3333333333333333, DX
    MOVQ DX, X2           // apply mask 0x3333333333333333 to X2
    MOVQ $0x0F0F0F0F0F0F0F0F, DX
    MOVQ DX, X3           // apply mask 0x0F0F0F0F0F0F0F0F to X3

sse_loop:
    // load 2 numbers (128 бит) to X4
    MOVUPS (SI), X4       // apply MOVUPS to load unaligned data

    ANDPS X1, X4          // X4 = X4 & 0x5555555555555555
    PSRLQ $1, X4          // X4 = X4 >> 1
    ANDPS X1, X4          // X4 = (X4 >> 1) & 0x5555555555555555
    PADDQ X4, X0          // X0 += X4

    ANDPS X2, X4          // X4 = X4 & 0x3333333333333333
    PSRLQ $2, X4          // X4 = X4 >> 2
    ANDPS X2, X4          // X4 = (X4 >> 2) & 0x3333333333333333
    PADDQ X4, X0          // X0 += X4

    ANDPS X3, X4          // X4 = X4 & 0x0F0F0F0F0F0F0F0F
    PSRLQ $4, X4          // X4 = X4 >> 4
    ANDPS X3, X4          // X4 = (X4 >> 4) & 0x0F0F0F0F0F0F0F0F
    PADDQ X4, X0          // X0 += X4

    // switch to next block
    ADDQ $16, SI          // SI += 16 (2 64-bit numbers)
    SUBQ $2, CX           // CX -= 2
    CMPQ CX, $2
    JGE  sse_loop         // repeat till CX >= 2

    // sum X0 to AX
    MOVQ X0, AX           // extract low 64 bits from X0 to AX
    PSHUFD $0b11101110, X0, X1  // move high 64 bits to low
    MOVQ X1, DX           // move high 64 bits to DX
    ADDQ DX, AX           // AX+DX

remainder:
    // process remain number (less than 2)
    CMPQ CX, $0
    JE   done

    // start loop to process remain numbers using POPCNT
    XORQ DX, DX
remainder_loop:
    POPCNTQ (SI), DX
    ADDQ DX, AX
    ADDQ $8, SI
    LOOP remainder_loop

done:
    MOVQ AX, ret+24(FP)
    RET
