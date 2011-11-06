// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package w32

import (
    "unsafe"
)

type (
    BOOL         int
    ATOM         uint16
    HANDLE       uintptr
    HINSTANCE    HANDLE
    HACCEL       HANDLE
    HCURSOR      HANDLE
    HDWP         HANDLE
    HICON        HANDLE
    HMENU        HANDLE
    HWND         HANDLE
    HBRUSH       HANDLE
    HRESULT      uint
    HFONT        HANDLE
    HDC          HANDLE
    HGDIOBJ      HANDLE
    HDROP        HANDLE
    HENHMETAFILE HANDLE
    HBITMAP      HANDLE
    HPEN         HANDLE
    HKEY         HANDLE
    HIMAGELIST   HANDLE
    COLORREF     uint
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
