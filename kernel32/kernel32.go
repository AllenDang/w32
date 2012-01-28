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

    procGetModuleHandle    uintptr
    procMulDiv             uintptr
    procGetCurrentThread   uintptr
    procGetUserDefaultLCID uintptr
    procLstrlen            uintptr
    procLstrcpy            uintptr
    procGlobalAlloc        uintptr
    procGlobalFree         uintptr
    procGlobalLock         uintptr
    procGlobalUnlock       uintptr
    procMoveMemory         uintptr
    procFindResource       uintptr
    procSizeofResource     uintptr
    procLockResource       uintptr
    procLoadResource       uintptr
    procCopyMemory         uintptr
)

func init() {
    lib = LoadLib("kernel32.dll")

    procGetModuleHandle = GetProcAddr(lib, "GetModuleHandleW")
    procMulDiv = GetProcAddr(lib, "MulDiv")
    procGetCurrentThread = GetProcAddr(lib, "GetCurrentThread")
    procGetUserDefaultLCID = GetProcAddr(lib, "GetUserDefaultLCID")
    procLstrlen = GetProcAddr(lib, "lstrlenW")
    procLstrcpy = GetProcAddr(lib, "lstrcpyW")
    procGlobalAlloc = GetProcAddr(lib, "GlobalAlloc")
    procGlobalFree = GetProcAddr(lib, "GlobalFree")
    procGlobalLock = GetProcAddr(lib, "GlobalLock")
    procGlobalUnlock = GetProcAddr(lib, "GlobalUnlock")
    procMoveMemory = GetProcAddr(lib, "RtlMoveMemory")
    procFindResource = GetProcAddr(lib, "FindResourceW")
    procSizeofResource = GetProcAddr(lib, "SizeofResource")
    procLockResource = GetProcAddr(lib, "LockResource")
    procLoadResource = GetProcAddr(lib, "LoadResource")
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

func GetUserDefaultLCID() uint32 {
    ret, _, _ := syscall.Syscall(procGetUserDefaultLCID, 0,
        0,
        0,
        0)

    return uint32(ret)
}

func Lstrlen(lpString *uint16) int {
    ret, _, _ := syscall.Syscall(procLstrlen, 1,
        uintptr(unsafe.Pointer(lpString)),
        0,
        0)

    return int(ret)
}

func Lstrcpy(buf []uint16, lpString *uint16) {
    syscall.Syscall(procLstrcpy, 2,
        uintptr(unsafe.Pointer(&buf[0])),
        uintptr(unsafe.Pointer(lpString)),
        0)
}

func GlobalAlloc(uFlags uint, dwBytes uint32) HGLOBAL {
    ret, _, _ := syscall.Syscall(procGlobalAlloc, 2,
        uintptr(uFlags),
        uintptr(dwBytes),
        0)

    if ret == 0 {
        panic("GlobalAlloc failed")
    }

    return HGLOBAL(ret)
}

func GlobalFree(hMem HGLOBAL) {
    ret, _, _ := syscall.Syscall(procGlobalFree, 1,
        uintptr(hMem),
        0,
        0)
    
    if ret != 0 {
        panic("GlobalFree failed")
    }
}

func GlobalLock(hMem HGLOBAL) unsafe.Pointer {
    ret, _, _ := syscall.Syscall(procGlobalLock, 1,
        uintptr(hMem),
        0,
        0)
    
    if ret == 0 {
        panic("GlobalLock failed")
    }

    return unsafe.Pointer(ret)
}

func GlobalUnlock(hMem HGLOBAL) bool {
    ret, _, _ := syscall.Syscall(procGlobalUnlock, 1,
        uintptr(hMem),
        0,
        0)

    return ret != 0
}

func MoveMemory(destination, source unsafe.Pointer, length uint32) {
    syscall.Syscall(procMoveMemory, 3,
        uintptr(unsafe.Pointer(destination)),
        uintptr(source),
        uintptr(length))
}

func FindResource(hModule HMODULE, lpName, lpType *uint16) (HRSRC, error) {
    ret, _, _ := syscall.Syscall(procFindResource, 3,
        uintptr(hModule),
        uintptr(unsafe.Pointer(lpName)),
        uintptr(unsafe.Pointer(lpType)))

    if ret == 0 {
        return 0, syscall.GetLastError()
    }

    return HRSRC(ret), nil
}

func SizeofResource(hModule HMODULE, hResInfo HRSRC) uint32 {
    ret, _, _ := syscall.Syscall(procSizeofResource, 2,
        uintptr(hModule),
        uintptr(hResInfo),
        0)
    
    if ret == 0 {
        panic("SizeofResource failed")
    }

    return uint32(ret)
}

func LockResource(hResData HGLOBAL) unsafe.Pointer {
    ret, _, _ := syscall.Syscall(procLockResource, 1,
        uintptr(hResData),
        0,
        0)
    
    if ret == 0 {
        panic("LockResource failed")
    }

    return unsafe.Pointer(ret)
}

func LoadResource(hModule HMODULE, hResInfo HRSRC) HGLOBAL {
    ret, _, _ := syscall.Syscall(procLoadResource, 2,
        uintptr(hModule),
        uintptr(hResInfo),
        0)

    if ret == 0 {
        panic("LoadResource failed")
    }

    return HGLOBAL(ret)
}