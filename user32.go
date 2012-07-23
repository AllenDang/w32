// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package w32

import (
    "fmt"
    "syscall"
    "unsafe"
)

var (
    moduser32 = syscall.NewLazyDLL("user32.dll")

    procRegisterClassEx               = moduser32.NewProc("RegisterClassExW")
    procLoadIcon                      = moduser32.NewProc("LoadIconW")
    procLoadCursor                    = moduser32.NewProc("LoadCursorW")
    procShowWindow                    = moduser32.NewProc("ShowWindow")
    procUpdateWindow                  = moduser32.NewProc("UpdateWindow")
    procCreateWindowEx                = moduser32.NewProc("CreateWindowExW")
    procAdjustWindowRect              = moduser32.NewProc("AdjustWindowRect")
    procAdjustWindowRectEx            = moduser32.NewProc("AdjustWindowRectEx")
    procDestroyWindow                 = moduser32.NewProc("DestroyWindow")
    procDefWindowProc                 = moduser32.NewProc("DefWindowProcW")
    procDefDlgProc                    = moduser32.NewProc("DefDlgProcW")
    procPostQuitMessage               = moduser32.NewProc("PostQuitMessage")
    procGetMessage                    = moduser32.NewProc("GetMessageW")
    procTranslateMessage              = moduser32.NewProc("TranslateMessage")
    procDispatchMessage               = moduser32.NewProc("DispatchMessageW")
    procSendMessage                   = moduser32.NewProc("SendMessageW")
    procPostMessage                   = moduser32.NewProc("PostMessageW")
    procSetWindowText                 = moduser32.NewProc("SetWindowTextW")
    procGetWindowTextLength           = moduser32.NewProc("GetWindowTextLengthW")
    procGetWindowText                 = moduser32.NewProc("GetWindowTextW")
    procGetWindowRect                 = moduser32.NewProc("GetWindowRect")
    procMoveWindow                    = moduser32.NewProc("MoveWindow")
    procScreenToClient                = moduser32.NewProc("ScreenToClient")
    procCallWindowProc                = moduser32.NewProc("CallWindowProcW")
    procSetWindowLongPtr              = moduser32.NewProc("SetWindowLongW")
    procEnableWindow                  = moduser32.NewProc("EnableWindow")
    procIsWindowEnabled               = moduser32.NewProc("IsWindowEnabled")
    procIsWindowVisible               = moduser32.NewProc("IsWindowVisible")
    procSetFocus                      = moduser32.NewProc("SetFocus")
    procInvalidateRect                = moduser32.NewProc("InvalidateRect")
    procGetClientRect                 = moduser32.NewProc("GetClientRect")
    procGetDC                         = moduser32.NewProc("GetDC")
    procReleaseDC                     = moduser32.NewProc("ReleaseDC")
    procSetCapture                    = moduser32.NewProc("SetCapture")
    procReleaseCapture                = moduser32.NewProc("ReleaseCapture")
    procGetWindowThreadProcessId      = moduser32.NewProc("GetWindowThreadProcessId")
    procMessageBox                    = moduser32.NewProc("MessageBoxW")
    procGetSystemMetrics              = moduser32.NewProc("GetSystemMetrics")
    procGetWindowLongPtr              = moduser32.NewProc("GetWindowLongW")
    procCopyRect                      = moduser32.NewProc("CopyRect")
    procEqualRect                     = moduser32.NewProc("EqualRect")
    procInflateRect                   = moduser32.NewProc("InflateRect")
    procIntersectRect                 = moduser32.NewProc("IntersectRect")
    procIsRectEmpty                   = moduser32.NewProc("IsRectEmpty")
    procOffsetRect                    = moduser32.NewProc("OffsetRect")
    procPtInRect                      = moduser32.NewProc("PtInRect")
    procSetRect                       = moduser32.NewProc("SetRect")
    procSetRectEmpty                  = moduser32.NewProc("SetRectEmpty")
    procSubtractRect                  = moduser32.NewProc("SubtractRect")
    procUnionRect                     = moduser32.NewProc("UnionRect")
    procCreateDialogParam             = moduser32.NewProc("CreateDialogParamW")
    procDialogBoxParam                = moduser32.NewProc("DialogBoxParamW")
    procGetDlgItem                    = moduser32.NewProc("GetDlgItem")
    procDrawIcon                      = moduser32.NewProc("DrawIcon")
    procClientToScreen                = moduser32.NewProc("ClientToScreen")
    procIsDialogMessage               = moduser32.NewProc("IsDialogMessageW")
    procIsWindow                      = moduser32.NewProc("IsWindow")
    procEndDialog                     = moduser32.NewProc("EndDialog")
    procSetWindowLong                 = moduser32.NewProc("SetWindowLongW")
    procPeekMessage                   = moduser32.NewProc("PeekMessageW")
    procTranslateAccelerator          = moduser32.NewProc("TranslateAcceleratorW")
    procDestroyIcon                   = moduser32.NewProc("DestroyIcon")
    procSetWindowPos                  = moduser32.NewProc("SetWindowPos")
    procFillRect                      = moduser32.NewProc("FillRect")
    procDrawText                      = moduser32.NewProc("DrawTextW")
    procAddClipboardFormatListener    = moduser32.NewProc("AddClipboardFormatListener")
    procRemoveClipboardFormatListener = moduser32.NewProc("RemoveClipboardFormatListener")
    procOpenClipboard                 = moduser32.NewProc("OpenClipboard")
    procCloseClipboard                = moduser32.NewProc("CloseClipboard")
    procEnumClipboardFormats          = moduser32.NewProc("EnumClipboardFormats")
    procGetClipboardData              = moduser32.NewProc("GetClipboardData")
    procSetClipboardData              = moduser32.NewProc("SetClipboardData")
    procEmptyClipboard                = moduser32.NewProc("EmptyClipboard")
    procGetClipboardFormatName        = moduser32.NewProc("GetClipboardFormatNameW")
    procIsClipboardFormatAvailable    = moduser32.NewProc("IsClipboardFormatAvailable")
    procBeginPaint                    = moduser32.NewProc("BeginPaint")
    procEndPaint                      = moduser32.NewProc("EndPaint")
    procGetKeyboardState              = moduser32.NewProc("GetKeyboardState")
    procMapVirtualKey                 = moduser32.NewProc("MapVirtualKeyExA")
    procToAscii                       = moduser32.NewProc("ToAscii")
)

func RegisterClassEx(wndClassEx *WNDCLASSEX) ATOM {
    ret, _, _ := procRegisterClassEx.Call(uintptr(unsafe.Pointer(wndClassEx)))
    return ATOM(ret)
}

func LoadIcon(instance HINSTANCE, iconName *uint16) HICON {
    ret, _, _ := procLoadIcon.Call(
        uintptr(instance),
        uintptr(unsafe.Pointer(iconName)))

    return HICON(ret)

}

func LoadCursor(instance HINSTANCE, cursorName *uint16) HCURSOR {
    ret, _, _ := procLoadCursor.Call(
        uintptr(instance),
        uintptr(unsafe.Pointer(cursorName)))

    return HCURSOR(ret)

}

func ShowWindow(hwnd HWND, cmdshow int) bool {
    ret, _, _ := procShowWindow.Call(
        uintptr(hwnd),
        uintptr(cmdshow))

    return ret != 0

}

func UpdateWindow(hwnd HWND) bool {
    ret, _, _ := procUpdateWindow.Call(
        uintptr(hwnd))
    return ret != 0
}

func CreateWindowEx(exStyle uint, className, windowName *uint16,
    style uint, x, y, width, height int, parent HWND, menu HMENU,
    instance HINSTANCE, param unsafe.Pointer) HWND {
    ret, _, _ := procCreateWindowEx.Call(
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

func AdjustWindowRectEx(rect *RECT, style uint, menu bool, exStyle uint) bool {
    ret, _, _ := procAdjustWindowRectEx.Call(
        uintptr(unsafe.Pointer(rect)),
        uintptr(style),
        uintptr(BoolToBOOL(menu)),
        uintptr(exStyle))

    return ret != 0 
}

func AdjustWindowRect(rect *RECT, style uint, menu bool) bool {
    ret, _, _ := procAdjustWindowRect.Call(
        uintptr(unsafe.Pointer(rect)),
        uintptr(style),
        uintptr(BoolToBOOL(menu)))

    return ret != 0 
}

func DestroyWindow(hwnd HWND) bool {
    ret, _, _ := procDestroyWindow.Call(
        uintptr(hwnd))

    return ret != 0
}

func DefWindowProc(hwnd HWND, msg uint, wParam, lParam uintptr) uintptr {
    ret, _, _ := procDefWindowProc.Call(
        uintptr(hwnd),
        uintptr(msg),
        wParam,
        lParam)

    return ret
}

func DefDlgProc(hwnd HWND, msg uint, wParam, lParam uintptr) uintptr {
    ret, _, _ := procDefDlgProc.Call(
        uintptr(hwnd),
        uintptr(msg),
        wParam,
        lParam)

    return ret
}

func PostQuitMessage(exitCode int) {
    procPostQuitMessage.Call(
        uintptr(exitCode))
}

func GetMessage(msg *MSG, hwnd HWND, msgFilterMin, msgFilterMax uint32) int {
    ret, _, _ := procGetMessage.Call(
        uintptr(unsafe.Pointer(msg)),
        uintptr(hwnd),
        uintptr(msgFilterMin),
        uintptr(msgFilterMax))

    return int(ret)
}

func TranslateMessage(msg *MSG) bool {
    ret, _, _ := procTranslateMessage.Call(
        uintptr(unsafe.Pointer(msg)))

    return ret != 0

}

func DispatchMessage(msg *MSG) uintptr {
    ret, _, _ := procDispatchMessage.Call(
        uintptr(unsafe.Pointer(msg)))

    return ret

}

func SendMessage(hwnd HWND, msg uint, wParam, lParam uintptr) uintptr {
    ret, _, _ := procSendMessage.Call(
        uintptr(hwnd),
        uintptr(msg),
        wParam,
        lParam)

    return ret
}

func PostMessage(hwnd HWND, msg uint, wParam, lParam uintptr) bool {
    ret, _, _ := procPostMessage.Call(
        uintptr(hwnd),
        uintptr(msg),
        wParam,
        lParam)

    return ret != 0
}

func SetWindowText(hwnd HWND, text string) {
    procSetWindowText.Call(
        uintptr(hwnd),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))))
}

func GetWindowTextLength(hwnd HWND) int {
    ret, _, _ := procGetWindowTextLength.Call(
        uintptr(hwnd))

    return int(ret)
}

func GetWindowText(hwnd HWND) string {
    textLen := GetWindowTextLength(hwnd) + 1

    buf := make([]uint16, textLen)
    procGetWindowText.Call(
        uintptr(hwnd),
        uintptr(unsafe.Pointer(&buf[0])),
        uintptr(textLen))

    return syscall.UTF16ToString(buf)
}

func GetWindowRect(hwnd HWND) *RECT {
    var rect RECT
    procGetWindowRect.Call(
        uintptr(hwnd),
        uintptr(unsafe.Pointer(&rect)))

    return &rect
}

func MoveWindow(hwnd HWND, x, y, width, height int, repaint bool) bool {
    ret, _, _ := procMoveWindow.Call(
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

    procScreenToClient.Call(
        uintptr(hwnd),
        uintptr(unsafe.Pointer(&pt)))

    return pt.X, pt.Y
}

func CallWindowProc(preWndProc uintptr, hwnd HWND, msg uint, wParam, lParam uintptr) uintptr {
    ret, _, _ := procCallWindowProc.Call(
        preWndProc,
        uintptr(hwnd),
        uintptr(msg),
        wParam,
        lParam)

    return ret
}

func SetWindowLongPtr(hwnd HWND, index int, value uintptr) uintptr {
    ret, _, _ := procSetWindowLongPtr.Call(
        uintptr(hwnd),
        uintptr(index),
        value)

    return ret
}

func EnableWindow(hwnd HWND, b bool) bool {
    ret, _, _ := procEnableWindow.Call(
        uintptr(hwnd),
        uintptr(BoolToBOOL(b)))
    return ret != 0
}

func IsWindowEnabled(hwnd HWND) bool {
    ret, _, _ := procIsWindowEnabled.Call(
        uintptr(hwnd))

    return ret != 0
}

func IsWindowVisible(hwnd HWND) bool {
    ret, _, _ := procIsWindowVisible.Call(
        uintptr(hwnd))

    return ret != 0
}

func SetFocus(hwnd HWND) HWND {
    ret, _, _ := procSetFocus.Call(
        uintptr(hwnd))

    return HWND(ret)
}

func InvalidateRect(hwnd HWND, rect *RECT, erase bool) bool {
    ret, _, _ := procInvalidateRect.Call(
        uintptr(hwnd),
        uintptr(unsafe.Pointer(rect)),
        uintptr(BoolToBOOL(erase)))

    return ret != 0
}

func GetClientRect(hwnd HWND) *RECT {
    var rect RECT
    ret, _, _ := procGetClientRect.Call(
        uintptr(hwnd),
        uintptr(unsafe.Pointer(&rect)))

    if ret == 0 {
        panic(fmt.Sprintf("GetClientRect(%d) failed", hwnd))
    }

    return &rect
}

func GetDC(hwnd HWND) HDC {
    ret, _, _ := procGetDC.Call(
        uintptr(hwnd))

    return HDC(ret)
}

func ReleaseDC(hwnd HWND, hDC HDC) bool {
    ret, _, _ := procReleaseDC.Call(
        uintptr(hwnd),
        uintptr(hDC))

    return ret != 0
}

func SetCapture(hwnd HWND) HWND {
    ret, _, _ := procSetCapture.Call(
        uintptr(hwnd))

    return HWND(ret)
}

func ReleaseCapture() bool {
    ret, _, _ := procReleaseCapture.Call()

    return ret != 0
}

func GetWindowThreadProcessId(hwnd HWND) (HANDLE, int) {
    var processId int
    ret, _, _ := procGetWindowThreadProcessId.Call(
        uintptr(hwnd),
        uintptr(unsafe.Pointer(&processId)))

    return HANDLE(ret), processId
}

func MessageBox(hwnd HWND, title, caption string, flags uint) int {
    ret, _, _ := procMessageBox.Call(
        uintptr(hwnd),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
        uintptr(flags))

    return int(ret)
}

func GetSystemMetrics(index int) int {
    ret, _, _ := procGetSystemMetrics.Call(
        uintptr(index))

    return int(ret)
}

func GetWindowLongPtr(hwnd HWND, index int) uintptr {
    ret, _, _ := procGetWindowLongPtr.Call(
        uintptr(hwnd),
        uintptr(index))

    return ret
}

func CopyRect(dst, src *RECT) bool {
    ret, _, _ := procCopyRect.Call(
        uintptr(unsafe.Pointer(dst)),
        uintptr(unsafe.Pointer(src)))

    return ret != 0
}

func EqualRect(rect1, rect2 *RECT) bool {
    ret, _, _ := procEqualRect.Call(
        uintptr(unsafe.Pointer(rect1)),
        uintptr(unsafe.Pointer(rect2)))

    return ret != 0
}

func InflateRect(rect *RECT, dx, dy int) bool {
    ret, _, _ := procInflateRect.Call(
        uintptr(unsafe.Pointer(rect)),
        uintptr(dx),
        uintptr(dy))

    return ret != 0
}

func IntersectRect(dst, src1, src2 *RECT) bool {
    ret, _, _ := procIntersectRect.Call(
        uintptr(unsafe.Pointer(dst)),
        uintptr(unsafe.Pointer(src1)),
        uintptr(unsafe.Pointer(src2)))

    return ret != 0
}

func IsRectEmpty(rect *RECT) bool {
    ret, _, _ := procIsRectEmpty.Call(
        uintptr(unsafe.Pointer(rect)))

    return ret != 0
}

func OffsetRect(rect *RECT, dx, dy int) bool {
    ret, _, _ := procOffsetRect.Call(
        uintptr(unsafe.Pointer(rect)),
        uintptr(dx),
        uintptr(dy))

    return ret != 0
}

func PtInRect(rect *RECT, x, y int) bool {
    pt := POINT{X: x, Y: y}
    ret, _, _ := procPtInRect.Call(
        uintptr(unsafe.Pointer(rect)),
        uintptr(unsafe.Pointer(&pt)))

    return ret != 0
}

func SetRect(rect *RECT, left, top, right, bottom int) bool {
    ret, _, _ := procSetRect.Call(
        uintptr(unsafe.Pointer(rect)),
        uintptr(left),
        uintptr(top),
        uintptr(right),
        uintptr(bottom))

    return ret != 0
}

func SetRectEmpty(rect *RECT) bool {
    ret, _, _ := procSetRectEmpty.Call(
        uintptr(unsafe.Pointer(rect)))

    return ret != 0
}

func SubtractRect(dst, src1, src2 *RECT) bool {
    ret, _, _ := procSubtractRect.Call(
        uintptr(unsafe.Pointer(dst)),
        uintptr(unsafe.Pointer(src1)),
        uintptr(unsafe.Pointer(src2)))

    return ret != 0
}

func UnionRect(dst, src1, src2 *RECT) bool {
    ret, _, _ := procUnionRect.Call(
        uintptr(unsafe.Pointer(dst)),
        uintptr(unsafe.Pointer(src1)),
        uintptr(unsafe.Pointer(src2)))

    return ret != 0
}

func CreateDialog(hInstance HINSTANCE, lpTemplate *uint16, hWndParent HWND, lpDialogProc uintptr) HWND {
    ret, _, _ := procCreateDialogParam.Call(
        uintptr(hInstance),
        uintptr(unsafe.Pointer(lpTemplate)),
        uintptr(hWndParent),
        lpDialogProc,
        0)

    return HWND(ret)
}

func DialogBox(hInstance HINSTANCE, lpTemplateName *uint16, hWndParent HWND, lpDialogProc uintptr) int {
    ret, _, _ := procDialogBoxParam.Call(
        uintptr(hInstance),
        uintptr(unsafe.Pointer(lpTemplateName)),
        uintptr(hWndParent),
        lpDialogProc,
        0)

    return int(ret)
}

func GetDlgItem(hDlg HWND, nIDDlgItem int) HWND {
    ret, _, _ := procGetDlgItem.Call(
        uintptr(unsafe.Pointer(hDlg)),
        uintptr(nIDDlgItem))

    return HWND(ret)
}

func DrawIcon(hDC HDC, x, y int, hIcon HICON) bool {
    ret, _, _ := procDrawIcon.Call(
        uintptr(unsafe.Pointer(hDC)),
        uintptr(x),
        uintptr(y),
        uintptr(unsafe.Pointer(hIcon)))

    return ret != 0
}

func ClientToScreen(hwnd HWND, x, y int) (int, int) {
    var pt POINT
    pt.X, pt.Y = x, y

    procClientToScreen.Call(
        uintptr(hwnd),
        uintptr(unsafe.Pointer(&pt)))

    return pt.X, pt.Y
}

func IsDialogMessage(hwnd HWND, msg *MSG) bool {
    ret, _, _ := procIsDialogMessage.Call(
        uintptr(hwnd),
        uintptr(unsafe.Pointer(msg)))

    return ret != 0
}

func IsWindow(hwnd HWND) bool {
    ret, _, _ := procIsWindow.Call(
        uintptr(hwnd))

    return ret != 0
}

func EndDialog(hwnd HWND, nResult uintptr) bool {
    ret, _, _ := procEndDialog.Call(
        uintptr(hwnd),
        nResult)

    return ret != 0
}

func SetWindowLong(hwnd HWND, nIndex int, dwNewLong uint32) uint32 {
    ret, _, _ := procSetWindowLong.Call(
        uintptr(hwnd),
        uintptr(nIndex),
        uintptr(dwNewLong))

    return uint32(ret)
}

func PeekMessage(lpMsg *MSG, hwnd HWND, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) bool {
    ret, _, _ := procPeekMessage.Call(
        uintptr(unsafe.Pointer(lpMsg)),
        uintptr(hwnd),
        uintptr(wMsgFilterMin),
        uintptr(wMsgFilterMax),
        uintptr(wRemoveMsg))

    return ret != 0
}

func TranslateAccelerator(hwnd HWND, hAccTable HACCEL, lpMsg *MSG) bool {
    ret, _, _ := procTranslateMessage.Call(
        uintptr(hwnd),
        uintptr(hAccTable),
        uintptr(unsafe.Pointer(lpMsg)))

    return ret != 0
}

func DestroyIcon(hIcon HICON) bool {
    ret, _, _ := procDestroyIcon.Call(
        uintptr(hIcon))

    return ret != 0
}

func SetWindowPos(hwnd, hWndInsertAfter HWND, x, y, cx, cy int, uFlags uint) bool {
    ret, _, _ := procSetWindowPos.Call(
        uintptr(hwnd),
        uintptr(hWndInsertAfter),
        uintptr(x),
        uintptr(y),
        uintptr(cx),
        uintptr(cy),
        uintptr(uFlags))

    return ret != 0
}

func FillRect(hDC HDC, lprc *RECT, hbr HBRUSH) bool {
    ret, _, _ := procFillRect.Call(
        uintptr(hDC),
        uintptr(unsafe.Pointer(lprc)),
        uintptr(hbr))

    return ret != 0
}

func DrawText(hDC HDC, text string, uCount int, lpRect *RECT, uFormat uint) int {
    ret, _, _ := procDrawText.Call(
        uintptr(hDC),
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(text))),
        uintptr(uCount),
        uintptr(unsafe.Pointer(lpRect)),
        uintptr(uFormat))

    return int(ret)
}

func AddClipboardFormatListener(hwnd HWND) bool {
    ret, _, _ := procAddClipboardFormatListener.Call(
        uintptr(hwnd))
    return ret != 0
}

func RemoveClipboardFormatListener(hwnd HWND) bool {
    ret, _, _ := procRemoveClipboardFormatListener.Call(
        uintptr(hwnd))
    return ret != 0
}

func OpenClipboard(hWndNewOwner HWND) bool {
    ret, _, _ := procOpenClipboard.Call(
        uintptr(hWndNewOwner))
    return ret != 0
}

func CloseClipboard() bool {
    ret, _, _ := procCloseClipboard.Call()
    return ret != 0
}

func EnumClipboardFormats(format uint) uint {
    ret, _, _ := procEnumClipboardFormats.Call(
        uintptr(format))
    return uint(ret)
}

func GetClipboardData(uFormat uint) HANDLE {
    ret, _, _ := procGetClipboardData.Call(
        uintptr(uFormat))
    return HANDLE(ret)
}

func SetClipboardData(uFormat uint, hMem HANDLE) HANDLE {
    ret, _, _ := procSetClipboardData.Call(
        uintptr(uFormat),
        uintptr(hMem))
    return HANDLE(ret)
}

func EmptyClipboard() bool {
    ret, _, _ := procEmptyClipboard.Call()
    return ret != 0
}

func GetClipboardFormatName(format uint) (string, bool) {
    cchMaxCount := 255
    buf := make([]uint16, cchMaxCount)
    ret, _, _ := procGetClipboardFormatName.Call(
        uintptr(format),
        uintptr(unsafe.Pointer(&buf[0])),
        uintptr(cchMaxCount))

    if ret > 0 {
        return syscall.UTF16ToString(buf), true
    }

    return "Requested format does not exist or is predefined", false
}

func IsClipboardFormatAvailable(format uint) bool {
    ret, _, _ := procIsClipboardFormatAvailable.Call(uintptr(format))
    return ret != 0
}

func BeginPaint(hwnd HWND, paint *PAINTSTRUCT) HDC {
    ret, _, _ := procBeginPaint.Call(
        uintptr(hwnd),
        uintptr(unsafe.Pointer(paint)))
    return HDC(ret)
}

func EndPaint(hwnd HWND, paint *PAINTSTRUCT) {
    procBeginPaint.Call(
        uintptr(hwnd),
        uintptr(unsafe.Pointer(paint)))
}

func GetKeyboardState(lpKeyState *[]byte) bool {
    ret, _, _ := procGetKeyboardState.Call(
        uintptr(unsafe.Pointer(&(*lpKeyState)[0])))
    return ret != 0
}

func MapVirtualKeyEx(uCode, uMapType uint, dwhkl HANDLE) uint {
    ret, _, _ := procMapVirtualKey.Call(
        uintptr(uCode),
        uintptr(uMapType),
        uintptr(dwhkl))
    return uint(ret)
}

func ToAscii(uVirtKey, uScanCode uint, lpKeyState *byte, lpChar *uint16, uFlags uint) int {
    ret, _, _ := procToAscii.Call(
        uintptr(uVirtKey),
        uintptr(uScanCode),
        uintptr(unsafe.Pointer(lpKeyState)),
        uintptr(unsafe.Pointer(lpChar)),
        uintptr(uFlags))
    return int(ret)
}
