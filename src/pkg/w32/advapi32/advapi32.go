// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package advapi32

import (
    "fmt"
    "syscall"
    "unsafe"
    . "w32"
)

var (
    lib uintptr

    procRegOpenKeyEx uintptr
    procRegCloseKey  uintptr
    procRegGetValue  uintptr
    procRegEnumKeyEx uintptr
)

func init() {
    lib = LoadLib("advapi32.dll")

    procRegOpenKeyEx = GetProcAddr(lib, "RegOpenKeyExW")
    procRegCloseKey = GetProcAddr(lib, "RegCloseKey")
    procRegGetValue = GetProcAddr(lib, "RegGetValueW")
    procRegEnumKeyEx = GetProcAddr(lib, "RegEnumKeyExW")
}

func RegOpenKeyEx(hKey HKEY, subKey string, samDesired uint32) HKEY {
    var result HKEY
    ret, _, _ := syscall.Syscall6(procRegOpenKeyEx, 5,
        uintptr(hKey),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))),
        uintptr(0),
        uintptr(samDesired),
        uintptr(unsafe.Pointer(&result)),
        0)

    if ret != ERROR_SUCCESS {
        panic(fmt.Sprintf("RegOpenKeyEx(%d, %s, %d) failed", hKey, subKey, samDesired))
    }
    return result
}

func RegCloseKey(hKey HKEY) {
    ret, _, _ := syscall.Syscall(procRegCloseKey, 1,
        uintptr(hKey),
        0,
        0)

    if ret != ERROR_SUCCESS {
        panic(fmt.Sprintf("RegCloseKey(%d) failed", hKey))
    }
}

func RegGetString(hKey HKEY, subKey string, value string) string {
    var bufLen uint32
    syscall.Syscall9(procRegGetValue, 7,
        uintptr(hKey),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(value))),
        uintptr(RRF_RT_REG_SZ),
        0,
        0,
        uintptr(unsafe.Pointer(&bufLen)),
        0,
        0)

    if bufLen == 0 {
        return ""
    }

    buf := make([]uint16, bufLen)
    ret, _, _ := syscall.Syscall9(procRegGetValue, 7,
        uintptr(hKey),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(value))),
        uintptr(RRF_RT_REG_SZ),
        0,
        uintptr(unsafe.Pointer(&buf[0])),
        uintptr(unsafe.Pointer(&bufLen)),
        0,
        0)

    if ret != ERROR_SUCCESS {
        return ""	
    }

    return syscall.UTF16ToString(buf)
}

func RegEnumKeyEx(hKey HKEY, index uint32) string {
    var bufLen uint32 = 255
    buf := make([]uint16, bufLen)
    syscall.Syscall9(procRegEnumKeyEx, 8,
        uintptr(hKey),
        uintptr(index),
        uintptr(unsafe.Pointer(&buf[0])),
        uintptr(unsafe.Pointer(&bufLen)),
        0,
        0,
        0,
        0,
        0)
    return syscall.UTF16ToString(buf)
}
