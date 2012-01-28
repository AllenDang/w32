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

    procGetDeviceCaps        uintptr
    procDeleteObject         uintptr
    procCreateFontIndirect   uintptr
    procAbortDoc             uintptr
    procBitBlt               uintptr
    procCloseEnhMetaFile     uintptr
    procCopyEnhMetaFile      uintptr
    procCreateBrushIndirect  uintptr
    procCreateCompatibleDC   uintptr
    procCreateDC             uintptr
    procCreateDIBSection     uintptr
    procCreateEnhMetaFile    uintptr
    procCreateIC             uintptr
    procDeleteDC             uintptr
    procDeleteEnhMetaFile    uintptr
    procEllipse              uintptr
    procEndDoc               uintptr
    procEndPage              uintptr
    procExtCreatePen         uintptr
    procGetEnhMetaFile       uintptr
    procGetEnhMetaFileHeader uintptr
    procGetObject            uintptr
    procGetStockObject       uintptr
    procGetTextExtentExPoint uintptr
    procGetTextExtentPoint32 uintptr
    procGetTextMetrics       uintptr
    procLineTo               uintptr
    procMoveToEx             uintptr
    procPlayEnhMetaFile      uintptr
    procRectangle            uintptr
    procResetDC              uintptr
    procSelectObject         uintptr
    procSetBkMode            uintptr
    procSetBrushOrgEx        uintptr
    procSetStretchBltMode    uintptr
    procSetTextColor         uintptr
    procSetBkColor           uintptr
    procStartDoc             uintptr
    procStartPage            uintptr
    procStretchBlt           uintptr
)

func init() {
    lib = LoadLib("gdi32.dll")

    procGetDeviceCaps = GetProcAddr(lib, "GetDeviceCaps")
    procDeleteObject = GetProcAddr(lib, "DeleteObject")
    procCreateFontIndirect = GetProcAddr(lib, "CreateFontIndirectW")
    procAbortDoc = GetProcAddr(lib, "AbortDoc")
    procBitBlt = GetProcAddr(lib, "BitBlt")
    procCloseEnhMetaFile = GetProcAddr(lib, "CloseEnhMetaFile")
    procCopyEnhMetaFile = GetProcAddr(lib, "CopyEnhMetaFileW")
    procCreateBrushIndirect = GetProcAddr(lib, "CreateBrushIndirect")
    procCreateCompatibleDC = GetProcAddr(lib, "CreateCompatibleDC")
    procCreateDC = GetProcAddr(lib, "CreateDCW")
    procCreateDIBSection = GetProcAddr(lib, "CreateDIBSection")
    procCreateEnhMetaFile = GetProcAddr(lib, "CreateEnhMetaFileW")
    procCreateIC = GetProcAddr(lib, "CreateICW")
    procDeleteDC = GetProcAddr(lib, "DeleteDC")
    procDeleteEnhMetaFile = GetProcAddr(lib, "DeleteEnhMetaFile")
    procEllipse = GetProcAddr(lib, "Ellipse")
    procEndDoc = GetProcAddr(lib, "EndDoc")
    procEndPage = GetProcAddr(lib, "EndPage")
    procExtCreatePen = GetProcAddr(lib, "ExtCreatePen")
    procGetEnhMetaFile = GetProcAddr(lib, "GetEnhMetaFileW")
    procGetEnhMetaFileHeader = GetProcAddr(lib, "GetEnhMetaFileHeader")
    procGetObject = GetProcAddr(lib, "GetObjectW")
    procGetStockObject = GetProcAddr(lib, "GetStockObject")
    procGetTextExtentExPoint = GetProcAddr(lib, "GetTextExtentExPointW")
    procGetTextExtentPoint32 = GetProcAddr(lib, "GetTextExtentPoint32W")
    procGetTextMetrics = GetProcAddr(lib, "GetTextMetricsW")
    procLineTo = GetProcAddr(lib, "LineTo")
    procMoveToEx = GetProcAddr(lib, "MoveToEx")
    procPlayEnhMetaFile = GetProcAddr(lib, "PlayEnhMetaFile")
    procRectangle = GetProcAddr(lib, "Rectangle")
    procResetDC = GetProcAddr(lib, "ResetDCW")
    procSelectObject = GetProcAddr(lib, "SelectObject")
    procSetBkMode = GetProcAddr(lib, "SetBkMode")
    procSetBrushOrgEx = GetProcAddr(lib, "SetBrushOrgEx")
    procSetStretchBltMode = GetProcAddr(lib, "SetStretchBltMode")
    procSetTextColor = GetProcAddr(lib, "SetTextColor")
    procSetBkColor = GetProcAddr(lib, "SetBkColor")
    procStartDoc = GetProcAddr(lib, "StartDocW")
    procStartPage = GetProcAddr(lib, "StartPage")
    procStretchBlt = GetProcAddr(lib, "StretchBlt")
}

func GetDeviceCaps(hdc HDC, index int) int {
    ret, _, _ := syscall.Syscall(procGetDeviceCaps, 2,
        uintptr(hdc),
        uintptr(index),
        0)

    return int(ret)
}

func DeleteObject(hObject HGDIOBJ) bool {
    ret, _, _ := syscall.Syscall(procDeleteObject, 1,
        uintptr(hObject),
        0,
        0)

    return ret != 0
}

func CreateFontIndirect(logFont *LOGFONT) HFONT {
    ret, _, _ := syscall.Syscall(procCreateFontIndirect, 1,
        uintptr(unsafe.Pointer(logFont)),
        0,
        0)

    return HFONT(ret)
}

func AbortDoc(hdc HDC) int {
    ret, _, _ := syscall.Syscall(procAbortDoc, 1,
        uintptr(hdc),
        0,
        0)

    return int(ret)
}

func BitBlt(hdcDest HDC, nXDest, nYDest, nWidth, nHeight int, hdcSrc HDC, nXSrc, nYSrc int, dwRop uint) {
    ret, _, _ := syscall.Syscall9(procBitBlt, 9,
        uintptr(hdcDest),
        uintptr(nXDest),
        uintptr(nYDest),
        uintptr(nWidth),
        uintptr(nHeight),
        uintptr(hdcSrc),
        uintptr(nXSrc),
        uintptr(nYSrc),
        uintptr(dwRop))

    if ret == 0 {
        panic("BitBlt failed")
    }
}

func CloseEnhMetaFile(hdc HDC) HENHMETAFILE {
    ret, _, _ := syscall.Syscall(procCloseEnhMetaFile, 1,
        uintptr(hdc),
        0,
        0)

    return HENHMETAFILE(ret)
}

func CopyEnhMetaFile(hemfSrc HENHMETAFILE, lpszFile *uint16) HENHMETAFILE {
    ret, _, _ := syscall.Syscall(procCopyEnhMetaFile, 2,
        uintptr(hemfSrc),
        uintptr(unsafe.Pointer(lpszFile)),
        0)

    return HENHMETAFILE(ret)
}

func CreateBrushIndirect(lplb *LOGBRUSH) HBRUSH {
    ret, _, _ := syscall.Syscall(procCreateBrushIndirect, 1,
        uintptr(unsafe.Pointer(lplb)),
        0,
        0)

    return HBRUSH(ret)
}

func CreateCompatibleDC(hdc HDC) HDC {
    ret, _, _ := syscall.Syscall(procCreateCompatibleDC, 1,
        uintptr(hdc),
        0,
        0)

    if ret == 0 {
        panic("Create compatible DC failed")
    }

    return HDC(ret)
}

func CreateDC(lpszDriver, lpszDevice, lpszOutput *uint16, lpInitData *DEVMODE) HDC {
    ret, _, _ := syscall.Syscall6(procCreateDC, 4,
        uintptr(unsafe.Pointer(lpszDriver)),
        uintptr(unsafe.Pointer(lpszDevice)),
        uintptr(unsafe.Pointer(lpszOutput)),
        uintptr(unsafe.Pointer(lpInitData)),
        0,
        0)

    return HDC(ret)
}

func CreateDIBSection(hdc HDC, pbmi *BITMAPINFO, iUsage uint, ppvBits *unsafe.Pointer, hSection HANDLE, dwOffset uint) HBITMAP {
    ret, _, _ := syscall.Syscall6(procCreateDIBSection, 6,
        uintptr(hdc),
        uintptr(unsafe.Pointer(pbmi)),
        uintptr(iUsage),
        uintptr(unsafe.Pointer(ppvBits)),
        uintptr(hSection),
        uintptr(dwOffset))

    return HBITMAP(ret)
}

func CreateEnhMetaFile(hdcRef HDC, lpFilename *uint16, lpRect *RECT, lpDescription *uint16) HDC {
    ret, _, _ := syscall.Syscall6(procCreateEnhMetaFile, 4,
        uintptr(hdcRef),
        uintptr(unsafe.Pointer(lpFilename)),
        uintptr(unsafe.Pointer(lpRect)),
        uintptr(unsafe.Pointer(lpDescription)),
        0,
        0)

    return HDC(ret)
}

func CreateIC(lpszDriver, lpszDevice, lpszOutput *uint16, lpdvmInit *DEVMODE) HDC {
    ret, _, _ := syscall.Syscall6(procCreateIC, 4,
        uintptr(unsafe.Pointer(lpszDriver)),
        uintptr(unsafe.Pointer(lpszDevice)),
        uintptr(unsafe.Pointer(lpszOutput)),
        uintptr(unsafe.Pointer(lpdvmInit)),
        0,
        0)

    return HDC(ret)
}

func DeleteDC(hdc HDC) bool {
    ret, _, _ := syscall.Syscall(procDeleteDC, 1,
        uintptr(hdc),
        0,
        0)

    return ret != 0
}

func DeleteEnhMetaFile(hemf HENHMETAFILE) bool {
    ret, _, _ := syscall.Syscall(procDeleteEnhMetaFile, 1,
        uintptr(hemf),
        0,
        0)

    return ret != 0
}

func Ellipse(hdc HDC, nLeftRect, nTopRect, nRightRect, nBottomRect int) bool {
    ret, _, _ := syscall.Syscall6(procEllipse, 5,
        uintptr(hdc),
        uintptr(nLeftRect),
        uintptr(nTopRect),
        uintptr(nRightRect),
        uintptr(nBottomRect),
        0)

    return ret != 0
}

func EndDoc(hdc HDC) int {
    ret, _, _ := syscall.Syscall(procEndDoc, 1,
        uintptr(hdc),
        0,
        0)

    return int(ret)
}

func EndPage(hdc HDC) int {
    ret, _, _ := syscall.Syscall(procEndPage, 1,
        uintptr(hdc),
        0,
        0)

    return int(ret)
}

func ExtCreatePen(dwPenStyle, dwWidth uint, lplb *LOGBRUSH, dwStyleCount uint, lpStyle *uint) HPEN {
    ret, _, _ := syscall.Syscall6(procExtCreatePen, 5,
        uintptr(dwPenStyle),
        uintptr(dwWidth),
        uintptr(unsafe.Pointer(lplb)),
        uintptr(dwStyleCount),
        uintptr(unsafe.Pointer(lpStyle)),
        0)

    return HPEN(ret)
}

func GetEnhMetaFile(lpszMetaFile *uint16) HENHMETAFILE {
    ret, _, _ := syscall.Syscall(procGetEnhMetaFile, 1,
        uintptr(unsafe.Pointer(lpszMetaFile)),
        0,
        0)

    return HENHMETAFILE(ret)
}

func GetEnhMetaFileHeader(hemf HENHMETAFILE, cbBuffer uint, lpemh *ENHMETAHEADER) uint {
    ret, _, _ := syscall.Syscall(procGetEnhMetaFileHeader, 3,
        uintptr(hemf),
        uintptr(cbBuffer),
        uintptr(unsafe.Pointer(lpemh)))

    return uint(ret)
}

func GetObject(hgdiobj HGDIOBJ, cbBuffer uintptr, lpvObject unsafe.Pointer) int {
    ret, _, _ := syscall.Syscall(procGetObject, 3,
        uintptr(hgdiobj),
        uintptr(cbBuffer),
        uintptr(lpvObject))

    return int(ret)
}

func GetStockObject(fnObject int) HGDIOBJ {
    ret, _, _ := syscall.Syscall(procGetDeviceCaps, 1,
        uintptr(fnObject),
        0,
        0)

    return HGDIOBJ(ret)
}

func GetTextExtentExPoint(hdc HDC, lpszStr *uint16, cchString, nMaxExtent int, lpnFit, alpDx *int, lpSize *SIZE) bool {
    ret, _, _ := syscall.Syscall9(procGetTextExtentExPoint, 7,
        uintptr(hdc),
        uintptr(unsafe.Pointer(lpszStr)),
        uintptr(cchString),
        uintptr(nMaxExtent),
        uintptr(unsafe.Pointer(lpnFit)),
        uintptr(unsafe.Pointer(alpDx)),
        uintptr(unsafe.Pointer(lpSize)),
        0,
        0)

    return ret != 0
}

func GetTextExtentPoint32(hdc HDC, lpString *uint16, c int, lpSize *SIZE) bool {
    ret, _, _ := syscall.Syscall6(procGetTextExtentPoint32, 4,
        uintptr(hdc),
        uintptr(unsafe.Pointer(lpString)),
        uintptr(c),
        uintptr(unsafe.Pointer(lpSize)),
        0,
        0)

    return ret != 0
}

func GetTextMetrics(hdc HDC, lptm *TEXTMETRIC) bool {
    ret, _, _ := syscall.Syscall(procGetTextMetrics, 2,
        uintptr(hdc),
        uintptr(unsafe.Pointer(lptm)),
        0)

    return ret != 0
}

func LineTo(hdc HDC, nXEnd, nYEnd int) bool {
    ret, _, _ := syscall.Syscall(procLineTo, 3,
        uintptr(hdc),
        uintptr(nXEnd),
        uintptr(nYEnd))

    return ret != 0
}

func MoveToEx(hdc HDC, x, y int, lpPoint *POINT) bool {
    ret, _, _ := syscall.Syscall6(procMoveToEx, 4,
        uintptr(hdc),
        uintptr(x),
        uintptr(y),
        uintptr(unsafe.Pointer(lpPoint)),
        0,
        0)

    return ret != 0
}

func PlayEnhMetaFile(hdc HDC, hemf HENHMETAFILE, lpRect *RECT) bool {
    ret, _, _ := syscall.Syscall(procPlayEnhMetaFile, 3,
        uintptr(hdc),
        uintptr(hemf),
        uintptr(unsafe.Pointer(lpRect)))

    return ret != 0
}

func Rectangle(hdc HDC, nLeftRect, nTopRect, nRightRect, nBottomRect int) bool {
    ret, _, _ := syscall.Syscall6(procRectangle, 5,
        uintptr(hdc),
        uintptr(nLeftRect),
        uintptr(nTopRect),
        uintptr(nRightRect),
        uintptr(nBottomRect),
        0)

    return ret != 0
}

func ResetDC(hdc HDC, lpInitData *DEVMODE) HDC {
    ret, _, _ := syscall.Syscall(procResetDC, 2,
        uintptr(hdc),
        uintptr(unsafe.Pointer(lpInitData)),
        0)

    return HDC(ret)
}

func SelectObject(hdc HDC, hgdiobj HGDIOBJ) HGDIOBJ {
    ret, _, _ := syscall.Syscall(procSelectObject, 2,
        uintptr(hdc),
        uintptr(hgdiobj),
        0)

    if ret == 0 {
        panic("SelectObject failed")
    }

    return HGDIOBJ(ret)
}

func SetBkMode(hdc HDC, iBkMode int) int {
    ret, _, _ := syscall.Syscall(procSetBkMode, 2,
        uintptr(hdc),
        uintptr(iBkMode),
        0)

    if ret == 0 {
        panic("SetBkMode failed")
    }

    return int(ret)
}

func SetBrushOrgEx(hdc HDC, nXOrg, nYOrg int, lppt *POINT) bool {
    ret, _, _ := syscall.Syscall6(procSetBrushOrgEx, 4,
        uintptr(hdc),
        uintptr(nXOrg),
        uintptr(nYOrg),
        uintptr(unsafe.Pointer(lppt)),
        0,
        0)

    return ret != 0
}

func SetStretchBltMode(hdc HDC, iStretchMode int) int {
    ret, _, _ := syscall.Syscall(procSetStretchBltMode, 2,
        uintptr(hdc),
        uintptr(iStretchMode),
        0)

    return int(ret)
}

func SetTextColor(hdc HDC, crColor COLORREF) COLORREF {
    ret, _, _ := syscall.Syscall(procSetTextColor, 2,
        uintptr(hdc),
        uintptr(crColor),
        0)

    if ret == CLR_INVALID {
        panic("SetTextColor failed")
    }

    return COLORREF(ret)
}

func SetBkColor(hdc HDC, crColor COLORREF) COLORREF {
    ret, _, _ := syscall.Syscall(procSetBkColor, 2,
        uintptr(hdc),
        uintptr(crColor),
        0)

    if ret == CLR_INVALID {
        panic("SetBkColor failed")
    }

    return COLORREF(ret)
}

func StartDoc(hdc HDC, lpdi *DOCINFO) int {
    ret, _, _ := syscall.Syscall(procStartDoc, 2,
        uintptr(hdc),
        uintptr(unsafe.Pointer(lpdi)),
        0)

    return int(ret)
}

func StartPage(hdc HDC) int {
    ret, _, _ := syscall.Syscall(procStartPage, 1,
        uintptr(hdc),
        0,
        0)

    return int(ret)
}

func StretchBlt(hdcDest HDC, nXOriginDest, nYOriginDest, nWidthDest, nHeightDest int, hdcSrc HDC, nXOriginSrc, nYOriginSrc, nWidthSrc, nHeightSrc int, dwRop uint) {
    ret, _, _ := syscall.Syscall12(procStretchBlt, 11,
        uintptr(hdcDest),
        uintptr(nXOriginDest),
        uintptr(nYOriginDest),
        uintptr(nWidthDest),
        uintptr(nHeightDest),
        uintptr(hdcSrc),
        uintptr(nXOriginSrc),
        uintptr(nYOriginSrc),
        uintptr(nWidthSrc),
        uintptr(nHeightSrc),
        uintptr(dwRop),
        0)

    if ret == 0 {
        panic("StretchBlt failed")
    }
}
