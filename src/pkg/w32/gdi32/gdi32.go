// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdi32

import (
    "syscall"
    "unsafe"
    . "w32"
)

var (
    lib uintptr

    procGetDeviceCaps     uintptr
    procDeleteObject      uintptr
    procCreateFontIndirect uintptr
)

func init() {
    lib = LoadLib("gdi32.dll")

    procGetDeviceCaps = GetProcAddr(lib, "GetDeviceCaps")
    procDeleteObject = GetProcAddr(lib, "DeleteObject")
    procCreateFontIndirect = GetProcAddr(lib, "CreateFontIndirectW")
}

func GetDeviceCaps(hdc HDC, index int) int {
    ret, _, _ := syscall.Syscall(procGetDeviceCaps, 2,
        uintptr(hdc),
        uintptr(index),
        0)

    return int(ret)
}

func DeleteObject(hObject HGDIOBJ) bool {
    ret, _, _ := syscall.Syscall(procDeleteObject, 1,
        uintptr(hObject),
        0,
        0)

    return ret != 0
}

func CreateFontIndirect(logFont *LOGFONT) HFONT {
    ret, _, _ := syscall.Syscall(procCreateFontIndirect, 1,
        uintptr(unsafe.Pointer(logFont)),
        0,
        0)

    return HFONT(ret)
}
