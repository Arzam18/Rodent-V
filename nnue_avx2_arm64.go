//go:build arm64

package main

// These placeholders satisfy references to x86 AVX2 kernels on ARM64.
// hasAVX2 is false on ARM64, so these functions must never be called.
// ARM64 currently uses the portable scalar NNUE implementation.

func addSingleAVX2_64(a, w *int16) {
	panic("AVX2 kernel called on ARM64")
}

func subSingleAVX2_64(a, w *int16) {
	panic("AVX2 kernel called on ARM64")
}

func addSingleAVX2_128(a, w *int16) {
	panic("AVX2 kernel called on ARM64")
}

func subSingleAVX2_128(a, w *int16) {
	panic("AVX2 kernel called on ARM64")
}

func addSingleAVX2_256(a, w *int16) {
	panic("AVX2 kernel called on ARM64")
}

func subSingleAVX2_256(a, w *int16) {
	panic("AVX2 kernel called on ARM64")
}

func addSingleAVX2_384(a, w *int16) {
	panic("AVX2 kernel called on ARM64")
}

func subSingleAVX2_384(a, w *int16) {
	panic("AVX2 kernel called on ARM64")
}

func addSingleAVX2_512(a, w *int16) {
	panic("AVX2 kernel called on ARM64")
}

func subSingleAVX2_512(a, w *int16) {
	panic("AVX2 kernel called on ARM64")
}

func addSingleAVX2_768(a, w *int16) {
	panic("AVX2 kernel called on ARM64")
}

func subSingleAVX2_768(a, w *int16) {
	panic("AVX2 kernel called on ARM64")
}

func addSingleAVX2_1024(a, w *int16) {
	panic("AVX2 kernel called on ARM64")
}

func subSingleAVX2_1024(a, w *int16) {
	panic("AVX2 kernel called on ARM64")
}

func captureAVX2_64(
	a0, a1 *int16,
	wTo0, wFrom0, wCap0 *int16,
	wTo1, wFrom1, wCap1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func captureAVX2_128(
	a0, a1 *int16,
	wTo0, wFrom0, wCap0 *int16,
	wTo1, wFrom1, wCap1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func captureAVX2_256(
	a0, a1 *int16,
	wTo0, wFrom0, wCap0 *int16,
	wTo1, wFrom1, wCap1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func captureAVX2_384(
	a0, a1 *int16,
	wTo0, wFrom0, wCap0 *int16,
	wTo1, wFrom1, wCap1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func captureAVX2_512(
	a0, a1 *int16,
	wTo0, wFrom0, wCap0 *int16,
	wTo1, wFrom1, wCap1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func captureAVX2_768(
	a0, a1 *int16,
	wTo0, wFrom0, wCap0 *int16,
	wTo1, wFrom1, wCap1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func captureAVX2_1024(
	a0, a1 *int16,
	wTo0, wFrom0, wCap0 *int16,
	wTo1, wFrom1, wCap1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func moveAVX2_64(
	a0, a1 *int16,
	wFrom0, wTo0 *int16,
	wFrom1, wTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func moveAVX2_128(
	a0, a1 *int16,
	wFrom0, wTo0 *int16,
	wFrom1, wTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func moveAVX2_256(
	a0, a1 *int16,
	wFrom0, wTo0 *int16,
	wFrom1, wTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func moveAVX2_384(
	a0, a1 *int16,
	wFrom0, wTo0 *int16,
	wFrom1, wTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func moveAVX2_512(
	a0, a1 *int16,
	wFrom0, wTo0 *int16,
	wFrom1, wTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func moveAVX2_768(
	a0, a1 *int16,
	wFrom0, wTo0 *int16,
	wFrom1, wTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func moveAVX2_1024(
	a0, a1 *int16,
	wFrom0, wTo0 *int16,
	wFrom1, wTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func castleAVX2_64(
	a0, a1 *int16,
	wKFrom0, wKTo0, wRFrom0, wRTo0 *int16,
	wKFrom1, wKTo1, wRFrom1, wRTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func castleAVX2_128(
	a0, a1 *int16,
	wKFrom0, wKTo0, wRFrom0, wRTo0 *int16,
	wKFrom1, wKTo1, wRFrom1, wRTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func castleAVX2_256(
	a0, a1 *int16,
	wKFrom0, wKTo0, wRFrom0, wRTo0 *int16,
	wKFrom1, wKTo1, wRFrom1, wRTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func castleAVX2_384(
	a0, a1 *int16,
	wKFrom0, wKTo0, wRFrom0, wRTo0 *int16,
	wKFrom1, wKTo1, wRFrom1, wRTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func castleAVX2_512(
	a0, a1 *int16,
	wKFrom0, wKTo0, wRFrom0, wRTo0 *int16,
	wKFrom1, wKTo1, wRFrom1, wRTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func castleAVX2_768(
	a0, a1 *int16,
	wKFrom0, wKTo0, wRFrom0, wRTo0 *int16,
	wKFrom1, wKTo1, wRFrom1, wRTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func castleAVX2_1024(
	a0, a1 *int16,
	wKFrom0, wKTo0, wRFrom0, wRTo0 *int16,
	wKFrom1, wKTo1, wRFrom1, wRTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func moveAVX2_512_3op(
	dst0, src0, dst1, src1 *int16,
	wFrom0, wTo0, wFrom1, wTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func captureAVX2_512_3op(
	dst0, src0, dst1, src1 *int16,
	wTo0, wFrom0, wCap0 *int16,
	wTo1, wFrom1, wCap1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func castleAVX2_512_3op(
	dst0, src0, dst1, src1 *int16,
	wKFrom0, wKTo0, wRFrom0, wRTo0 *int16,
	wKFrom1, wKTo1, wRFrom1, wRTo1 *int16,
) {
	panic("AVX2 kernel called on ARM64")
}

func getEvalAVX2_64(
	a0, a1 *int16,
	w0, w1 *int16,
	sum *int32,
) {
	panic("AVX2 kernel called on ARM64")
}

func getEvalAVX2_128(
	a0, a1 *int16,
	w0, w1 *int16,
	sum *int32,
) {
	panic("AVX2 kernel called on ARM64")
}

func getEvalAVX2_256(
	a0, a1 *int16,
	w0, w1 *int16,
	sum *int32,
) {
	panic("AVX2 kernel called on ARM64")
}

func getEvalAVX2_384(
	a0, a1 *int16,
	w0, w1 *int16,
	sum *int32,
) {
	panic("AVX2 kernel called on ARM64")
}

func getEvalAVX2_512(
	a0, a1 *int16,
	w0, w1 *int16,
	sum *int32,
) {
	panic("AVX2 kernel called on ARM64")
}

func getEvalAVX2_768(
	a0, a1 *int16,
	w0, w1 *int16,
	sum *int32,
) {
	panic("AVX2 kernel called on ARM64")
}

func getEvalAVX2_1024(
	a0, a1 *int16,
	w0, w1 *int16,
	sum *int32,
) {
	panic("AVX2 kernel called on ARM64")
}

func getEvalMultilayerAVX2(
	a0, a1 *int16,
	l1w *int8,
	l1b *float32,
	l2w *float32,
	l2b *float32,
	l3w *float32,
	l3b float32,
	scale float32,
	sum *int32,
) {
	panic("AVX2 kernel called on ARM64")
}
