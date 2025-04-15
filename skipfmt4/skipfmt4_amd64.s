#include "textflag.h"

// Whitespace mask (space, tab, cr, lf)
DATA ·whitespaceMask<>+0(SB)/16, $" \t\n\r"
GLOBL ·whitespaceMask<>(SB), RODATA, $16

// func skipfmt4(p []byte) int
TEXT ·skipfmt4SSE2(SB), NOSPLIT, $0-32
    MOVQ p_base+0(FP), SI     // SI = pointer to byte array
    MOVQ p_len+8(FP), CX      // CX = length of array
    XORQ AX, AX               // AX = counter (position)

    CMPQ CX, $0
    JE   not_found            // empty slice, return -1

    // Load mask of whitespace chars (space, tab, cr, lf)
    MOVOU ·whitespaceMask<>(SB), X0

block_loop:
    CMPQ CX, $16
    JL   tail_processing      // if less than 16 bytes left

    // Load 16 bytes
    MOVOU (SI)(AX*1), X1

    // Compare with whitespace chars
    PCMPEQB X0, X1

    // Get mask of matches
    PMOVMSKB X1, BX

    // Check if all 16 are whitespace
    CMPL BX, $0xFFFF
    JNE  find_non_ws

    // All 16 matched, continue
    ADDQ $16, AX
    SUBQ $16, CX
    JMP  block_loop

find_non_ws:
    // Find first non-matching bit (0-based)
    BSFW BX, BX
    // Convert bit position to byte offset
    ADDQ BX, AX
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

    // Found non-whitespace
    MOVQ AX, ret+24(FP)
    RET

next_tail:
    INCQ AX
    DECQ CX
    JMP  tail_processing

not_found:
    MOVQ $-1, ret+24(FP)
    RET
