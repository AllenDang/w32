// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kernel32

import (
    . "github.com/AllenDang/w32"
    "syscall"
    "unsafe"
)

var (
    modkernel32 = syscall.NewLazyDLL("kernel32.dll")

    procGetModuleHandle    = modkernel32.NewProc("GetModuleHandleW")
    procMulDiv             = modkernel32.NewProc("MulDiv")
    procGetCurrentThread   = modkernel32.NewProc("GetCurrentThread")
    procGetUserDefaultLCID = modkernel32.NewProc("GetUserDefaultLCID")
    procLstrlen            = modkernel32.NewProc("lstrlenW")
    procLstrcpy            = modkernel32.NewProc("lstrcpyW")
    procGlobalAlloc        = modkernel32.NewProc("GlobalAlloc")
    procGlobalFree         = modkernel32.NewProc("GlobalFree")
    procGlobalLock         = modkernel32.NewProc("GlobalLock")
    procGlobalUnlock       = modkernel32.NewProc("GlobalUnlock")
    procMoveMemory         = modkernel32.NewProc("RtlMoveMemory")
    procFindResource       = modkernel32.NewProc("FindResourceW")
    procSizeofResource     = modkernel32.NewProc("SizeofResource")
    procLockResource       = modkernel32.NewProc("LockResource")
    procLoadResource       = modkernel32.NewProc("LoadResource")
    procGetLastError       = modkernel32.NewProc("GetLastError")
)

func GetModuleHandle(modulename string) HINSTANCE {
    var mn uintptr
    if modulename == "" {
        mn = 0
    } else {
        mn = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(modulename)))
    }
    ret, _, _ := procGetModuleHandle.Call(mn)
    return HINSTANCE(ret)
}

func MulDiv(number, numerator, denominator int) int {
    ret, _, _ := procMulDiv.Call(
        uintptr(number),
        uintptr(numerator),
        uintptr(denominator))

    return int(ret)
}

func GetCurrentThread() HANDLE {
    ret, _, _ := procGetCurrentThread.Call()

    return HANDLE(ret)
}

func GetUserDefaultLCID() uint32 {
    ret, _, _ := procGetUserDefaultLCID.Call()

    return uint32(ret)
}

func Lstrlen(lpString *uint16) int {
    ret, _, _ := procLstrlen.Call(uintptr(unsafe.Pointer(lpString)))

    return int(ret)
}

func Lstrcpy(buf []uint16, lpString *uint16) {
    procLstrcpy.Call(
        uintptr(unsafe.Pointer(&buf[0])),
        uintptr(unsafe.Pointer(lpString)))
}

func GlobalAlloc(uFlags uint, dwBytes uint32) HGLOBAL {
    ret, _, _ := procGlobalAlloc.Call(
        uintptr(uFlags),
        uintptr(dwBytes))

    if ret == 0 {
        panic("GlobalAlloc failed")
    }

    return HGLOBAL(ret)
}

func GlobalFree(hMem HGLOBAL) {
    ret, _, _ := procGlobalFree.Call(uintptr(hMem))

    if ret != 0 {
        panic("GlobalFree failed")
    }
}

func GlobalLock(hMem HGLOBAL) unsafe.Pointer {
    ret, _, _ := procGlobalLock.Call(uintptr(hMem))

    if ret == 0 {
        panic("GlobalLock failed")
    }

    return unsafe.Pointer(ret)
}

func GlobalUnlock(hMem HGLOBAL) bool {
    ret, _, _ := procGlobalUnlock.Call(uintptr(hMem))

    return ret != 0
}

func MoveMemory(destination, source unsafe.Pointer, length uint32) {
    procMoveMemory.Call(
        uintptr(unsafe.Pointer(destination)),
        uintptr(source),
        uintptr(length))
}

func FindResource(hModule HMODULE, lpName, lpType *uint16) (HRSRC, error) {
    ret, _, _ := procFindResource.Call(
        uintptr(hModule),
        uintptr(unsafe.Pointer(lpName)),
        uintptr(unsafe.Pointer(lpType)))

    if ret == 0 {
        return 0, syscall.GetLastError()
    }

    return HRSRC(ret), nil
}

func SizeofResource(hModule HMODULE, hResInfo HRSRC) uint32 {
    ret, _, _ := procSizeofResource.Call(
        uintptr(hModule),
        uintptr(hResInfo))

    if ret == 0 {
        panic("SizeofResource failed")
    }

    return uint32(ret)
}

func LockResource(hResData HGLOBAL) unsafe.Pointer {
    ret, _, _ := procLockResource.Call(uintptr(hResData))

    if ret == 0 {
        panic("LockResource failed")
    }

    return unsafe.Pointer(ret)
}

func LoadResource(hModule HMODULE, hResInfo HRSRC) HGLOBAL {
    ret, _, _ := procLoadResource.Call(
        uintptr(hModule),
        uintptr(hResInfo))

    if ret == 0 {
        panic("LoadResource failed")
    }

    return HGLOBAL(ret)
}

func GetLastError() uint32 {
    ret, _, _ := procGetLastError.Call()
    return uint32(ret)
}
