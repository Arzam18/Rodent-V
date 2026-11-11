//go:build !arm64

package main

// These stubs keep the common dispatch code portable. hasNEON is false on
// non-AArch64 targets, so none of these functions should be reached.
func addSingleNEON(a, w *int16)                          { panic("NEON kernel called on non-ARM64") }
func subSingleNEON(a, w *int16)                          { panic("NEON kernel called on non-ARM64") }
func moveNEON(a0, a1, wFrom0, wTo0, wFrom1, wTo1 *int16) { panic("NEON kernel called on non-ARM64") }
func captureNEON(a0, a1, wTo0, wFrom0, wCap0, wTo1, wFrom1, wCap1 *int16) {
	panic("NEON kernel called on non-ARM64")
}
func castleNEON(a0, a1, wKFrom0, wKTo0, wRFrom0, wRTo0, wKFrom1, wKTo1, wRFrom1, wRTo1 *int16) {
	panic("NEON kernel called on non-ARM64")
}
func moveNEON3(dst0, src0, dst1, src1, wFrom0, wTo0, wFrom1, wTo1 *int16) {
	panic("NEON kernel called on non-ARM64")
}
func captureNEON3(dst0, src0, dst1, src1, wTo0, wFrom0, wCap0, wTo1, wFrom1, wCap1 *int16) {
	panic("NEON kernel called on non-ARM64")
}
func castleNEON3(dst0, src0, dst1, src1, wKFrom0, wKTo0, wRFrom0, wRTo0, wKFrom1, wKTo1, wRFrom1, wRTo1 *int16) {
	panic("NEON kernel called on non-ARM64")
}
func dotprodL1(sums *int32, activations, tiled *int8) {
	panic("DotProd kernel called on non-ARM64")
}
func (acc *Accumulator) getEvalMultilayerARM64(bucket, stm int) int {
	panic("ARM64 NNUE kernel called on non-ARM64")
}

func evalNEON(a0, a1, w0, w1 *int16, sum *int32) { evalScalar(a0, a1, w0, w1, sum) }
