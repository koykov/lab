#include "textflag.h"

// func skipfmt4SIMD(p []byte) int
TEXT ·skipfmt4SIMD(SB), NOSPLIT, $0-32
    MOVQ p_base+0(FP), SI     // SI = pointer to byte array
    MOVQ p_len+8(FP), CX      // CX = length of array
    XORQ AX, AX               // AX = counter (position)

    CMPQ CX, $0
    JE   not_found            // empty slice, return -1

    // Load masks for each whitespace char
    MOVOU ·spaceMask<>(SB), X0  // ' '
    MOVOU ·tabMask<>(SB), X1    // '\t'
    MOVOU ·crMask<>(SB), X2     // '\r'
    MOVOU ·lfMask<>(SB), X3     // '\n'

block_loop:
    CMPQ CX, $16
    JL   tail_processing

    // Load 16 bytes
    MOVOU (SI)(AX*1), X4

    // Compare against each whitespace char
    MOVOU X4, X5
    PCMPEQB X0, X5            // Compare with ' '

    MOVOU X4, X6
    PCMPEQB X1, X6            // Compare with '\t'
    POR X5, X6

    MOVOU X4, X7
    PCMPEQB X2, X7            // Compare with '\r'
    POR X6, X7

    MOVOU X4, X8
    PCMPEQB X3, X8            // Compare with '\n'
    POR X7, X8

    // Get match mask (1 = whitespace)
    PMOVMSKB X8, DX

    // Check if any non-whitespace exists
    CMPL DX, $0xFFFF
    JNE  find_non_ws

    // All 16 are whitespace, continue
    ADDQ $16, AX
    SUBQ $16, CX
    JMP  block_loop

find_non_ws:
    // Invert mask (now 1 = non-whitespace)
    NOTL DX
    // Find first non-whitespace
    BSFW DX, DX
    ADDQ DX, AX
    MOVQ AX, ret+24(FP)
    RET

tail_processing:
    CMPQ CX, $0
    JE   not_found

    MOVB (SI)(AX*1), BL
    CMPB BL, $' '
    JE   next_tail
    CMPB BL, $'\t'
    JE   next_tail
    CMPB BL, $'\n'
    JE   next_tail
    CMPB BL, $'\r'
    JE   next_tail

    MOVQ AX, ret+24(FP)
    RET

next_tail:
    INCQ AX
    DECQ CX
    JMP  tail_processing

not_found:
    MOVQ $-1, ret+24(FP)
    RET

// Whitespace masks (each contains 16 copies of the character)
DATA ·spaceMask<>+0(SB)/16, $"                "  // 16 spaces
DATA ·tabMask<>+0(SB)/16, $"\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t"
DATA ·crMask<>+0(SB)/16, $"\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r\r"
DATA ·lfMask<>+0(SB)/16, $"\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n"
GLOBL ·spaceMask<>(SB), RODATA, $16
GLOBL ·tabMask<>(SB), RODATA, $16
GLOBL ·crMask<>(SB), RODATA, $16
GLOBL ·lfMask<>(SB), RODATA, $16
