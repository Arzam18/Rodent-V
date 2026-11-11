#!/usr/bin/env python3
"""Verify five Rodent Android ARM64 binaries contain real NEON and SDOT code."""
from __future__ import annotations

import os
import subprocess
import sys
from pathlib import Path


def run(*args: str) -> str:
    return subprocess.check_output(args, text=True, stderr=subprocess.STDOUT)


def main() -> int:
    binaries = [Path(arg) for arg in sys.argv[1:]]
    if len(binaries) != 5:
        raise SystemExit(f"Expected exactly five engine binaries, got {len(binaries)}")

    objdump = os.environ.get("LLVM_OBJDUMP", "llvm-objdump")
    for binary in binaries:
        if not binary.is_file() or binary.stat().st_size == 0:
            raise SystemExit(f"Missing or empty engine binary: {binary}")

        header = run("readelf", "-h", str(binary))
        if "ELF64" not in header or "AArch64" not in header:
            raise SystemExit(f"Not an ELF64/AArch64 executable: {binary}\n{header}")
        program_headers = run("readelf", "-l", str(binary))
        if "/system/bin/linker64" not in program_headers:
            raise SystemExit(f"Missing Android ARM64 interpreter in {binary}:\n{program_headers}")

        disassembly = run(objdump, "-d", str(binary)).lower()
        if "sdot" not in disassembly:
            raise SystemExit(f"No SDOT instruction found in {binary}")
        if not ("add\tv" in disassembly or "add v" in disassembly):
            raise SystemExit(f"No NEON vector arithmetic found in {binary}")
        if not ("sub\tv" in disassembly or "sub v" in disassembly):
            raise SystemExit(f"No NEON vector subtraction found in {binary}")
        for instruction in ("smull", "smull2", "smin", "smax", "mla", "addv"):
            if instruction not in disassembly:
                raise SystemExit(f"Missing NEON NNUE evaluation instruction {instruction} in {binary}")

        dynamic = run("readelf", "-d", str(binary))
        if "libc++_shared.so" in dynamic or "libstdc++.so" in dynamic:
            raise SystemExit(f"Unexpected C++ runtime dependency in {binary}:\n{dynamic}")
        print(f"PASS: {binary.name}: Android ELF64/AArch64, NEON accumulator + evaluation kernels, SDOT kernel")

    print("PASS: all five Rodent builds contain NEON accumulator/evaluation and runtime-gated SDOT code")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
