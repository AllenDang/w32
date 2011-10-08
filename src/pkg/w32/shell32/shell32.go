// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shell32

import (
    "syscall"
    "unsafe"
    "os"
    . "w32"
)

var (
    lib uintptr

    procSHBrowseForFolder   uintptr
    procSHGetPathFromIDList uintptr
    procDragAcceptFiles     uintptr
    procDragQueryFile       uintptr
    procDragQueryPoint      uintptr
    procDragFinish          uintptr
    procShellExecute        uintptr
    procExtractIcon         uintptr
)

func init() {
    lib = LoadLib("shell32.dll")

    procSHBrowseForFolder = GetProcAddr(lib, "SHBrowseForFolderW")
    procSHGetPathFromIDList = GetProcAddr(lib, "SHGetPathFromIDListW")
    procDragAcceptFiles = GetProcAddr(lib, "DragAcceptFiles")
    procDragQueryFile = GetProcAddr(lib, "DragQueryFileW")
    procDragQueryPoint = GetProcAddr(lib, "DragQueryPoint")
    procDragFinish = GetProcAddr(lib, "DragFinish")
    procShellExecute = GetProcAddr(lib, "ShellExecuteW")
    procExtractIcon = GetProcAddr(lib, "ExtractIconW")
}

func SHBrowseForFolder(bi *BROWSEINFO) uintptr {
    ret, _, _ := syscall.Syscall(procSHBrowseForFolder, 1,
        uintptr(unsafe.Pointer(bi)),
        0,
        0)

    return ret
}

func SHGetPathFromIDList(idl uintptr) string {
    buf := make([]uint16, 1024)
    syscall.Syscall(procSHGetPathFromIDList, 2,
        idl,
        uintptr(unsafe.Pointer(&buf[0])),
        0)

    return syscall.UTF16ToString(buf)
}

func DragAcceptFiles(hwnd HWND, accept bool) {
    syscall.Syscall(procDragAcceptFiles, 2,
        uintptr(hwnd),
        uintptr(BoolToBOOL(accept)),
        0)
}

func DragQueryFile(hDrop HDROP, iFile uint) (fileName string, fileCount uint) {
    ret, _, _ := syscall.Syscall6(procDragQueryFile, 4,
        uintptr(hDrop),
        uintptr(iFile),
        0,
        0,
        0,
        0)

    fileCount = uint(ret)

    if iFile != 0xFFFFFFFF {
        buf := make([]uint16, fileCount+1)

        ret, _, _ := syscall.Syscall6(procDragQueryFile, 4,
            uintptr(hDrop),
            uintptr(iFile),
            uintptr(unsafe.Pointer(&buf[0])),
            uintptr(fileCount+1),
            0,
            0)

        if ret == 0 {
            panic("Invoke DragQueryFile error.")
        }

        fileName = syscall.UTF16ToString(buf)
    }

    return
}

func DragQueryPoint(hDrop HDROP) (x, y int, isClientArea bool) {
    var pt POINT
    ret, _, _ := syscall.Syscall(procDragQueryPoint, 2,
        uintptr(hDrop),
        uintptr(unsafe.Pointer(&pt)),
        0)

    isClientArea = false
    if ret == 1 {
        isClientArea = true
    }

    x = pt.X
    y = pt.Y

    return
}

func DragFinish(hDrop HDROP) {
    syscall.Syscall(procDragFinish, 1,
        uintptr(hDrop),
        0,
        0)
}

func ShellExecute(hwnd HWND, lpOperation, lpFile, lpParameters, lpDirectory string, nShowCmd int) (int, os.Error) {
    var op, param, directory uintptr
    if len(lpOperation) != 0 {
        op = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpOperation)))
    }
    if len(lpParameters) != 0 {
        param = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpParameters)))
    }
    if len(lpDirectory) != 0 {
        directory = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpDirectory)))
    }

    ret, _, _ := syscall.Syscall6(procShellExecute, 6,
        uintptr(hwnd),
        op,
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpFile))),
        param,
        directory,
        uintptr(nShowCmd))

    errorMsg := ""
    if ret != 0 {
        switch int(ret) {
        case ERROR_FILE_NOT_FOUND:
            errorMsg = "The specified file was not found."
        case ERROR_PATH_NOT_FOUND:
            errorMsg = "The specified path was not found."
        case ERROR_BAD_FORMAT:
            errorMsg = "The .exe file is invalid (non-Win32 .exe or error in .exe image)."
        case SE_ERR_ACCESSDENIED:
            errorMsg = "The operating system denied access to the specified file."
        case SE_ERR_ASSOCINCOMPLETE:
            errorMsg = "The file name association is incomplete or invalid."
        case SE_ERR_DDEBUSY:
            errorMsg = "The DDE transaction could not be completed because other DDE transactions were being processed."
        case SE_ERR_DDEFAIL:
            errorMsg = "The DDE transaction failed."
        case SE_ERR_DDETIMEOUT:
            errorMsg = "The DDE transaction could not be completed because the request timed out."
        case SE_ERR_DLLNOTFOUND:
            errorMsg = "The specified DLL was not found."
        case SE_ERR_NOASSOC:
            errorMsg = "There is no application associated with the given file name extension. This error will also be returned if you attempt to print a file that is not printable."
        case SE_ERR_OOM:
            errorMsg = "There was not enough memory to complete the operation."
        case SE_ERR_SHARE:
            errorMsg = "A sharing violation occurred."
        }
    }

    return int(ret), os.NewError(errorMsg)
}

func ExtractIcon(lpszExeFileName string, nIconIndex uint) HICON {
    ret, _, _ := syscall.Syscall(procExtractIcon, 3,
        0,
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpszExeFileName))),
        uintptr(nIconIndex))

    return HICON(ret)
}
