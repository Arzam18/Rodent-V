//go:build arm64

package main

// ARM64 NEON kernels. Advanced SIMD is mandatory in AArch64.
//
//go:noescape
func addSingleNEON(a, w *int16)

//go:noescape
func subSingleNEON(a, w *int16)

//go:noescape
func moveNEON(a0, a1, wFrom0, wTo0, wFrom1, wTo1 *int16)

//go:noescape
func captureNEON(a0, a1, wTo0, wFrom0, wCap0, wTo1, wFrom1, wCap1 *int16)

//go:noescape
func castleNEON(a0, a1, wKFrom0, wKTo0, wRFrom0, wRTo0, wKFrom1, wKTo1, wRFrom1, wRTo1 *int16)

//go:noescape
func moveNEON3(dst0, src0, dst1, src1, wFrom0, wTo0, wFrom1, wTo1 *int16)

//go:noescape
func captureNEON3(dst0, src0, dst1, src1, wTo0, wFrom0, wCap0, wTo1, wFrom1, wCap1 *int16)

//go:noescape
func castleNEON3(dst0, src0, dst1, src1, wKFrom0, wKTo0, wRFrom0, wRTo0, wKFrom1, wKTo1, wRFrom1, wRTo1 *int16)

// dotprodL1 computes signed int8 dot products for 128 groups of four inputs.
// activations contains (u8-128), repeated four times in each 16-byte vector.
// The caller adds 128*sum(weights) to restore exact unsigned-input arithmetic.
//
//go:noescape
func dotprodL1(sums *int32, activations, tiled *int8)

func (acc *Accumulator) getEvalMultilayerARM64(bucket, stm int) int {
	var pw [NNUEHiddenSize]uint8
	stmAcc := &acc.values[stm]
	ntmAcc := &acc.values[stm^1]
	const halfSize = NNUEHiddenSize / 2
	for i := 0; i < halfSize; i++ {
		s0 := stmAcc[i]
		if s0 < 0 {
			s0 = 0
		} else if s0 > 255 {
			s0 = 255
		}
		s1 := stmAcc[i+halfSize]
		if s1 < 0 {
			s1 = 0
		} else if s1 > 255 {
			s1 = 255
		}
		pw[i] = uint8((int32(s0) * int32(s1)) >> 8)

		n0 := ntmAcc[i]
		if n0 < 0 {
			n0 = 0
		} else if n0 > 255 {
			n0 = 255
		}
		n1 := ntmAcc[i+halfSize]
		if n1 < 0 {
			n1 = 0
		} else if n1 > 255 {
			n1 = 255
		}
		pw[halfSize+i] = uint8((int32(n0) * int32(n1)) >> 8)
	}

	var intSums [NNUEMultilayerL1Size]int32
	if hasDotProd {
		var dotInputs [NNUEHiddenSize / 4][16]int8
		for k := 0; k < NNUEHiddenSize/4; k++ {
			for repeat := 0; repeat < 4; repeat++ {
				for lane := 0; lane < 4; lane++ {
					dotInputs[k][repeat*4+lane] = int8(int(pw[k*4+lane]) - 128)
				}
			}
		}
		dotprodL1(&intSums[0], &dotInputs[0][0], &nnueParams.L1WeightsTiled[bucket][0][0][0])
		for j := 0; j < NNUEMultilayerL1Size; j++ {
			intSums[j] += 128 * nnueParams.L1WeightSums[bucket][j]
		}
	} else {
		l1wRaw := &nnueParams.L1WeightsRaw[bucket]
		for i := 0; i < NNUEHiddenSize; i++ {
			pwi := int32(pw[i])
			if pwi == 0 {
				continue
			}
			for j := 0; j < NNUEMultilayerL1Size; j++ {
				intSums[j] += pwi * int32(l1wRaw[i][j])
			}
		}
	}

	const l1Factor = float32(4.0 / 65025.0)
	var l1Out [NNUEMultilayerL2Inputs]float32
	for j := 0; j < NNUEMultilayerL1Size; j++ {
		sum := nnueParams.L1Biases[bucket][j] + float32(intSums[j])*l1Factor
		if sum < 0 {
			sum = 0
		} else if sum > 1 {
			sum = 1
		}
		screlu := sum * sum
		if nnueParams.IsDualActivation {
			l1Out[j] = sum
			l1Out[j+NNUEMultilayerL1Size] = screlu
		} else {
			l1Out[j] = screlu
		}
	}

	l2Inputs := NNUEMultilayerL1Size
	if nnueParams.IsDualActivation {
		l2Inputs *= 2
	}
	var l2Out [NNUEMultilayerL2Size]float32
	copy(l2Out[:], nnueParams.L2Biases[bucket][:])
	for j := 0; j < l2Inputs; j++ {
		act := l1Out[j]
		if act == 0 {
			continue
		}
		for k := 0; k < NNUEMultilayerL2Size; k++ {
			l2Out[k] += act * nnueParams.L2Weights[bucket][j][k]
		}
	}
	for k := 0; k < NNUEMultilayerL2Size; k++ {
		if l2Out[k] < 0 {
			l2Out[k] = 0
		} else if l2Out[k] > 1 {
			l2Out[k] = 1
		}
	}

	score := nnueParams.L3Biases[bucket]
	for k := 0; k < NNUEMultilayerL2Size; k++ {
		score += l2Out[k] * nnueParams.L3Weights[bucket][k]
	}
	return int(score * float32(singleOptionValue[NnueScale]))
}

//go:noescape
func evalNEON(a0, a1, w0, w1 *int16, sum *int32)
