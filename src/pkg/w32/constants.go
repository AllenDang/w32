// Copyright 2010 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package w32

const (
    FALSE = 0
    TRUE  = 1
)

const (
    CW_USEDEFAULT = ^0x7fffffff
)

// ShowWindow constants
const (
    SW_HIDE            = 0
    SW_NORMAL          = 1
    SW_SHOWNORMAL      = 1
    SW_SHOWMINIMIZED   = 2
    SW_MAXIMIZE        = 3
    SW_SHOWMAXIMIZED   = 3
    SW_SHOWNOACTIVATE  = 4
    SW_SHOW            = 5
    SW_MINIMIZE        = 6
    SW_SHOWMINNOACTIVE = 7
    SW_SHOWNA          = 8
    SW_RESTORE         = 9
    SW_SHOWDEFAULT     = 10
    SW_FORCEMINIMIZE   = 11
)

// Window class styles
const (
    CS_VREDRAW         = 0x00000001
    CS_HREDRAW         = 0x00000002
    CS_KEYCVTWINDOW    = 0x00000004
    CS_DBLCLKS         = 0x00000008
    CS_OWNDC           = 0x00000020
    CS_CLASSDC         = 0x00000040
    CS_PARENTDC        = 0x00000080
    CS_NOKEYCVT        = 0x00000100
    CS_NOCLOSE         = 0x00000200
    CS_SAVEBITS        = 0x00000800
    CS_BYTEALIGNCLIENT = 0x00001000
    CS_BYTEALIGNWINDOW = 0x00002000
    CS_GLOBALCLASS     = 0x00004000
    CS_IME             = 0x00010000
    CS_DROPSHADOW      = 0x00020000
)

// Predefined cursor constants
const (
    IDC_ARROW       = 32512
    IDC_IBEAM       = 32513
    IDC_WAIT        = 32514
    IDC_CROSS       = 32515
    IDC_UPARROW     = 32516
    IDC_SIZENWSE    = 32642
    IDC_SIZENESW    = 32643
    IDC_SIZEWE      = 32644
    IDC_SIZENS      = 32645
    IDC_SIZEALL     = 32646
    IDC_NO          = 32648
    IDC_HAND        = 32649
    IDC_APPSTARTING = 32650
    IDC_HELP        = 32651
    IDC_ICON        = 32641
    IDC_SIZE        = 32640
)

// Predefined icon constants
const (
    IDI_APPLICATION = 32512
    IDI_HAND        = 32513
    IDI_QUESTION    = 32514
    IDI_EXCLAMATION = 32515
    IDI_ASTERISK    = 32516
    IDI_WINLOGO     = 32517
    IDI_WARNING     = IDI_EXCLAMATION
    IDI_ERROR       = IDI_HAND
    IDI_INFORMATION = IDI_ASTERISK
)

// Button style constants
const (
    BS_3STATE          = 5
    BS_AUTO3STATE      = 6
    BS_AUTOCHECKBOX    = 3
    BS_AUTORADIOBUTTON = 9
    BS_BITMAP          = 128
    BS_BOTTOM          = 0X800
    BS_CENTER          = 0X300
    BS_CHECKBOX        = 2
    BS_DEFPUSHBUTTON   = 1
    BS_GROUPBOX        = 7
    BS_ICON            = 64
    BS_LEFT            = 256
    BS_LEFTTEXT        = 32
    BS_MULTILINE       = 0X2000
    BS_NOTIFY          = 0X4000
    BS_OWNERDRAW       = 0XB
    BS_PUSHBUTTON      = 0
    BS_PUSHLIKE        = 4096
    BS_RADIOBUTTON     = 4
    BS_RIGHT           = 512
    BS_RIGHTBUTTON     = 32
    BS_TEXT            = 0
    BS_TOP             = 0X400
    BS_USERBUTTON      = 8
    BS_VCENTER         = 0XC00
    BS_FLAT            = 0X8000
)

// Button state constants
const (
    BST_CHECKED       = 1
    BST_INDETERMINATE = 2
    BST_UNCHECKED     = 0
    BST_FOCUS         = 8
    BST_PUSHED        = 4
)

// Predefined brushes constants
const (
    COLOR_3DDKSHADOW              = 21
    COLOR_3DFACE                  = 15
    COLOR_3DHILIGHT               = 20
    COLOR_3DHIGHLIGHT             = 20
    COLOR_3DLIGHT                 = 22
    COLOR_BTNHILIGHT              = 20
    COLOR_3DSHADOW                = 16
    COLOR_ACTIVEBORDER            = 10
    COLOR_ACTIVECAPTION           = 2
    COLOR_APPWORKSPACE            = 12
    COLOR_BACKGROUND              = 1
    COLOR_DESKTOP                 = 1
    COLOR_BTNFACE                 = 15
    COLOR_BTNHIGHLIGHT            = 20
    COLOR_BTNSHADOW               = 16
    COLOR_BTNTEXT                 = 18
    COLOR_CAPTIONTEXT             = 9
    COLOR_GRAYTEXT                = 17
    COLOR_HIGHLIGHT               = 13
    COLOR_HIGHLIGHTTEXT           = 14
    COLOR_INACTIVEBORDER          = 11
    COLOR_INACTIVECAPTION         = 3
    COLOR_INACTIVECAPTIONTEXT     = 19
    COLOR_INFOBK                  = 24
    COLOR_INFOTEXT                = 23
    COLOR_MENU                    = 4
    COLOR_MENUTEXT                = 7
    COLOR_SCROLLBAR               = 0
    COLOR_WINDOW                  = 5
    COLOR_WINDOWFRAME             = 6
    COLOR_WINDOWTEXT              = 8
    COLOR_HOTLIGHT                = 26
    COLOR_GRADIENTACTIVECAPTION   = 27
    COLOR_GRADIENTINACTIVECAPTION = 28
)

// Button message constants
const (
    BM_CLICK    = 245
    BM_GETCHECK = 240
    BM_GETIMAGE = 246
    BM_GETSTATE = 242
    BM_SETCHECK = 241
    BM_SETIMAGE = 247
    BM_SETSTATE = 243
    BM_SETSTYLE = 244
)

// Button notifications
const (
    BN_CLICKED       = 0
    BN_PAINT         = 1
    BN_HILITE        = 2
    BN_PUSHED        = BN_HILITE
    BN_UNHILITE      = 3
    BN_UNPUSHED      = BN_UNHILITE
    BN_DISABLE       = 4
    BN_DOUBLECLICKED = 5
    BN_DBLCLK        = BN_DOUBLECLICKED
    BN_SETFOCUS      = 6
    BN_KILLFOCUS     = 7
)

// GetWindowLong and GetWindowLongPtr constants
const (
    GWL_EXSTYLE     = -20
    GWL_STYLE       = -16
    GWL_WNDPROC     = -4
    GWLP_WNDPROC    = -4
    GWL_HINSTANCE   = -6
    GWLP_HINSTANCE  = -6
    GWL_HWNDPARENT  = -8
    GWLP_HWNDPARENT = -8
    GWL_ID          = -12
    GWLP_ID         = -12
    GWL_USERDATA    = -21
    GWLP_USERDATA   = -21
)

// Window style constants
const (
    WS_OVERLAPPED       = 0X00000000
    WS_POPUP            = 0X80000000
    WS_CHILD            = 0X40000000
    WS_MINIMIZE         = 0X20000000
    WS_VISIBLE          = 0X10000000
    WS_DISABLED         = 0X08000000
    WS_CLIPSIBLINGS     = 0X04000000
    WS_CLIPCHILDREN     = 0X02000000
    WS_MAXIMIZE         = 0X01000000
    WS_CAPTION          = 0X00C00000
    WS_BORDER           = 0X00800000
    WS_DLGFRAME         = 0X00400000
    WS_VSCROLL          = 0X00200000
    WS_HSCROLL          = 0X00100000
    WS_SYSMENU          = 0X00080000
    WS_THICKFRAME       = 0X00040000
    WS_GROUP            = 0X00020000
    WS_TABSTOP          = 0X00010000
    WS_MINIMIZEBOX      = 0X00020000
    WS_MAXIMIZEBOX      = 0X00010000
    WS_TILED            = 0X00000000
    WS_ICONIC           = 0X20000000
    WS_SIZEBOX          = 0X00040000
    WS_OVERLAPPEDWINDOW = 0X00000000 | 0X00C00000 | 0X00080000 | 0X00040000 | 0X00020000 | 0X00010000
    WS_POPUPWINDOW      = 0X80000000 | 0X00800000 | 0X00080000
    WS_CHILDWINDOW      = 0X40000000
)

// Extended window style constants
const (
    WS_EX_DLGMODALFRAME    = 0X00000001
    WS_EX_NOPARENTNOTIFY   = 0X00000004
    WS_EX_TOPMOST          = 0X00000008
    WS_EX_ACCEPTFILES      = 0X00000010
    WS_EX_TRANSPARENT      = 0X00000020
    WS_EX_MDICHILD         = 0X00000040
    WS_EX_TOOLWINDOW       = 0X00000080
    WS_EX_WINDOWEDGE       = 0X00000100
    WS_EX_CLIENTEDGE       = 0X00000200
    WS_EX_CONTEXTHELP      = 0X00000400
    WS_EX_RIGHT            = 0X00001000
    WS_EX_LEFT             = 0X00000000
    WS_EX_RTLREADING       = 0X00002000
    WS_EX_LTRREADING       = 0X00000000
    WS_EX_LEFTSCROLLBAR    = 0X00004000
    WS_EX_RIGHTSCROLLBAR   = 0X00000000
    WS_EX_CONTROLPARENT    = 0X00010000
    WS_EX_STATICEDGE       = 0X00020000
    WS_EX_APPWINDOW        = 0X00040000
    WS_EX_OVERLAPPEDWINDOW = 0X00000100 | 0X00000200
    WS_EX_PALETTEWINDOW    = 0X00000100 | 0X00000080 | 0X00000008
    WS_EX_LAYERED          = 0X00080000
    WS_EX_NOINHERITLAYOUT  = 0X00100000
    WS_EX_LAYOUTRTL        = 0X00400000
    WS_EX_NOACTIVATE       = 0X08000000
)

// Window message constants
const (
    WM_APP                    = 32768
    WM_ACTIVATE               = 6
    WM_ACTIVATEAPP            = 28
    WM_AFXFIRST               = 864
    WM_AFXLAST                = 895
    WM_ASKCBFORMATNAME        = 780
    WM_CANCELJOURNAL          = 75
    WM_CANCELMODE             = 31
    WM_CAPTURECHANGED         = 533
    WM_CHANGECBCHAIN          = 781
    WM_CHAR                   = 258
    WM_CHARTOITEM             = 47
    WM_CHILDACTIVATE          = 34
    WM_CLEAR                  = 771
    WM_CLOSE                  = 16
    WM_COMMAND                = 273
    WM_COMMNOTIFY             = 68 /* OBSOLETE */
    WM_COMPACTING             = 65
    WM_COMPAREITEM            = 57
    WM_CONTEXTMENU            = 123
    WM_COPY                   = 769
    WM_COPYDATA               = 74
    WM_CREATE                 = 1
    WM_CTLCOLORBTN            = 309
    WM_CTLCOLORDLG            = 310
    WM_CTLCOLOREDIT           = 307
    WM_CTLCOLORLISTBOX        = 308
    WM_CTLCOLORMSGBOX         = 306
    WM_CTLCOLORSCROLLBAR      = 311
    WM_CTLCOLORSTATIC         = 312
    WM_CUT                    = 768
    WM_DEADCHAR               = 259
    WM_DELETEITEM             = 45
    WM_DESTROY                = 2
    WM_DESTROYCLIPBOARD       = 775
    WM_DEVICECHANGE           = 537
    WM_DEVMODECHANGE          = 27
    WM_DISPLAYCHANGE          = 126
    WM_DRAWCLIPBOARD          = 776
    WM_DRAWITEM               = 43
    WM_DROPFILES              = 563
    WM_ENABLE                 = 10
    WM_ENDSESSION             = 22
    WM_ENTERIDLE              = 289
    WM_ENTERMENULOOP          = 529
    WM_ENTERSIZEMOVE          = 561
    WM_ERASEBKGND             = 20
    WM_EXITMENULOOP           = 530
    WM_EXITSIZEMOVE           = 562
    WM_FONTCHANGE             = 29
    WM_GETDLGCODE             = 135
    WM_GETFONT                = 49
    WM_GETHOTKEY              = 51
    WM_GETICON                = 127
    WM_GETMINMAXINFO          = 36
    WM_GETTEXT                = 13
    WM_GETTEXTLENGTH          = 14
    WM_HANDHELDFIRST          = 856
    WM_HANDHELDLAST           = 863
    WM_HELP                   = 83
    WM_HOTKEY                 = 786
    WM_HSCROLL                = 276
    WM_HSCROLLCLIPBOARD       = 782
    WM_ICONERASEBKGND         = 39
    WM_INITDIALOG             = 272
    WM_INITMENU               = 278
    WM_INITMENUPOPUP          = 279
    WM_INPUT                  = 0X00FF
    WM_INPUTLANGCHANGE        = 81
    WM_INPUTLANGCHANGEREQUEST = 80
    WM_KEYDOWN                = 256
    WM_KEYUP                  = 257
    WM_KILLFOCUS              = 8
    WM_MDIACTIVATE            = 546
    WM_MDICASCADE             = 551
    WM_MDICREATE              = 544
    WM_MDIDESTROY             = 545
    WM_MDIGETACTIVE           = 553
    WM_MDIICONARRANGE         = 552
    WM_MDIMAXIMIZE            = 549
    WM_MDINEXT                = 548
    WM_MDIREFRESHMENU         = 564
    WM_MDIRESTORE             = 547
    WM_MDISETMENU             = 560
    WM_MDITILE                = 550
    WM_MEASUREITEM            = 44
    WM_GETOBJECT              = 0X003D
    WM_CHANGEUISTATE          = 0X0127
    WM_UPDATEUISTATE          = 0X0128
    WM_QUERYUISTATE           = 0X0129
    WM_UNINITMENUPOPUP        = 0X0125
    WM_MENURBUTTONUP          = 290
    WM_MENUCOMMAND            = 0X0126
    WM_MENUGETOBJECT          = 0X0124
    WM_MENUDRAG               = 0X0123
    WM_APPCOMMAND             = 0X0319
    WM_MENUCHAR               = 288
    WM_MENUSELECT             = 287
    WM_MOVE                   = 3
    WM_MOVING                 = 534
    WM_NCACTIVATE             = 134
    WM_NCCALCSIZE             = 131
    WM_NCCREATE               = 129
    WM_NCDESTROY              = 130
    WM_NCHITTEST              = 132
    WM_NCLBUTTONDBLCLK        = 163
    WM_NCLBUTTONDOWN          = 161
    WM_NCLBUTTONUP            = 162
    WM_NCMBUTTONDBLCLK        = 169
    WM_NCMBUTTONDOWN          = 167
    WM_NCMBUTTONUP            = 168
    WM_NCXBUTTONDOWN          = 171
    WM_NCXBUTTONUP            = 172
    WM_NCXBUTTONDBLCLK        = 173
    WM_NCMOUSEHOVER           = 0X02A0
    WM_NCMOUSELEAVE           = 0X02A2
    WM_NCMOUSEMOVE            = 160
    WM_NCPAINT                = 133
    WM_NCRBUTTONDBLCLK        = 166
    WM_NCRBUTTONDOWN          = 164
    WM_NCRBUTTONUP            = 165
    WM_NEXTDLGCTL             = 40
    WM_NEXTMENU               = 531
    WM_NOTIFY                 = 78
    WM_NOTIFYFORMAT           = 85
    WM_NULL                   = 0
    WM_PAINT                  = 15
    WM_PAINTCLIPBOARD         = 777
    WM_PAINTICON              = 38
    WM_PALETTECHANGED         = 785
    WM_PALETTEISCHANGING      = 784
    WM_PARENTNOTIFY           = 528
    WM_PASTE                  = 770
    WM_PENWINFIRST            = 896
    WM_PENWINLAST             = 911
    WM_POWER                  = 72
    WM_POWERBROADCAST         = 536
    WM_PRINT                  = 791
    WM_PRINTCLIENT            = 792
    WM_QUERYDRAGICON          = 55
    WM_QUERYENDSESSION        = 17
    WM_QUERYNEWPALETTE        = 783
    WM_QUERYOPEN              = 19
    WM_QUEUESYNC              = 35
    WM_QUIT                   = 18
    WM_RENDERALLFORMATS       = 774
    WM_RENDERFORMAT           = 773
    WM_SETCURSOR              = 32
    WM_SETFOCUS               = 7
    WM_SETFONT                = 48
    WM_SETHOTKEY              = 50
    WM_SETICON                = 128
    WM_SETREDRAW              = 11
    WM_SETTEXT                = 12
    WM_SETTINGCHANGE          = 26
    WM_SHOWWINDOW             = 24
    WM_SIZE                   = 5
    WM_SIZECLIPBOARD          = 779
    WM_SIZING                 = 532
    WM_SPOOLERSTATUS          = 42
    WM_STYLECHANGED           = 125
    WM_STYLECHANGING          = 124
    WM_SYSCHAR                = 262
    WM_SYSCOLORCHANGE         = 21
    WM_SYSCOMMAND             = 274
    WM_SYSDEADCHAR            = 263
    WM_SYSKEYDOWN             = 260
    WM_SYSKEYUP               = 261
    WM_TCARD                  = 82
    WM_THEMECHANGED           = 794
    WM_TIMECHANGE             = 30
    WM_TIMER                  = 275
    WM_UNDO                   = 772
    WM_USER                   = 1024
    WM_USERCHANGED            = 84
    WM_VKEYTOITEM             = 46
    WM_VSCROLL                = 277
    WM_VSCROLLCLIPBOARD       = 778
    WM_WINDOWPOSCHANGED       = 71
    WM_WINDOWPOSCHANGING      = 70
    WM_WININICHANGE           = 26
    WM_KEYFIRST               = 256
    WM_KEYLAST                = 264
    WM_SYNCPAINT              = 136
    WM_MOUSEACTIVATE          = 33
    WM_MOUSEMOVE              = 512
    WM_LBUTTONDOWN            = 513
    WM_LBUTTONUP              = 514
    WM_LBUTTONDBLCLK          = 515
    WM_RBUTTONDOWN            = 516
    WM_RBUTTONUP              = 517
    WM_RBUTTONDBLCLK          = 518
    WM_MBUTTONDOWN            = 519
    WM_MBUTTONUP              = 520
    WM_MBUTTONDBLCLK          = 521
    WM_MOUSEWHEEL             = 522
    WM_MOUSEFIRST             = 512
    WM_XBUTTONDOWN            = 523
    WM_XBUTTONUP              = 524
    WM_XBUTTONDBLCLK          = 525
    WM_MOUSELAST              = 525
    WM_MOUSEHOVER             = 0X2A1
    WM_MOUSELEAVE             = 0X2A3
)

const LF_FACESIZE = 32

// Font weight constants
const (
    FW_DONTCARE   = 0
    FW_THIN       = 100
    FW_EXTRALIGHT = 200
    FW_ULTRALIGHT = FW_EXTRALIGHT
    FW_LIGHT      = 300
    FW_NORMAL     = 400
    FW_REGULAR    = 400
    FW_MEDIUM     = 500
    FW_SEMIBOLD   = 600
    FW_DEMIBOLD   = FW_SEMIBOLD
    FW_BOLD       = 700
    FW_EXTRABOLD  = 800
    FW_ULTRABOLD  = FW_EXTRABOLD
    FW_HEAVY      = 900
    FW_BLACK      = FW_HEAVY
)

// Charset constants
const (
    ANSI_CHARSET        = 0
    DEFAULT_CHARSET     = 1
    SYMBOL_CHARSET      = 2
    SHIFTJIS_CHARSET    = 128
    HANGEUL_CHARSET     = 129
    HANGUL_CHARSET      = 129
    GB2312_CHARSET      = 134
    CHINESEBIG5_CHARSET = 136
    GREEK_CHARSET       = 161
    TURKISH_CHARSET     = 162
    HEBREW_CHARSET      = 177
    ARABIC_CHARSET      = 178
    BALTIC_CHARSET      = 186
    RUSSIAN_CHARSET     = 204
    THAI_CHARSET        = 222
    EASTEUROPE_CHARSET  = 238
    OEM_CHARSET         = 255
    JOHAB_CHARSET       = 130
    VIETNAMESE_CHARSET  = 163
    MAC_CHARSET         = 77
)

// Font output precision constants
const (
    OUT_DEFAULT_PRECIS   = 0
    OUT_STRING_PRECIS    = 1
    OUT_CHARACTER_PRECIS = 2
    OUT_STROKE_PRECIS    = 3
    OUT_TT_PRECIS        = 4
    OUT_DEVICE_PRECIS    = 5
    OUT_RASTER_PRECIS    = 6
    OUT_TT_ONLY_PRECIS   = 7
    OUT_OUTLINE_PRECIS   = 8
    OUT_PS_ONLY_PRECIS   = 10
)

// Font clipping precision constants
const (
    CLIP_DEFAULT_PRECIS   = 0
    CLIP_CHARACTER_PRECIS = 1
    CLIP_STROKE_PRECIS    = 2
    CLIP_MASK             = 15
    CLIP_LH_ANGLES        = 16
    CLIP_TT_ALWAYS        = 32
    CLIP_EMBEDDED         = 128
)

// Font output quality constants
const (
    DEFAULT_QUALITY        = 0
    DRAFT_QUALITY          = 1
    PROOF_QUALITY          = 2
    NONANTIALIASED_QUALITY = 3
    ANTIALIASED_QUALITY    = 4
    CLEARTYPE_QUALITY      = 5
)

// Font pitch constants
const (
    DEFAULT_PITCH  = 0
    FIXED_PITCH    = 1
    VARIABLE_PITCH = 2
)

// Font family constants
const (
    FF_DECORATIVE = 80
    FF_DONTCARE   = 0
    FF_MODERN     = 48
    FF_ROMAN      = 16
    FF_SCRIPT     = 64
    FF_SWISS      = 32
)

// DeviceCapabilities capabilities
const (
    DC_FIELDS            = 1
    DC_PAPERS            = 2
    DC_PAPERSIZE         = 3
    DC_MINEXTENT         = 4
    DC_MAXEXTENT         = 5
    DC_BINS              = 6
    DC_DUPLEX            = 7
    DC_SIZE              = 8
    DC_EXTRA             = 9
    DC_VERSION           = 10
    DC_DRIVER            = 11
    DC_BINNAMES          = 12
    DC_ENUMRESOLUTIONS   = 13
    DC_FILEDEPENDENCIES  = 14
    DC_TRUETYPE          = 15
    DC_PAPERNAMES        = 16
    DC_ORIENTATION       = 17
    DC_COPIES            = 18
    DC_BINADJUST         = 19
    DC_EMF_COMPLIANT     = 20
    DC_DATATYPE_PRODUCED = 21
    DC_COLLATE           = 22
    DC_MANUFACTURER      = 23
    DC_MODEL             = 24
    DC_PERSONALITY       = 25
    DC_PRINTRATE         = 26
    DC_PRINTRATEUNIT     = 27
    DC_PRINTERMEM        = 28
    DC_MEDIAREADY        = 29
    DC_STAPLE            = 30
    DC_PRINTRATEPPM      = 31
    DC_COLORDEVICE       = 32
    DC_NUP               = 33
    DC_MEDIATYPENAMES    = 34
    DC_MEDIATYPES        = 35
)

// GetDeviceCaps index constants
const (
    DRIVERVERSION   = 0
    TECHNOLOGY      = 2
    HORZSIZE        = 4
    VERTSIZE        = 6
    HORZRES         = 8
    VERTRES         = 10
    LOGPIXELSX      = 88
    LOGPIXELSY      = 90
    BITSPIXEL       = 12
    PLANES          = 14
    NUMBRUSHES      = 16
    NUMPENS         = 18
    NUMFONTS        = 22
    NUMCOLORS       = 24
    NUMMARKERS      = 20
    ASPECTX         = 40
    ASPECTY         = 42
    ASPECTXY        = 44
    PDEVICESIZE     = 26
    CLIPCAPS        = 36
    SIZEPALETTE     = 104
    NUMRESERVED     = 106
    COLORRES        = 108
    PHYSICALWIDTH   = 110
    PHYSICALHEIGHT  = 111
    PHYSICALOFFSETX = 112
    PHYSICALOFFSETY = 113
    SCALINGFACTORX  = 114
    SCALINGFACTORY  = 115
    VREFRESH        = 116
    DESKTOPHORZRES  = 118
    DESKTOPVERTRES  = 117
    BLTALIGNMENT    = 119
    SHADEBLENDCAPS  = 120
    COLORMGMTCAPS   = 121
    RASTERCAPS      = 38
    CURVECAPS       = 28
    LINECAPS        = 30
    POLYGONALCAPS   = 32
    TEXTCAPS        = 34
)

// GetDeviceCaps TECHNOLOGY constants
const (
    DT_PLOTTER    = 0
    DT_RASDISPLAY = 1
    DT_RASPRINTER = 2
    DT_RASCAMERA  = 3
    DT_CHARSTREAM = 4
    DT_METAFILE   = 5
    DT_DISPFILE   = 6
)

// GetDeviceCaps SHADEBLENDCAPS constants
const (
    SB_NONE          = 0x00
    SB_CONST_ALPHA   = 0x01
    SB_PIXEL_ALPHA   = 0x02
    SB_PREMULT_ALPHA = 0x04
    SB_GRAD_RECT     = 0x10
    SB_GRAD_TRI      = 0x20
)

// GetDeviceCaps COLORMGMTCAPS constants
const (
    CM_NONE       = 0x00
    CM_DEVICE_ICM = 0x01
    CM_GAMMA_RAMP = 0x02
    CM_CMYK_COLOR = 0x04
)

// GetDeviceCaps RASTERCAPS constants
const (
    RC_BANDING      = 2
    RC_BITBLT       = 1
    RC_BITMAP64     = 8
    RC_DI_BITMAP    = 128
    RC_DIBTODEV     = 512
    RC_FLOODFILL    = 4096
    RC_GDI20_OUTPUT = 16
    RC_PALETTE      = 256
    RC_SCALING      = 4
    RC_STRETCHBLT   = 2048
    RC_STRETCHDIB   = 8192
    RC_DEVBITS      = 0x8000
    RC_OP_DX_OUTPUT = 0x4000
)

// GetDeviceCaps CURVECAPS constants
const (
    CC_NONE       = 0
    CC_CIRCLES    = 1
    CC_PIE        = 2
    CC_CHORD      = 4
    CC_ELLIPSES   = 8
    CC_WIDE       = 16
    CC_STYLED     = 32
    CC_WIDESTYLED = 64
    CC_INTERIORS  = 128
    CC_ROUNDRECT  = 256
)

// GetDeviceCaps LINECAPS constants
const (
    LC_NONE       = 0
    LC_POLYLINE   = 2
    LC_MARKER     = 4
    LC_POLYMARKER = 8
    LC_WIDE       = 16
    LC_STYLED     = 32
    LC_WIDESTYLED = 64
    LC_INTERIORS  = 128
)

// GetDeviceCaps POLYGONALCAPS constants
const (
    PC_NONE        = 0
    PC_POLYGON     = 1
    PC_POLYPOLYGON = 256
    PC_PATHS       = 512
    PC_RECTANGLE   = 2
    PC_WINDPOLYGON = 4
    PC_SCANLINE    = 8
    PC_TRAPEZOID   = 4
    PC_WIDE        = 16
    PC_STYLED      = 32
    PC_WIDESTYLED  = 64
    PC_INTERIORS   = 128
)

// GetDeviceCaps TEXTCAPS constants
const (
    TC_OP_CHARACTER = 1
    TC_OP_STROKE    = 2
    TC_CP_STROKE    = 4
    TC_CR_90        = 8
    TC_CR_ANY       = 16
    TC_SF_X_YINDEP  = 32
    TC_SA_DOUBLE    = 64
    TC_SA_INTEGER   = 128
    TC_SA_CONTIN    = 256
    TC_EA_DOUBLE    = 512
    TC_IA_ABLE      = 1024
    TC_UA_ABLE      = 2048
    TC_SO_ABLE      = 4096
    TC_RA_ABLE      = 8192
    TC_VA_ABLE      = 16384
    TC_RESERVED     = 32768
    TC_SCROLLBLT    = 65536
)

// Static control styles
const (
    SS_BITMAP          = 14
    SS_BLACKFRAME      = 7
    SS_BLACKRECT       = 4
    SS_CENTER          = 1
    SS_CENTERIMAGE     = 512
    SS_EDITCONTROL     = 0x2000
    SS_ENHMETAFILE     = 15
    SS_ETCHEDFRAME     = 18
    SS_ETCHEDHORZ      = 16
    SS_ETCHEDVERT      = 17
    SS_GRAYFRAME       = 8
    SS_GRAYRECT        = 5
    SS_ICON            = 3
    SS_LEFT            = 0
    SS_LEFTNOWORDWRAP  = 0xc
    SS_NOPREFIX        = 128
    SS_NOTIFY          = 256
    SS_OWNERDRAW       = 0xd
    SS_REALSIZECONTROL = 0x040
    SS_REALSIZEIMAGE   = 0x800
    SS_RIGHT           = 2
    SS_RIGHTJUST       = 0x400
    SS_SIMPLE          = 11
    SS_SUNKEN          = 4096
    SS_WHITEFRAME      = 9
    SS_WHITERECT       = 6
    SS_USERITEM        = 10
    SS_TYPEMASK        = 0x0000001F
    SS_ENDELLIPSIS     = 0x00004000
    SS_PATHELLIPSIS    = 0x00008000
    SS_WORDELLIPSIS    = 0x0000C000
    SS_ELLIPSISMASK    = 0x0000C000
)

// Edit styles
const (
    ES_LEFT        = 0x0000
    ES_CENTER      = 0x0001
    ES_RIGHT       = 0x0002
    ES_MULTILINE   = 0x0004
    ES_UPPERCASE   = 0x0008
    ES_LOWERCASE   = 0x0010
    ES_PASSWORD    = 0x0020
    ES_AUTOVSCROLL = 0x0040
    ES_AUTOHSCROLL = 0x0080
    ES_NOHIDESEL   = 0x0100
    ES_OEMCONVERT  = 0x0400
    ES_READONLY    = 0x0800
    ES_WANTRETURN  = 0x1000
    ES_NUMBER      = 0x2000
)

// Edit notifications
const (
    EN_SETFOCUS     = 0x0100
    EN_KILLFOCUS    = 0x0200
    EN_CHANGE       = 0x0300
    EN_UPDATE       = 0x0400
    EN_ERRSPACE     = 0x0500
    EN_MAXTEXT      = 0x0501
    EN_HSCROLL      = 0x0601
    EN_VSCROLL      = 0x0602
    EN_ALIGN_LTR_EC = 0x0700
    EN_ALIGN_RTL_EC = 0x0701
)

// Edit messages
const (
    EM_GETSEL              = 0x00B0
    EM_SETSEL              = 0x00B1
    EM_GETRECT             = 0x00B2
    EM_SETRECT             = 0x00B3
    EM_SETRECTNP           = 0x00B4
    EM_SCROLL              = 0x00B5
    EM_LINESCROLL          = 0x00B6
    EM_SCROLLCARET         = 0x00B7
    EM_GETMODIFY           = 0x00B8
    EM_SETMODIFY           = 0x00B9
    EM_GETLINECOUNT        = 0x00BA
    EM_LINEINDEX           = 0x00BB
    EM_SETHANDLE           = 0x00BC
    EM_GETHANDLE           = 0x00BD
    EM_GETTHUMB            = 0x00BE
    EM_LINELENGTH          = 0x00C1
    EM_REPLACESEL          = 0x00C2
    EM_GETLINE             = 0x00C4
    EM_LIMITTEXT           = 0x00C5
    EM_CANUNDO             = 0x00C6
    EM_UNDO                = 0x00C7
    EM_FMTLINES            = 0x00C8
    EM_LINEFROMCHAR        = 0x00C9
    EM_SETTABSTOPS         = 0x00CB
    EM_SETPASSWORDCHAR     = 0x00CC
    EM_EMPTYUNDOBUFFER     = 0x00CD
    EM_GETFIRSTVISIBLELINE = 0x00CE
    EM_SETREADONLY         = 0x00CF
    EM_SETWORDBREAKPROC    = 0x00D0
    EM_GETWORDBREAKPROC    = 0x00D1
    EM_GETPASSWORDCHAR     = 0x00D2
    EM_SETMARGINS          = 0x00D3
    EM_GETMARGINS          = 0x00D4
    EM_SETLIMITTEXT        = EM_LIMITTEXT
    EM_GETLIMITTEXT        = 0x00D5
    EM_POSFROMCHAR         = 0x00D6
    EM_CHARFROMPOS         = 0x00D7
    EM_SETIMESTATUS        = 0x00D8
    EM_GETIMESTATUS        = 0x00D9
    EM_SETCUEBANNER        = 0x1501
    EM_GETCUEBANNER        = 0x1502
)

const (
    CCM_FIRST            = 0x2000
    CCM_LAST             = CCM_FIRST + 0x200
    CCM_SETBKCOLOR       = 8193
    CCM_SETCOLORSCHEME   = 8194
    CCM_GETCOLORSCHEME   = 8195
    CCM_GETDROPTARGET    = 8196
    CCM_SETUNICODEFORMAT = 8197
    CCM_GETUNICODEFORMAT = 8198
    CCM_SETVERSION       = 0x2007
    CCM_GETVERSION       = 0x2008
    CCM_SETNOTIFYWINDOW  = 0x2009
    CCM_SETWINDOWTHEME   = 0x200b
    CCM_DPISCALE         = 0x200c
)

// Common controls styles
const (
    CCS_TOP           = 1
    CCS_NOMOVEY       = 2
    CCS_BOTTOM        = 3
    CCS_NORESIZE      = 4
    CCS_NOPARENTALIGN = 8
    CCS_ADJUSTABLE    = 32
    CCS_NODIVIDER     = 64
    CCS_VERT          = 128
    CCS_LEFT          = 129
    CCS_NOMOVEX       = 130
    CCS_RIGHT         = 131
)

// ProgressBar messages
const (
    PROGRESS_CLASS  = "msctls_progress32"
    PBM_SETPOS      = WM_USER + 2
    PBM_DELTAPOS    = WM_USER + 3
    PBM_SETSTEP     = WM_USER + 4
    PBM_STEPIT      = WM_USER + 5
    PBM_SETRANGE32  = 1030
    PBM_GETRANGE    = 1031
    PBM_GETPOS      = 1032
    PBM_SETBARCOLOR = 1033
    PBM_SETBKCOLOR  = CCM_SETBKCOLOR
    PBS_SMOOTH      = 1
    PBS_VERTICAL    = 4
)

// GetOpenFileName and GetSaveFileName extended flags
const (
    OFN_EX_NOPLACESBAR = 0x00000001
)

// GetOpenFileName and GetSaveFileName flags
const (
    OFN_ALLOWMULTISELECT     = 0x00000200
    OFN_CREATEPROMPT         = 0x00002000
    OFN_DONTADDTORECENT      = 0x02000000
    OFN_ENABLEHOOK           = 0x00000020
    OFN_ENABLEINCLUDENOTIFY  = 0x00400000
    OFN_ENABLESIZING         = 0x00800000
    OFN_ENABLETEMPLATE       = 0x00000040
    OFN_ENABLETEMPLATEHANDLE = 0x00000080
    OFN_EXPLORER             = 0x00080000
    OFN_EXTENSIONDIFFERENT   = 0x00000400
    OFN_FILEMUSTEXIST        = 0x00001000
    OFN_FORCESHOWHIDDEN      = 0x10000000
    OFN_HIDEREADONLY         = 0x00000004
    OFN_LONGNAMES            = 0x00200000
    OFN_NOCHANGEDIR          = 0x00000008
    OFN_NODEREFERENCELINKS   = 0x00100000
    OFN_NOLONGNAMES          = 0x00040000
    OFN_NONETWORKBUTTON      = 0x00020000
    OFN_NOREADONLYRETURN     = 0x00008000
    OFN_NOTESTFILECREATE     = 0x00010000
    OFN_NOVALIDATE           = 0x00000100
    OFN_OVERWRITEPROMPT      = 0x00000002
    OFN_PATHMUSTEXIST        = 0x00000800
    OFN_READONLY             = 0x00000001
    OFN_SHAREAWARE           = 0x00004000
    OFN_SHOWHELP             = 0x00000010
)

//SHBrowseForFolder flags
const (
    BIF_RETURNONLYFSDIRS    = 0x00000001
    BIF_DONTGOBELOWDOMAIN   = 0x00000002
    BIF_STATUSTEXT          = 0x00000004
    BIF_RETURNFSANCESTORS   = 0x00000008
    BIF_EDITBOX             = 0x00000010
    BIF_VALIDATE            = 0x00000020
    BIF_NEWDIALOGSTYLE      = 0x00000040
    BIF_BROWSEINCLUDEURLS   = 0x00000080
    BIF_USENEWUI            = BIF_EDITBOX | BIF_NEWDIALOGSTYLE
    BIF_UAHINT              = 0x00000100
    BIF_NONEWFOLDERBUTTON   = 0x00000200
    BIF_NOTRANSLATETARGETS  = 0x00000400
    BIF_BROWSEFORCOMPUTER   = 0x00001000
    BIF_BROWSEFORPRINTER    = 0x00002000
    BIF_BROWSEINCLUDEFILES  = 0x00004000
    BIF_SHAREABLE           = 0x00008000
    BIF_BROWSEFILEJUNCTIONS = 0x00010000
)

//MessageBox flags
const (
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND
	MB_DEFBUTTON1        = 0x00000000
	MB_DEFBUTTON2        = 0x00000100
	MB_DEFBUTTON3        = 0x00000200
	MB_DEFBUTTON4        = 0x00000300
)

//COM 
const (
    E_INVALIDARG  = 0x80070057
    E_OUTOFMEMORY = 0x8007000E
    E_UNEXPECTED  = 0x8000FFFF
)

const (
    S_OK               = 0
    S_FALSE            = 0x0001
    RPC_E_CHANGED_MODE = 0x80010106
)

// GetSystemMetrics constants
const (
	SM_CXSCREEN             = 0
	SM_CYSCREEN             = 1
	SM_CXVSCROLL            = 2
	SM_CYHSCROLL            = 3
	SM_CYCAPTION            = 4
	SM_CXBORDER             = 5
	SM_CYBORDER             = 6
	SM_CXDLGFRAME           = 7
	SM_CYDLGFRAME           = 8
	SM_CYVTHUMB             = 9
	SM_CXHTHUMB             = 10
	SM_CXICON               = 11
	SM_CYICON               = 12
	SM_CXCURSOR             = 13
	SM_CYCURSOR             = 14
	SM_CYMENU               = 15
	SM_CXFULLSCREEN         = 16
	SM_CYFULLSCREEN         = 17
	SM_CYKANJIWINDOW        = 18
	SM_MOUSEPRESENT         = 19
	SM_CYVSCROLL            = 20
	SM_CXHSCROLL            = 21
	SM_DEBUG                = 22
	SM_SWAPBUTTON           = 23
	SM_RESERVED1            = 24
	SM_RESERVED2            = 25
	SM_RESERVED3            = 26
	SM_RESERVED4            = 27
	SM_CXMIN                = 28
	SM_CYMIN                = 29
	SM_CXSIZE               = 30
	SM_CYSIZE               = 31
	SM_CXFRAME              = 32
	SM_CYFRAME              = 33
	SM_CXMINTRACK           = 34
	SM_CYMINTRACK           = 35
	SM_CXDOUBLECLK          = 36
	SM_CYDOUBLECLK          = 37
	SM_CXICONSPACING        = 38
	SM_CYICONSPACING        = 39
	SM_MENUDROPALIGNMENT    = 40
	SM_PENWINDOWS           = 41
	SM_DBCSENABLED          = 42
	SM_CMOUSEBUTTONS        = 43
	SM_CXFIXEDFRAME         = SM_CXDLGFRAME
	SM_CYFIXEDFRAME         = SM_CYDLGFRAME
	SM_CXSIZEFRAME          = SM_CXFRAME
	SM_CYSIZEFRAME          = SM_CYFRAME
	SM_SECURE               = 44
	SM_CXEDGE               = 45
	SM_CYEDGE               = 46
	SM_CXMINSPACING         = 47
	SM_CYMINSPACING         = 48
	SM_CXSMICON             = 49
	SM_CYSMICON             = 50
	SM_CYSMCAPTION          = 51
	SM_CXSMSIZE             = 52
	SM_CYSMSIZE             = 53
	SM_CXMENUSIZE           = 54
	SM_CYMENUSIZE           = 55
	SM_ARRANGE              = 56
	SM_CXMINIMIZED          = 57
	SM_CYMINIMIZED          = 58
	SM_CXMAXTRACK           = 59
	SM_CYMAXTRACK           = 60
	SM_CXMAXIMIZED          = 61
	SM_CYMAXIMIZED          = 62
	SM_NETWORK              = 63
	SM_CLEANBOOT            = 67
	SM_CXDRAG               = 68
	SM_CYDRAG               = 69
	SM_SHOWSOUNDS           = 70
	SM_CXMENUCHECK          = 71
	SM_CYMENUCHECK          = 72
	SM_SLOWMACHINE          = 73
	SM_MIDEASTENABLED       = 74
	SM_MOUSEWHEELPRESENT    = 75
	SM_XVIRTUALSCREEN       = 76
	SM_YVIRTUALSCREEN       = 77
	SM_CXVIRTUALSCREEN      = 78
	SM_CYVIRTUALSCREEN      = 79
	SM_CMONITORS            = 80
	SM_SAMEDISPLAYFORMAT    = 81
	SM_IMMENABLED           = 82
	SM_CXFOCUSBORDER        = 83
	SM_CYFOCUSBORDER        = 84
	SM_TABLETPC             = 86
	SM_MEDIACENTER          = 87
	SM_STARTER              = 88
	SM_SERVERR2             = 89
	SM_CMETRICS             = 91
	SM_REMOTESESSION        = 0x1000
	SM_SHUTTINGDOWN         = 0x2000
	SM_REMOTECONTROL        = 0x2001
	SM_CARETBLINKINGENABLED = 0x2002
)

