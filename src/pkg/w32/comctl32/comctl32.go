// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comctl32

import (
    "syscall"
    "unsafe"
    . "w32"
)

var (
    lib uintptr

    procInitCommonControlsEx    uintptr
    procImageList_Create        uintptr
    procImageList_Destroy       uintptr
    procImageList_GetImageCount uintptr
    procImageList_SetImageCount uintptr
    procImageList_Add           uintptr
    procImageList_ReplaceIcon   uintptr
    procImageList_Remove        uintptr
)

func init() {
    lib = LoadLib("comctl32.dll")

    procInitCommonControlsEx = GetProcAddr(lib, "InitCommonControlsEx")
    procImageList_Create = GetProcAddr(lib, "ImageList_Create")
    procImageList_Destroy = GetProcAddr(lib, "ImageList_Destroy")
    procImageList_GetImageCount = GetProcAddr(lib, "ImageList_GetImageCount")
    procImageList_SetImageCount = GetProcAddr(lib, "ImageList_SetImageCount")
    procImageList_Add = GetProcAddr(lib, "ImageList_Add")
    procImageList_ReplaceIcon = GetProcAddr(lib, "ImageList_ReplaceIcon")
    procImageList_Remove = GetProcAddr(lib, "ImageList_Remove")
}

func InitCommonControlsEx(lpInitCtrls *INITCOMMONCONTROLSEX) bool {
    ret, _, _ := syscall.Syscall(procInitCommonControlsEx, 1,
        uintptr(unsafe.Pointer(lpInitCtrls)),
        0,
        0)

    return ret != 0
}

func ImageList_Create(cx, cy int, flags uint, cInitial, cGrow int) HIMAGELIST {
    ret, _, _ := syscall.Syscall6(procImageList_Create, 5,
        uintptr(cx),
        uintptr(cy),
        uintptr(flags),
        uintptr(cInitial),
        uintptr(cGrow),
        0)

    if ret == 0 {
        panic("Create image list failed")
    }

    return HIMAGELIST(ret)
}

func ImageList_Destroy(himl HIMAGELIST) bool {
    ret, _, _ := syscall.Syscall(procImageList_Destroy, 1,
        uintptr(himl),
        0,
        0)

    return ret != 0
}

func ImageList_GetImageCount(himl HIMAGELIST) int {
    ret, _, _ := syscall.Syscall(procImageList_GetImageCount, 1,
        uintptr(himl),
        0,
        0)

    return int(ret)
}

func ImageList_SetImageCount(himl HIMAGELIST, uNewCount uint) bool {
    ret, _, _ := syscall.Syscall(procImageList_SetImageCount, 2,
        uintptr(himl),
        uintptr(uNewCount),
        0)

    return ret != 0
}

func ImageList_Add(himl HIMAGELIST, hbmImage, hbmMask HBITMAP) int {
    ret, _, _ := syscall.Syscall(procImageList_Add, 3,
        uintptr(himl),
        uintptr(hbmImage),
        uintptr(hbmMask))

    return int(ret)
}

func ImageList_ReplaceIcon(himl HIMAGELIST, i int, hicon HICON) int {
    ret, _, _ := syscall.Syscall(procImageList_ReplaceIcon, 3,
        uintptr(himl),
        uintptr(i),
        uintptr(hicon))

    return int(ret)
}

func ImageList_AddIcon(himl HIMAGELIST, hicon HICON) int {
    return ImageList_ReplaceIcon(himl, -1, hicon)
}

func ImageList_Remove(himl HIMAGELIST, i int) bool {
    ret, _, _ := syscall.Syscall(procImageList_Remove, 2,
        uintptr(himl),
        uintptr(i),
        0)

    return ret != 0
}

func ImageList_RemoveAll(himl HIMAGELIST) bool {
    return ImageList_Remove(himl, -1)
}
