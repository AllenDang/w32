package w32

import (
	"fmt"
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

// TODO Make sure all parameter names are non-Hungarian notation.

var msftedit syscall.Handle

// Init is required to use the rich edit control. Call Close when at the end of
// your application.
func Init() {
	msftedit, _ = syscall.LoadLibrary("msftedit.dll")
}

// Close must be called at the end of your program if you have called Init.
func Close() {
	syscall.FreeLibrary(msftedit)
}

var (
	user32   = syscall.NewLazyDLL("user32.dll")
	kernel32 = syscall.NewLazyDLL("kernel32.dll")
	ntdll    = syscall.NewLazyDLL("ntdll.dll")
	gdi32    = syscall.NewLazyDLL("gdi32.dll")
	comctl32 = syscall.NewLazyDLL("comctl32.dll")
	dwmapi   = syscall.NewLazyDLL("dwmapi.dll")
	winmm    = syscall.NewLazyDLL("winmm.dll")

	registerClassEx               = user32.NewProc("RegisterClassExW")
	unregisterClass               = user32.NewProc("UnregisterClassW")
	defWindowProc                 = user32.NewProc("DefWindowProcW")
	postQuitMessage               = user32.NewProc("PostQuitMessage")
	createWindowEx                = user32.NewProc("CreateWindowExW")
	destroyWindow                 = user32.NewProc("DestroyWindow")
	getMessage                    = user32.NewProc("GetMessageW")
	peekMessage                   = user32.NewProc("PeekMessageW")
	translateAccelerator          = user32.NewProc("TranslateAcceleratorW")
	translateMessage              = user32.NewProc("TranslateMessage")
	dispatchMessage               = user32.NewProc("DispatchMessageW")
	loadCursor                    = user32.NewProc("LoadCursorW")
	getSysColorBrush              = user32.NewProc("GetSysColorBrush")
	loadIcon                      = user32.NewProc("LoadIconW")
	loadImage                     = user32.NewProc("LoadImageW")
	createIcon                    = user32.NewProc("CreateIcon")
	createIconIndirect            = user32.NewProc("CreateIconIndirect")
	createIconFromResource        = user32.NewProc("CreateIconFromResource")
	createIconFromResourceEx      = user32.NewProc("CreateIconFromResourceEx")
	destroyIcon                   = user32.NewProc("DestroyIcon")
	lookupIconIdFromDirectoryEx   = user32.NewProc("LookupIconIdFromDirectoryEx")
	beginPaint                    = user32.NewProc("BeginPaint")
	endPaint                      = user32.NewProc("EndPaint")
	fillRect                      = user32.NewProc("FillRect")
	getClientRect                 = user32.NewProc("GetClientRect")
	sendMessage                   = user32.NewProc("SendMessageW")
	messageBox                    = user32.NewProc("MessageBoxW")
	setCursorPos                  = user32.NewProc("SetCursorPos")
	getCursorPos                  = user32.NewProc("GetCursorPos")
	sendInput                     = user32.NewProc("SendInput")
	openClipboard                 = user32.NewProc("OpenClipboard")
	closeClipboard                = user32.NewProc("CloseClipboard")
	getClipboardData              = user32.NewProc("GetClipboardData")
	setClipboardData              = user32.NewProc("SetClipboardData")
	emptyClipboard                = user32.NewProc("EmptyClipboard")
	addClipboardFormatListener    = user32.NewProc("AddClipboardFormatListener")
	removeClipboardFormatListener = user32.NewProc("RemoveClipboardFormatListener")
	setWindowsHookEx              = user32.NewProc("SetWindowsHookExW")
	unhookWindowsHookEx           = user32.NewProc("UnhookWindowsHookEx")
	callNextHookEx                = user32.NewProc("CallNextHookEx")
	getForegroundWindow           = user32.NewProc("GetForegroundWindow")
	setForegroundWindow           = user32.NewProc("SetForegroundWindow")
	getDC                         = user32.NewProc("GetDC")
	releaseDC                     = user32.NewProc("ReleaseDC")
	mapVirtualKey                 = user32.NewProc("MapVirtualKeyW")
	mapVirtualKeyEx               = user32.NewProc("MapVirtualKeyExW")
	loadKeyboardLayout            = user32.NewProc("LoadKeyboardLayoutW")
	showWindow                    = user32.NewProc("ShowWindow")
	showWindowAsync               = user32.NewProc("ShowWindowAsync")
	clientToScreen                = user32.NewProc("ClientToScreen")
	screenToClient                = user32.NewProc("ScreenToClient")
	getWindowLong                 = user32.NewProc("GetWindowLongW")
	getWindowLongPtr              = user32.NewProc("GetWindowLongPtrW")
	setWindowLong                 = user32.NewProc("SetWindowLongW")
	setWindowLongPtr              = user32.NewProc("SetWindowLongPtrW")
	getMenu                       = user32.NewProc("GetMenu")
	adjustWindowRectEx            = user32.NewProc("AdjustWindowRectEx")
	setWindowPos                  = user32.NewProc("SetWindowPos")
	getWindowRect                 = user32.NewProc("GetWindowRect")
	getClassName                  = user32.NewProc("GetClassNameW")
	getWindowPlacement            = user32.NewProc("GetWindowPlacement")
	setWindowPlacement            = user32.NewProc("SetWindowPlacement")
	isWindowVisible               = user32.NewProc("IsWindowVisible")
	getWindowTextLength           = user32.NewProc("GetWindowTextLengthW")
	monitorFromPoint              = user32.NewProc("MonitorFromPoint")
	printWindow                   = user32.NewProc("PrintWindow")
	animateWindow                 = user32.NewProc("AnimateWindow")
	invalidateRect                = user32.NewProc("InvalidateRect")
	enumWindows                   = user32.NewProc("EnumWindows")
	enumDisplayMonitors           = user32.NewProc("EnumDisplayMonitors")
	getMonitorInfo                = user32.NewProc("GetMonitorInfoW")
	setFocus                      = user32.NewProc("SetFocus")
	messageBeep                   = user32.NewProc("MessageBeep")
	setWindowText                 = user32.NewProc("SetWindowTextW")
	getWindowText                 = user32.NewProc("GetWindowTextW")
	setTimer                      = user32.NewProc("SetTimer")
	killTimer                     = user32.NewProc("KillTimer")
	enableWindow                  = user32.NewProc("EnableWindow")
	createAcceleratorTable        = user32.NewProc("CreateAcceleratorTableW")
	destroyAcceleratorTable       = user32.NewProc("DestroyAcceleratorTable")
	getFocus                      = user32.NewProc("GetFocus")
	getWindowThreadProcessId      = user32.NewProc("GetWindowThreadProcessId")
	setScrollPos                  = user32.NewProc("SetScrollPos")
	getScrollPos                  = user32.NewProc("GetScrollPos")
	getSystemMetrics              = user32.NewProc("GetSystemMetrics")
	getSystemMetricsForDpi        = user32.NewProc("GetSystemMetricsForDpi")
	getAsyncKeyState              = user32.NewProc("GetAsyncKeyState")

	getModuleHandle     = kernel32.NewProc("GetModuleHandleW")
	globalAlloc         = kernel32.NewProc("GlobalAlloc")
	globalFree          = kernel32.NewProc("GlobalFree")
	globalLock          = kernel32.NewProc("GlobalLock")
	globalUnlock        = kernel32.NewProc("GlobalUnlock")
	formatMessage       = kernel32.NewProc("FormatMessageW")
	beep                = kernel32.NewProc("Beep")
	createActCtx        = kernel32.NewProc("CreateActCtxW")
	activateActCtx      = kernel32.NewProc("ActivateActCtx")
	deactivateActCtx    = kernel32.NewProc("DeactivateActCtx")
	getConsoleWindow    = kernel32.NewProc("GetConsoleWindow")
	getCurrentProcessId = kernel32.NewProc("GetCurrentProcessId")

	moveMemory = ntdll.NewProc("RtlMoveMemory")

	createSolidBrush       = gdi32.NewProc("CreateSolidBrush")
	deleteObject           = gdi32.NewProc("DeleteObject")
	selectObject           = gdi32.NewProc("SelectObject")
	deleteDC               = gdi32.NewProc("DeleteDC")
	createCompatibleDC     = gdi32.NewProc("CreateCompatibleDC")
	createCompatibleBitmap = gdi32.NewProc("CreateCompatibleBitmap")
	bitBlt                 = gdi32.NewProc("BitBlt")
	getDIBits              = gdi32.NewProc("GetDIBits")
	enumFontFamiliesEx     = gdi32.NewProc("EnumFontFamiliesExW")
	createFontIndirect     = gdi32.NewProc("CreateFontIndirectW")
	getTextExtentPoint32   = gdi32.NewProc("GetTextExtentPoint32W")

	initCommonControlsEx   = comctl32.NewProc("InitCommonControlsEx")
	imageList_Create       = comctl32.NewProc("ImageList_Create")
	imageList_Destroy      = comctl32.NewProc("ImageList_Destroy")
	imageList_Add          = comctl32.NewProc("ImageList_Add")
	imageList_AddMasked    = comctl32.NewProc("ImageList_AddMasked")
	imageList_ReplaceIcon  = comctl32.NewProc("ImageList_ReplaceIcon")
	imageList_GetIcon      = comctl32.NewProc("ImageList_GetIcon")
	imageList_Remove       = comctl32.NewProc("ImageList_Remove")
	imageList_Draw         = comctl32.NewProc("ImageList_Draw")
	imageList_DrawEx       = comctl32.NewProc("ImageList_DrawEx")
	imageList_DrawIndirect = comctl32.NewProc("ImageList_DrawIndirect")
	defSubclassProc        = comctl32.NewProc("DefSubclassProc")
	setWindowSubclass      = comctl32.NewProc("SetWindowSubclass")

	dwmGetWindowAttribute = dwmapi.NewProc("DwmGetWindowAttribute")
	dwmSetWindowAttribute = dwmapi.NewProc("DwmSetWindowAttribute")

	playSound = winmm.NewProc("PlaySoundW")
)

// String converts a Go string to a Windows compatible UTF-16 string.
func String(s string) UTF16String {
	return syscall.StringToUTF16Ptr(s)
}

// StringLengthInCharacters returns the number of UTF-16 characters in the
// given string. Some Windows functions need to be given this number in
// addition to the string itself.
func StringLengthInCharacters(s string) int {
	utf16, err := syscall.UTF16FromString(s)
	if err != nil {
		return 0
	}
	// The array ends in a trailing zero byte which is not part of the actual
	// string, thus we subtract 1.
	return len(utf16) - 1
}

// StringAtom converts the atom to the string type. Some functions will accept
// either a UTF-16 string or an ATOM for the same parameter. E.g.
// CreateWindowEx can be given a class name as a string or the atom returned by
// RegisterClassEx.
func StringAtom(atom ATOM) UTF16String {
	return UTF16String(unsafe.Pointer(atom))
}

// RGB creates a Windows compatible BGR color from the  given red, green and
// blue color channels.
func RGB(r, g, b byte) COLORREF {
	return COLORREF(r) | COLORREF(g)<<8 | COLORREF(b)<<16
}

// GetRValue returns the red channel of the given color.
func GetRValue(color COLORREF) byte {
	return byte(color & 0xFF)
}

// GetGValue returns the green channel of the given color.
func GetGValue(color COLORREF) byte {
	return byte((color & 0xFF00) >> 8)
}

// GetBValue returns the blue channel of the given color.
func GetBValue(color COLORREF) byte {
	return byte((color & 0xFF0000) >> 16)
}

// https://learn.microsoft.com/windows/win32/api/winnt/nf-winnt-makelangid
func MAKELANGID(language, subLanguage uint32) uint32 {
	return (subLanguage << 10) | language
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-makeintresourcew
func MakeIntResource(id int) UTF16String {
	return (UTF16String)(unsafe.Pointer(uintptr(id)))
}

// NewWindowProcedure converts the given function to a callback that can be
// used for WNDCLASSEX.WndProc.
func NewWindowProcedure(f func(window HWND, message uint32, wParam, lParam uintptr) uintptr) uintptr {
	return syscall.NewCallback(f)
}

// makeErr returns nil if isErr is false.
// If isErr is true, makeErr returns an error containing the message and
// optionally the lastError.
// lastError must be a syscall.Errno returned as the third return value from a
// call to syscall.LazyProc.Call.
// If lastError is ERROR_SUCCESS, i.e. 0, it is not included in the message.
// The scenario that a function fails and does not set the last error is
// possible, e.g. if passing an invalid path to LoadImage, it will set the last
// error to "file not found". But if we pass a valid PNG file when LoadImage
// when it expects a BMP file, it will return NULL (invalid image) but NOT set
// the last error.
func makeErr(isErr bool, message string, lastError error) error {
	if !isErr {
		return nil
	}
	if lastError == nil || lastError.(syscall.Errno) == 0 {
		return fmt.Errorf(message)
	}
	return fmt.Errorf(message+": %w", lastError)
}

// RegisterClassEx sets the Size of the WNDCLASSEX automatically.
//
// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-registerclassexw
func RegisterClassEx(wndClassEx *WNDCLASSEX) (ATOM, error) {
	if wndClassEx != nil {
		wndClassEx.Size = uint32(unsafe.Sizeof(*wndClassEx))
	}
	ret, _, err := registerClassEx.Call(uintptr(unsafe.Pointer(wndClassEx)))
	return ATOM(ret), makeErr(ret == 0, "w32.RegisterClassEx returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-unregisterclassw
func UnregisterClass(class UTF16String, instance HINSTANCE) error {
	ret, _, err := unregisterClass.Call(
		uintptr(unsafe.Pointer(class)),
		uintptr(instance),
	)
	return makeErr(ret == 0, "w32.UnregisterClass returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-defwindowprocw
func DefWindowProc(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := defWindowProc.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam,
	)
	return ret
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-postquitmessage
func PostQuitMessage(exitCode int) {
	postQuitMessage.Call(uintptr(exitCode))
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-createwindowexw
func CreateWindowEx(
	exStyle uint32,
	className UTF16String,
	windowName UTF16String,
	style uint32,
	x int,
	y int,
	width int,
	height int,
	parent HWND,
	menu HMENU,
	instance HINSTANCE,
	param unsafe.Pointer,
) (HWND, error) {
	ret, _, err := createWindowEx.Call(
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
		uintptr(param),
	)
	return HWND(ret), makeErr(ret == 0, "w32.CreateWindowEx returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-destroywindow
func DestroyWindow(window HWND) error {
	ret, _, err := destroyWindow.Call(uintptr(window))
	return makeErr(ret == 0, "w32.DestroyWindow returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getmessage
func GetMessage(msg *MSG, hwnd HWND, msgFilterMin, msgFilterMax uint32) (bool, error) {
	ret, _, err := getMessage.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax),
	)
	return ret != 0, makeErr(int(ret) == -1, "w32.getMessage returned -1", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-peekmessagew
func PeekMessage(msg *MSG, hwnd HWND, wMsgFilterMin, wMsgFilterMax, remove uint32) bool {
	ret, _, _ := peekMessage.Call(
		uintptr(unsafe.Pointer(msg)),
		uintptr(hwnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(remove),
	)
	return ret != 0
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-translateacceleratorw
func TranslateAccelerator(hwnd HWND, hAccTable HACCEL, lpMsg *MSG) error {
	ret, _, err := translateAccelerator.Call(
		uintptr(hwnd),
		uintptr(hAccTable),
		uintptr(unsafe.Pointer(lpMsg)),
	)
	return makeErr(ret == 0, "w32.TranslateAccelerator returned 0", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-translatemessage
func TranslateMessage(msg *MSG) bool {
	ret, _, _ := translateMessage.Call(uintptr(unsafe.Pointer(msg)))
	return ret != 0
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-dispatchmessage
func DispatchMessage(msg *MSG) uintptr {
	ret, _, _ := dispatchMessage.Call(uintptr(unsafe.Pointer(msg)))
	return ret

}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-loadcursorw
func LoadCursor(instance HINSTANCE, cursorName UTF16String) (HCURSOR, error) {
	ret, _, err := loadCursor.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(cursorName)),
	)
	return HCURSOR(ret), makeErr(ret == 0, "w32.LoadCursor returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getsyscolorbrush
func GetSysColorBrush(index int) (HBRUSH, error) {
	ret, _, _ := getSysColorBrush.Call(uintptr(index))
	return HBRUSH(ret), makeErr(ret == 0, "w32.GetSysColorBrush returned nil", nil)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-loadiconw
func LoadIcon(instance HINSTANCE, iconName UTF16String) (HICON, error) {
	ret, _, err := loadIcon.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(iconName)),
	)
	return HICON(ret), makeErr(ret == 0, "w32.LoadIcon returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-loadimagew
func LoadImage(
	instance HINSTANCE,
	name UTF16String,
	typ uint32,
	desiredWidth, desiredHeight int,
	load uint32,
) (HANDLE, error) {
	ret, _, err := loadImage.Call(
		uintptr(instance),
		uintptr(unsafe.Pointer(name)),
		uintptr(typ),
		uintptr(desiredWidth),
		uintptr(desiredHeight),
		uintptr(load),
	)
	return HANDLE(ret), makeErr(ret == 0, "w32.LoadImage returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-createicon
func CreateIcon(
	instance HINSTANCE,
	width int,
	height int,
	planes byte,
	bitsPerPixel byte,
	andBits,
	xorBits *byte,
) (HICON, error) {
	ret, _, err := createIcon.Call(
		uintptr(instance),
		uintptr(width),
		uintptr(height),
		uintptr(planes),
		uintptr(bitsPerPixel),
		uintptr(unsafe.Pointer(andBits)),
		uintptr(unsafe.Pointer(xorBits)),
	)
	return HICON(ret), makeErr(ret == 0, "w32.CreateIcon returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-createiconindirect
func CreateIconIndirect(info *ICONINFO) (HICON, error) {
	ret, _, err := createIconIndirect.Call(uintptr(unsafe.Pointer(info)))
	return HICON(ret), makeErr(ret == 0, "w32.CreateIconIndirect returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-createiconfromresource
func CreateIconFromResource(bits []byte, icon bool, version uint32) (HICON, error) {
	var ptr uintptr
	if len(bits) > 0 {
		ptr = uintptr(unsafe.Pointer(&bits[0]))
	}
	ret, _, err := createIconFromResource.Call(
		ptr,
		uintptr(len(bits)),
		uintptr(BoolToBOOL(icon)),
		uintptr(version),
	)
	return HICON(ret), makeErr(ret == 0, "w32.CreateIconFromResource returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-createiconfromresourceex
func CreateIconFromResourceEx(bits []byte, icon bool, version uint32, width, height int, flags uint32) (HICON, error) {
	var ptr uintptr
	if len(bits) > 0 {
		ptr = uintptr(unsafe.Pointer(&bits[0]))
	}
	ret, _, err := createIconFromResourceEx.Call(
		ptr,
		uintptr(len(bits)),
		uintptr(BoolToBOOL(icon)),
		uintptr(version),
		uintptr(width),
		uintptr(height),
		uintptr(flags),
	)
	return HICON(ret), makeErr(ret == 0, "w32.CreateIconFromResourceEx returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-destroyicon
func DestroyIcon(icon HICON) error {
	ret, _, err := destroyIcon.Call(uintptr(icon))
	return makeErr(ret == 0, "w32.DestroyIcon returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-lookupiconidfromdirectoryex
func LookupIconIdFromDirectoryEx(bits []byte, icon bool, width, height int, flags uint32) (int, error) {
	var ptr uintptr
	if len(bits) > 0 {
		ptr = uintptr(unsafe.Pointer(&bits[0]))
	}
	ret, _, err := lookupIconIdFromDirectoryEx.Call(
		ptr,
		uintptr(BoolToBOOL(icon)),
		uintptr(width),
		uintptr(height),
		uintptr(flags),
	)
	return int(ret), makeErr(ret == 0, "w32.LookupIconIdFromDirectoryEx returned 0", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-beginpaint
func BeginPaint(hwnd HWND, paint *PAINTSTRUCT) (HDC, error) {
	ret, _, _ := beginPaint.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(paint)),
	)
	return HDC(ret), makeErr(ret == 0, "w32.BeginPaint returned nil", nil)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-endpaint
func EndPaint(hwnd HWND, paint *PAINTSTRUCT) {
	endPaint.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(paint)),
	)
}

// https://learn.microsoft.com/windows/win32/api/libloaderapi/nf-libloaderapi-getmodulehandlew
func GetModuleHandle(moduleName UTF16String) (HINSTANCE, error) {
	ret, _, err := getModuleHandle.Call(uintptr(unsafe.Pointer(moduleName)))
	return HINSTANCE(ret), makeErr(ret == 0, "w32.GetModuleHandle returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winbase/nf-winbase-globalalloc
func GlobalAlloc(flags uint32, sizeInBytes uint) (HGLOBAL, error) {
	ret, _, err := globalAlloc.Call(
		uintptr(flags),
		uintptr(sizeInBytes),
	)
	return HGLOBAL(ret), makeErr(ret == 0, "w32.GlobalAlloc returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winbase/nf-winbase-globalfree
func GlobalFree(mem HGLOBAL) error {
	ret, _, err := globalFree.Call(uintptr(mem))
	// NOTE that this fails if the return value is NOT 0.
	return makeErr(ret != 0, "w32.GlobalFree failed", err)
}

// https://learn.microsoft.com/windows/win32/api/winbase/nf-winbase-globallock
func GlobalLock(mem HGLOBAL) (unsafe.Pointer, error) {
	ret, _, err := globalLock.Call(uintptr(mem))
	return unsafe.Pointer(ret), makeErr(ret == 0, "w32.GlobalLock returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winbase/nf-winbase-globalunlock
func GlobalUnlock(mem HGLOBAL) (int, error) {
	ret, _, err := globalUnlock.Call(uintptr(mem))
	return int(ret), makeErr(
		ret == 0 && err.(syscall.Errno) != 0,
		"w32.GlobalUnlock returned 0",
		err,
	)
}

// https://learn.microsoft.com/windows/win32/api/winbase/nf-winbase-formatmessagew
func FormatMessage(
	flags uint32,
	source unsafe.Pointer,
	messageID uint32,
	languageID uint32,
	buffer UTF16String,
	bufferSizeInCharacters uint32,
	Arguments unsafe.Pointer,
) (uint32, error) {
	ret, _, err := formatMessage.Call(
		uintptr(flags),
		uintptr(source),
		uintptr(messageID),
		uintptr(languageID),
		uintptr(unsafe.Pointer(buffer)),
		uintptr(bufferSizeInCharacters),
		uintptr(Arguments),
	)
	return uint32(ret), makeErr(ret == 0, "w32.FormatMessage returned nil", err)
}

// https://learn.microsoft.com/windows/win32/devnotes/rtlmovememory
func MoveMemory(destination, source unsafe.Pointer, length uint) {
	moveMemory.Call(
		uintptr(destination),
		uintptr(source),
		uintptr(length),
	)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-fillrect
func FillRect(hdc HDC, r *RECT, brush HBRUSH) error {
	ret, _, _ := fillRect.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(r)),
		uintptr(brush),
	)
	return makeErr(ret == 0, "w32.FillRect returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-createsolidbrush
func CreateSolidBrush(color COLORREF) (HBRUSH, error) {
	ret, _, _ := createSolidBrush.Call(uintptr(color))
	return HBRUSH(ret), makeErr(ret == 0, "w32.CreateSolidBrush returned nil", nil)
}

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-deleteobject
func DeleteObject(hObject HGDIOBJ) error {
	ret, _, _ := deleteObject.Call(uintptr(hObject))
	return makeErr(ret == 0, "w32.DeleteObject returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-selectobject
func SelectObject(hdc HDC, hgdiobj HGDIOBJ) HGDIOBJ {
	ret, _, _ := selectObject.Call(
		uintptr(hdc),
		uintptr(hgdiobj),
	)
	return HGDIOBJ(ret)
}

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-deletedc
func DeleteDC(hdc HDC) error {
	ret, _, _ := deleteDC.Call(uintptr(hdc))
	return makeErr(ret == 0, "w32.DeleteDC returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-createcompatibledc
func CreateCompatibleDC(hdc HDC) (HDC, error) {
	ret, _, _ := createCompatibleDC.Call(uintptr(hdc))
	return HDC(ret), makeErr(ret == 0, "w32.CreateCompatibleDC returned nil", nil)
}

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-createcompatiblebitmap
func CreateCompatibleBitmap(hdc HDC, width int, height int) (HBITMAP, error) {
	ret, _, _ := createCompatibleBitmap.Call(
		uintptr(hdc),
		uintptr(width),
		uintptr(height),
	)
	return HBITMAP(ret), makeErr(ret == 0, "w32.CreateCompatibleBitmap returned nil", nil)
}

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-bitblt
func BitBlt(hdc HDC, x, y, width, height int, hdcSource HDC, x1, y1 int, rop uint32) error {
	ret, _, err := bitBlt.Call(
		uintptr(hdc),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(hdcSource),
		uintptr(x1),
		uintptr(y1),
		uintptr(rop),
	)
	return makeErr(ret == 0, "w32.BitBlt returned FALSE", err)
}

// GetDIBits sets the BITMAPINFO's header size.
//
// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-getdibits
func GetDIBits(
	hdc HDC,
	bitmap HBITMAP,
	start uint32,
	lines uint32,
	bits unsafe.Pointer,
	bitmapInfo *BITMAPINFO,
	usage uint32,
) (int, error) {
	if bitmapInfo != nil {
		bitmapInfo.Header.Size = uint32(unsafe.Sizeof(bitmapInfo.Header))
	}
	ret, _, _ := getDIBits.Call(
		uintptr(hdc),
		uintptr(bitmap),
		uintptr(start),
		uintptr(lines),
		uintptr(bits),
		uintptr(unsafe.Pointer(bitmapInfo)),
		uintptr(usage),
	)
	return int(int32(ret)), makeErr(ret == 0, "w32.GetDIBits returned nil", nil)
}

// EnumFontFamiliesEx takes a function returned from NewFontFamilyExCallback as
// the proc parameter.
//
// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-enumfontfamiliesexw
func EnumFontFamiliesEx(hdc HDC, font *LOGFONT, proc, l uintptr, flags uint32) int {
	ret, _, _ := enumFontFamiliesEx.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(font)),
		uintptr(proc),
		uintptr(l),
		uintptr(flags),
	)
	return int(int32(ret))
}

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-createfontindirectw
func CreateFontIndirect(font *LOGFONT) (HFONT, error) {
	ret, _, _ := createFontIndirect.Call(uintptr(unsafe.Pointer(font)))
	return HFONT(ret), makeErr(ret == 0, "w32.CreateFontIndirect returned nil", nil)
}

func StringLength(s UTF16String) int {
	if s == nil {
		return 0
	}
	length := 0
	for *s != 0 {
		length++
		s = UTF16String(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + 2))
	}
	return length
}

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-gettextextentpoint32w
func GetTextExtentPoint32(hdc HDC, text UTF16String) (SIZE, error) {
	var size SIZE
	ret, _, _ := getTextExtentPoint32.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(text)),
		uintptr(StringLength(text)),
		uintptr(unsafe.Pointer(&size)),
	)
	return size, makeErr(ret == 0, "w32.GetTextExtentPoint32 returned FALSE", nil)
}

// NewFontFamilyExCallback returns a new callback that can be passed to
// EnumFontFamiliesEx.
//
// TODO Make all the function names for the New...Callback functions consistent
// with each other.
func NewFontFamilyExCallback(f func(
	font *ENUMLOGFONTEXDV,
	metric *ENUMTEXTMETRIC,
	fontType uint32,
	lParam uintptr,
) bool) uintptr {
	return syscall.NewCallback(func(
		font *ENUMLOGFONTEXDV,
		metric *ENUMTEXTMETRIC,
		fontType uint32,
		lParam uintptr,
	) uintptr {
		if f(font, metric, fontType, lParam) {
			return 1
		}
		return 0
	})
}

// f func(font *ENUMLOGFONTEX, metric *ENUMTEXTMETRIC, fontType FontType) bool)

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getclientrect
func GetClientRect(window HWND) (RECT, error) {
	var r RECT
	ret, _, err := getClientRect.Call(
		uintptr(window),
		uintptr(unsafe.Pointer(&r)),
	)
	return r, makeErr(ret == 0, "w32.GetClientRect returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-sendmessagew
func SendMessage(hwnd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	ret, _, _ := sendMessage.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam,
	)
	return ret
}

// InitCommonControlsEx sets the Size of the INITCOMMONCONTROLSEX
// automatically.
//
// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-initcommoncontrolsex
func InitCommonControlsEx(controls *INITCOMMONCONTROLSEX) error {
	if controls != nil {
		controls.Size = uint32(unsafe.Sizeof(*controls))
	}
	ret, _, _ := initCommonControlsEx.Call(uintptr(unsafe.Pointer(controls)))
	return makeErr(ret == 0, "w32.InitCommonControlsEx returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-indextooverlaymask
func INDEXTOOVERLAYMASK(index uint32) uint32 {
	return index << 8
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-indextostateimagemask
func INDEXTOSTATEIMAGEMASK(index uint32) uint32 {
	return index << 12
}

// TreeView_SetImageList returns the handle to the previous image list. If
// there was no image list, it returns 0.
//
// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_setimagelist
func TreeView_SetImageList(control HWND, imageList HIMAGELIST, imageType int) HIMAGELIST {
	return HIMAGELIST(SendMessage(
		control,
		TVM_SETIMAGELIST,
		uintptr(imageType),
		uintptr(imageList),
	))
}

// TreeView_GetImageList returns the image list associated with the control. It
// returns 0 if the control has no image list.
//
// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_getimagelist
func TreeView_GetImageList(control HWND, imageType int) HIMAGELIST {
	return HIMAGELIST(SendMessage(control, TVM_GETIMAGELIST, uintptr(imageType), 0))
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-imagelist_create
func ImageList_Create(width, height int, flags uint32, imageCount, grow int) (HIMAGELIST, error) {
	ret, _, _ := imageList_Create.Call(
		uintptr(width),
		uintptr(height),
		uintptr(flags),
		uintptr(imageCount),
		uintptr(grow),
	)
	return HIMAGELIST(ret), makeErr(ret == 0, "w32.ImageList_Create returned nil", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-imagelist_destroy
func ImageList_Destroy(list HIMAGELIST) error {
	ret, _, _ := imageList_Destroy.Call(uintptr(list))
	return makeErr(ret == 0, "w32.ImageList_Destroy returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-imagelist_add
func ImageList_Add(list HIMAGELIST, image, mask HBITMAP) (int, error) {
	ret, _, _ := imageList_Add.Call(
		uintptr(list),
		uintptr(image),
		uintptr(mask),
	)
	index := int32(ret)
	return int(index), makeErr(index == -1, "w32.ImageList_Add returned -1", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-imagelist_addmasked
func ImageList_AddMasked(list HIMAGELIST, image HBITMAP, mask COLORREF) (int, error) {
	ret, _, _ := imageList_AddMasked.Call(
		uintptr(list),
		uintptr(image),
		uintptr(mask),
	)
	index := int32(ret)
	return int(index), makeErr(index == -1, "w32.ImageList_AddMasked returned -1", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-imagelist_replaceicon
func ImageList_ReplaceIcon(list HIMAGELIST, atIndex int, icon HICON) (int, error) {
	ret, _, _ := imageList_ReplaceIcon.Call(
		uintptr(list),
		uintptr(atIndex),
		uintptr(icon),
	)
	index := int32(ret)
	return int(index), makeErr(index == -1, "w32.ImageList_ReplaceIcon returned -1", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-imagelist_addicon
func ImageList_AddIcon(list HIMAGELIST, icon HICON) (int, error) {
	i, err := ImageList_ReplaceIcon(list, -1, icon)
	if err != nil {
		return i, fmt.Errorf("w32.ImageList_AddIcon: %w", err)
	}
	return i, nil
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-imagelist_geticon
func ImageList_GetIcon(list HIMAGELIST, index int, flags uint32) (HICON, error) {
	ret, _, _ := imageList_GetIcon.Call(
		uintptr(list),
		uintptr(index),
		uintptr(flags),
	)
	return HICON(ret), makeErr(ret == 0, "w32.ImageList_GetIcon returned nil", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-imagelist_remove
func ImageList_Remove(list HIMAGELIST, index int) error {
	ret, _, _ := imageList_Remove.Call(
		uintptr(list),
		uintptr(index),
	)
	return makeErr(ret == 0, "w32.ImageList_Remove returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-imagelist_removeall
func ImageList_RemoveAll(list HIMAGELIST) error {
	err := ImageList_Remove(list, -1)
	if err != nil {
		return fmt.Errorf("w32.ImageList_RemoveAll: %w", err)
	}
	return nil
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-imagelist_draw
func ImageList_Draw(list HIMAGELIST, index int, dest HDC, x, y int, style uint32) error {
	ret, _, _ := imageList_Draw.Call(
		uintptr(list),
		uintptr(index),
		uintptr(dest),
		uintptr(x),
		uintptr(y),
		uintptr(style),
	)
	return makeErr(ret == 0, "w32.ImageList_Draw returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-imagelist_drawex
func ImageList_DrawEx(
	list HIMAGELIST,
	index int,
	dest HDC,
	x, y, dx, dy int,
	background, foreground COLORREF,
	style uint32,
) error {
	ret, _, _ := imageList_DrawEx.Call(
		uintptr(list),
		uintptr(index),
		uintptr(dest),
		uintptr(x),
		uintptr(y),
		uintptr(dx),
		uintptr(dy),
		uintptr(background),
		uintptr(foreground),
		uintptr(style),
	)
	return makeErr(ret == 0, "w32.ImageList_DrawEx returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-imagelist_drawindirect
func ImageList_DrawIndirect(params *IMAGELISTDRAWPARAMS) error {
	ret, _, _ := imageList_DrawIndirect.Call(uintptr(unsafe.Pointer(params)))
	return makeErr(ret == 0, "w32.ImageList_DrawIndirect returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-openclipboard
func OpenClipboard(window HWND) error {
	ret, _, err := openClipboard.Call(uintptr(window))
	return makeErr(ret == 0, "w32.OpenClipboard returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-closeclipboard
func CloseClipboard() error {
	ret, _, err := closeClipboard.Call()
	return makeErr(ret == 0, "w32.CloseClipboard returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getclipboarddata
func GetClipboardData(format uint32) (HANDLE, error) {
	ret, _, err := getClipboardData.Call(uintptr(format))
	return HANDLE(ret), makeErr(ret == 0, "w32.GetClipboardData returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-setclipboarddata
func SetClipboardData(format uint32, mem HANDLE) (HANDLE, error) {
	ret, _, err := setClipboardData.Call(
		uintptr(format),
		uintptr(mem),
	)
	return HANDLE(ret), makeErr(ret == 0, "w32.SetClipboardData returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-emptyclipboard
func EmptyClipboard() error {
	ret, _, err := emptyClipboard.Call()
	return makeErr(ret == 0, "w32.EmptyClipboard returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-addclipboardformatlistener
func AddClipboardFormatListener(window HWND) error {
	ret, _, err := addClipboardFormatListener.Call(uintptr(window))
	return makeErr(ret == 0, "w32.AddClipboardFormatListener returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-removeclipboardformatlistener
func RemoveClipboardFormatListener(winodw HWND) error {
	ret, _, err := removeClipboardFormatListener.Call(uintptr(winodw))
	return makeErr(ret == 0, "w32.RemoveClipboardFormatListener returned FALSE", err)
}

// NewHookProcedure converts the given function to a callback that can be used
// as the hookFunc in SetWindowsHookEx.
func NewHookProcedure(f func(code int32, wParam, lParam uintptr) uintptr) uintptr {
	return syscall.NewCallback(f)
}

// Use NewHookProcedure to create the hookFunc.
//
// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-setwindowshookexw
func SetWindowsHookEx(idHook int32, hookFunc uintptr, module HINSTANCE, threadID uint32) (HHOOK, error) {
	ret, _, err := setWindowsHookEx.Call(
		uintptr(idHook),
		hookFunc,
		uintptr(module),
		uintptr(threadID),
	)
	return HHOOK(ret), makeErr(ret == 0, "w32.SetWindowsHookEx returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-unhookwindowshookex
func UnhookWindowsHookEx(hook HHOOK) error {
	ret, _, err := unhookWindowsHookEx.Call(uintptr(hook))
	return makeErr(ret == 0, "w32.UnhookWindowsHookEx returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-callnexthookex
func CallNextHookEx(hook HHOOK, code int32, wParam uintptr, lParam uintptr) uintptr {
	ret, _, _ := callNextHookEx.Call(
		uintptr(hook),
		uintptr(code),
		wParam,
		lParam,
	)
	return ret
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getforegroundwindow
func GetForegroundWindow() HWND {
	ret, _, _ := getForegroundWindow.Call()
	return HWND(ret)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-setforegroundwindow
func SetForegroundWindow(hWnd HWND) error {
	ret, _, _ := setForegroundWindow.Call(uintptr(hWnd))
	return makeErr(ret == 0, "w32.SetForegroundWindow returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getdc
func GetDC(window HWND) (HDC, error) {
	ret, _, _ := getDC.Call(uintptr(window))
	return HDC(ret), makeErr(ret == 0, "w32.GetDC returned nil", nil)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-releasedc
func ReleaseDC(window HWND, hdc HDC) error {
	ret, _, _ := releaseDC.Call(
		uintptr(window),
		uintptr(hdc),
	)
	return makeErr(ret == 0, "w32.ReleaseDC returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-mapvirtualkeyw
func MapVirtualKey(code, mapType uint32) uint32 {
	ret, _, _ := mapVirtualKey.Call(
		uintptr(code),
		uintptr(mapType),
	)
	return uint32(ret)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-mapvirtualkeyexw
func MapVirtualKeyEx(code, mapType uint32, locale HKL) uint32 {
	ret, _, _ := mapVirtualKeyEx.Call(
		uintptr(code),
		uintptr(mapType),
		uintptr(locale),
	)
	return uint32(ret)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-loadkeyboardlayoutw
func LoadKeyboardLayout(localID UTF16String, flags uint32) (HKL, error) {
	ret, _, err := loadKeyboardLayout.Call(
		uintptr(unsafe.Pointer(localID)),
		uintptr(flags),
	)
	return HKL(ret), makeErr(ret == 0, "w32.LoadKeyboardLayout returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-showwindow
func ShowWindow(window HWND, showCommand int32) bool {
	ret, _, _ := showWindow.Call(
		uintptr(window),
		uintptr(showCommand),
	)
	return ret != 0
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-showwindowasync
func ShowWindowAsync(window HWND, showCommand int32) bool {
	ret, _, _ := showWindowAsync.Call(
		uintptr(window),
		uintptr(showCommand),
	)
	return ret != 0
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-clienttoscreen
func ClientToScreen(window HWND, p POINT) (POINT, error) {
	ret, _, _ := clientToScreen.Call(
		uintptr(window),
		uintptr(unsafe.Pointer(&p)),
	)
	return p, makeErr(ret == 0, "w32.ClientToScreen returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-screentoclient
func ScreenToClient(window HWND, p POINT) (POINT, error) {
	ret, _, _ := screenToClient.Call(
		uintptr(window),
		uintptr(unsafe.Pointer(&p)),
	)
	return p, makeErr(ret == 0, "w32.ScreenToClient returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getwindowlongw
func GetWindowLong(window HWND, index int32) (int32, error) {
	ret, _, err := getWindowLong.Call(
		uintptr(window),
		uintptr(index),
	)
	return int32(ret), makeErr(ret == 0, "w32.GetWindowLong returned 0", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
func GetWindowLongPtr(window HWND, index int32) (uintptr, error) {
	ret, _, err := getWindowLongPtr.Call(
		uintptr(window),
		uintptr(index),
	)
	return ret, makeErr(ret == 0, "w32.GetWindowLongPtr returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-setwindowlongw
func SetWindowLong(window HWND, index int32, newValue int) (int32, error) {
	ret, _, err := setWindowLong.Call(
		uintptr(window),
		uintptr(index),
		uintptr(newValue),
	)
	return int32(ret), makeErr(ret == 0, "w32.SetWindowLong returned 0", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-setwindowlongptrw
func SetWindowLongPtr(window HWND, index int32, newValue uintptr) (uintptr, error) {
	ret, _, err := setWindowLongPtr.Call(
		uintptr(window),
		uintptr(index),
		uintptr(newValue),
	)
	return ret, makeErr(ret == 0, "w32.SetWindowLongPtr returned nil", err)
}

// GetMenu returns the window's menu. It might be nil.
//
// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getmenu
func GetMenu(window HWND) HMENU {
	ret, _, _ := getMenu.Call(uintptr(window))
	return HMENU(ret)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-adjustwindowrectex
func AdjustWindowRectEx(
	r RECT,
	style uint32,
	hasMenu bool,
	exStyle uint32,
) (RECT, error) {
	ret, _, err := adjustWindowRectEx.Call(
		uintptr(unsafe.Pointer(&r)),
		uintptr(style),
		uintptr(BoolToBOOL(hasMenu)),
		uintptr(exStyle),
	)
	return r, makeErr(ret == 0, "w32.AdjustWindowRectEx returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-setwindowpos
func SetWindowPos(
	window,
	insertAfter HWND,
	x, y, width, height int32,
	flags uint32,
) error {
	ret, _, err := setWindowPos.Call(
		uintptr(window),
		uintptr(insertAfter),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(flags),
	)
	return makeErr(ret == 0, "w32.SetWindowPos returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getwindowrect
func GetWindowRect(window HWND) (RECT, error) {
	var r RECT
	ret, _, err := getWindowRect.Call(
		uintptr(window),
		uintptr(unsafe.Pointer(&r)),
	)
	return r, makeErr(ret == 0, "w32.GetWindowRect returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getclassnamew
func GetClassName(window HWND) (string, error) {
	var output [256]uint16
	ret, _, err := getClassName.Call(
		uintptr(window),
		uintptr(unsafe.Pointer(&output[0])),
		uintptr(len(output)),
	)
	return syscall.UTF16ToString(output[:]),
		makeErr(ret == 0, "w32.GetClassName returned 0", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getwindowplacement
func GetWindowPlacement(winodw HWND) (WINDOWPLACEMENT, error) {
	var placement WINDOWPLACEMENT
	placement.Length = uint32(unsafe.Sizeof(placement))
	ret, _, err := getWindowPlacement.Call(
		uintptr(winodw),
		uintptr(unsafe.Pointer(&placement)),
	)
	return placement, makeErr(ret == 0, "w32.GetWindowPlacement returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-setwindowplacement
func SetWindowPlacement(window HWND, placement WINDOWPLACEMENT) error {
	ret, _, err := setWindowPlacement.Call(
		uintptr(window),
		uintptr(unsafe.Pointer(&placement)),
	)
	return makeErr(ret == 0, "w32.SetWindowPlacement returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-iswindowvisible
func IsWindowVisible(window HWND) bool {
	ret, _, _ := isWindowVisible.Call(uintptr(window))
	return ret != 0
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-monitorfrompoint
func MonitorFromPoint(p POINT, flags uint32) HMONITOR {
	ret, _, _ := monitorFromPoint.Call(
		// TODO This seems to need 32 and 64 bit specific code.
		// uintptr(p),
		uintptr(flags),
	)
	return HMONITOR(ret)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-printwindow
func PrintWindow(window HWND, destination HDC, flags uint32) error {
	ret, _, _ := printWindow.Call(
		uintptr(window),
		uintptr(destination),
		uintptr(flags),
	)
	return makeErr(ret == 0, "w32.PrintWindow returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-animatewindow
func AnimateWindow(window HWND, time, flags uint32) error {
	ret, _, err := animateWindow.Call(
		uintptr(window),
		uintptr(time),
		uintptr(flags),
	)
	return makeErr(ret == 0, "w32.AnimateWindow returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-invalidaterect
func InvalidateRect(window HWND, r *RECT, erase bool) error {
	bErase := BoolToBOOL(erase)
	ret, _, _ := invalidateRect.Call(
		uintptr(window),
		uintptr(unsafe.Pointer(r)),
		uintptr(bErase),
	)
	return makeErr(ret == 0, "w32.InvalidateRect returned FALSE", nil)
}

// Use NewWindowProcedure to create the callback function.
//
// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-enumwindows
func EnumWindows(callback, lParam uintptr) error {
	ret, _, err := enumWindows.Call(callback, lParam)
	return makeErr(ret == 0, "w32.EnumWindows returned FALSE", err)
}

// NewEnumWindowsCallback is used in EnumWindows to create the callback.
func NewEnumWindowsCallback(callback func(window HWND, lParam uintptr) bool) uintptr {
	return syscall.NewCallback(func(w, l uintptr) uintptr {
		if callback(HWND(w), l) {
			return 1
		}
		return 0
	})
}

// Use NewEnumDisplayMonitorsCallback to create the callback function.
//
// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors
func EnumDisplayMonitors(hdc HDC, clip *RECT, lpfnEnum, data uintptr) error {
	ret, _, _ := enumDisplayMonitors.Call(
		uintptr(hdc),
		uintptr(unsafe.Pointer(clip)),
		lpfnEnum,
		data,
	)
	return makeErr(ret == 0, "w32.EnumDisplayMonitors returned FALSE", nil)
}

// NewEnumDisplayMonitorsCallback is used in EnumDisplayMonitors to create the
// callback.
func NewEnumDisplayMonitorsCallback(
	callback func(monitor HMONITOR, hdc HDC, bounds RECT, lParam uintptr) bool,
) uintptr {
	return syscall.NewCallback(
		func(monitor HMONITOR, hdc HDC, bounds *RECT, lParam uintptr) uintptr {
			var r RECT
			if bounds != nil {
				r = *bounds
			}
			if callback(monitor, hdc, r, lParam) {
				return 1
			}
			return 0
		},
	)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getmonitorinfow
func GetMonitorInfo(monitor HMONITOR) (MONITORINFOEX, error) {
	var info MONITORINFOEX
	info.Size = uint32(unsafe.Sizeof(info))
	ret, _, _ := getMonitorInfo.Call(
		uintptr(monitor),
		uintptr(unsafe.Pointer(&info)),
	)
	return info, makeErr(ret == 0, "w32.GetMonitorInfo returned FALSE", nil)
}

// SendInput sends mouse, keyboards and hardware input to the system.
//
// Use functions MouseInput, KeyboardInput and HardwareInput to create these
// INPUTs.
//
// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-sendinput
func SendInput(inputs ...INPUT) (uint32, error) {
	if len(inputs) == 0 {
		return 0, nil
	}
	ret, _, err := sendInput.Call(
		uintptr(len(inputs)),
		uintptr(unsafe.Pointer(&inputs[0])),
		unsafe.Sizeof(inputs[0]),
	)
	return uint32(ret), makeErr(ret == 0, "w32.SendInput returned 0", err)
}

// MouseInput converts a MOUSEINPUT to an INPUT for use in SendInput.
func MouseInput(input MOUSEINPUT) INPUT {
	return INPUT{
		Type:  INPUT_MOUSE,
		mouse: input,
	}
}

// KeyboardInput converts a KEYBDINPUT to an INPUT for use in SendInput.
func KeyboardInput(input KEYBDINPUT) INPUT {
	return INPUT{
		Type:  INPUT_KEYBOARD,
		mouse: *((*MOUSEINPUT)(unsafe.Pointer(&input))),
	}
}

// HardwareInput converts a HARDWAREINPUT to an INPUT for use in SendInput.
func HardwareInput(input HARDWAREINPUT) INPUT {
	return INPUT{
		Type:  INPUT_HARDWARE,
		mouse: *((*MOUSEINPUT)(unsafe.Pointer(&input))),
	}
}

func hresultToError(callingFunc string, result uintptr) error {
	if result == S_OK {
		return nil
	}

	var msgBuf [4096]uint16
	n, err := FormatMessage(
		FORMAT_MESSAGE_FROM_SYSTEM,
		nil,
		uint32(result),
		MAKELANGID(LANG_NEUTRAL, SUBLANG_DEFAULT),
		&msgBuf[0],
		uint32(len(msgBuf)),
		nil,
	)
	if err != nil {
		// FormatMessage failed so we just report the error code as a number.
		return fmt.Errorf(callingFunc+" returned error code %v", uint32(result))
	}

	// Make sure the UTF-16 is NULL-terminated.
	if n >= uint32(len(msgBuf)) {
		n = uint32(len(msgBuf)) - 1
	}
	msgBuf[n] = 0

	msg := syscall.UTF16ToString(msgBuf[:n+1])

	// Message will often end in a new line which is not what we want for Go
	// errors.
	msg = strings.TrimSuffix(msg, "\r\n")

	return fmt.Errorf(callingFunc + ": " + msg)
}

// DwmGetWindowAttribute uses a pointer to raw memory which is inconvenient.
// Use the specific variations, like DwmGetWindowAttributeBOOL or
// DwmGetWindowAttributeUint32, to get the right types back.
//
// https://learn.microsoft.com/windows/win32/api/dwmapi/nf-dwmapi-dwmgetwindowattribute
func DwmGetWindowAttribute(
	window HWND,
	attribute uint32,
	output unsafe.Pointer,
	outputSizeInBytes uint32,
) error {
	ret, _, _ := dwmGetWindowAttribute.Call(
		uintptr(window),
		uintptr(attribute),
		uintptr(output),
		uintptr(outputSizeInBytes),
	)
	return hresultToError("w32.DwmGetWindowAttribute", ret)
}

// See DwmGetWindowAttribute.
func DwmGetWindowAttributeBOOL(window HWND, attribute uint32) (bool, error) {
	var b BOOL
	err := DwmGetWindowAttribute(
		window,
		attribute,
		unsafe.Pointer(&b),
		uint32(unsafe.Sizeof(b)),
	)
	return b != 0, err
}

// See DwmGetWindowAttribute.
func DwmGetWindowAttributeRECT(window HWND, attribute uint32) (RECT, error) {
	var r RECT
	err := DwmGetWindowAttribute(
		window,
		attribute,
		unsafe.Pointer(&r),
		uint32(unsafe.Sizeof(r)),
	)
	return r, err
}

// See DwmGetWindowAttribute.
func DwmGetWindowAttributeUint32(window HWND, attribute uint32) (uint32, error) {
	var v uint32
	err := DwmGetWindowAttribute(
		window,
		attribute,
		unsafe.Pointer(&v),
		uint32(unsafe.Sizeof(v)),
	)
	return v, err
}

// DwmSetWindowAttribute takes an enumeration for the attribute and a value of
// one of a number of different types for the value. These values are
// abstracted with the Buffer type. Use functions BufferUint32, BufferBOOL and
// BufferCOLORREF.
//
// https://learn.microsoft.com/windows/win32/api/dwmapi/nf-dwmapi-dwmsetwindowattribute
func DwmSetWindowAttribute(
	window HWND,
	attribute uint32,
	buf Buffer,
) error {
	ret, _, _ := dwmSetWindowAttribute.Call(
		uintptr(window),
		uintptr(attribute),
		uintptr(buf.Memory),
		uintptr(buf.SizeInBytes),
	)
	return hresultToError("w32.DwmSetWindowAttribute", ret)
}

// Buffer represents an arbitrarily typed value. Memory points to the value and
// SizeInBytes is the value's size.
// This is used in DwmSetWindowAttribute.
type Buffer struct {
	Memory      unsafe.Pointer
	SizeInBytes uintptr
}

// See Buffer.
func BufferUint32(value uint32) Buffer {
	return Buffer{
		Memory:      unsafe.Pointer(&value),
		SizeInBytes: unsafe.Sizeof(value),
	}
}

// See Buffer.
func BufferBOOL(value bool) Buffer {
	b := BoolToBOOL(value)
	return Buffer{
		Memory:      unsafe.Pointer(&b),
		SizeInBytes: unsafe.Sizeof(b),
	}
}

// See Buffer.
func BufferCOLORREF(value COLORREF) Buffer {
	return Buffer{
		Memory:      unsafe.Pointer(&value),
		SizeInBytes: unsafe.Sizeof(value),
	}
}

// TODO Return the HRESULT as a wrapped error so the user can get to it from
// the returned error.

// TODO These were all double-checked for correctness and can be used as is:
// Once the above section is empty, this file is good to be checked in.

// BoolToBOOL converts the Go bool into a Windows BOOL, which is represented as
// a signed 32 bit integer.
func BoolToBOOL(b bool) BOOL {
	if b {
		return TRUE
	}
	return FALSE
}

// SetString fills the given slice dest with a NULL-terminated UTF-16 string
// created from s. If s is too long to place in dest, it is cut off at the end.
//
// Example:
//
//	var f w32.LOGFONT
//	w32.SetString(f.FaceName[:], "Tahoma")
func SetString(dest []uint16, s string) {
	if len(dest) == 0 {
		return
	}

	s16 := utf16.Encode([]rune(s))
	copy(dest, s16)

	// Add the NULL-termination.
	if len(s16) < len(dest) {
		dest[len(s16)] = 0
	} else {
		// s is too long, cut it off at the end.
		dest[len(dest)-1] = 0
	}
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-setfocus
func SetFocus(window HWND) (HWND, error) {
	ret, _, err := setFocus.Call(uintptr(window))
	return HWND(ret), makeErr(ret == 0, "w32.SetFocus returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_select
func TreeView_Select(tree HWND, item HTREEITEM, flag uint32) error {
	ret := SendMessage(tree, TVM_SELECTITEM, uintptr(flag), uintptr(item))
	return makeErr(ret == 0, "w32.TreeView_Select returned FALSE", nil)
}

// TODO It seems the documentation is false, the parent of the selected item is
// not expanded automatically. Confirm this and write an issue.
//
// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_selectdroptarget
func TreeView_SelectDropTarget(tree HWND, item HTREEITEM) error {
	ret := SendMessage(tree, TVM_SELECTITEM, TVGN_DROPHILITE, uintptr(item))
	return makeErr(ret == 0, "w32.TreeView_SelectDropTarget returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_selectitem
func TreeView_SelectItem(tree HWND, item HTREEITEM) error {
	ret := SendMessage(tree, TVM_SELECTITEM, TVGN_CARET, uintptr(item))
	return makeErr(ret == 0, "w32.TreeView_SelectItem returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_selectsetfirstvisible
func TreeView_SelectSetFirstVisible(tree HWND, item HTREEITEM) error {
	ret := SendMessage(tree, TVM_SELECTITEM, TVGN_FIRSTVISIBLE, uintptr(item))
	return makeErr(ret == 0, "w32.TreeView_SelectSetFirstVisible returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_setitem
func TreeView_SetItem(tree HWND, item *TVITEMEX) error {
	ret := SendMessage(tree, TVM_SETITEM, 0, uintptr(unsafe.Pointer(item)))
	return makeErr(ret == 0, "w32.TreeView_SetItem returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_showinfotip
func TreeView_ShowInfoTip(tree HWND, item HTREEITEM) {
	SendMessage(tree, TVM_SHOWINFOTIP, 0, uintptr(item))
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_deleteallitems
func TreeView_DeleteAllItems(tree HWND) error {
	ret := SendMessage(tree, TVM_DELETEITEM, 0, 0)
	return makeErr(ret == 0, "w32.TreeView_DeleteAllItems returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_deleteitem
func TreeView_DeleteItem(tree HWND, item HTREEITEM) error {
	ret := SendMessage(tree, TVM_DELETEITEM, 0, uintptr(item))
	return makeErr(ret == 0, "w32.TreeView_DeleteItem returned FALSE", nil)
}

// TreeView_Expand returns an error when collapsing an already collapsed node
// or when expanding an already expanded node. This is in contrast to what the
// documentation says.
//
// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_expand
func TreeView_Expand(tree HWND, item HTREEITEM, code uint32) error {
	ret := SendMessage(tree, TVM_EXPAND, uintptr(code), uintptr(item))
	return makeErr(ret == 0, "w32.TreeView_Expand returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_editlabel
func TreeView_EditLabel(tree HWND, item HTREEITEM) (HWND, error) {
	ret := SendMessage(tree, TVM_EDITLABEL, 0, uintptr(item))
	return HWND(ret), makeErr(ret == 0, "w32.TreeView_EditLabel returned nil", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_endeditlabelnow
func TreeView_EndEditLabelNow(tree HWND, cancel bool) error {
	ret := SendMessage(tree, TVM_ENDEDITLABELNOW, uintptr(BoolToBOOL(cancel)), 0)
	return makeErr(ret == 0, "w32.TreeView_EndEditLabelNow returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_ensurevisible
func TreeView_EnsureVisible(tree HWND, item HTREEITEM) bool {
	ret := SendMessage(tree, TVM_ENSUREVISIBLE, 0, uintptr(item))
	return ret != 0
}

// Intead of returning negaitve numbers like the documentation says, this
// version of TreeView_GetCount returns positive values up to 65535. If you add
// more than that, the node will not appear correctly.
//
// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_getcount
func TreeView_GetCount(tree HWND) int {
	ret := SendMessage(tree, TVM_GETCOUNT, 0, 0)
	return int(uint16(ret))
}

// TreeView_HitTest returns the tree item under the given point and TVHT_*
// flags describing the result.
//
// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_hittest
func TreeView_HitTest(tree HWND, p POINT) (HTREEITEM, uint32) {
	info := TVHITTESTINFO{Pt: p}
	SendMessage(tree, TVM_HITTEST, 0, uintptr(unsafe.Pointer(&info)))
	return info.Item, info.Flags
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_insertitem
func TreeView_InsertItem(tree HWND, item *TVINSERTSTRUCT) (HTREEITEM, error) {
	ret := SendMessage(tree, TVM_INSERTITEM, 0, uintptr(unsafe.Pointer(item)))
	return HTREEITEM(ret), makeErr(ret == 0, "w32.TreeView_InsertItem returned nil", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_sethot
func TreeView_SetHot(tree HWND, item HTREEITEM) error {
	ret := SendMessage(tree, TVM_SETHOT, 0, uintptr(item))
	return makeErr(ret == 0, "w32.TreeView_SetHot returned FALSE", nil)
}

// TreeView_SortChildren sorts the direct children of parent alphabetically.
//
// Note that the ignored argument, which should make sorting recurse down
// through all children, actually does nothing. See:
//
//	https://github.com/microsoft/win32metadata/issues/1823
//
// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_sortchildren
func TreeView_SortChildren(tree HWND, parent HTREEITEM, ignored bool) error {
	ret := SendMessage(
		tree,
		TVM_SORTCHILDREN,
		uintptr(BoolToBOOL(ignored)),
		uintptr(parent),
	)
	return makeErr(ret == 0, "w32.TreeView_SortChildren returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-treeview_sortchildrencb
func TreeView_SortChildrenCB(tree HWND, cb *TVSORTCB) error {
	ok := SendMessage(tree, TVM_SORTCHILDRENCB, 0, uintptr(unsafe.Pointer(cb)))
	if ok == 0 {
		return fmt.Errorf("w32.TreeView_SortChildrenCB returned FALSE")
	}
	return nil
}

// NewTreeSortCallback creates a new callback to be used in
// TreeView_SortChildrenCB. Use this as the Compare of the TVSORTCB.
//
// a and b are the LParam values of the tree items. You specify them when
// creating the items (set TVIF_PARAM and the LParam of the TVINSERTSTRUCT).
//
// sort is the LParam that you specify in the TVSORTCB.
func NewTreeSortCallback(f func(a, b, sort uintptr) int) uintptr {
	return syscall.NewCallback(f)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getcursorpos
func GetCursorPos() (POINT, error) {
	var p POINT
	ret, _, err := getCursorPos.Call(uintptr(unsafe.Pointer(&p)))
	return p, makeErr(ret == 0, "w32.GetCursorPos returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-setcursorpos
func SetCursorPos(x, y int32) error {
	ret, _, err := setCursorPos.Call(uintptr(x), uintptr(y))
	return makeErr(ret == 0, "w32.SetCursorPos returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-messagebeep
func MessageBeep(typ uint32) error {
	ret, _, err := messageBeep.Call(uintptr(typ))
	return makeErr(ret == 0, "w32.MessageBeep returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/utilapiset/nf-utilapiset-beep
func Beep(frequencyInHz, durationInMs uint32) error {
	ret, _, err := beep.Call(
		uintptr(frequencyInHz),
		uintptr(durationInMs),
	)
	return makeErr(ret == 0, "w32.Beep returned FALSE", err)
}

// CreateActCtx sets the Size field of the given ACTCTX.
//
// https://learn.microsoft.com/windows/win32/api/winbase/nf-winbase-createactctxw
func CreateActCtx(activisionCtx *ACTCTX) (HANDLE, error) {
	if activisionCtx != nil {
		activisionCtx.Size = uint32(unsafe.Sizeof(*activisionCtx))
	}
	ret, _, err := createActCtx.Call(uintptr(unsafe.Pointer(activisionCtx)))
	return HANDLE(ret), makeErr(
		HANDLE(ret) == INVALID_HANDLE_VALUE,
		"w32.CreateActCtx returned INVALID_HANDLE_VALUE",
		err,
	)
}

// https://learn.microsoft.com/windows/win32/api/winbase/nf-winbase-activateactctx
func ActivateActCtx(activationCtx HANDLE) (uintptr, error) {
	var cookie uintptr
	ret, _, err := activateActCtx.Call(
		uintptr(activationCtx),
		uintptr(unsafe.Pointer(&cookie)),
	)
	return cookie, makeErr(ret == 0, "w32.ActivateActCtx returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winbase/nf-winbase-deactivateactctx
func DeactivateActCtx(flags uint32, cookie uintptr) error {
	ret, _, err := deactivateActCtx.Call(uintptr(flags), cookie)
	return makeErr(ret == 0, "w32.DeactivateActCtx returned FALSE", err)
}

// https://learn.microsoft.com/previous-versions/dd743680(v=vs.85)
func PlaySound(sound UTF16String, module HINSTANCE, flags uint32) error {
	ret, _, _ := playSound.Call(
		uintptr(unsafe.Pointer(sound)),
		uintptr(module),
		uintptr(flags),
	)
	return makeErr(ret == 0, "w32.PlaySound returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-messageboxw
func MessageBox(owner HWND, text, caption UTF16String, typ uint32) (int32, error) {
	ret, _, err := messageBox.Call(
		uintptr(owner),
		uintptr(unsafe.Pointer(text)),
		uintptr(unsafe.Pointer(caption)),
		uintptr(typ),
	)
	return int32(ret), makeErr(ret == 0, "w32.MessageBox returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-setwindowtextw
func SetWindowText(window HWND, text UTF16String) error {
	ret, _, err := setWindowText.Call(
		uintptr(window),
		uintptr(unsafe.Pointer(text)),
	)
	return makeErr(ret == 0, "w32.SetWindowText returned FALSE", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getwindowtextw
func GetWindowText(window HWND) (string, error) {
	textLen, err := GetWindowTextLength(window)
	if err != nil {
		return "", fmt.Errorf("w32.GetWindowText: %w", err)
	}
	textLen++ // Make room for trailing 0 byte.
	buf := make([]uint16, textLen)
	len, _, err := getWindowText.Call(
		uintptr(window),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen),
	)
	return syscall.UTF16ToString(buf[:len]), makeErr(
		len == 0 && err.(syscall.Errno) != 0,
		"w32.GetWindowText returned 0: %w",
		err,
	)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getwindowtextlengthw
func GetWindowTextLength(window HWND) (int, error) {
	ret, _, err := getWindowTextLength.Call(uintptr(window))
	return int(ret), makeErr(
		ret == 0 && err.(syscall.Errno) != 0,
		"w32.GetWindowTextLength returned 0",
		err,
	)
}

// If you pass 0 for the callback, a WM_TIMER message will be sent to the
// message loop. If you want a custom callback, create one with NewTimerProc.
//
// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-settimer
func SetTimer(window HWND, id uintptr, ms uint32, callback uintptr) (uintptr, error) {
	ret, _, err := setTimer.Call(uintptr(window), id, uintptr(ms), callback)
	return ret, makeErr(ret == 0, "w32.SetTimer returned 0", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-killtimer
func KillTimer(window HWND, id uintptr) error {
	ret, _, err := killTimer.Call(uintptr(window), id)
	return makeErr(ret == 0, "w32.KillTimer returned FALSE", err)
}

// NewTimerProc returns a callback to be used in SetTimer.
func NewTimerProc(f func(window HWND, message uint32, id uintptr, time uint32)) uintptr {
	return syscall.NewCallback(func(window HWND, message uint32, id uintptr, time uint32) uintptr {
		f(window, message, id, time)
		return 0
	})
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-enablewindow
func EnableWindow(window HWND, enable bool) bool {
	ret, _, _ := enableWindow.Call(
		uintptr(window),
		uintptr(BoolToBOOL(enable)),
	)
	return ret == 0
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-createacceleratortablew
func CreateAcceleratorTable(accelerators []ACCEL) (HACCEL, error) {
	// NOTE that there seems to be a bug in this function. Whenever we pass an
	// odd number of accelerators, it fails and returns nil. As a work-around
	// we make sure the number is even by adding an empty accelerator, which
	// will not trigger.
	// This bug was not reproducible in a C program.
	if len(accelerators)%2 == 1 {
		accelerators = append(accelerators, ACCEL{})
	}

	var ptr uintptr
	if len(accelerators) > 0 {
		ptr = uintptr(unsafe.Pointer(&accelerators[0]))
	}
	ret, _, err := createAcceleratorTable.Call(ptr, uintptr(len(accelerators)))
	return HACCEL(ret), makeErr(ret == 0, "w32.CreateAcceleratorTable returned nil", err)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-destroyacceleratortable
func DestroyAcceleratorTable(hAccel HACCEL) error {
	ret, _, _ := destroyAcceleratorTable.Call(
		uintptr(hAccel),
	)
	return makeErr(ret == 0, "w32.DestroyAcceleratorTable returned FALSE", nil)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getfocus
func GetFocus() HWND {
	ret, _, _ := getFocus.Call()
	return HWND(ret)
}

// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-defsubclassproc
func DefSubclassProc(window HWND, message uint32, w, l uintptr) uintptr {
	ret, _, _ := defSubclassProc.Call(uintptr(window), uintptr(message), w, l)
	return ret
}

// Use NewWindowSubclassProc to create the proc.
//
// https://learn.microsoft.com/windows/win32/api/commctrl/nf-commctrl-setwindowsubclass
func SetWindowSubclass(window HWND, proc, subclassID, refData uintptr) error {
	ret, _, _ := setWindowSubclass.Call(
		uintptr(window),
		proc,
		subclassID,
		refData,
	)
	return makeErr(ret == 0, "w32.SetWindowSubclass returned FALSE", nil)
}

// NewWindowSubclassProc returns a new callback to be used in
// SetWindowSubclass.
func NewWindowSubclassProc(f func(window HWND, message uint32, w, l, subclassID, refData uintptr) uintptr) uintptr {
	return syscall.NewCallback(f)
}

// https://learn.microsoft.com/windows/console/getconsolewindow
func GetConsoleWindow() HWND {
	ret, _, _ := getConsoleWindow.Call()
	return HWND(ret)
}

// GetWindowThreadProcessId returns the ID of the thread that created the
// window as the first return value, and the ID of the process that created the
// window as the second return value.
//
// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getwindowthreadprocessid
func GetWindowThreadProcessId(window HWND) (uint32, uint32, error) {
	var processID uint32
	ret, _, err := getWindowThreadProcessId.Call(
		uintptr(window),
		uintptr(unsafe.Pointer(&processID)),
	)
	return uint32(ret), processID,
		makeErr(ret == 0, "w32.GetWindowThreadProcessId returned 0", err)
}

// https://learn.microsoft.com/windows/win32/api/processthreadsapi/nf-processthreadsapi-getcurrentprocessid
func GetCurrentProcessId() uint32 {
	ret, _, _ := getCurrentProcessId.Call()
	return uint32(ret)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-setscrollpos
func SetScrollPos(window HWND, bar, pos int32, redraw bool) (int32, error) {
	ret, _, err := setScrollPos.Call(
		uintptr(window),
		uintptr(bar),
		uintptr(pos),
		uintptr(BoolToBOOL(redraw)),
	)
	return int32(ret), makeErr(
		ret == 0 && err.(syscall.Errno) != 0,
		"w32.SetScrollPos returned nil",
		err,
	)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getscrollpos
func GetScrollPos(window HWND, bar int32) (int32, error) {
	ret, _, err := getScrollPos.Call(
		uintptr(window),
		uintptr(bar),
	)
	return int32(ret), makeErr(
		ret == 0 && err.(syscall.Errno) != 0,
		"w32.GetScrollPos returned nil",
		err,
	)
}

// TODO
func Edit_GetFirstVisibleLine(edit HWND) int32 {
	return int32(SendMessage(edit, EM_GETFIRSTVISIBLELINE, 0, 0))
}

// TODO
func Edit_LimitText(edit HWND, maxChars int32) int32 {
	return int32(SendMessage(edit, EM_LIMITTEXT, uintptr(maxChars), 0))
}

// TODO
func Edit_GetLineCount(edit HWND) uint32 {
	return uint32(SendMessage(edit, EM_GETLINECOUNT, 0, 0))
}

// TODO: https://github.com/microsoft/win32metadata/issues/1846
//
// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getsystemmetrics
func GetSystemMetrics(index int32) (int32, error) {
	ret, _, err := getSystemMetrics.Call(uintptr(index))
	return int32(ret), makeErr(
		ret == 0 && err.(syscall.Errno) != 0,
		"w32.GetSystemMetricsForDpi returned 0",
		err,
	)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getsystemmetricsfordpi
func GetSystemMetricsForDpi(index int32, dpi uint32) (int32, error) {
	ret, _, err := getSystemMetricsForDpi.Call(
		uintptr(index),
		uintptr(dpi),
	)
	return int32(ret), makeErr(
		ret == 0 && err.(syscall.Errno) != 0,
		"w32.GetSystemMetricsForDpi returned 0",
		err,
	)
}

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-getasynckeystate
func GetAsyncKeyState(key int32) uint16 {
	ret, _, _ := getAsyncKeyState.Call(uintptr(key))
	return uint16(ret)
}
