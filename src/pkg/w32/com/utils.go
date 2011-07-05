package com

import (
    "syscall"
    "unsafe"
    . "w32"
    . "w32/kernel32"
)

func ComAddRef(unknown *IUnknown) int32 {
    ret, _, _ := syscall.Syscall(unknown.lpVtbl.pAddRef, 1,
        uintptr(unsafe.Pointer(unknown)),
        0,
        0)
    return int32(ret)
}

func ComRelease(unknown *IUnknown) int32 {
    ret, _, _ := syscall.Syscall(unknown.lpVtbl.pRelease, 1,
        uintptr(unsafe.Pointer(unknown)),
        0,
        0)
    return int32(ret)
}

func ComQueryInterface(unknown *IUnknown, id *GUID) *IDispatch {
    var disp *IDispatch
    hr, _, _ := syscall.Syscall(unknown.lpVtbl.pQueryInterface, 3,
        uintptr(unsafe.Pointer(unknown)),
        uintptr(unsafe.Pointer(id)),
        uintptr(unsafe.Pointer(&disp)))
    if hr != 0 {
        panic("Invoke QieryInterface error.")
    }
    return disp
}

func ComGetIDsOfName(disp *IDispatch, names []string) []int32 {
    wnames := make([]*uint16, len(names))
    dispid := make([]int32, len(names))
    for i := 0; i < len(names); i++ {
        wnames[i] = syscall.StringToUTF16Ptr(names[i])
    }
    hr, _, _ := syscall.Syscall6(disp.lpVtbl.pGetIDsOfNames, 6,
        uintptr(unsafe.Pointer(disp)),
        uintptr(unsafe.Pointer(IID_NULL)),
        uintptr(unsafe.Pointer(&wnames[0])),
        uintptr(len(names)),
        uintptr(GetUserDefaultLCID()),
        uintptr(unsafe.Pointer(&dispid[0])))
    if hr != 0 {
        panic("Invoke GetIDsOfName error.")
    }
    return dispid
}
