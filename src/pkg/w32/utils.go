// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package w32

import (
    "fmt"
    "syscall"
    "unsafe"
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
    if err != 0 {
        panic(fmt.Sprintf("syscall.LoadLibrary('%s') failed: %s",
            name, syscall.Errstr(err)))
    }
    return uintptr(lib)
}

func GetProcAddr(lib uintptr, name string) uintptr {
    addr, err := syscall.GetProcAddress(syscall.Handle(lib), name)
    if err != 0 {
        panic(fmt.Sprintf("syscal.GetProcAddress(%d, '%s') failed: %s",
            lib, name, syscall.Errstr(err)))
    }
    return uintptr(addr)
}

func BoolToBOOL(value bool) BOOL {
	if value {
		return 1
	}

	return 0
}

func UTF16PtrToString(s *uint16) string {
	return syscall.UTF16ToString((*[1 << 30]uint16)(unsafe.Pointer(s))[0:])
}
