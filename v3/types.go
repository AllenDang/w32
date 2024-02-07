package w32

// TODO Remove the de-de and en-us from the URLs.
// TODO Check that no type field uses Hungarian notation.

import (
	"syscall"
	"unsafe"
)

type (
	BOOL        int32
	HRESULT     int32 // TODO Have a public HRESULTToError function.
	HANDLE      uintptr
	HINSTANCE   uintptr
	HICON       uintptr
	HCURSOR     uintptr
	HBRUSH      uintptr
	HWND        uintptr
	HMENU       uintptr
	HACCEL      uintptr
	HBITMAP     uintptr
	HDC         uintptr
	HGDIOBJ     uintptr
	HTREEITEM   uintptr
	HIMAGELIST  uintptr
	HHOOK       uintptr
	HMONITOR    uintptr
	HKL         uintptr
	HGLOBAL     uintptr
	HFONT       uintptr
	ATOM        uintptr
	UTF16String *uint16
	COLORREF    uint32
)

// http://msdn.microsoft.com/en-us/library/windows/desktop/ms633577.aspx
type WNDCLASSEX struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   HINSTANCE
	Icon       HICON
	Cursor     HCURSOR
	Background HBRUSH
	MenuName   UTF16String
	ClassName  UTF16String
	IconSm     HICON
}

// https://learn.microsoft.com/de-de/windows/win32/api/winuser/ns-winuser-msg
type MSG struct {
	Hwnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
	private uint32
}

// https://learn.microsoft.com/en-us/windows/win32/api/windef/ns-windef-point
type POINT struct {
	X, Y int32
}

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-iconinfo
type ICONINFO struct {
	Icon     BOOL // TRUE for icon, FALSE for cursor.
	XHotspot uint32
	YHotspot uint32
	Mask     HBITMAP
	Color    HBITMAP
}

// https://learn.microsoft.com/en-us/windows/win32/api/winuser/ns-winuser-paintstruct
type PAINTSTRUCT struct {
	Hdc       HDC
	Erase     BOOL
	Paint     RECT
	restore   BOOL
	incUpdate BOOL
	reserved  [32]byte
}

// https://learn.microsoft.com/en-us/windows/win32/api/windef/ns-windef-rect
type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

// https://learn.microsoft.com/en-us/windows/win32/api/commctrl/ns-commctrl-initcommoncontrolsex
type INITCOMMONCONTROLSEX struct {
	Size uint32
	ICC  uint32
}

// https://learn.microsoft.com/en-us/windows/win32/api/richedit/ns-richedit-nmhdr
type NMHDR struct {
	HwndFrom HWND
	IdFrom   uint32
	Code     uint32
}

// http://msdn.microsoft.com/en-us/library/windows/desktop/dd183371.aspx
type BITMAP struct {
	Type       int32
	Width      int32
	Height     int32
	WidthBytes int32
	Planes     uint16
	BitsPixel  uint16
	Bits       unsafe.Pointer
}

// INPUT is used in SendInput. To create a concrete INPUT type, use the helper
// functions MouseInput, KeyboardInput and HardwareInput. These are necessary
// because the C API uses a union, which Go does not provide.
type INPUT struct {
	Type uint32
	// This is a union of MOUSEINPUT, KEYBDINPUT and HARDWAREINPUT in C. Since
	// MOUSEINPUT is the largest of those, we use it to hold the data.
	mouse MOUSEINPUT
}

// Use MouseInput to create an INPUT from a MOUSEINPUT.
//
// https://learn.microsoft.com/windows/win32/api/winuser/ns-winuser-mouseinput
type MOUSEINPUT struct {
	Dx        int32
	Dy        int32
	MouseData uint32
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

// Use KeyboardInput to create an INPUT from a KEYBDINPUT.
//
// https://learn.microsoft.com/windows/win32/api/winuser/ns-winuser-keybdinput
type KEYBDINPUT struct {
	Vk        uint16
	Scan      uint16
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

// Use HardwareInput to create an INPUT from HARDWAREINPUT.
//
// https://learn.microsoft.com/windows/win32/api/winuser/ns-winuser-hardwareinput
type HARDWAREINPUT struct {
	Msg    uint32
	ParamL uint16
	ParamH uint16
}

// https://learn.microsoft.com/windows/win32/api/winuser/ns-winuser-kbdllhookstruct
type KBDLLHOOKSTRUCT struct {
	VkCode    uint32
	ScanCode  uint32
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

// https://learn.microsoft.com/windows/win32/api/winuser/ns-winuser-msllhookstruct
type MSLLHOOKSTRUCT struct {
	Pt        POINT
	MouseData uint32
	Flags     uint32
	Time      uint32
	ExtraInfo uintptr
}

// https://learn.microsoft.com/windows/win32/api/winuser/ns-winuser-windowplacement
type WINDOWPLACEMENT struct {
	Length         uint32
	Flags          uint32
	ShowCmd        uint32
	MinPosition    POINT
	MaxPosition    POINT
	NormalPosition RECT
	Device         RECT
}

type MONITORINFO struct {
	Size    uint32
	Monitor RECT
	Work    RECT
	Flags   uint32
}

type MONITORINFOEX struct {
	MONITORINFO
	Device [CCHDEVICENAME]uint16
}

func (m MONITORINFOEX) DeviceName() string {
	return syscall.UTF16ToString(m.Device[:])
}

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-bitmapinfoheader
type BITMAPINFOHEADER struct {
	Size          uint32
	Width         int32
	Height        int32
	Planes        uint16
	BitCount      uint16
	Compression   uint32
	SizeImage     uint32
	XPelsPerMeter int32
	YPelsPerMeter int32
	ClrUsed       uint32
	ClrImportant  uint32
}

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-bitmapinfo
type BITMAPINFO struct {
	Header BITMAPINFOHEADER
	Colors *RGBQUAD
}

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-rgbquad
type RGBQUAD struct {
	Blue     byte
	Green    byte
	Red      byte
	Reserved byte
}

// TODO These were all double-checked for correctness and can be used as is:
// Once the above section is empty, this file is good to be checked in.

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-nmtreevieww
type NMTREEVIEW struct {
	Hdr     NMHDR
	Action  uint32
	ItemOld TVITEM
	ItemNew TVITEM
	Drag    POINT
}

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-nmtvasyncdraw
type NMTVASYNCDRAW struct {
	Hdr           NMHDR
	Imldp         *IMAGELISTDRAWPARAMS
	Hr            HRESULT
	Item          HTREEITEM
	LParam        uintptr
	RetFlags      uint32
	RetImageIndex int32
}

// https://learn.microsoft.com/windows/win32/api/commoncontrols/ns-commoncontrols-imagelistdrawparams
type IMAGELISTDRAWPARAMS struct {
	Size    uint32
	Himl    HIMAGELIST
	I       int32
	HdcDst  HDC
	X       int32
	Y       int32
	Cx      int32
	Cy      int32
	XBitmap int32
	YBitmap int32
	Bk      COLORREF
	Fg      COLORREF
	Style   uint32
	Rop     uint32
	State   uint32
	Frame   uint32
	Effect  COLORREF
}

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-nmtvcustomdraw
type NMTVCUSTOMDRAW struct {
	Nmcd   NMCUSTOMDRAW
	Text   COLORREF
	TextBk COLORREF
	Level  int32
}

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-nmcustomdraw
type NMCUSTOMDRAW struct {
	Hdr        NMHDR
	DrawStage  uint32
	Hdc        HDC
	Rc         RECT
	ItemSpec   uintptr
	ItemState  uint32
	ItemlParam uintptr
}

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-nmtvdispinfow
type NMTVDISPINFO struct {
	Hdr  NMHDR
	Item TVITEM
}

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-nmtvdispinfoexw
type NMTVDISPINFOEX struct {
	Hdr  NMHDR
	Item TVITEMEX
}

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-nmtvgetinfotipw
type NMTVGETINFOTIP struct {
	Hdr     NMHDR
	Text    UTF16String
	TextMax int32
	Item    HTREEITEM
	LParam  uintptr
}

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-nmtvitemchange
type NMTVITEMCHANGE struct {
	Hdr      NMHDR
	Changed  uint32
	Item     HTREEITEM
	StateNew uint32
	StateOld uint32
	LParam   uintptr
}

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-nmtvkeydown
type NMTVKEYDOWN struct {
	Hdr   NMHDR
	VKey  uint16
	Flags uint32
}

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-nmtvstateimagechanging
type NMTVSTATEIMAGECHANGING struct {
	Hdr                NMHDR
	Hti                HTREEITEM
	OldStateImageIndex int32
	NewStateImageIndex int32
}

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-tvgetitempartrectinfo
type TVGETITEMPARTRECTINFO struct {
	Hti    HTREEITEM
	Rc     *RECT
	PartID uint32
}

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-tvhittestinfo
type TVHITTESTINFO struct {
	Pt    POINT
	Flags uint32
	Item  HTREEITEM
}

// TVSORTCB's Compare is a callback. Create it with NewTreeSortCallback.
//
// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-tvsortcb
type TVSORTCB struct {
	HParent HTREEITEM
	Compare uintptr
	LParam  uintptr
}

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-tvinsertstructw
type TVINSERTSTRUCT struct {
	Parent      HTREEITEM
	InsertAfter HTREEITEM
	ItemEx      TVITEMEX
}

// TODO
// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-tvitemw
type TVITEM struct {
	Mask          uint32
	Item          HTREEITEM
	State         uint32
	StateMask     uint32
	Text          UTF16String
	TextMax       int32
	Image         int32
	SelectedImage int32
	Children      int32
	LParam        uintptr
}

// TODO
// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-tvitemexw
type TVITEMEX struct {
	Mask          uint32
	Item          HTREEITEM
	State         uint32
	StateMask     uint32
	Text          UTF16String
	TextMax       int32
	Image         int32
	SelectedImage int32
	Children      int32
	LParam        uintptr
	Integral      int32
	StateEx       uint32
	Hwnd          HWND
	ExpandedImage int32
	reserved      int32
}

// https://learn.microsoft.com/windows/win32/api/winbase/ns-winbase-actctxw
type ACTCTX struct {
	Size                  uint32
	Flags                 uint32
	Source                UTF16String
	ProcessorArchitecture uint16
	LangId                uint16
	AssemblyDirectory     UTF16String
	ResourceName          UTF16String
	ApplicationName       UTF16String
	Module                HINSTANCE
}

// https://learn.microsoft.com/windows/win32/api/shtypes/ns-shtypes-logfontw
type LOGFONT struct {
	Height         int32
	Width          int32
	Escapement     int32
	Orientation    int32
	Weight         int32
	Italic         byte
	Underline      byte
	StrikeOut      byte
	CharSet        byte
	OutPrecision   byte
	ClipPrecision  byte
	Quality        byte
	PitchAndFamily byte
	FaceName       [LF_FACESIZE]uint16
}

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-enumlogfontexdvw
type ENUMLOGFONTEXDV struct {
	EnumLogfontEx ENUMLOGFONTEX
	DesignVector  DESIGNVECTOR
}

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-enumlogfontexw
type ENUMLOGFONTEX struct {
	LogFont  LOGFONT
	FullName [LF_FULLFACESIZE]uint16
	Style    [LF_FULLFACESIZE]uint16
	Script   [LF_FULLFACESIZE]uint16
}

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-designvector
type DESIGNVECTOR struct {
	dvReserved uint32
	dvNumAxes  uint32
	dvValues   [MM_MAX_NUMAXES]int32
}

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-enumtextmetricw
type ENUMTEXTMETRIC struct {
	NewTextMetricEx NEWTEXTMETRICEX
	AxesList        AXESLIST
}

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-newtextmetricexw
type NEWTEXTMETRICEX struct {
	Tm      NEWTEXTMETRIC
	FontSig FONTSIGNATURE
}

// Reserved must be STAMP_AXESLIST.
//
// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-axeslistw
type AXESLIST struct {
	Reserved uint32
	NumAxes  uint32
	AxisInfo [MM_MAX_NUMAXES]AXISINFO
}

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-newtextmetricw
type NEWTEXTMETRIC struct {
	Height           int32
	Ascent           int32
	Descent          int32
	InternalLeading  int32
	ExternalLeading  int32
	AveCharWidth     int32
	MaxCharWidth     int32
	Weight           int32
	Overhang         int32
	DigitizedAspectX int32
	DigitizedAspectY int32
	FirstChar        uint16
	LastChar         uint16
	DefaultChar      uint16
	BreakChar        uint16
	Italic           byte
	Underlined       byte
	StruckOut        byte
	PitchAndFamily   byte
	CharSet          byte
	Flags            uint32
	SizeEM           uint32
	CellHeight       uint32
	AvgWidth         uint32
}

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-fontsignature
type FONTSIGNATURE struct {
	Usb [4]uint32
	Csb [2]uint32
}

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-axisinfow
type AXISINFO struct {
	MinValue int32
	MaxValue int32
	AxisName [MM_MAX_AXES_NAMELEN]uint16
}

// https://learn.microsoft.com/windows/win32/api/windef/ns-windef-size
type SIZE struct {
	Cx int32
	Cy int32
}

// https://learn.microsoft.com/windows/win32/api/winuser/ns-winuser-accel
type ACCEL struct {
	Virt byte
	Key  uint16
	Cmd  uint16
}
