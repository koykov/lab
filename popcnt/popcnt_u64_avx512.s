#include "textflag.h"

TEXT ·popcntU64AVX512s(SB), NOSPLIT, $0-32
    MOVQ data+0(FP), SI   // point to slice start (SI = &data[0])
    MOVQ len+8(FP), CX    // slice len (CX = len(data))
    XORQ AX, AX           // reset acc (AX = 0)

    // check if slice len is less than 8
    CMPQ CX, $8
    JL   remainder        // go to remainder label

    // prepare AVX-512 regs
    VPXORQ Z0, Z0, Z0     // clean reg Z0 (acc)

avx512_loop:
    // load 8 numbers (512 бит) to Z1
    VMOVDQU64 (SI), Z1

    // count bits using VPOPCNTQ
    VPOPCNTQ Z1, Z1       // Z1 = nmber of bits in each part of Z1

    // sum to Z0
    VPADDQ Z1, Z0, Z0     // Z0 += Z1

    // switch to next block
    ADDQ $64, SI          // SI += 64 (8 64-bit numbers)
    SUBQ $8, CX           // CX -= 8
    CMPQ CX, $8
    JGE  avx512_loop      // repeat till CX >= 8

    // sum Z0 to AX
    VEXTRACTI64X4 $1, Z0, Y1  // extract high 256 bits from Z0 to Y1
    VPADDQ Y0, Y1, Y0         // Y0 + Y1
    VEXTRACTI128 $1, Y0, X1   // extract high 128 bits from Y0 to X1
    VPADDQ X0, X1, X0         // X0 + X1
    VPSHUFD $0b11101110, X0, X1  // move high 64 bits to low
    VPADDQ X0, X1, X0         // X0 + X1
    VMOVQ X0, AX              // store result to AX

remainder:
    // process remain number (less than 8)
    CMPQ CX, $0
    JE   done

    // start loop to process remain numbers using POPCNTQ
    XORQ DX, DX
remainder_loop:
    POPCNTQ (SI), DX
    ADDQ DX, AX
    ADDQ $8, SI
    LOOP remainder_loop

done:
    MOVQ AX, ret+24(FP)
    RET
