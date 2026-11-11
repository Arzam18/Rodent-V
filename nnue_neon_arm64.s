#include "textflag.h"

// All accumulator kernels operate on the active 512-neuron network.
TEXT ·addSingleNEON(SB), NOSPLIT, $0-16
	MOVD a+0(FP), R0
	MOVD w+8(FP), R1
	MOVD $64, R2
add_single_loop:
	VLD1 (R0), [V0.H8]
	VLD1 (R1), [V1.H8]
	VADD V1.H8, V0.H8, V0.H8
	VST1 [V0.H8], (R0)
	ADD $16, R0
	ADD $16, R1
	SUB $1, R2
	CBNZ R2, add_single_loop
	RET

TEXT ·subSingleNEON(SB), NOSPLIT, $0-16
	MOVD a+0(FP), R0
	MOVD w+8(FP), R1
	MOVD $64, R2
sub_single_loop:
	VLD1 (R0), [V0.H8]
	VLD1 (R1), [V1.H8]
	VSUB V1.H8, V0.H8, V0.H8
	VST1 [V0.H8], (R0)
	ADD $16, R0
	ADD $16, R1
	SUB $1, R2
	CBNZ R2, sub_single_loop
	RET

// a0 += to0-from0; a1 += to1-from1
TEXT ·moveNEON(SB), NOSPLIT, $0-48
	MOVD a0+0(FP), R0
	MOVD a1+8(FP), R1
	MOVD wFrom0+16(FP), R2
	MOVD wTo0+24(FP), R3
	MOVD wFrom1+32(FP), R4
	MOVD wTo1+40(FP), R5
	MOVD $64, R6
move_loop:
	VLD1 (R0), [V0.H8]
	VLD1 (R2), [V1.H8]
	VLD1 (R3), [V2.H8]
	VSUB V1.H8, V2.H8, V2.H8
	VADD V2.H8, V0.H8, V0.H8
	VST1 [V0.H8], (R0)
	VLD1 (R1), [V3.H8]
	VLD1 (R4), [V4.H8]
	VLD1 (R5), [V5.H8]
	VSUB V4.H8, V5.H8, V5.H8
	VADD V5.H8, V3.H8, V3.H8
	VST1 [V3.H8], (R1)
	ADD $16, R0; ADD $16, R1
	ADD $16, R2; ADD $16, R3
	ADD $16, R4; ADD $16, R5
	SUB $1, R6
	CBNZ R6, move_loop
	RET

// a += to-from-capture for both perspectives.
TEXT ·captureNEON(SB), NOSPLIT, $0-64
	MOVD a0+0(FP), R0
	MOVD a1+8(FP), R1
	MOVD wTo0+16(FP), R2
	MOVD wFrom0+24(FP), R3
	MOVD wCap0+32(FP), R4
	MOVD wTo1+40(FP), R5
	MOVD wFrom1+48(FP), R6
	MOVD wCap1+56(FP), R7
	MOVD $64, R8
capture_loop:
	VLD1 (R0), [V0.H8]
	VLD1 (R2), [V1.H8]
	VLD1 (R3), [V2.H8]
	VLD1 (R4), [V3.H8]
	VSUB V2.H8, V1.H8, V1.H8
	VSUB V3.H8, V1.H8, V1.H8
	VADD V1.H8, V0.H8, V0.H8
	VST1 [V0.H8], (R0)
	VLD1 (R1), [V4.H8]
	VLD1 (R5), [V5.H8]
	VLD1 (R6), [V6.H8]
	VLD1 (R7), [V7.H8]
	VSUB V6.H8, V5.H8, V5.H8
	VSUB V7.H8, V5.H8, V5.H8
	VADD V5.H8, V4.H8, V4.H8
	VST1 [V4.H8], (R1)
	ADD $16, R0; ADD $16, R1
	ADD $16, R2; ADD $16, R3; ADD $16, R4
	ADD $16, R5; ADD $16, R6; ADD $16, R7
	SUB $1, R8
	CBNZ R8, capture_loop
	RET

// a += kingTo-kingFrom+rookTo-rookFrom for both perspectives.
TEXT ·castleNEON(SB), NOSPLIT, $0-80
	MOVD a0+0(FP), R0
	MOVD a1+8(FP), R1
	MOVD wKFrom0+16(FP), R2
	MOVD wKTo0+24(FP), R3
	MOVD wRFrom0+32(FP), R4
	MOVD wRTo0+40(FP), R5
	MOVD wKFrom1+48(FP), R6
	MOVD wKTo1+56(FP), R7
	MOVD wRFrom1+64(FP), R8
	MOVD wRTo1+72(FP), R9
	MOVD $64, R10
castle_loop:
	VLD1 (R0), [V0.H8]
	VLD1 (R2), [V1.H8]
	VLD1 (R3), [V2.H8]
	VLD1 (R4), [V3.H8]
	VLD1 (R5), [V4.H8]
	VSUB V1.H8, V2.H8, V2.H8
	VSUB V3.H8, V4.H8, V4.H8
	VADD V4.H8, V2.H8, V2.H8
	VADD V2.H8, V0.H8, V0.H8
	VST1 [V0.H8], (R0)
	VLD1 (R1), [V5.H8]
	VLD1 (R6), [V6.H8]
	VLD1 (R7), [V7.H8]
	VLD1 (R8), [V8.H8]
	VLD1 (R9), [V9.H8]
	VSUB V6.H8, V7.H8, V7.H8
	VSUB V8.H8, V9.H8, V9.H8
	VADD V9.H8, V7.H8, V7.H8
	VADD V7.H8, V5.H8, V5.H8
	VST1 [V5.H8], (R1)
	ADD $16, R0; ADD $16, R1
	ADD $16, R2; ADD $16, R3; ADD $16, R4; ADD $16, R5
	ADD $16, R6; ADD $16, R7; ADD $16, R8; ADD $16, R9
	SUB $1, R10
	CBNZ R10, castle_loop
	RET

// Three-operand forms: dst = src + delta.
TEXT ·moveNEON3(SB), NOSPLIT, $0-64
	MOVD dst0+0(FP), R0
	MOVD src0+8(FP), R1
	MOVD dst1+16(FP), R2
	MOVD src1+24(FP), R3
	MOVD wFrom0+32(FP), R4
	MOVD wTo0+40(FP), R5
	MOVD wFrom1+48(FP), R6
	MOVD wTo1+56(FP), R7
	MOVD $64, R8
move3_loop:
	VLD1 (R1), [V0.H8]
	VLD1 (R4), [V1.H8]
	VLD1 (R5), [V2.H8]
	VSUB V1.H8, V2.H8, V2.H8
	VADD V2.H8, V0.H8, V0.H8
	VST1 [V0.H8], (R0)
	VLD1 (R3), [V3.H8]
	VLD1 (R6), [V4.H8]
	VLD1 (R7), [V5.H8]
	VSUB V4.H8, V5.H8, V5.H8
	VADD V5.H8, V3.H8, V3.H8
	VST1 [V3.H8], (R2)
	ADD $16, R0; ADD $16, R1; ADD $16, R2; ADD $16, R3
	ADD $16, R4; ADD $16, R5; ADD $16, R6; ADD $16, R7
	SUB $1, R8
	CBNZ R8, move3_loop
	RET

TEXT ·captureNEON3(SB), NOSPLIT, $0-80
	MOVD dst0+0(FP), R0
	MOVD src0+8(FP), R1
	MOVD dst1+16(FP), R2
	MOVD src1+24(FP), R3
	MOVD wTo0+32(FP), R4
	MOVD wFrom0+40(FP), R5
	MOVD wCap0+48(FP), R6
	MOVD wTo1+56(FP), R7
	MOVD wFrom1+64(FP), R8
	MOVD wCap1+72(FP), R9
	MOVD $64, R10
capture3_loop:
	VLD1 (R1), [V0.H8]
	VLD1 (R4), [V1.H8]
	VLD1 (R5), [V2.H8]
	VLD1 (R6), [V3.H8]
	VSUB V2.H8, V1.H8, V1.H8
	VSUB V3.H8, V1.H8, V1.H8
	VADD V1.H8, V0.H8, V0.H8
	VST1 [V0.H8], (R0)
	VLD1 (R3), [V4.H8]
	VLD1 (R7), [V5.H8]
	VLD1 (R8), [V6.H8]
	VLD1 (R9), [V7.H8]
	VSUB V6.H8, V5.H8, V5.H8
	VSUB V7.H8, V5.H8, V5.H8
	VADD V5.H8, V4.H8, V4.H8
	VST1 [V4.H8], (R2)
	ADD $16, R0; ADD $16, R1; ADD $16, R2; ADD $16, R3
	ADD $16, R4; ADD $16, R5; ADD $16, R6; ADD $16, R7; ADD $16, R8; ADD $16, R9
	SUB $1, R10
	CBNZ R10, capture3_loop
	RET

TEXT ·castleNEON3(SB), NOSPLIT, $0-96
	MOVD dst0+0(FP), R0
	MOVD src0+8(FP), R1
	MOVD dst1+16(FP), R2
	MOVD src1+24(FP), R3
	MOVD wKFrom0+32(FP), R4
	MOVD wKTo0+40(FP), R5
	MOVD wRFrom0+48(FP), R6
	MOVD wRTo0+56(FP), R7
	MOVD wKFrom1+64(FP), R8
	MOVD wKTo1+72(FP), R9
	MOVD wRFrom1+80(FP), R10
	MOVD wRTo1+88(FP), R11
	MOVD $64, R12
castle3_loop:
	VLD1 (R1), [V0.H8]
	VLD1 (R4), [V1.H8]
	VLD1 (R5), [V2.H8]
	VLD1 (R6), [V3.H8]
	VLD1 (R7), [V4.H8]
	VSUB V1.H8, V2.H8, V2.H8
	VSUB V3.H8, V4.H8, V4.H8
	VADD V4.H8, V2.H8, V2.H8
	VADD V2.H8, V0.H8, V0.H8
	VST1 [V0.H8], (R0)
	VLD1 (R3), [V5.H8]
	VLD1 (R8), [V6.H8]
	VLD1 (R9), [V7.H8]
	VLD1 (R10), [V8.H8]
	VLD1 (R11), [V9.H8]
	VSUB V6.H8, V7.H8, V7.H8
	VSUB V8.H8, V9.H8, V9.H8
	VADD V9.H8, V7.H8, V7.H8
	VADD V7.H8, V5.H8, V5.H8
	VST1 [V5.H8], (R2)
	ADD $16, R0; ADD $16, R1; ADD $16, R2; ADD $16, R3
	ADD $16, R4; ADD $16, R5; ADD $16, R6; ADD $16, R7
	ADD $16, R8; ADD $16, R9; ADD $16, R10; ADD $16, R11
	SUB $1, R12
	CBNZ R12, castle3_loop
	RET

// Exact signed-byte dot products for the u8 x i8 NNUE layer. The Go caller
// converts each activation to (u8-128), repeated four times per 16-byte
// vector, then adds 128*sum(weight) to recover the original unsigned dot.
TEXT ·dotprodL1(SB), NOSPLIT, $0-24
	MOVD sums+0(FP), R0
	MOVD activations+8(FP), R1
	MOVD tiled+16(FP), R2
	MOVD $128, R3
	VEOR V21.B16, V21.B16, V21.B16
	VEOR V22.B16, V22.B16, V22.B16
	VEOR V23.B16, V23.B16, V23.B16
	VEOR V24.B16, V24.B16, V24.B16
dotprod_l1_loop:
	VLD1.P 16(R1), [V16.B16]
	VLD1.P 16(R2), [V17.B16]
	WORD $0x4e919615 // SDOT V21.4S, V16.16B, V17.16B
	VLD1.P 16(R2), [V18.B16]
	WORD $0x4e929616 // SDOT V22.4S, V16.16B, V18.16B
	VLD1.P 16(R2), [V19.B16]
	WORD $0x4e939617 // SDOT V23.4S, V16.16B, V19.16B
	VLD1.P 16(R2), [V20.B16]
	WORD $0x4e949618 // SDOT V24.4S, V16.16B, V20.16B
	SUB $1, R3
	CMP $0, R3
	BNE dotprod_l1_loop
	VST1.P [V21.S4], 16(R0)
	VST1.P [V22.S4], 16(R0)
	VST1.P [V23.S4], 16(R0)
	VST1.P [V24.S4], 16(R0)
	RET

// Sum SCReLU(a0)*w0 + SCReLU(a1)*w1 across the 512 accumulator lanes.
// The raw words below are AArch64 NEON SMIN/SMAX/SMULL/SMULL2/SSHLL/SSHLL2/MLA.
TEXT ·evalNEON(SB), NOSPLIT, $0-40
	MOVD a0+0(FP), R0
	MOVD a1+8(FP), R1
	MOVD w0+16(FP), R2
	MOVD w1+24(FP), R3
	MOVD sum+32(FP), R4
	MOVD $64, R5
	WORD $0x4f008400 // MOVI V0.8H, #0
	WORD $0x4f0787e1 // MOVI V1.8H, #255
	WORD $0x4f000402 // MOVI V2.4S, #0
	WORD $0x4f000403 // MOVI V3.4S, #0
eval_neon_loop:
	VLD1.P 16(R0), [V4.H8]
	VLD1.P 16(R1), [V5.H8]
	VLD1.P 16(R2), [V6.H8]
	VLD1.P 16(R3), [V17.H8]
	WORD $0x4e616c84 // SMIN V4.8H, V4.8H, V1.8H
	WORD $0x4e616ca5 // SMIN V5.8H, V5.8H, V1.8H
	WORD $0x4e606484 // SMAX V4.8H, V4.8H, V0.8H
	WORD $0x4e6064a5 // SMAX V5.8H, V5.8H, V0.8H
	WORD $0x0e64c087 // SMULL V7.4S, V4.4H, V4.4H
	WORD $0x4e64c084 // SMULL2 V4.4S, V4.8H, V4.8H
	WORD $0x0e65c0b2 // SMULL V18.4S, V5.4H, V5.4H
	WORD $0x4e65c0a5 // SMULL2 V5.4S, V5.8H, V5.8H
	WORD $0x0f10a4d0 // SSHLL V16.4S, V6.4H, #0
	WORD $0x4f10a4c6 // SSHLL2 V6.4S, V6.8H, #0
	WORD $0x4eb094e2 // MLA V2.4S, V7.4S, V16.4S
	WORD $0x4ea69483 // MLA V3.4S, V4.4S, V6.4S
	WORD $0x0f10a624 // SSHLL V4.4S, V17.4H, #0
	WORD $0x4f10a626 // SSHLL2 V6.4S, V17.8H, #0
	WORD $0x4ea49642 // MLA V2.4S, V18.4S, V4.4S
	WORD $0x4ea694a3 // MLA V3.4S, V5.4S, V6.4S
	SUB $1, R5
	CMP $0, R5
	BNE eval_neon_loop
	VADD V2.S4, V3.S4, V0.S4
	VADDV V0.S4, V0
	WORD $0xbd000080 // STR S0, [X4]
	RET
