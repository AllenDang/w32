// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comdlg32

import (
    "syscall"
    "unsafe"
    . "github.com/AllenDang/w32"
)

var (
    lib uintptr

    procGetOpenFileName      uintptr
    procGetSaveFileName      uintptr
    procCommDlgExtendedError uintptr
)

func init() {
    lib = LoadLib("comdlg32.dll")

    procGetSaveFileName = GetProcAddr(lib, "GetSaveFileNameW")
    procGetOpenFileName = GetProcAddr(lib, "GetOpenFileNameW")
    procCommDlgExtendedError = GetProcAddr(lib, "CommDlgExtendedError")
}

func GetOpenFileName(ofn *OPENFILENAME) bool {
    ret, _, _ := syscall.Syscall(procGetOpenFileName, 1,
        uintptr(unsafe.Pointer(ofn)),
        0,
        0)

    return ret != 0
}

func GetSaveFileName(ofn *OPENFILENAME) bool {
    ret, _, _ := syscall.Syscall(procGetSaveFileName, 1,
        uintptr(unsafe.Pointer(ofn)),
        0,
        0)

    return ret != 0
}

func CommDlgExtendedError() uint {
    ret, _, _ := syscall.Syscall(procCommDlgExtendedError, 0,
        0,
        0,
        0)

    return uint(ret)
}
