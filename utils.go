// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package w32

import (
    "fmt"
    "syscall"
    "unsafe"
    "unicode/utf16"
)

func MakeIntResource(id uint16) *uint16 {
    return (*uint16)(unsafe.Pointer(uintptr(id)))
}

func LOWORD(dw uint) uint16 {
    return uint16(dw)
}

func HIWORD(dw uint) uint16 {
    return uint16(dw >> 16 & 0xffff)
}

func LoadLib(name string) uintptr {
    lib, err := syscall.LoadLibrary(name)
    if err != nil {
        panic(fmt.Sprintf("syscall.LoadLibrary('%s') failed: %s",
            name, err.Error()))
    }
    return uintptr(lib)
}

func GetProcAddr(lib uintptr, name string) uintptr {
    addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
    if err != nil {
        panic(fmt.Sprintf("syscal.GetProcAddress(%d, '%s') failed: %s",
            lib, name, err.Error()))
    }
    return uintptr(addr)
}

func BoolToBOOL(value bool) BOOL {
	if value {
		return 1
	}

	return 0
}

func UTF16PtrToString(cstr *uint16) string {
    if cstr != nil {
        us := make([]uint16, 0, 256)
        for p := uintptr(unsafe.Pointer(cstr)); ; p += 2 {
            u := *(*uint16)(unsafe.Pointer(p))
            if u == 0 {
                return string(utf16.Decode(us))
            }
            us = append(us, u)
        }
    }

    return ""
}
