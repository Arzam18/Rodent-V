//go:build arm64

package main

import (
	"math/rand"
	"testing"
)

func TestNEONAccumulatorKernelsMatchScalar(t *testing.T) {
	rng := rand.New(rand.NewSource(0x524f44454e54))
	fill := func(dst *[NNUEHiddenSize]int16) {
		for i := range dst {
			dst[i] = int16(rng.Intn(65536) - 32768)
		}
	}
	var a0, a1, w0, w1, w2, w3, w4, w5, w6, w7 [NNUEHiddenSize]int16
	for _, v := range []*[NNUEHiddenSize]int16{&a0, &a1, &w0, &w1, &w2, &w3, &w4, &w5, &w6, &w7} {
		fill(v)
	}

	want, got := a0, a0
	addSingleScalar(&want[0], &w0[0])
	addSingleNEON(&got[0], &w0[0])
	if want != got {
		t.Fatal("NEON addSingle differs from scalar")
	}
	want, got = a0, a0
	subSingleScalar(&want[0], &w0[0])
	subSingleNEON(&got[0], &w0[0])
	if want != got {
		t.Fatal("NEON subSingle differs from scalar")
	}

	want0, want1, got0, got1 := a0, a1, a0, a1
	moveScalar(&want0[0], &want1[0], &w0[0], &w1[0], &w2[0], &w3[0])
	moveNEON(&got0[0], &got1[0], &w0[0], &w1[0], &w2[0], &w3[0])
	if want0 != got0 || want1 != got1 {
		t.Fatal("NEON move differs from scalar")
	}

	want0, want1, got0, got1 = a0, a1, a0, a1
	captureScalar(&want0[0], &want1[0], &w0[0], &w1[0], &w2[0], &w3[0], &w4[0], &w5[0])
	captureNEON(&got0[0], &got1[0], &w0[0], &w1[0], &w2[0], &w3[0], &w4[0], &w5[0])
	if want0 != got0 || want1 != got1 {
		t.Fatal("NEON capture differs from scalar")
	}

	want0, want1, got0, got1 = a0, a1, a0, a1
	castleScalar(&want0[0], &want1[0], &w0[0], &w1[0], &w2[0], &w3[0], &w4[0], &w5[0], &w6[0], &w7[0])
	castleNEON(&got0[0], &got1[0], &w0[0], &w1[0], &w2[0], &w3[0], &w4[0], &w5[0], &w6[0], &w7[0])
	if want0 != got0 || want1 != got1 {
		t.Fatal("NEON castle differs from scalar")
	}

	var d0, d1, e0, e1 [NNUEHiddenSize]int16
	moveScalar3(&d0[0], &a0[0], &d1[0], &a1[0], &w0[0], &w1[0], &w2[0], &w3[0])
	moveNEON3(&e0[0], &a0[0], &e1[0], &a1[0], &w0[0], &w1[0], &w2[0], &w3[0])
	if d0 != e0 || d1 != e1 {
		t.Fatal("NEON move3 differs from scalar")
	}
	captureScalar3(&d0[0], &a0[0], &d1[0], &a1[0], &w0[0], &w1[0], &w2[0], &w3[0], &w4[0], &w5[0])
	captureNEON3(&e0[0], &a0[0], &e1[0], &a1[0], &w0[0], &w1[0], &w2[0], &w3[0], &w4[0], &w5[0])
	if d0 != e0 || d1 != e1 {
		t.Fatal("NEON capture3 differs from scalar")
	}
	castleScalar3(&d0[0], &a0[0], &d1[0], &a1[0], &w0[0], &w1[0], &w2[0], &w3[0], &w4[0], &w5[0], &w6[0], &w7[0])
	castleNEON3(&e0[0], &a0[0], &e1[0], &a1[0], &w0[0], &w1[0], &w2[0], &w3[0], &w4[0], &w5[0], &w6[0], &w7[0])
	if d0 != e0 || d1 != e1 {
		t.Fatal("NEON castle3 differs from scalar")
	}
}

func TestDotProdL1KernelMatchesScalar(t *testing.T) {
	if !hasDotProd {
		t.Skip("CPU does not advertise ASIMDDP")
	}
	rng := rand.New(rand.NewSource(0x53444f54))
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
		var got [NNUEMultilayerL1Size]int32
		dotprodL1(&got[0], &signedInputs[0][0], &tiled[0][0][0])
		var want [NNUEMultilayerL1Size]int32
		for i, a := range activations {
			for j := range want {
				want[j] += int32(a) * int32(raw[i][j])
			}
		}
		for j := range got {
			got[j] += 128 * weightSums[j]
			if got[j] != want[j] {
				t.Fatalf("trial %d output %d: SDOT got %d, scalar %d", trial, j, got[j], want[j])
			}
		}
	}
}

func TestNEONEvalMatchesScalar(t *testing.T) {
	rng := rand.New(rand.NewSource(0x4e454f4e))
	for trial := 0; trial < 100; trial++ {
		var a0, a1, w0, w1 [NNUEHiddenSize]int16
		for i := 0; i < NNUEHiddenSize; i++ {
			a0[i] = int16(rng.Intn(65536) - 32768)
			a1[i] = int16(rng.Intn(65536) - 32768)
			w0[i] = int16(rng.Intn(65536) - 32768)
			w1[i] = int16(rng.Intn(65536) - 32768)
		}
		var want, got int32
		evalScalar(&a0[0], &a1[0], &w0[0], &w1[0], &want)
		evalNEON(&a0[0], &a1[0], &w0[0], &w1[0], &got)
		if got != want {
			t.Fatalf("trial %d: NEON eval got %d, scalar %d", trial, got, want)
		}
	}
}
