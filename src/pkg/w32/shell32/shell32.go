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
    procDragAcceptFiles     uintptr
    procDragQueryFile       uintptr
    procDragQueryPoint      uintptr
    procDragFinish          uintptr
)

func init() {
    lib = LoadLib("shell32.dll")

    procSHBrowseForFolder = GetProcAddr(lib, "SHBrowseForFolderW")
    procSHGetPathFromIDList = GetProcAddr(lib, "SHGetPathFromIDListW")
    procDragAcceptFiles = GetProcAddr(lib, "DragAcceptFiles")
    procDragQueryFile = GetProcAddr(lib, "DragQueryFileW")
    procDragQueryPoint = GetProcAddr(lib, "DragQueryPoint")
    procDragFinish = GetProcAddr(lib, "DragFinish")
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

func DragAcceptFiles(hwnd HWND, accept bool) {
    syscall.Syscall(procDragAcceptFiles, 2,
        uintptr(hwnd),
        uintptr(BoolToBOOL(accept)),
        0)
}

func DragQueryFile(hDrop HDROP, iFile uint) (fileName string, fileCount uint) {
    ret, _, _ := syscall.Syscall6(procDragQueryFile, 4,
        uintptr(hDrop),
        uintptr(iFile),
        0,
        0,
        0,
        0)

    fileCount = uint(ret)

    if iFile != 0xFFFFFFFF {
        buf := make([]uint16, fileCount+1)

        ret, _, _ := syscall.Syscall6(procDragQueryFile, 4,
            uintptr(hDrop),
            uintptr(iFile),
            uintptr(unsafe.Pointer(&buf[0])),
            uintptr(fileCount+1),
            0,
            0)

        if ret == 0 {
            panic("Invoke DragQueryFile error.")
        }

        fileName = syscall.UTF16ToString(buf)
    }

    return
}

func DragQueryPoint(hDrop HDROP) (x, y int, isClientArea bool) {
    var pt POINT
    ret, _, _ := syscall.Syscall(procDragQueryPoint, 2,
        uintptr(hDrop),
        uintptr(unsafe.Pointer(&pt)),
        0)

    isClientArea = false
    if ret == 1 {
        isClientArea = true
    }

    x = pt.X
    y = pt.Y

    return
}

func DragFinish(hDrop HDROP) {
    syscall.Syscall(procDragFinish, 1,
        uintptr(hDrop),
        0,
        0)
}
