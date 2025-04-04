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

// func memclr64AVX2(p []byte)
TEXT ·memclr64AVX2(SB),NOSPLIT,$0-24
    MOVQ    p_data+0(FP), DI   // DI = pointer
    MOVQ    p_len+8(FP), CX    // CX = length

    // Fast path для маленьких блоков (<=128 байт)
    CMPQ    CX, $128
    JBE     small

    // Выравнивание до 32 байт
    MOVQ    DI, AX
    ANDQ    $31, AX
    JZ      aligned
    MOVQ    $32, BX
    SUBQ    AX, BX

    XORQ    AX, AX
align_loop:
    MOVB    AX, (DI)
    INCQ    DI
    DECQ    CX
    DECQ    BX
    JNZ     align_loop

aligned:
    // Основной цикл с разверткой 4x
    MOVQ    CX, BX
    SHRQ    $7, BX           // BX = количество 128-байтных блоков (4x32)
    JZ      tail

    VZEROUPPER
    VPXOR   Y0, Y0, Y0
    VPXOR   Y1, Y1, Y1      // Используем второй регистр для параллелизации

avx_loop:
    VMOVDQU Y0, 0(DI)       // Развертка 4x32 байта
    VMOVDQU Y1, 32(DI)
    VMOVDQU Y0, 64(DI)
    VMOVDQU Y1, 96(DI)
    ADDQ    $128, DI        // Уменьшаем количество ADDQ
    DECQ    BX
    JNZ     avx_loop

    VZEROUPPER

tail:
    // Обработка остатка (0-127 байт)
    ANDQ    $127, CX
    JZ      done

small:
    // Оптимизированная очистка остатка
    XORQ    AX, AX
    CMPQ    CX, $32
    JB      very_small

    // Очистка 32-байтными блоками
    VMOVDQU Y0, (DI)
    ADDQ    $32, DI
    SUBQ    $32, CX
    JMP     small

very_small:
    TESTQ   CX, CX
    JZ      done
    MOVQ    AX, (DI)        // Очистка 8 байт за раз
    MOVQ    AX, -8(DI)(CX*1)
done:
    RET
