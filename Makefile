EXE = rodent_v

ifeq ($(OS),Windows_NT)
	EXE := $(EXE).exe
endif

.PHONY: all build clean AVX2 x86-64-bmi2 x86-64-avx2 x86-64-modern

export GOAMD64 = v3

all: build

# OpenBench builds using cpuflags or ISA as targets (e.g. 'make AVX2')
AVX2: build
x86-64-bmi2: build
x86-64-avx2: build
x86-64-modern: build

# Default build (uses host OS)
build:
	go build -ldflags="-s -w" -o $(EXE)

clean:
	go clean
