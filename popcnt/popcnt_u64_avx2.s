#include "textflag.h"

TEXT ·popcntU64AVX2s(SB), NOSPLIT, $0-32
    MOVQ data+0(FP), SI   // point to slice start (SI = &data[0])
    MOVQ len+8(FP), CX    // slice len (CX = len(data))
    XORQ AX, AX           // reset acc (AX = 0)

    // check if slice len is less than 4
    CMPQ CX, $4
    JL   remainder        // go to remainder label

    // prepare AVX-2 regs
    VPXOR Y0, Y0, Y0     // clean reg Y0 (acc)
    MOVQ $0x5555555555555555, DX
    VPBROADCASTQ DX, Y1  // apply mask 0x5555555555555555 to Y1
    MOVQ $0x3333333333333333, DX
    VPBROADCASTQ DX, Y2  // apply mask 0x3333333333333333 to Y2
    MOVQ $0x0F0F0F0F0F0F0F0F, DX
    VPBROADCASTQ DX, Y3  // apply mask 0x0F0F0F0F0F0F0F0F to Y3

avx2_loop:
    // load 4 numbers (256 бит) to Y4
    VMOVDQU (SI), Y4

    VPAND Y4, Y1, Y5     // Y5 = Y4 & 0x5555555555555555
    VPSRLQ $1, Y4, Y6    // Y6 = Y4 >> 1
    VPAND Y6, Y1, Y6     // Y6 = (Y4 >> 1) & 0x5555555555555555
    VPADDQ Y5, Y6, Y4    // Y4 = Y5 + Y6

    VPAND Y4, Y2, Y5     // Y5 = Y4 & 0x3333333333333333
    VPSRLQ $2, Y4, Y6    // Y6 = Y4 >> 2
    VPAND Y6, Y2, Y6     // Y6 = (Y4 >> 2) & 0x3333333333333333
    VPADDQ Y5, Y6, Y4    // Y4 = Y5 + Y6

    VPAND Y4, Y3, Y5     // Y5 = Y4 & 0x0F0F0F0F0F0F0F0F
    VPSRLQ $4, Y4, Y6    // Y6 = Y4 >> 4
    VPAND Y6, Y3, Y6     // Y6 = (Y4 >> 4) & 0x0F0F0F0F0F0F0F0F
    VPADDQ Y5, Y6, Y4    // Y4 = Y5 + Y6

    // sum result to Y0
    VPADDQ Y4, Y0, Y0    // Y0 += Y4

    // switch to next block
    ADDQ $32, SI         // SI += 32 (4 64-bit numbers)
    SUBQ $4, CX          // CX -= 4
    CMPQ CX, $4
    JGE  avx2_loop       // repeat till CX >= 4

    // sum Y0 to AX
    VEXTRACTI128 $1, Y0, X1  // extract high 128 bits from Y0 to X1
    VPADDQ X0, X1, X0        // X0+X1
    VPSHUFD $0b11101110, X0, X1  // move high 64 bits to low
    VPADDQ X0, X1, X0        // X0+X1
    VMOVQ X0, AX             // move result to AX

remainder:
    // process remain number (less than 4)
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
