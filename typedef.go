// Copyright 2010-2012 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package w32

import (
	"unsafe"
)

// From MSDN: Windows Data Types
// http://msdn.microsoft.com/en-us/library/windows/desktop/aa383751%28v=vs.85%29.aspx
// ATOM                  WORD
// BOOL                  int32
// BOOLEAN               byte
// BYTE                  byte
// CCHAR                 int8
// CHAR                  int8
// COLORREF              DWORD
// DWORD                 uint32
// DWORDLONG             ULONGLONG
// DWORD_PTR             ULONG_PTR
// DWORD32               uint32
// DWORD64               uint64
// FLOAT                 float32
// HACCEL                HANDLE
// HALF_PTR              struct{} // ???
// HANDLE                PVOID
// HBITMAP               HANDLE
// HBRUSH                HANDLE
// HCOLORSPACE           HANDLE
// HCONV                 HANDLE
// HCONVLIST             HANDLE
// HCURSOR               HANDLE
// HDC                   HANDLE
// HDDEDATA              HANDLE
// HDESK                 HANDLE
// HDROP                 HANDLE
// HDWP                  HANDLE
// HENHMETAFILE          HANDLE
// HFILE                 HANDLE
// HFONT                 HANDLE
// HGDIOBJ               HANDLE
// HGLOBAL               HANDLE
// HHOOK                 HANDLE
// HICON                 HANDLE
// HINSTANCE             HANDLE
// HKEY                  HANDLE
// HKL                   HANDLE
// HLOCAL                HANDLE
// HMENU                 HANDLE
// HMETAFILE             HANDLE
// HMODULE               HANDLE
// HPALETTE              HANDLE
// HPEN                  HANDLE
// HRESULT               int32
// HRGN                  HANDLE
// HSZ                   HANDLE
// HWINSTA               HANDLE
// HWND                  HANDLE
// INT                   int32
// INT_PTR               uintptr
// INT8                  int8
// INT16                 int16
// INT32                 int32
// INT64                 int64
// LANGID                WORD
// LCID                  DWORD
// LCTYPE                DWORD
// LGRPID                DWORD
// LONG                  int32
// LONGLONG              int64
// LONG_PTR              uintptr
// LONG32                int32
// LONG64                int64
// LPARAM                LONG_PTR
// LPBOOL                *BOOL
// LPBYTE                *BYTE
// LPCOLORREF            *COLORREF
// LPCSTR                *int8
// LPCTSTR               LPCWSTR
// LPCVOID               unsafe.Pointer
// LPCWSTR               *WCHAR
// LPDWORD               *DWORD
// LPHANDLE              *HANDLE
// LPINT                 *INT
// LPLONG                *LONG
// LPSTR                 *CHAR
// LPTSTR                LPWSTR
// LPVOID                unsafe.Pointer
// LPWORD                *WORD
// LPWSTR                *WCHAR
// LRESULT               LONG_PTR
// PBOOL                 *BOOL
// PBOOLEAN              *BOOLEAN
// PBYTE                 *BYTE
// PCHAR                 *CHAR
// PCSTR                 *CHAR
// PCTSTR                PCWSTR
// PCWSTR                *WCHAR
// PDWORD                *DWORD
// PDWORDLONG            *DWORDLONG
// PDWORD_PTR            *DWORD_PTR
// PDWORD32              *DWORD32
// PDWORD64              *DWORD64
// PFLOAT                *FLOAT
// PHALF_PTR             *HALF_PTR
// PHANDLE               *HANDLE
// PHKEY                 *HKEY
// PINT_PTR              *INT_PTR
// PINT8                 *INT8
// PINT16                *INT16
// PINT32                *INT32
// PINT64                *INT64
// PLCID                 *LCID
// PLONG                 *LONG
// PLONGLONG             *LONGLONG
// PLONG_PTR             *LONG_PTR
// PLONG32               *LONG32
// PLONG64               *LONG64
// POINTER_32            struct{} // ???
// POINTER_64            struct{} // ???
// POINTER_SIGNED        uintptr
// POINTER_UNSIGNED      uintptr
// PSHORT                *SHORT
// PSIZE_T               *SIZE_T
// PSSIZE_T              *SSIZE_T
// PSTR                  *CHAR
// PTBYTE                *TBYTE
// PTCHAR                *TCHAR
// PTSTR                 PWSTR
// PUCHAR                *UCHAR
// PUHALF_PTR            *UHALF_PTR
// PUINT                 *UINT
// PUINT_PTR             *UINT_PTR
// PUINT8                *UINT8
// PUINT16               *UINT16
// PUINT32               *UINT32
// PUINT64               *UINT64
// PULONG                *ULONG
// PULONGLONG            *ULONGLONG
// PULONG_PTR            *ULONG_PTR
// PULONG32              *ULONG32
// PULONG64              *ULONG64
// PUSHORT               *USHORT
// PVOID                 unsafe.Pointer
// PWCHAR                *WCHAR
// PWORD                 *WORD
// PWSTR                 *WCHAR
// QWORD                 uint64
// SC_HANDLE             HANDLE
// SC_LOCK               LPVOID
// SERVICE_STATUS_HANDLE HANDLE
// SHORT                 int16
// SIZE_T                ULONG_PTR
// SSIZE_T               LONG_PTR
// TBYTE                 WCHAR
// TCHAR                 WCHAR
// UCHAR                 uint8
// UHALF_PTR             struct{} // ???
// UINT                  uint32
// UINT_PTR              uintptr
// UINT8                 uint8
// UINT16                uint16
// UINT32                uint32
// UINT64                uint64
// ULONG                 uint32
// ULONGLONG             uint64
// ULONG_PTR             uintptr
// ULONG32               uint32
// ULONG64               uint64
// USHORT                uint16
// USN                   LONGLONG
// WCHAR                 uint16
// WORD                  uint16
// WPARAM                UINT_PTR
type (
	ATOM            uint16
	BOOL            int32
	COLORREF        DWORD
	DWM_FRAME_COUNT uint64
	DWORD           uint32
	HACCEL          HANDLE
	HANDLE          uintptr
	HBITMAP         HANDLE
	HBRUSH          HANDLE
	HCURSOR         HANDLE
	HDC             HANDLE
	HDROP           HANDLE
	HDWP            HANDLE
	HENHMETAFILE    HANDLE
	HFONT           HANDLE
	HGDIOBJ         HANDLE
	HGLOBAL         HANDLE
	HICON           HANDLE
	HIMAGELIST      HANDLE
	HINSTANCE       HANDLE
	HKEY            HANDLE
	HMENU           HANDLE
	HMODULE         HANDLE
	HPEN            HANDLE
	HRESULT         int32
	HRGN            HANDLE
	HRSRC           HANDLE
	HTHUMBNAIL      HANDLE
	HWND            HANDLE
	LPCVOID         unsafe.Pointer
	PVOID           unsafe.Pointer
	QPC_TIME        uint64
)

type POINT struct {
	X, Y int
}

type RECT struct {
	Left, Top, Right, Bottom int
}

type WNDCLASSEX struct {
	Size       uint
	Style      uint
	WndProc    uintptr
	ClsExtra   int
	WndExtra   int
	Instance   HINSTANCE
	Icon       HICON
	Cursor     HCURSOR
	Background HBRUSH
	MenuName   *uint16
	ClassName  *uint16
	IconSm     HICON
}

type MSG struct {
	Hwnd    HWND
	Message uint
	WParam  uintptr
	LParam  uintptr
	Time    uint
	Pt      POINT
}

type LOGFONT struct {
	Height         int
	Width          int
	Escapement     int
	Orientation    int
	Weight         int
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

type OPENFILENAME struct {
	StructSize      uint
	Owner           HWND
	Instance        HINSTANCE
	Filter          *uint16
	CustomFilter    *uint16
	MaxCustomFilter uint
	FilterIndex     uint
	File            *uint16
	MaxFile         uint
	FileTitle       *uint16
	MaxFileTitle    uint
	InitialDir      *uint16
	Title           *uint16
	Flags           uint
	FileOffset      uint16
	FileExtension   uint16
	DefExt          *uint16
	CustData        uintptr
	FnHook          uintptr
	TemplateName    *uint16
	PvReserved      unsafe.Pointer
	DwReserved      uint
	FlagsEx         uint
}

type BROWSEINFO struct {
	Owner        HWND
	Root         *uint16
	DisplayName  *uint16
	Title        *uint16
	Flags        uint
	CallbackFunc uintptr
	LParam       uintptr
	Image        int
}

type GUID struct {
	Data1 uint
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type VARIANT struct {
	VT         uint16 //  2
	WReserved1 uint16 //  4
	WReserved2 uint16 //  6
	WReserved3 uint16 //  8
	Val        int64  // 16
}

type DISPPARAMS struct {
	Rgvarg            uintptr
	RgdispidNamedArgs uintptr
	CArgs             uint
	CNamedArgs        uint
}

type EXCEPINFO struct {
	WCode             uint16
	WReserved         uint16
	BstrSource        *uint16
	BstrDescription   *uint16
	BstrHelpFile      *uint16
	DwHelpContext     uint
	PvReserved        uintptr
	PfnDeferredFillIn uintptr
	Scode             int
}

type LOGBRUSH struct {
	LbStyle uint
	LbColor COLORREF
	LbHatch uintptr
}

type DEVMODE struct {
	DmDeviceName       [CCHDEVICENAME]uint16
	DmSpecVersion      uint16
	DmDriverVersion    uint16
	DmSize             uint16
	DmDriverExtra      uint16
	DmFields           uint
	DmOrientation      int16
	DmPaperSize        int16
	DmPaperLength      int16
	DmPaperWidth       int16
	DmScale            int16
	DmCopies           int16
	DmDefaultSource    int16
	DmPrintQuality     int16
	DmColor            int16
	DmDuplex           int16
	DmYResolution      int16
	DmTTOption         int16
	DmCollate          int16
	DmFormName         [CCHFORMNAME]uint16
	DmLogPixels        uint16
	DmBitsPerPel       uint
	DmPelsWidth        uint
	DmPelsHeight       uint
	DmDisplayFlags     uint
	DmDisplayFrequency uint
	DmICMMethod        uint
	DmICMIntent        uint
	DmMediaType        uint
	DmDitherType       uint
	DmReserved1        uint
	DmReserved2        uint
	DmPanningWidth     uint
	DmPanningHeight    uint
}

type BITMAPINFOHEADER struct {
	BiSize          uint
	BiWidth         int
	BiHeight        int
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint
	BiSizeImage     uint
	BiXPelsPerMeter int
	BiYPelsPerMeter int
	BiClrUsed       uint
	BiClrImportant  uint
}

type RGBQUAD struct {
	RgbBlue     byte
	RgbGreen    byte
	RgbRed      byte
	RgbReserved byte
}

type BITMAPINFO struct {
	BmiHeader BITMAPINFOHEADER
	BmiColors *RGBQUAD
}

type BITMAP struct {
	BmType       int32
	BmWidth      int32
	BmHeight     int32
	BmWidthBytes int32
	BmPlanes     uint16
	BmBitsPixel  uint16
	BmBits       unsafe.Pointer
}

type DIBSECTION struct {
	DsBm        BITMAP
	DsBmih      BITMAPINFOHEADER
	DsBitfields [3]uint32
	DshSection  HANDLE
	DsOffset    uint32
}

type ENHMETAHEADER struct {
	IType          uint
	NSize          uint
	RclBounds      RECT
	RclFrame       RECT
	DSignature     uint
	NVersion       uint
	NBytes         uint
	NRecords       uint
	NHandles       uint16
	SReserved      uint16
	NDescription   uint
	OffDescription uint
	NPalEntries    uint
	SzlDevice      SIZE
	SzlMillimeters SIZE
	CbPixelFormat  uint
	OffPixelFormat uint
	BOpenGL        uint
	SzlMicrometers SIZE
}

type SIZE struct {
	CX, CY int
}

type TEXTMETRIC struct {
	TmHeight           int
	TmAscent           int
	TmDescent          int
	TmInternalLeading  int
	TmExternalLeading  int
	TmAveCharWidth     int
	TmMaxCharWidth     int
	TmWeight           int
	TmOverhang         int
	TmDigitizedAspectX int
	TmDigitizedAspectY int
	TmFirstChar        uint16
	TmLastChar         uint16
	TmDefaultChar      uint16
	TmBreakChar        uint16
	TmItalic           byte
	TmUnderlined       byte
	TmStruckOut        byte
	TmPitchAndFamily   byte
	TmCharSet          byte
}

type DOCINFO struct {
	CbSize       int
	LpszDocName  *uint16
	LpszOutput   *uint16
	LpszDatatype *uint16
	FwType       uint
}

type NMHDR struct {
	HwndFrom HWND
	IdFrom   uintptr
	Code     uint
}

type LVCOLUMN struct {
	Mask       uint
	Fmt        int
	Cx         int
	PszText    *uint16
	CchTextMax int
	ISubItem   int
	IImage     int
	IOrder     int
}

type LVITEM struct {
	Mask       uint
	IItem      int
	ISubItem   int
	State      uint
	StateMask  uint
	PszText    *uint16
	CchTextMax int
	IImage     int
	LParam     uintptr
	IIndent    int
	IGroupId   int
	CColumns   uint
	PuColumns  uint
}

type LVHITTESTINFO struct {
	Pt       POINT
	Flags    uint
	IItem    int
	ISubItem int
	IGroup   int
}

type NMITEMACTIVATE struct {
	Hdr       NMHDR
	IItem     int
	ISubItem  int
	UNewState uint
	UOldState uint
	UChanged  uint
	PtAction  POINT
	LParam    uintptr
	UKeyFlags uint
}

type NMLISTVIEW struct {
	Hdr       NMHDR
	IItem     int
	ISubItem  int
	UNewState uint
	UOldState uint
	UChanged  uint
	PtAction  POINT
	LParam    uintptr
}

type NMLVDISPINFO struct {
	Hdr  NMHDR
	Item LVITEM
}

type INITCOMMONCONTROLSEX struct {
	DwSize uint
	DwICC  uint
}

type TOOLINFO struct {
	CbSize     uint
	UFlags     uint
	Hwnd       HWND
	UId        uintptr
	Rect       RECT
	Hinst      HINSTANCE
	LpszText   *uint16
	LParam     uintptr
	LpReserved unsafe.Pointer
}

type TRACKMOUSEEVENT struct {
	CbSize      uint
	DwFlags     uint
	HwndTrack   HWND
	DwHoverTime uint
}

type GdiplusStartupInput struct {
	GdiplusVersion           uint32
	DebugEventCallback       uintptr
	SuppressBackgroundThread BOOL
	SuppressExternalCodecs   BOOL
}

type GdiplusStartupOutput struct {
	NotificationHook   uintptr
	NotificationUnhook uintptr
}

type PAINTSTRUCT struct {
	Hdc         HDC
	FErase      BOOL
	RcPaint     RECT
	FRestore    BOOL
	FIncUpdate  BOOL
	RgbReserved [32]byte
}

type EVENTLOGRECORD struct {
	Length              uint32
	Reserved            uint32
	RecordNumber        uint32
	TimeGenerated       uint32
	TimeWritten         uint32
	EventID             uint32
	EventType           uint16
	NumStrings          uint16
	EventCategory       uint16
	ReservedFlags       uint16
	ClosingRecordNumber uint32
	StringOffset        uint32
	UserSidLength       uint32
	UserSidOffset       uint32
	DataLength          uint32
	DataOffset          uint32
}

type SERVICE_STATUS struct {
	DwServiceType             DWORD
	DwCurrentState            DWORD
	DwControlsAccepted        DWORD
	DwWin32ExitCode           DWORD
	DwServiceSpecificExitCode DWORD
	DwCheckPoint              DWORD
	DwWaitHint                DWORD
}

type MODULEENTRY32 struct {
	Size         uint32
	ModuleID     uint32
	ProcessID    uint32
	GlblcntUsage uint32
	ProccntUsage uint32
	ModBaseAddr  *uint8
	ModBaseSize  uint32
	HModule      HMODULE
	SzModule     [MAX_MODULE_NAME32 + 1]uint16
	SzExePath    [MAX_PATH]uint16
}

type FILETIME struct {
	DwLowDateTime  DWORD
	DwHighDateTime DWORD
}

type COORD struct {
	X, Y int16
}

type SMALL_RECT struct {
	Left, Top, Right, Bottom int16
}

type CONSOLE_SCREEN_BUFFER_INFO struct {
	DwSize              COORD
	DwCursorPosition    COORD
	WAttributes         uint16
	SrWindow            SMALL_RECT
	DwMaximumWindowSize COORD
}

type MARGINS struct {
	cxLeftWidth    int
	cxRightWidth   int
	cyTopHeight    int
	cyBottomHeight int
}

type DWM_BLURBEHIND struct {
	DwFlags                DWORD
	fEnable                BOOL
	hRgnBlur               HRGN
	fTransitionOnMaximized BOOL
}

type DWM_PRESENT_PARAMETERS struct {
	cbSize             uint32
	fQueue             BOOL
	cRefreshStart      DWM_FRAME_COUNT
	cBuffer            uint
	fUseSourceRate     BOOL
	rateSource         UNSIGNED_RATIO
	cRefreshesPerFrame uint
	eSampling          DWM_SOURCE_FRAME_SAMPLING
}

type DWM_THUMBNAIL_PROPERTIES struct {
	dwFlags               DWORD
	rcDestination         RECT
	rcSource              RECT
	opacity               byte
	fVisible              BOOL
	fSourceClientAreaOnly BOOL
}

type DWM_TIMING_INFO struct {
	cbSize                 uint32
	rateRefresh            UNSIGNED_RATIO
	qpcRefreshPeriod       QPC_TIME
	rateCompose            UNSIGNED_RATIO
	qpcVBlank              QPC_TIME
	cRefresh               DWM_FRAME_COUNT
	cDXRefresh             uint
	qpcCompose             QPC_TIME
	cFrame                 DWM_FRAME_COUNT
	cDXPresent             uint
	cRefreshFrame          DWM_FRAME_COUNT
	cFrameSubmitted        DWM_FRAME_COUNT
	cDXPresentSubmitted    uint
	cFrameConfirmed        DWM_FRAME_COUNT
	cDXPresentConfirmed    uint
	cRefreshConfirmed      DWM_FRAME_COUNT
	cDXRefreshConfirmed    uint
	cFramesLate            DWM_FRAME_COUNT
	cFramesOutstanding     uint
	cFrameDisplayed        DWM_FRAME_COUNT
	qpcFrameDisplayed      QPC_TIME
	cRefreshFrameDisplayed DWM_FRAME_COUNT
	cFrameComplete         DWM_FRAME_COUNT
	qpcFrameComplete       QPC_TIME
	cFramePending          DWM_FRAME_COUNT
	qpcFramePending        QPC_TIME
	cFramesDisplayed       DWM_FRAME_COUNT
	cFramesComplete        DWM_FRAME_COUNT
	cFramesPending         DWM_FRAME_COUNT
	cFramesAvailable       DWM_FRAME_COUNT
	cFramesDropped         DWM_FRAME_COUNT
	cFramesMissed          DWM_FRAME_COUNT
	cRefreshNextDisplayed  DWM_FRAME_COUNT
	cRefreshNextPresented  DWM_FRAME_COUNT
	cRefreshesDisplayed    DWM_FRAME_COUNT
	cRefreshesPresented    DWM_FRAME_COUNT
	cRefreshStarted        DWM_FRAME_COUNT
	cPixelsReceived        uint64
	cPixelsDrawn           uint64
	cBuffersEmpty          DWM_FRAME_COUNT
}

type MilMatrix3x2D struct {
	S_11 float64
	S_12 float64
	S_21 float64
	S_22 float64
	DX   float64
	DY   float64
}

type UNSIGNED_RATIO struct {
	uiNumerator   uint32
	uiDenominator uint32
}
