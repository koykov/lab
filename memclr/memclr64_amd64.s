#include "textflag.h"

// func memclr64SSE2(p []uint64)
TEXT ·memclr64SSE2(SB),NOSPLIT,$0
    MOVQ    p_data+0(FP), DI        // point to slice start (SI = &data[0])
    MOVQ    p_len+8(FP), CX         // slice len (CX = len(data))

    SHRQ    $2, CX                  // CX = unroll loop by 4
    PXOR    X0, X0                  // init XMM0 with 128 zero bits

loop:
    MOVUPS  X0, (DI)                // write lo 16 bytes (2 uint64 numbers) at once
    MOVUPS  X0, 16(DI)              // write hi 16 bytes (4 uint64 per iteration in total)
    ADDQ    $32, DI                 // switch to next 32 bytes
    DECQ    CX
    JNZ     loop

    RET
