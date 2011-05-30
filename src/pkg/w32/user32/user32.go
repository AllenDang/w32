// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package user32

import (
    "fmt"
    "syscall"
    "unsafe"
    . "w32"
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

func GetWindowRect(hwnd HWND) RECT {
    var rect RECT
    syscall.Syscall(procGetWindowRect, 2,
        uintptr(hwnd),
        uintptr(unsafe.Pointer(&rect)),
        0)

    return rect
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

func MessageBox(hwnd HWND, text, caption string, flags uint) int {
    ret, _, _ := syscall.Syscall6(procMessageBox, 4,
        uintptr(hwnd),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
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
