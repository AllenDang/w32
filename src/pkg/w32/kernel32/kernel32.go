// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kernel32

import (
    "syscall"
    "unsafe"
    . "w32"
)

var (
    lib uintptr

    procGetModuleHandle  uintptr
    procMulDiv           uintptr
    procGetCurrentThread uintptr
)

func init() {
    lib = LoadLib("kernel32.dll")

    procGetModuleHandle = GetProcAddr(lib, "GetModuleHandleW")
    procMulDiv = GetProcAddr(lib, "MulDiv")
    procGetCurrentThread = GetProcAddr(lib, "GetCurrentThread")
}

func GetModuleHandle(modulename string) HINSTANCE {
    var mn uintptr
    if modulename == "" {
        mn = 0
    } else {
        mn = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(modulename)))
    }
    ret, _, _ := syscall.Syscall(procGetModuleHandle, 1,
        mn,
        0,
        0)
    return HINSTANCE(ret)
}

func MulDiv(number, numerator, denominator int) int {
    ret, _, _ := syscall.Syscall(procMulDiv, 3,
        uintptr(number),
        uintptr(numerator),
        uintptr(denominator))

    return int(ret)
}

func GetCurrentThread() HANDLE {
    ret, _, _ := syscall.Syscall(procGetCurrentThread, 0,
        0,
        0,
        0)

    return HANDLE(ret)
}
