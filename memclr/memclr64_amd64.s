#include "textflag.h"

// func memclr64SSE2(p []uint64)
TEXT ·memclr64SSE2(SB), NOSPLIT, $0-24
    MOVQ    p_data+0(FP), DI   
    MOVQ    p_len+8(FP), CX    

    
    CMPQ    CX, $8
    JB      small

    
    XORPS   X0, X0            
    XORPS   X1, X1            
    XORPS   X2, X2            
    XORPS   X3, X3            

    
    MOVQ    CX, DX
    SHRQ    $3, DX            
    JZ      tail_sse2

sse2_loop:
    MOVUPS  X0, (DI)          
    MOVUPS  X1, 16(DI)        
    MOVUPS  X2, 32(DI)        
    MOVUPS  X3, 48(DI)        
    ADDQ    $64, DI           
    DECQ    DX                
    JNZ     sse2_loop         

    
tail_sse2:
    ANDQ    $7, CX            
    JZ      done

    
    CMPQ    CX, $4
    JB      lt4
    MOVUPS  X0, (DI)
    MOVUPS  X1, 16(DI)
    ADDQ    $32, DI
    SUBQ    $4, CX

lt4:
    
    CMPQ    CX, $2
    JB      lt2
    MOVQ    X0, (DI)
    MOVQ    X0, 8(DI)
    ADDQ    $16, DI
    SUBQ    $2, CX

lt2:
    
    TESTQ   CX, CX
    JZ      done
    MOVQ    $0, (DI)

done:
    RET

    
small:
    TESTQ   CX, CX
    JZ      done

    
    MOVQ    CX, DX
    SHRQ    $2, DX
    JZ      tail_small

loop_small:
    MOVQ    $0, (DI)
    MOVQ    $0, 8(DI)
    MOVQ    $0, 16(DI)
    MOVQ    $0, 24(DI)
    ADDQ    $32, DI
    DECQ    DX
    JNZ     loop_small

tail_small:
    ANDQ    $3, CX
    JZ      done

    
    CMPQ    CX, $2
    JB      lt2_small
    MOVQ    $0, (DI)
    MOVQ    $0, 8(DI)
    ADDQ    $16, DI
    SUBQ    $2, CX

lt2_small:
    TESTQ   CX, CX
    JZ      done
    MOVQ    $0, (DI)
    RET

// func memclr64AVX2(p []uint64)
TEXT ·memclr64AVX2(SB), NOSPLIT, $0-24
    
    MOVQ    p_data+0(FP), DI   
    MOVQ    p_len+8(FP), CX    

    
    CMPQ    CX, $2097152       
    JAE     huge_clearing

    
    CMPQ    CX, $16
    JB      small

    
    VZEROUPPER
    VPXOR   Y0, Y0, Y0        
    VPXOR   Y1, Y1, Y1        
    VPXOR   Y2, Y2, Y2        
    VPXOR   Y3, Y3, Y3        

    
    MOVQ    CX, DX
    SHRQ    $4, DX            
    JZ      tail_avx2

avx2_loop:
    VMOVDQU Y0, (DI)          
    VMOVDQU Y1, 32(DI)        
    VMOVDQU Y2, 64(DI)        
    VMOVDQU Y3, 96(DI)        
    ADDQ    $128, DI          
    DECQ    DX                
    JNZ     avx2_loop         

    
tail_avx2:
    ANDQ    $15, CX           
    JZ      done_avx2         

    
    CMPQ    CX, $8
    JB      lt8_avx2
    VMOVDQU Y0, (DI)
    VMOVDQU Y1, 32(DI)
    ADDQ    $64, DI
    SUBQ    $8, CX

lt8_avx2:
    
    CMPQ    CX, $4
    JB      lt4_avx2
    VMOVDQU Y0, (DI)
    ADDQ    $32, DI
    SUBQ    $4, CX

lt4_avx2:
    
    CMPQ    CX, $2
    JB      lt2_avx2
    VMOVQ   X0, (DI)
    VMOVQ   X0, 8(DI)
    ADDQ    $16, DI
    SUBQ    $2, CX

lt2_avx2:
    
    TESTQ   CX, CX
    JZ      done_avx2
    MOVQ    $0, (DI)

done_avx2:
    VZEROUPPER
    RET

    
huge_clearing:
    VZEROUPPER
    VPXOR   Y0, Y0, Y0        
    VPXOR   Y1, Y1, Y1        

    
    MOVQ    CX, DX
    SHRQ    $5, DX            
    JZ      huge_tail

huge_loop:
    PREFETCHNTA 1024(DI)      
    VMOVNTDQ Y0, (DI)         
    VMOVNTDQ Y1, 32(DI)
    VMOVNTDQ Y0, 64(DI)
    VMOVNTDQ Y1, 96(DI)
    VMOVNTDQ Y0, 128(DI)
    VMOVNTDQ Y1, 160(DI)
    VMOVNTDQ Y0, 192(DI)
    VMOVNTDQ Y1, 224(DI)
    ADDQ    $256, DI          
    DECQ    DX                
    JNZ     huge_loop         

    SFENCE                    
    ANDQ    $31, CX           
    JZ      done

    
huge_tail:
    VMOVDQU Y0, (DI)
    VMOVDQU Y1, 32(DI)
    VMOVDQU Y0, 64(DI)
    VMOVDQU Y1, 96(DI)
    JMP     done

    
small:
    TESTQ   CX, CX
    JZ      done

    
    MOVQ    CX, DX
    SHRQ    $3, DX
    JZ      tail_small

loop_small:
    MOVQ    $0, (DI)
    MOVQ    $0, 8(DI)
    MOVQ    $0, 16(DI)
    MOVQ    $0, 24(DI)
    MOVQ    $0, 32(DI)
    MOVQ    $0, 40(DI)
    MOVQ    $0, 48(DI)
    MOVQ    $0, 56(DI)
    ADDQ    $64, DI
    DECQ    DX
    JNZ     loop_small

tail_small:
    ANDQ    $7, CX
    JZ      done

    
    CMPQ    CX, $4
    JB      lt4_small
    MOVQ    $0, (DI)
    MOVQ    $0, 8(DI)
    MOVQ    $0, 16(DI)
    MOVQ    $0, 24(DI)
    ADDQ    $32, DI
    SUBQ    $4, CX

lt4_small:
    CMPQ    CX, $2
    JB      lt2_small
    MOVQ    $0, (DI)
    MOVQ    $0, 8(DI)
    ADDQ    $16, DI
    SUBQ    $2, CX

lt2_small:
    TESTQ   CX, CX
    JZ      done
    MOVQ    $0, (DI)

done:
    RET

// func memclr64AVX512(p []uint64)
TEXT ·memclr64AVX512(SB), NOSPLIT, $0-24
    
    MOVQ    p_data+0(FP), DI   
    MOVQ    p_len+8(FP), CX    

    
    CMPQ    CX, $2097152       
    JAE     huge_clearing

    
    CMPQ    CX, $32
    JB      small

    
    VPXORQ  Z0, Z0, Z0        
    VPXORQ  Z1, Z1, Z1        
    VPXORQ  Z2, Z2, Z2        
    VPXORQ  Z3, Z3, Z3        

    
    MOVQ    CX, DX
    SHRQ    $5, DX            
    JZ      tail_avx512

avx512_loop:
    VMOVDQU64 Z0, (DI)        
    VMOVDQU64 Z1, 64(DI)
    VMOVDQU64 Z2, 128(DI)
    VMOVDQU64 Z3, 192(DI)
    ADDQ    $256, DI          
    DECQ    DX                
    JNZ     avx512_loop       

    
tail_avx512:
    ANDQ    $31, CX           
    JZ      done_avx512       

    
    CMPQ    CX, $16
    JB      lt16_avx512
    VMOVDQU64 Z0, (DI)
    VMOVDQU64 Z1, 64(DI)
    ADDQ    $128, DI
    SUBQ    $16, CX

lt16_avx512:
    
    CMPQ    CX, $8
    JB      lt8_avx512
    VMOVDQU64 Z0, (DI)
    ADDQ    $64, DI
    SUBQ    $8, CX

lt8_avx512:
    
    CMPQ    CX, $4
    JB      lt4_avx512
    VMOVDQU Y0, (DI)
    ADDQ    $32, DI
    SUBQ    $4, CX

lt4_avx512:
    
    CMPQ    CX, $2
    JB      lt2_avx512
    VMOVQ   X0, (DI)
    VMOVQ   X0, 8(DI)
    ADDQ    $16, DI
    SUBQ    $2, CX

lt2_avx512:
    
    TESTQ   CX, CX
    JZ      done_avx512
    MOVQ    $0, (DI)

done_avx512:
    VZEROUPPER
    RET

    
huge_clearing:
    VPXORQ  Z0, Z0, Z0        
    VPXORQ  Z1, Z1, Z1        

    
    MOVQ    CX, DX
    SHRQ    $6, DX            
    JZ      huge_tail

huge_loop:
    PREFETCHNTA 2048(DI)      
    VMOVNTDQ Z0, (DI)         
    VMOVNTDQ Z1, 64(DI)
    VMOVNTDQ Z0, 128(DI)
    VMOVNTDQ Z1, 192(DI)
    VMOVNTDQ Z0, 256(DI)
    VMOVNTDQ Z1, 320(DI)
    VMOVNTDQ Z0, 384(DI)
    VMOVNTDQ Z1, 448(DI)
    ADDQ    $512, DI          
    DECQ    DX                
    JNZ     huge_loop         

    SFENCE                    
    ANDQ    $63, CX           
    JZ      done

    
huge_tail:
    CMPQ    CX, $32
    JB      huge_lt32
    VMOVDQU64 Z0, (DI)
    VMOVDQU64 Z1, 64(DI)
    VMOVDQU64 Z0, 128(DI)
    VMOVDQU64 Z1, 192(DI)
    ADDQ    $256, DI
    SUBQ    $32, CX

huge_lt32:
    JMP     tail_avx512       

    
small:
    TESTQ   CX, CX
    JZ      done

    
    MOVQ    CX, DX
    SHRQ    $4, DX
    JZ      tail_small

loop_small:
    MOVQ    $0, (DI)
    MOVQ    $0, 8(DI)
    MOVQ    $0, 16(DI)
    MOVQ    $0, 24(DI)
    MOVQ    $0, 32(DI)
    MOVQ    $0, 40(DI)
    MOVQ    $0, 48(DI)
    MOVQ    $0, 56(DI)
    MOVQ    $0, 64(DI)
    MOVQ    $0, 72(DI)
    MOVQ    $0, 80(DI)
    MOVQ    $0, 88(DI)
    MOVQ    $0, 96(DI)
    MOVQ    $0, 104(DI)
    MOVQ    $0, 112(DI)
    MOVQ    $0, 120(DI)
    ADDQ    $128, DI
    DECQ    DX
    JNZ     loop_small

tail_small:
    ANDQ    $15, CX
    JZ      done

    
    CMPQ    CX, $8
    JB      lt8_small
    MOVQ    $0, (DI)
    MOVQ    $0, 8(DI)
    MOVQ    $0, 16(DI)
    MOVQ    $0, 24(DI)
    MOVQ    $0, 32(DI)
    MOVQ    $0, 40(DI)
    MOVQ    $0, 48(DI)
    MOVQ    $0, 56(DI)
    ADDQ    $64, DI
    SUBQ    $8, CX

lt8_small:
    CMPQ    CX, $4
    JB      lt4_small
    MOVQ    $0, (DI)
    MOVQ    $0, 8(DI)
    MOVQ    $0, 16(DI)
    MOVQ    $0, 24(DI)
    ADDQ    $32, DI
    SUBQ    $4, CX

lt4_small:
    CMPQ    CX, $2
    JB      lt2_small
    MOVQ    $0, (DI)
    MOVQ    $0, 8(DI)
    ADDQ    $16, DI
    SUBQ    $2, CX

lt2_small:
    TESTQ   CX, CX
    JZ      done
    MOVQ    $0, (DI)

done:
    VZEROUPPER
    RET
