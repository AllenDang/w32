// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package advapi32

import (
    "fmt"
    . "github.com/AllenDang/w32"
    "syscall"
    "unsafe"
)

var (
    modadvapi32 = syscall.NewLazyDLL("advapi32.dll")

    procRegOpenKeyEx   = modadvapi32.NewProc("RegOpenKeyExW")
    procRegCloseKey    = modadvapi32.NewProc("RegCloseKey")
    procRegGetValue    = modadvapi32.NewProc("RegGetValueW")
    procRegEnumKeyEx   = modadvapi32.NewProc("RegEnumKeyExW")
    procRegSetKeyValue = modadvapi32.NewProc("RegSetKeyValueW")
)

func RegOpenKeyEx(hKey HKEY, subKey string, samDesired uint32) HKEY {
    var result HKEY
    ret, _, _ := procRegOpenKeyEx.Call(
        uintptr(hKey),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))),
        uintptr(0),
        uintptr(samDesired),
        uintptr(unsafe.Pointer(&result)))

    if ret != ERROR_SUCCESS {
        panic(fmt.Sprintf("RegOpenKeyEx(%d, %s, %d) failed", hKey, subKey, samDesired))
    }
    return result
}

func RegCloseKey(hKey HKEY) {
    ret, _, _ := procRegCloseKey.Call(
        uintptr(hKey))

    if ret != ERROR_SUCCESS {
        panic(fmt.Sprintf("RegCloseKey(%d) failed", hKey))
    }
}

func RegGetString(hKey HKEY, subKey string, value string) string {
    var bufLen uint32
    procRegGetValue.Call(
        uintptr(hKey),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(value))),
        uintptr(RRF_RT_REG_SZ),
        0,
        0,
        uintptr(unsafe.Pointer(&bufLen)))

    if bufLen == 0 {
        return ""
    }

    buf := make([]uint16, bufLen)
    ret, _, _ := procRegGetValue.Call(
        uintptr(hKey),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(value))),
        uintptr(RRF_RT_REG_SZ),
        0,
        uintptr(unsafe.Pointer(&buf[0])),
        uintptr(unsafe.Pointer(&bufLen)))

    if ret != ERROR_SUCCESS {
        return ""
    }

    return syscall.UTF16ToString(buf)
}

func RegSetKeyValue(hKey HKEY, subKey string, valueName string, dwType uint32, data uintptr, cbData uint32) (errno int) {
    ret, _, _ := procRegSetKeyValue.Call(
        uintptr(hKey),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(subKey))),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(valueName))),
        uintptr(dwType),
        data,
        uintptr(cbData))

    return int(ret)
}

func RegEnumKeyEx(hKey HKEY, index uint32) string {
    var bufLen uint32 = 255
    buf := make([]uint16, bufLen)
    procRegEnumKeyEx.Call(
        uintptr(hKey),
        uintptr(index),
        uintptr(unsafe.Pointer(&buf[0])),
        uintptr(unsafe.Pointer(&bufLen)),
        0,
        0,
        0,
        0)
    return syscall.UTF16ToString(buf)
}
