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
    procRegCloseKey uintptr
)

func init() {
    lib = LoadLib("advapi32.dll")

    procRegOpenKeyEx = GetProcAddr(lib, "RegOpenKeyExW")
    procRegCloseKey = GetProcAddr(lib, "RegCloseKey")
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