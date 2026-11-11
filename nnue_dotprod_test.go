package main

import (
	"math/rand"
	"testing"
)

// The arithmetic identity used by the ARM SDOT kernel is exact:
// sum(u8*w) = sum((u8-128)*w) + 128*sum(w).
func TestDotProdPackedU8I8MatchesScalar(t *testing.T) {
	rng := rand.New(rand.NewSource(0x524f4445))
	for trial := 0; trial < 100; trial++ {
		var activations [NNUEHiddenSize]uint8
		var raw [NNUEHiddenSize][NNUEMultilayerL1Size]int8
		var tiled [NNUEHiddenSize / 4][2][32]int8
		var weightSums [NNUEMultilayerL1Size]int32
		for i := range activations {
			activations[i] = uint8(rng.Intn(256))
			for j := 0; j < NNUEMultilayerL1Size; j++ {
				w := int8(rng.Intn(256) - 128)
				raw[i][j] = w
				weightSums[j] += int32(w)
			}
		}
		for k := 0; k < NNUEHiddenSize/4; k++ {
			for m := 0; m < 8; m++ {
				for c := 0; c < 4; c++ {
					tiled[k][0][m*4+c] = raw[4*k+c][m]
					tiled[k][1][m*4+c] = raw[4*k+c][8+m]
				}
			}
		}
		var signedInputs [NNUEHiddenSize / 4][16]int8
		for k := 0; k < NNUEHiddenSize/4; k++ {
			for repeat := 0; repeat < 4; repeat++ {
				for lane := 0; lane < 4; lane++ {
					signedInputs[k][repeat*4+lane] = int8(int(activations[k*4+lane]) - 128)
				}
			}
		}

		var scalar [NNUEMultilayerL1Size]int32
		for i, activation := range activations {
			for j := range scalar {
				scalar[j] += int32(activation) * int32(raw[i][j])
			}
		}
		var signed [NNUEMultilayerL1Size]int32
		for k := 0; k < NNUEHiddenSize/4; k++ {
			for group := 0; group < 4; group++ {
				block, offset := group/2, (group%2)*16
				for neuron := 0; neuron < 4; neuron++ {
					j := group*4 + neuron
					for lane := 0; lane < 4; lane++ {
						w := tiled[k][block][offset+neuron*4+lane]
						signed[j] += int32(signedInputs[k][neuron*4+lane]) * int32(w)
					}
				}
			}
		}
		for j := range scalar {
			if got := signed[j] + 128*weightSums[j]; got != scalar[j] {
				t.Fatalf("trial %d output %d: SDOT identity got %d, scalar %d", trial, j, got, scalar[j])
			}
		}
	}
}
