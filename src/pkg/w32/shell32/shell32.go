// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shell32

import (
    "syscall"
    "unsafe"
    . "w32"
)

var (
    lib uintptr

    procSHBrowseForFolder   uintptr
    procSHGetPathFromIDList uintptr
)

func init() {
    lib = LoadLib("shell32.dll")

    procSHBrowseForFolder = GetProcAddr(lib, "SHBrowseForFolderW")
    procSHGetPathFromIDList = GetProcAddr(lib, "SHGetPathFromIDListW")
}

func SHBrowseForFolder(bi *BROWSEINFO) uintptr {
    ret, _, _ := syscall.Syscall(procSHBrowseForFolder, 1,
        uintptr(unsafe.Pointer(bi)),
        0,
        0)

    return ret
}

func SHGetPathFromIDList(idl uintptr) string {
    buf := make([]uint16, 1024)
    syscall.Syscall(procSHGetPathFromIDList, 2,
        idl,
        uintptr(unsafe.Pointer(&buf[0])),
        0)

    return syscall.UTF16ToString(buf)
}
