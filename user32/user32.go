// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package user32

import (
    "fmt"
    "syscall"
    "unsafe"
    . "github.com/AllenDang/w32"
)

var (
    lib uintptr

    procRegisterClassEx          uintptr
    procLoadIcon                 uintptr
    procLoadCursor               uintptr
    procShowWindow               uintptr
    procUpdateWindow             uintptr
    procCreateWindowEx           uintptr
    procDestroyWindow            uintptr
    procDefWindowProc            uintptr
    procDefDlgProc               uintptr
    procPostQuitMessage          uintptr
    procGetMessage               uintptr
    procTranslateMessage         uintptr
    procDispatchMessage          uintptr
    procSendMessage              uintptr
    procPostMessage              uintptr
    procSetWindowText            uintptr
    procGetWindowTextLength      uintptr
    procGetWindowText            uintptr
    procGetWindowRect            uintptr
    procMoveWindow               uintptr
    procScreenToClient           uintptr
    procCallWindowProc           uintptr
    procSetWindowLongPtr         uintptr
    procEnableWindow             uintptr
    procIsWindowEnabled          uintptr
    procIsWindowVisible          uintptr
    procSetFocus                 uintptr
    procInvalidateRect           uintptr
    procGetClientRect            uintptr
    procGetDC                    uintptr
    procReleaseDC                uintptr
    procSetCapture               uintptr
    procReleaseCapture           uintptr
    procGetWindowThreadProcessId uintptr
    procMessageBox               uintptr
    procGetSystemMetrics         uintptr
    procGetWindowLongPtr         uintptr
    procCopyRect                 uintptr
    procEqualRect                uintptr
    procInflateRect              uintptr
    procIntersectRect            uintptr
    procIsRectEmpty              uintptr
    procOffsetRect               uintptr
    procPtInRect                 uintptr
    procSetRect                  uintptr
    procSetRectEmpty             uintptr
    procSubtractRect             uintptr
    procUnionRect                uintptr
    procCreateDialogParam        uintptr
    procDialogBoxParam           uintptr
    procGetDlgItem               uintptr
    procDrawIcon                 uintptr
    procClientToScreen           uintptr
    procIsDialogMessage          uintptr
    procIsWindow                 uintptr
    procEndDialog                uintptr
    procSetWindowLong            uintptr
    procPeekMessage              uintptr
    procTranslateAccelerator     uintptr
    procDestroyIcon              uintptr
    procSetWindowPos             uintptr
    procFillRect                 uintptr
    procDrawText                 uintptr
)

func init() {
    lib = LoadLib("user32.dll")

    procRegisterClassEx = GetProcAddr(lib, "RegisterClassExW")
    procLoadIcon = GetProcAddr(lib, "LoadIconW")
    procLoadCursor = GetProcAddr(lib, "LoadCursorW")
    procShowWindow = GetProcAddr(lib, "ShowWindow")
    procUpdateWindow = GetProcAddr(lib, "UpdateWindow")
    procCreateWindowEx = GetProcAddr(lib, "CreateWindowExW")
    procDestroyWindow = GetProcAddr(lib, "DestroyWindow")
    procDefWindowProc = GetProcAddr(lib, "DefWindowProcW")
    procDefDlgProc = GetProcAddr(lib, "DefDlgProcW")
    procPostQuitMessage = GetProcAddr(lib, "PostQuitMessage")
    procGetMessage = GetProcAddr(lib, "GetMessageW")
    procTranslateMessage = GetProcAddr(lib, "TranslateMessage")
    procDispatchMessage = GetProcAddr(lib, "DispatchMessageW")
    procSendMessage = GetProcAddr(lib, "SendMessageW")
    procPostMessage = GetProcAddr(lib, "PostMessageW")
    procSetWindowText = GetProcAddr(lib, "SetWindowTextW")
    procGetWindowTextLength = GetProcAddr(lib, "GetWindowTextLengthW")
    procGetWindowText = GetProcAddr(lib, "GetWindowTextW")
    procGetWindowRect = GetProcAddr(lib, "GetWindowRect")
    procMoveWindow = GetProcAddr(lib, "MoveWindow")
    procScreenToClient = GetProcAddr(lib, "ScreenToClient")
    procCallWindowProc = GetProcAddr(lib, "CallWindowProcW")
    procSetWindowLongPtr = GetProcAddr(lib, "SetWindowLongW")
    procEnableWindow = GetProcAddr(lib, "EnableWindow")
    procIsWindowEnabled = GetProcAddr(lib, "IsWindowEnabled")
    procIsWindowVisible = GetProcAddr(lib, "IsWindowVisible")
    procSetFocus = GetProcAddr(lib, "SetFocus")
    procInvalidateRect = GetProcAddr(lib, "InvalidateRect")
    procGetClientRect = GetProcAddr(lib, "GetClientRect")
    procGetDC = GetProcAddr(lib, "GetDC")
    procReleaseDC = GetProcAddr(lib, "ReleaseDC")
    procSetCapture = GetProcAddr(lib, "SetCapture")
    procReleaseCapture = GetProcAddr(lib, "ReleaseCapture")
    procGetWindowThreadProcessId = GetProcAddr(lib, "GetWindowThreadProcessId")
    procMessageBox = GetProcAddr(lib, "MessageBoxW")
    procGetSystemMetrics = GetProcAddr(lib, "GetSystemMetrics")
    procGetWindowLongPtr = GetProcAddr(lib, "GetWindowLongW")
    procCopyRect = GetProcAddr(lib, "CopyRect")
    procEqualRect = GetProcAddr(lib, "EqualRect")
    procInflateRect = GetProcAddr(lib, "InflateRect")
    procIntersectRect = GetProcAddr(lib, "IntersectRect")
    procIsRectEmpty = GetProcAddr(lib, "IsRectEmpty")
    procOffsetRect = GetProcAddr(lib, "OffsetRect")
    procPtInRect = GetProcAddr(lib, "PtInRect")
    procSetRect = GetProcAddr(lib, "SetRect")
    procSetRectEmpty = GetProcAddr(lib, "SetRectEmpty")
    procSubtractRect = GetProcAddr(lib, "SubtractRect")
    procUnionRect = GetProcAddr(lib, "UnionRect")
    procCreateDialogParam = GetProcAddr(lib, "CreateDialogParamW")
    procDialogBoxParam = GetProcAddr(lib, "DialogBoxParamW")
    procGetDlgItem = GetProcAddr(lib, "GetDlgItem")
    procDrawIcon = GetProcAddr(lib, "DrawIcon")
    procClientToScreen = GetProcAddr(lib, "ClientToScreen")
    procIsDialogMessage = GetProcAddr(lib, "IsDialogMessageW")
    procIsWindow = GetProcAddr(lib, "IsWindow")
    procEndDialog = GetProcAddr(lib, "EndDialog")
    procSetWindowLong = GetProcAddr(lib, "SetWindowLongW")
    procPeekMessage = GetProcAddr(lib, "PeekMessageW")
    procTranslateAccelerator = GetProcAddr(lib, "TranslateAcceleratorW")
    procDestroyIcon = GetProcAddr(lib, "DestroyIcon")
    procSetWindowPos = GetProcAddr(lib, "SetWindowPos")
    procFillRect = GetProcAddr(lib, "FillRect")
    procDrawText = GetProcAddr(lib, "DrawTextW")
}

func RegisterClassEx(wndClassEx *WNDCLASSEX) ATOM {
    ret, _, _ := syscall.Syscall(procRegisterClassEx, 1,
        uintptr(unsafe.Pointer(wndClassEx)),
        0,
        0)
    return ATOM(ret)
}

func LoadIcon(instance HINSTANCE, iconName *uint16) HICON {
    ret, _, _ := syscall.Syscall(procLoadIcon, 2,
        uintptr(instance),
        uintptr(unsafe.Pointer(iconName)),
        0)

    return HICON(ret)

}

func LoadCursor(instance HINSTANCE, cursorName *uint16) HCURSOR {
    ret, _, _ := syscall.Syscall(procLoadCursor, 2,
        uintptr(instance),
        uintptr(unsafe.Pointer(cursorName)),
        0)

    return HCURSOR(ret)

}

func ShowWindow(hwnd HWND, cmdshow int) bool {
    ret, _, _ := syscall.Syscall(procShowWindow, 2,
        uintptr(hwnd),
        uintptr(cmdshow),
        0)

    return ret != 0

}

func UpdateWindow(hwnd HWND) bool {
    ret, _, _ := syscall.Syscall(procUpdateWindow, 1,
        uintptr(hwnd),
        0,
        0)
    return ret != 0
}

func CreateWindowEx(exStyle uint, className, windowName *uint16,
style uint, x, y, width, height int, parent HWND, menu HMENU,
instance HINSTANCE, param unsafe.Pointer) HWND {
    ret, _, _ := syscall.Syscall12(procCreateWindowEx, 12,
        uintptr(exStyle),
        uintptr(unsafe.Pointer(className)),
        uintptr(unsafe.Pointer(windowName)),
        uintptr(style),
        uintptr(x),
        uintptr(y),
        uintptr(width),
        uintptr(height),
        uintptr(parent),
        uintptr(menu),
        uintptr(instance),
        uintptr(param))

    return HWND(ret)
}

func DestroyWindow(hwnd HWND) bool {
    ret, _, _ := syscall.Syscall(procDestroyWindow, 1,
        uintptr(hwnd),
        0,
        0)

    return ret != 0
}

func DefWindowProc(hwnd HWND, msg uint, wParam, lParam uintptr) uintptr {
    ret, _, _ := syscall.Syscall6(procDefWindowProc, 4,
        uintptr(hwnd),
        uintptr(msg),
        wParam,
        lParam,
        0,
        0)

    return ret
}

func DefDlgProc(hwnd HWND, msg uint, wParam, lParam uintptr) uintptr {
    ret, _, _ := syscall.Syscall6(procDefDlgProc, 4,
        uintptr(hwnd),
        uintptr(msg),
        wParam,
        lParam,
        0,
        0)

    return ret
}

func PostQuitMessage(exitCode int) {
    syscall.Syscall(procPostQuitMessage, 1,
        uintptr(exitCode),
        0,
        0)
}

func GetMessage(msg *MSG, hwnd HWND, msgFilterMin, msgFilterMax uint32) int {
    ret, _, _ := syscall.Syscall6(procGetMessage, 4,
        uintptr(unsafe.Pointer(msg)),
        uintptr(hwnd),
        uintptr(msgFilterMin),
        uintptr(msgFilterMax),
        0,
        0)

    return int(ret)
}

func TranslateMessage(msg *MSG) bool {
    ret, _, _ := syscall.Syscall(procTranslateMessage, 1,
        uintptr(unsafe.Pointer(msg)),
        0,
        0)

    return ret != 0

}

func DispatchMessage(msg *MSG) uintptr {
    ret, _, _ := syscall.Syscall(procDispatchMessage, 1,
        uintptr(unsafe.Pointer(msg)),
        0,
        0)

    return ret

}

func SendMessage(hwnd HWND, msg uint, wParam, lParam uintptr) uintptr {
    ret, _, _ := syscall.Syscall6(procSendMessage, 4,
        uintptr(hwnd),
        uintptr(msg),
        wParam,
        lParam,
        0,
        0)

    return ret
}

func PostMessage(hwnd HWND, msg uint, wParam, lParam uintptr) bool {
    ret, _, _ := syscall.Syscall6(procPostMessage, 4,
        uintptr(hwnd),
        uintptr(msg),
        wParam,
        lParam,
        0,
        0)

    return ret != 0
}

func SetWindowText(hwnd HWND, text string) {
    syscall.Syscall(procSetWindowText, 2,
        uintptr(hwnd),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
        0)
}

func GetWindowTextLength(hwnd HWND) int {
    ret, _, _ := syscall.Syscall(procGetWindowTextLength, 1,
        uintptr(hwnd),
        0,
        0)

    return int(ret)
}

func GetWindowText(hwnd HWND) string {
    textLen := GetWindowTextLength(hwnd) + 1

    buf := make([]uint16, textLen)
    syscall.Syscall(procGetWindowText, 3,
        uintptr(hwnd),
        uintptr(unsafe.Pointer(&buf[0])),
        uintptr(textLen))

    return syscall.UTF16ToString(buf)
}

func GetWindowRect(hwnd HWND) *RECT {
    var rect RECT
    syscall.Syscall(procGetWindowRect, 2,
        uintptr(hwnd),
        uintptr(unsafe.Pointer(&rect)),
        0)

    return &rect
}

func MoveWindow(hwnd HWND, x, y, width, height int, repaint bool) bool {
    ret, _, _ := syscall.Syscall6(procMoveWindow, 6,
        uintptr(hwnd),
        uintptr(x),
        uintptr(y),
        uintptr(width),
        uintptr(height),
        uintptr(BoolToBOOL(repaint)))

    return ret != 0

}

func ScreenToClient(hwnd HWND, x, y int) (int, int) {
    var pt POINT
    pt.X = x
    pt.Y = y

    syscall.Syscall(procScreenToClient, 2,
        uintptr(hwnd),
        uintptr(unsafe.Pointer(&pt)),
        0)

    return pt.X, pt.Y
}

func CallWindowProc(preWndProc uintptr, hwnd HWND, msg uint, wParam, lParam uintptr) uintptr {
    ret, _, _ := syscall.Syscall6(procCallWindowProc, 5,
        preWndProc,
        uintptr(hwnd),
        uintptr(msg),
        wParam,
        lParam,
        0)

    return ret
}

func SetWindowLongPtr(hwnd HWND, index int, value uintptr) uintptr {
    ret, _, _ := syscall.Syscall(procSetWindowLongPtr, 3,
        uintptr(hwnd),
        uintptr(index),
        value)

    return ret
}

func EnableWindow(hwnd HWND, b bool) bool {
    ret, _, _ := syscall.Syscall(procEnableWindow, 2,
        uintptr(hwnd),
        uintptr(BoolToBOOL(b)),
        0)
    return ret != 0
}

func IsWindowEnabled(hwnd HWND) bool {
    ret, _, _ := syscall.Syscall(procIsWindowEnabled, 1,
        uintptr(hwnd),
        0,
        0)

    return ret != 0
}

func IsWindowVisible(hwnd HWND) bool {
    ret, _, _ := syscall.Syscall(procIsWindowVisible, 1,
        uintptr(hwnd),
        0,
        0)

    return ret != 0
}

func SetFocus(hwnd HWND) HWND {
    ret, _, _ := syscall.Syscall(procSetFocus, 1,
        uintptr(hwnd),
        0,
        0)

    return HWND(ret)
}

func InvalidateRect(hwnd HWND, rect *RECT, erase bool) bool {
    ret, _, _ := syscall.Syscall(procInvalidateRect, 3,
        uintptr(hwnd),
        uintptr(unsafe.Pointer(rect)),
        uintptr(BoolToBOOL(erase)))

    return ret != 0
}

func GetClientRect(hwnd HWND) *RECT {
    var rect RECT
    ret, _, _ := syscall.Syscall(procGetWindowRect, 2,
        uintptr(hwnd),
        uintptr(unsafe.Pointer(&rect)),
        0)

    if ret == 0 {
        panic(fmt.Sprintf("GetClientRect(%d) failed", hwnd))
    }

    return &rect
}

func GetDC(hwnd HWND) HDC {
    ret, _, _ := syscall.Syscall(procGetDC, 1,
        uintptr(hwnd),
        0,
        0)

    return HDC(ret)
}

func ReleaseDC(hwnd HWND, hDC HDC) bool {
    ret, _, _ := syscall.Syscall(procReleaseDC, 2,
        uintptr(hwnd),
        uintptr(hDC),
        0)

    return ret != 0
}

func SetCapture(hwnd HWND) HWND {
    ret, _, _ := syscall.Syscall(procSetCapture, 1,
        uintptr(hwnd),
        0,
        0)

    return HWND(ret)
}

func ReleaseCapture() bool {
    ret, _, _ := syscall.Syscall(procReleaseCapture, 0,
        0,
        0,
        0)

    return ret != 0
}

func GetWindowThreadProcessId(hwnd HWND) (HANDLE, int) {
    var processId int
    ret, _, _ := syscall.Syscall(procGetWindowThreadProcessId, 2,
        uintptr(hwnd),
        uintptr(unsafe.Pointer(&processId)),
        0)

    return HANDLE(ret), processId
}

func MessageBox(hwnd HWND, title, caption string, flags uint) int {
    ret, _, _ := syscall.Syscall6(procMessageBox, 4,
        uintptr(hwnd),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
        uintptr(flags),
        0,
        0)

    return int(ret)
}

func GetSystemMetrics(index int) int {
    ret, _, _ := syscall.Syscall(procGetSystemMetrics, 1,
        uintptr(index),
        0,
        0)

    return int(ret)
}

func GetWindowLongPtr(hwnd HWND, index int) uintptr {
    ret, _, _ := syscall.Syscall(procGetWindowLongPtr, 2,
        uintptr(hwnd),
        uintptr(index),
        0)

    return ret
}

func CopyRect(dst, src *RECT) bool {
    ret, _, _ := syscall.Syscall(procCopyRect, 2,
        uintptr(unsafe.Pointer(dst)),
        uintptr(unsafe.Pointer(src)),
        0)

    return ret != 0
}

func EqualRect(rect1, rect2 *RECT) bool {
    ret, _, _ := syscall.Syscall(procEqualRect, 2,
        uintptr(unsafe.Pointer(rect1)),
        uintptr(unsafe.Pointer(rect2)),
        0)

    return ret != 0
}

func InflateRect(rect *RECT, dx, dy int) bool {
    ret, _, _ := syscall.Syscall(procInflateRect, 3,
        uintptr(unsafe.Pointer(rect)),
        uintptr(dx),
        uintptr(dy))

    return ret != 0
}

func IntersectRect(dst, src1, src2 *RECT) bool {
    ret, _, _ := syscall.Syscall(procIntersectRect, 3,
        uintptr(unsafe.Pointer(dst)),
        uintptr(unsafe.Pointer(src1)),
        uintptr(unsafe.Pointer(src2)))

    return ret != 0
}

func IsRectEmpty(rect *RECT) bool {
    ret, _, _ := syscall.Syscall(procIsRectEmpty, 1,
        uintptr(unsafe.Pointer(rect)),
        0,
        0)

    return ret != 0
}

func OffsetRect(rect *RECT, dx, dy int) bool {
    ret, _, _ := syscall.Syscall(procOffsetRect, 3,
        uintptr(unsafe.Pointer(rect)),
        uintptr(dx),
        uintptr(dy))

    return ret != 0
}

func PtInRect(rect *RECT, x, y int) bool {
    pt := POINT{X: x, Y: y}
    ret, _, _ := syscall.Syscall(procPtInRect, 2,
        uintptr(unsafe.Pointer(rect)),
        uintptr(unsafe.Pointer(&pt)),
        0)

    return ret != 0
}

func SetRect(rect *RECT, left, top, right, bottom int) bool {
    ret, _, _ := syscall.Syscall6(procSetRect, 5,
        uintptr(unsafe.Pointer(rect)),
        uintptr(left),
        uintptr(top),
        uintptr(right),
        uintptr(bottom),
        0)

    return ret != 0
}

func SetRectEmpty(rect *RECT) bool {
    ret, _, _ := syscall.Syscall(procSetRectEmpty, 1,
        uintptr(unsafe.Pointer(rect)),
        0,
        0)

    return ret != 0
}

func SubtractRect(dst, src1, src2 *RECT) bool {
    ret, _, _ := syscall.Syscall(procSubtractRect, 3,
        uintptr(unsafe.Pointer(dst)),
        uintptr(unsafe.Pointer(src1)),
        uintptr(unsafe.Pointer(src2)))

    return ret != 0
}

func UnionRect(dst, src1, src2 *RECT) bool {
    ret, _, _ := syscall.Syscall(procUnionRect, 3,
        uintptr(unsafe.Pointer(dst)),
        uintptr(unsafe.Pointer(src1)),
        uintptr(unsafe.Pointer(src2)))

    return ret != 0
}

func CreateDialog(hInstance HINSTANCE, lpTemplate *uint16, hWndParent HWND, lpDialogProc uintptr) HWND {
    ret, _, _ := syscall.Syscall6(procCreateDialogParam, 5,
        uintptr(hInstance),
        uintptr(unsafe.Pointer(lpTemplate)),
        uintptr(hWndParent),
        lpDialogProc,
        0,
        0)

    return HWND(ret)
}

func DialogBox(hInstance HINSTANCE, lpTemplateName *uint16, hWndParent HWND, lpDialogProc uintptr) int {
    ret, _, _ := syscall.Syscall6(procDialogBoxParam, 5,
        uintptr(hInstance),
        uintptr(unsafe.Pointer(lpTemplateName)),
        uintptr(hWndParent),
        lpDialogProc,
        0,
        0)

    return int(ret)
}

func GetDlgItem(hDlg HWND, nIDDlgItem int) HWND {
    ret, _, _ := syscall.Syscall(procGetDlgItem, 2,
        uintptr(unsafe.Pointer(hDlg)),
        uintptr(nIDDlgItem),
        0)

    return HWND(ret)
}

func DrawIcon(hDC HDC, x, y int, hIcon HICON) bool {
    ret, _, _ := syscall.Syscall6(procDrawIcon, 4,
        uintptr(unsafe.Pointer(hDC)),
        uintptr(x),
        uintptr(y),
        uintptr(unsafe.Pointer(hIcon)),
        0,
        0)

    return ret != 0
}

func ClientToScreen(hwnd HWND, x, y int) (int, int) {
    var pt POINT
    pt.X, pt.Y = x, y

    syscall.Syscall(procClientToScreen, 2,
        uintptr(hwnd),
        uintptr(unsafe.Pointer(&pt)),
        0)

    return pt.X, pt.Y
}

func IsDialogMessage(hwnd HWND, msg *MSG) bool {
    ret, _, _ := syscall.Syscall(procIsDialogMessage, 2,
        uintptr(hwnd),
        uintptr(unsafe.Pointer(msg)),
        0)

    return ret != 0
}

func IsWindow(hwnd HWND) bool {
    ret, _, _ := syscall.Syscall(procIsWindow, 1,
        uintptr(hwnd),
        0,
        0)

    return ret != 0
}

func EndDialog(hwnd HWND, nResult uintptr) bool {
    ret, _, _ := syscall.Syscall(procEndDialog, 2,
        uintptr(hwnd),
        nResult,
        0)

    return ret != 0
}

func SetWindowLong(hwnd HWND, nIndex int, dwNewLong uint32) uint32 {
    ret, _, _ := syscall.Syscall(procSetWindowLong, 3,
        uintptr(hwnd),
        uintptr(nIndex),
        uintptr(dwNewLong))

    return uint32(ret)
}

func PeekMessage(lpMsg *MSG, hwnd HWND, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) bool {
    ret, _, _ := syscall.Syscall6(procPeekMessage, 5,
        uintptr(unsafe.Pointer(lpMsg)),
        uintptr(hwnd),
        uintptr(wMsgFilterMin),
        uintptr(wMsgFilterMax),
        uintptr(wRemoveMsg),
        0)

    return ret != 0
}

func TranslateAccelerator(hwnd HWND, hAccTable HACCEL, lpMsg *MSG) bool {
    ret, _, _ := syscall.Syscall(procTranslateMessage, 3,
        uintptr(hwnd),
        uintptr(hAccTable),
        uintptr(unsafe.Pointer(lpMsg)))

    return ret != 0
}

func DestroyIcon(hIcon HICON) bool {
    ret, _, _ := syscall.Syscall(procDestroyIcon, 1,
        uintptr(hIcon),
        0,
        0)

    return ret != 0
}

func SetWindowPos(hwnd, hWndInsertAfter HWND, x, y, cx, cy int, uFlags uint) bool {
    ret, _, _ := syscall.Syscall9(procSetWindowPos, 7,
        uintptr(hwnd), 
        uintptr(hWndInsertAfter),
        uintptr(x),
        uintptr(y),
        uintptr(cx),
        uintptr(cy),
        uintptr(uFlags),
        0,
        0)
    
    return ret != 0
}

func FillRect(hDC HDC, lprc *RECT, hbr HBRUSH) bool {
    ret, _, _ := syscall.Syscall(procFillRect, 3, 
        uintptr(hDC), 
        uintptr(unsafe.Pointer(lprc)), 
        uintptr(hbr))

    return ret != 0
}

func DrawText(hDC HDC, text string, uCount int, lpRect *RECT, uFormat uint) int {
    ret, _, _ := syscall.Syscall6(procDrawText, 5, 
        uintptr(hDC), 
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
        uintptr(uCount),
        uintptr(unsafe.Pointer(lpRect)),
        uintptr(uFormat),
        0)

    return int(ret)
}
