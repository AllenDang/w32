// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdi32

import (
	"syscall"
	"unsafe"
    . "w32"
)

var (
    lib uintptr

    procInitCommonControlsEx uintptr
)

func init() {
    lib = LoadLib("comctl32.dll")

    procInitCommonControlsEx = GetProcAddr(lib, "InitCommonControlsEx")

    // Initialize the common controls
	var initCtrls INITCOMMONCONTROLSEX
	initCtrls.DwSize = uint32(unsafe.Sizeof(initCtrls))
	initCtrls.DwICC = ICC_LISTVIEW_CLASSES | ICC_PROGRESS_CLASS | ICC_TAB_CLASSES | ICC_TREEVIEW_CLASSES

	InitCommonControlsEx(&initCtrls)
}

func InitCommonControlsEx(lpInitCtrls *INITCOMMONCONTROLSEX) bool {
	ret, _, _ := syscall.Syscall(procInitCommonControlsEx, 1,
		uintptr(unsafe.Pointer(lpInitCtrls)),
		0,
		0)
	
	return ret != 0
}