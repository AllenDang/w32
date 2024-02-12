package w32

// TODO Remove the de-de and en-us from the URLs.

import "unsafe"

// Windows' BOOL is really a int32, 0 is false, everything else is considered
// true. If you must set a BOOL field or parameter, use TRUE or FALSE.
//
// https://learn.microsoft.com/en-us/windows/win32/winprog/windows-data-types
const (
	TRUE  = 1
	FALSE = 0
)

// These color constants are mostly deprecated from Windows 10 onwards. They
// are still used to set a window class' background color, e.g. to
// w32.GetSysColorBrush(w32.COLOR_BTNFACE).
//
// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getsyscolor
const (
	COLOR_SCROLLBAR               = 0
	COLOR_DESKTOP                 = 1
	COLOR_BACKGROUND              = 1
	COLOR_ACTIVECAPTION           = 2
	COLOR_INACTIVECAPTION         = 3
	COLOR_MENU                    = 4
	COLOR_WINDOW                  = 5
	COLOR_WINDOWFRAME             = 6
	COLOR_MENUTEXT                = 7
	COLOR_WINDOWTEXT              = 8
	COLOR_CAPTIONTEXT             = 9
	COLOR_ACTIVEBORDER            = 10
	COLOR_INACTIVEBORDER          = 11
	COLOR_APPWORKSPACE            = 12
	COLOR_HIGHLIGHT               = 13
	COLOR_HIGHLIGHTTEXT           = 14
	COLOR_3DFACE                  = 15
	COLOR_BTNFACE                 = 15
	COLOR_BTNSHADOW               = 16
	COLOR_3DSHADOW                = 16
	COLOR_GRAYTEXT                = 17
	COLOR_BTNTEXT                 = 18
	COLOR_INACTIVECAPTIONTEXT     = 19
	COLOR_3DHIGHLIGHT             = 20
	COLOR_3DHILIGHT               = 20
	COLOR_BTNHIGHLIGHT            = 20
	COLOR_BTNHILIGHT              = 20
	COLOR_3DDKSHADOW              = 21
	COLOR_3DLIGHT                 = 22
	COLOR_INFOTEXT                = 23
	COLOR_INFOBK                  = 24
	COLOR_HOTLIGHT                = 26
	COLOR_GRADIENTACTIVECAPTION   = 27
	COLOR_GRADIENTINACTIVECAPTION = 28
	COLOR_MENUHILIGHT             = 29
	COLOR_MENUBAR                 = 30
)

// https://learn.microsoft.com/en-us/windows/win32/winmsg/window-class-styles
const (
	CS_VREDRAW         = 1
	CS_HREDRAW         = 2
	CS_DBLCLKS         = 8
	CS_OWNDC           = 32
	CS_CLASSDC         = 64
	CS_PARENTDC        = 128
	CS_NOCLOSE         = 512
	CS_SAVEBITS        = 2048
	CS_BYTEALIGNCLIENT = 4096
	CS_BYTEALIGNWINDOW = 8192
	CS_GLOBALCLASS     = 16384
	CS_IME             = 65536
	CS_DROPSHADOW      = 131072
)

// Window messages are sent to the window class' window procedure. They
// identify the type of event that happened.
//
// https://learn.microsoft.com/en-us/windows/win32/learnwin32/window-messages
const (
	WM_CTLCOLOR                                  = 25
	WM_MOUSEHOVER                                = 673
	WM_MOUSELEAVE                                = 675
	WM_CAP_START                                 = 1024
	WM_CAP_UNICODE_START                         = 1124
	WM_CAP_GET_CAPSTREAMPTR                      = 1025
	WM_CAP_SET_CALLBACK_ERRORW                   = 1126
	WM_CAP_SET_CALLBACK_STATUSW                  = 1127
	WM_CAP_SET_CALLBACK_ERRORA                   = 1026
	WM_CAP_SET_CALLBACK_STATUSA                  = 1027
	WM_CAP_SET_CALLBACK_ERROR                    = 1126
	WM_CAP_SET_CALLBACK_STATUS                   = 1127
	WM_CAP_SET_CALLBACK_YIELD                    = 1028
	WM_CAP_SET_CALLBACK_FRAME                    = 1029
	WM_CAP_SET_CALLBACK_VIDEOSTREAM              = 1030
	WM_CAP_SET_CALLBACK_WAVESTREAM               = 1031
	WM_CAP_GET_USER_DATA                         = 1032
	WM_CAP_SET_USER_DATA                         = 1033
	WM_CAP_DRIVER_CONNECT                        = 1034
	WM_CAP_DRIVER_DISCONNECT                     = 1035
	WM_CAP_DRIVER_GET_NAMEA                      = 1036
	WM_CAP_DRIVER_GET_VERSIONA                   = 1037
	WM_CAP_DRIVER_GET_NAMEW                      = 1136
	WM_CAP_DRIVER_GET_VERSIONW                   = 1137
	WM_CAP_DRIVER_GET_NAME                       = 1136
	WM_CAP_DRIVER_GET_VERSION                    = 1137
	WM_CAP_DRIVER_GET_CAPS                       = 1038
	WM_CAP_FILE_SET_CAPTURE_FILEA                = 1044
	WM_CAP_FILE_GET_CAPTURE_FILEA                = 1045
	WM_CAP_FILE_SAVEASA                          = 1047
	WM_CAP_FILE_SAVEDIBA                         = 1049
	WM_CAP_FILE_SET_CAPTURE_FILEW                = 1144
	WM_CAP_FILE_GET_CAPTURE_FILEW                = 1145
	WM_CAP_FILE_SAVEASW                          = 1147
	WM_CAP_FILE_SAVEDIBW                         = 1149
	WM_CAP_FILE_SET_CAPTURE_FILE                 = 1144
	WM_CAP_FILE_GET_CAPTURE_FILE                 = 1145
	WM_CAP_FILE_SAVEAS                           = 1147
	WM_CAP_FILE_SAVEDIB                          = 1149
	WM_CAP_FILE_ALLOCATE                         = 1046
	WM_CAP_FILE_SET_INFOCHUNK                    = 1048
	WM_CAP_EDIT_COPY                             = 1054
	WM_CAP_SET_AUDIOFORMAT                       = 1059
	WM_CAP_GET_AUDIOFORMAT                       = 1060
	WM_CAP_DLG_VIDEOFORMAT                       = 1065
	WM_CAP_DLG_VIDEOSOURCE                       = 1066
	WM_CAP_DLG_VIDEODISPLAY                      = 1067
	WM_CAP_GET_VIDEOFORMAT                       = 1068
	WM_CAP_SET_VIDEOFORMAT                       = 1069
	WM_CAP_DLG_VIDEOCOMPRESSION                  = 1070
	WM_CAP_SET_PREVIEW                           = 1074
	WM_CAP_SET_OVERLAY                           = 1075
	WM_CAP_SET_PREVIEWRATE                       = 1076
	WM_CAP_SET_SCALE                             = 1077
	WM_CAP_GET_STATUS                            = 1078
	WM_CAP_SET_SCROLL                            = 1079
	WM_CAP_GRAB_FRAME                            = 1084
	WM_CAP_GRAB_FRAME_NOSTOP                     = 1085
	WM_CAP_SEQUENCE                              = 1086
	WM_CAP_SEQUENCE_NOFILE                       = 1087
	WM_CAP_SET_SEQUENCE_SETUP                    = 1088
	WM_CAP_GET_SEQUENCE_SETUP                    = 1089
	WM_CAP_SET_MCI_DEVICEA                       = 1090
	WM_CAP_GET_MCI_DEVICEA                       = 1091
	WM_CAP_SET_MCI_DEVICEW                       = 1190
	WM_CAP_GET_MCI_DEVICEW                       = 1191
	WM_CAP_SET_MCI_DEVICE                        = 1190
	WM_CAP_GET_MCI_DEVICE                        = 1191
	WM_CAP_STOP                                  = 1092
	WM_CAP_ABORT                                 = 1093
	WM_CAP_SINGLE_FRAME_OPEN                     = 1094
	WM_CAP_SINGLE_FRAME_CLOSE                    = 1095
	WM_CAP_SINGLE_FRAME                          = 1096
	WM_CAP_PAL_OPENA                             = 1104
	WM_CAP_PAL_SAVEA                             = 1105
	WM_CAP_PAL_OPENW                             = 1204
	WM_CAP_PAL_SAVEW                             = 1205
	WM_CAP_PAL_OPEN                              = 1204
	WM_CAP_PAL_SAVE                              = 1205
	WM_CAP_PAL_PASTE                             = 1106
	WM_CAP_PAL_AUTOCREATE                        = 1107
	WM_CAP_PAL_MANUALCREATE                      = 1108
	WM_CAP_SET_CALLBACK_CAPCONTROL               = 1109
	WM_CAP_UNICODE_END                           = 1205
	WM_CAP_END                                   = 1205
	WM_CPL_LAUNCH                                = 2024
	WM_CPL_LAUNCHED                              = 2025
	WM_TABLET_DEFBASE                            = 704
	WM_TABLET_MAXOFFSET                          = 32
	WM_TABLET_ADDED                              = 712
	WM_TABLET_DELETED                            = 713
	WM_TABLET_FLICK                              = 715
	WM_TABLET_QUERYSYSTEMGESTURESTATUS           = 716
	WM_FI_FILENAME                               = 900
	WM_SampleExtension_ContentType_Size          = 1
	WM_SampleExtension_PixelAspectRatio_Size     = 2
	WM_SampleExtension_Timecode_Size             = 14
	WM_SampleExtension_SampleDuration_Size       = 2
	WM_SampleExtension_ChromaLocation_Size       = 1
	WM_SampleExtension_ColorSpaceInfo_Size       = 3
	WM_CT_REPEAT_FIRST_FIELD                     = 16
	WM_CT_BOTTOM_FIELD_FIRST                     = 32
	WM_CT_TOP_FIELD_FIRST                        = 64
	WM_CT_INTERLACED                             = 128
	WM_CL_INTERLACED420                          = 0
	WM_CL_PROGRESSIVE420                         = 1
	WM_MAX_VIDEO_STREAMS                         = 63
	WM_MAX_STREAMS                               = 63
	WM_ADSPROP_NOTIFY_PAGEINIT                   = 2125
	WM_ADSPROP_NOTIFY_PAGEHWND                   = 2126
	WM_ADSPROP_NOTIFY_CHANGE                     = 2127
	WM_ADSPROP_NOTIFY_APPLY                      = 2128
	WM_ADSPROP_NOTIFY_SETFOCUS                   = 2129
	WM_ADSPROP_NOTIFY_FOREGROUND                 = 2130
	WM_ADSPROP_NOTIFY_EXIT                       = 2131
	WM_ADSPROP_NOTIFY_ERROR                      = 2134
	WM_RASDIALEVENT                              = 52429
	WM_DDE_FIRST                                 = 992
	WM_DDE_INITIATE                              = 992
	WM_DDE_TERMINATE                             = 993
	WM_DDE_ADVISE                                = 994
	WM_DDE_UNADVISE                              = 995
	WM_DDE_ACK                                   = 996
	WM_DDE_DATA                                  = 997
	WM_DDE_REQUEST                               = 998
	WM_DDE_POKE                                  = 999
	WM_DDE_EXECUTE                               = 1000
	WM_DDE_LAST                                  = 1000
	WM_IME_REPORT                                = 640
	WM_WNT_CONVERTREQUESTEX                      = 265
	WM_CONVERTREQUEST                            = 266
	WM_CONVERTRESULT                             = 267
	WM_INTERIM                                   = 268
	WM_IMEKEYDOWN                                = 656
	WM_IMEKEYUP                                  = 657
	WM_CHOOSEFONT_GETLOGFONT                     = 1025
	WM_CHOOSEFONT_SETLOGFONT                     = 1125
	WM_CHOOSEFONT_SETFLAGS                       = 1126
	WM_PSD_FULLPAGERECT                          = 1025
	WM_PSD_MINMARGINRECT                         = 1026
	WM_PSD_MARGINRECT                            = 1027
	WM_PSD_GREEKTEXTRECT                         = 1028
	WM_PSD_ENVSTAMPRECT                          = 1029
	WM_PSD_YAFULLPAGERECT                        = 1030
	WM_CONTEXTMENU                               = 123
	WM_UNICHAR                                   = 265
	WM_PRINTCLIENT                               = 792
	WM_NOTIFY                                    = 78
	WM_DEVICECHANGE                              = 537
	WM_NULL                                      = 0
	WM_CREATE                                    = 1
	WM_DESTROY                                   = 2
	WM_MOVE                                      = 3
	WM_SIZE                                      = 5
	WM_ACTIVATE                                  = 6
	WM_SETFOCUS                                  = 7
	WM_KILLFOCUS                                 = 8
	WM_ENABLE                                    = 10
	WM_SETREDRAW                                 = 11
	WM_SETTEXT                                   = 12
	WM_GETTEXT                                   = 13
	WM_GETTEXTLENGTH                             = 14
	WM_PAINT                                     = 15
	WM_CLOSE                                     = 16
	WM_QUERYENDSESSION                           = 17
	WM_QUERYOPEN                                 = 19
	WM_ENDSESSION                                = 22
	WM_QUIT                                      = 18
	WM_ERASEBKGND                                = 20
	WM_SYSCOLORCHANGE                            = 21
	WM_SHOWWINDOW                                = 24
	WM_WININICHANGE                              = 26
	WM_SETTINGCHANGE                             = 26
	WM_DEVMODECHANGE                             = 27
	WM_ACTIVATEAPP                               = 28
	WM_FONTCHANGE                                = 29
	WM_TIMECHANGE                                = 30
	WM_CANCELMODE                                = 31
	WM_SETCURSOR                                 = 32
	WM_MOUSEACTIVATE                             = 33
	WM_CHILDACTIVATE                             = 34
	WM_QUEUESYNC                                 = 35
	WM_GETMINMAXINFO                             = 36
	WM_PAINTICON                                 = 38
	WM_ICONERASEBKGND                            = 39
	WM_NEXTDLGCTL                                = 40
	WM_SPOOLERSTATUS                             = 42
	WM_DRAWITEM                                  = 43
	WM_MEASUREITEM                               = 44
	WM_DELETEITEM                                = 45
	WM_VKEYTOITEM                                = 46
	WM_CHARTOITEM                                = 47
	WM_SETFONT                                   = 48
	WM_GETFONT                                   = 49
	WM_SETHOTKEY                                 = 50
	WM_GETHOTKEY                                 = 51
	WM_QUERYDRAGICON                             = 55
	WM_COMPAREITEM                               = 57
	WM_GETOBJECT                                 = 61
	WM_COMPACTING                                = 65
	WM_COMMNOTIFY                                = 68
	WM_WINDOWPOSCHANGING                         = 70
	WM_WINDOWPOSCHANGED                          = 71
	WM_POWER                                     = 72
	WM_COPYDATA                                  = 74
	WM_CANCELJOURNAL                             = 75
	WM_INPUTLANGCHANGEREQUEST                    = 80
	WM_INPUTLANGCHANGE                           = 81
	WM_TCARD                                     = 82
	WM_HELP                                      = 83
	WM_USERCHANGED                               = 84
	WM_NOTIFYFORMAT                              = 85
	WM_STYLECHANGING                             = 124
	WM_STYLECHANGED                              = 125
	WM_DISPLAYCHANGE                             = 126
	WM_GETICON                                   = 127
	WM_SETICON                                   = 128
	WM_NCCREATE                                  = 129
	WM_NCDESTROY                                 = 130
	WM_NCCALCSIZE                                = 131
	WM_NCHITTEST                                 = 132
	WM_NCPAINT                                   = 133
	WM_NCACTIVATE                                = 134
	WM_GETDLGCODE                                = 135
	WM_SYNCPAINT                                 = 136
	WM_NCMOUSEMOVE                               = 160
	WM_NCLBUTTONDOWN                             = 161
	WM_NCLBUTTONUP                               = 162
	WM_NCLBUTTONDBLCLK                           = 163
	WM_NCRBUTTONDOWN                             = 164
	WM_NCRBUTTONUP                               = 165
	WM_NCRBUTTONDBLCLK                           = 166
	WM_NCMBUTTONDOWN                             = 167
	WM_NCMBUTTONUP                               = 168
	WM_NCMBUTTONDBLCLK                           = 169
	WM_NCXBUTTONDOWN                             = 171
	WM_NCXBUTTONUP                               = 172
	WM_NCXBUTTONDBLCLK                           = 173
	WM_INPUT_DEVICE_CHANGE                       = 254
	WM_INPUT                                     = 255
	WM_KEYFIRST                                  = 256
	WM_KEYDOWN                                   = 256
	WM_KEYUP                                     = 257
	WM_CHAR                                      = 258
	WM_DEADCHAR                                  = 259
	WM_SYSKEYDOWN                                = 260
	WM_SYSKEYUP                                  = 261
	WM_SYSCHAR                                   = 262
	WM_SYSDEADCHAR                               = 263
	WM_KEYLAST                                   = 265
	WM_IME_STARTCOMPOSITION                      = 269
	WM_IME_ENDCOMPOSITION                        = 270
	WM_IME_COMPOSITION                           = 271
	WM_IME_KEYLAST                               = 271
	WM_INITDIALOG                                = 272
	WM_COMMAND                                   = 273
	WM_SYSCOMMAND                                = 274
	WM_TIMER                                     = 275
	WM_HSCROLL                                   = 276
	WM_VSCROLL                                   = 277
	WM_INITMENU                                  = 278
	WM_INITMENUPOPUP                             = 279
	WM_GESTURE                                   = 281
	WM_GESTURENOTIFY                             = 282
	WM_MENUSELECT                                = 287
	WM_MENUCHAR                                  = 288
	WM_ENTERIDLE                                 = 289
	WM_MENURBUTTONUP                             = 290
	WM_MENUDRAG                                  = 291
	WM_MENUGETOBJECT                             = 292
	WM_UNINITMENUPOPUP                           = 293
	WM_MENUCOMMAND                               = 294
	WM_CHANGEUISTATE                             = 295
	WM_UPDATEUISTATE                             = 296
	WM_QUERYUISTATE                              = 297
	WM_CTLCOLORMSGBOX                            = 306
	WM_CTLCOLOREDIT                              = 307
	WM_CTLCOLORLISTBOX                           = 308
	WM_CTLCOLORBTN                               = 309
	WM_CTLCOLORDLG                               = 310
	WM_CTLCOLORSCROLLBAR                         = 311
	WM_CTLCOLORSTATIC                            = 312
	WM_MOUSEFIRST                                = 512
	WM_MOUSEMOVE                                 = 512
	WM_LBUTTONDOWN                               = 513
	WM_LBUTTONUP                                 = 514
	WM_LBUTTONDBLCLK                             = 515
	WM_RBUTTONDOWN                               = 516
	WM_RBUTTONUP                                 = 517
	WM_RBUTTONDBLCLK                             = 518
	WM_MBUTTONDOWN                               = 519
	WM_MBUTTONUP                                 = 520
	WM_MBUTTONDBLCLK                             = 521
	WM_MOUSEWHEEL                                = 522
	WM_XBUTTONDOWN                               = 523
	WM_XBUTTONUP                                 = 524
	WM_XBUTTONDBLCLK                             = 525
	WM_MOUSEHWHEEL                               = 526
	WM_MOUSELAST                                 = 526
	WM_PARENTNOTIFY                              = 528
	WM_ENTERMENULOOP                             = 529
	WM_EXITMENULOOP                              = 530
	WM_NEXTMENU                                  = 531
	WM_SIZING                                    = 532
	WM_CAPTURECHANGED                            = 533
	WM_MOVING                                    = 534
	WM_POWERBROADCAST                            = 536
	WM_MDICREATE                                 = 544
	WM_MDIDESTROY                                = 545
	WM_MDIACTIVATE                               = 546
	WM_MDIRESTORE                                = 547
	WM_MDINEXT                                   = 548
	WM_MDIMAXIMIZE                               = 549
	WM_MDITILE                                   = 550
	WM_MDICASCADE                                = 551
	WM_MDIICONARRANGE                            = 552
	WM_MDIGETACTIVE                              = 553
	WM_MDISETMENU                                = 560
	WM_ENTERSIZEMOVE                             = 561
	WM_EXITSIZEMOVE                              = 562
	WM_DROPFILES                                 = 563
	WM_MDIREFRESHMENU                            = 564
	WM_POINTERDEVICECHANGE                       = 568
	WM_POINTERDEVICEINRANGE                      = 569
	WM_POINTERDEVICEOUTOFRANGE                   = 570
	WM_TOUCH                                     = 576
	WM_NCPOINTERUPDATE                           = 577
	WM_NCPOINTERDOWN                             = 578
	WM_NCPOINTERUP                               = 579
	WM_POINTERUPDATE                             = 581
	WM_POINTERDOWN                               = 582
	WM_POINTERUP                                 = 583
	WM_POINTERENTER                              = 585
	WM_POINTERLEAVE                              = 586
	WM_POINTERACTIVATE                           = 587
	WM_POINTERCAPTURECHANGED                     = 588
	WM_TOUCHHITTESTING                           = 589
	WM_POINTERWHEEL                              = 590
	WM_POINTERHWHEEL                             = 591
	WM_POINTERROUTEDTO                           = 593
	WM_POINTERROUTEDAWAY                         = 594
	WM_POINTERROUTEDRELEASED                     = 595
	WM_IME_SETCONTEXT                            = 641
	WM_IME_NOTIFY                                = 642
	WM_IME_CONTROL                               = 643
	WM_IME_COMPOSITIONFULL                       = 644
	WM_IME_SELECT                                = 645
	WM_IME_CHAR                                  = 646
	WM_IME_REQUEST                               = 648
	WM_IME_KEYDOWN                               = 656
	WM_IME_KEYUP                                 = 657
	WM_NCMOUSEHOVER                              = 672
	WM_NCMOUSELEAVE                              = 674
	WM_WTSSESSION_CHANGE                         = 689
	WM_TABLET_FIRST                              = 704
	WM_TABLET_LAST                               = 735
	WM_DPICHANGED                                = 736
	WM_DPICHANGED_BEFOREPARENT                   = 738
	WM_DPICHANGED_AFTERPARENT                    = 739
	WM_GETDPISCALEDSIZE                          = 740
	WM_CUT                                       = 768
	WM_COPY                                      = 769
	WM_PASTE                                     = 770
	WM_CLEAR                                     = 771
	WM_UNDO                                      = 772
	WM_RENDERFORMAT                              = 773
	WM_RENDERALLFORMATS                          = 774
	WM_DESTROYCLIPBOARD                          = 775
	WM_DRAWCLIPBOARD                             = 776
	WM_PAINTCLIPBOARD                            = 777
	WM_VSCROLLCLIPBOARD                          = 778
	WM_SIZECLIPBOARD                             = 779
	WM_ASKCBFORMATNAME                           = 780
	WM_CHANGECBCHAIN                             = 781
	WM_HSCROLLCLIPBOARD                          = 782
	WM_QUERYNEWPALETTE                           = 783
	WM_PALETTEISCHANGING                         = 784
	WM_PALETTECHANGED                            = 785
	WM_HOTKEY                                    = 786
	WM_PRINT                                     = 791
	WM_APPCOMMAND                                = 793
	WM_THEMECHANGED                              = 794
	WM_CLIPBOARDUPDATE                           = 797
	WM_DWMCOMPOSITIONCHANGED                     = 798
	WM_DWMNCRENDERINGCHANGED                     = 799
	WM_DWMCOLORIZATIONCOLORCHANGED               = 800
	WM_DWMWINDOWMAXIMIZEDCHANGE                  = 801
	WM_DWMSENDICONICTHUMBNAIL                    = 803
	WM_DWMSENDICONICLIVEPREVIEWBITMAP            = 806
	WM_GETTITLEBARINFOEX                         = 831
	WM_HANDHELDFIRST                             = 856
	WM_HANDHELDLAST                              = 863
	WM_AFXFIRST                                  = 864
	WM_AFXLAST                                   = 895
	WM_PENWINFIRST                               = 896
	WM_PENWINLAST                                = 911
	WM_APP                                       = 32768
	WM_USER                                      = 1024
	WM_TOOLTIPDISMISS                            = 837
	WM_SF_CLEANPOINT                             = 1
	WM_SF_DISCONTINUITY                          = 2
	WM_SF_DATALOSS                               = 4
	WM_SFEX_NOTASYNCPOINT                        = 2
	WM_SFEX_DATALOSS                             = 4
	WM_DM_NOTINTERLACED                          = 0
	WM_DM_DEINTERLACE_NORMAL                     = 1
	WM_DM_DEINTERLACE_HALFSIZE                   = 2
	WM_DM_DEINTERLACE_HALFSIZEDOUBLERATE         = 3
	WM_DM_DEINTERLACE_INVERSETELECINE            = 4
	WM_DM_DEINTERLACE_VERTICALHALFSIZEDOUBLERATE = 5
	WM_DM_IT_DISABLE_COHERENT_MODE               = 0
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_AA_TOP       = 1
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_BB_TOP       = 2
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_BC_TOP       = 3
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_CD_TOP       = 4
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_DD_TOP       = 5
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_AA_BOTTOM    = 6
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_BB_BOTTOM    = 7
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_BC_BOTTOM    = 8
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_CD_BOTTOM    = 9
	WM_DM_IT_FIRST_FRAME_IN_CLIP_IS_DD_BOTTOM    = 10
	WM_PLAYBACK_DRC_HIGH                         = 0
	WM_PLAYBACK_DRC_MEDIUM                       = 1
	WM_PLAYBACK_DRC_LOW                          = 2
	WM_AETYPE_INCLUDE                            = 105
	WM_AETYPE_EXCLUDE                            = 101
)

// CW_USEDEFAULT can be used for x, y, width and height in CreateWindowEx to
// let Windows decide the position and size of a window.
const CW_USEDEFAULT = -2147483648

// Window styles can be used in CreateWindowEx to give a window specific
// properties, like borders and min/max buttons.
//
// https://learn.microsoft.com/en-us/windows/win32/winmsg/window-styles
const (
	WS_OVERLAPPED       = 0
	WS_POPUP            = 2147483648
	WS_CHILD            = 1073741824
	WS_MINIMIZE         = 536870912
	WS_VISIBLE          = 268435456
	WS_DISABLED         = 134217728
	WS_CLIPSIBLINGS     = 67108864
	WS_CLIPCHILDREN     = 33554432
	WS_MAXIMIZE         = 16777216
	WS_CAPTION          = 12582912
	WS_BORDER           = 8388608
	WS_DLGFRAME         = 4194304
	WS_VSCROLL          = 2097152
	WS_HSCROLL          = 1048576
	WS_SYSMENU          = 524288
	WS_THICKFRAME       = 262144
	WS_GROUP            = 131072
	WS_TABSTOP          = 65536
	WS_MINIMIZEBOX      = 131072
	WS_MAXIMIZEBOX      = 65536
	WS_TILED            = 0
	WS_ICONIC           = 536870912
	WS_SIZEBOX          = 262144
	WS_TILEDWINDOW      = 13565952
	WS_OVERLAPPEDWINDOW = 13565952
	WS_POPUPWINDOW      = 2156396544
	WS_CHILDWINDOW      = 1073741824
	WS_ACTIVECAPTION    = 1
)

// Extended window styles are used for the first parameter to CreateWindowEx.
// They allow more control over the window properties, in addition to the
// non-extended window styles.
//
// https://learn.microsoft.com/en-us/windows/win32/winmsg/extended-window-styles
const (
	WS_EX_DLGMODALFRAME       = 1
	WS_EX_NOPARENTNOTIFY      = 4
	WS_EX_TOPMOST             = 8
	WS_EX_ACCEPTFILES         = 16
	WS_EX_TRANSPARENT         = 32
	WS_EX_MDICHILD            = 64
	WS_EX_TOOLWINDOW          = 128
	WS_EX_WINDOWEDGE          = 256
	WS_EX_CLIENTEDGE          = 512
	WS_EX_CONTEXTHELP         = 1024
	WS_EX_RIGHT               = 4096
	WS_EX_LEFT                = 0
	WS_EX_RTLREADING          = 8192
	WS_EX_LTRREADING          = 0
	WS_EX_LEFTSCROLLBAR       = 16384
	WS_EX_RIGHTSCROLLBAR      = 0
	WS_EX_CONTROLPARENT       = 65536
	WS_EX_STATICEDGE          = 131072
	WS_EX_APPWINDOW           = 262144
	WS_EX_OVERLAPPEDWINDOW    = 768
	WS_EX_PALETTEWINDOW       = 392
	WS_EX_LAYERED             = 524288
	WS_EX_NOINHERITLAYOUT     = 1048576
	WS_EX_NOREDIRECTIONBITMAP = 2097152
	WS_EX_LAYOUTRTL           = 4194304
	WS_EX_COMPOSITED          = 33554432
	WS_EX_NOACTIVATE          = 134217728
)

// PeekMessage takes these constants as its last parameter. They define what
// messages are peeked and whether they remain in the queue or not.
//
// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-peekmessagew
const (
	PM_NOREMOVE       = 0
	PM_REMOVE         = 1
	PM_NOYIELD        = 2
	PM_QS_INPUT       = 67567616
	PM_QS_POSTMESSAGE = 9961472
	PM_QS_PAINT       = 2097152
	PM_QS_SENDMESSAGE = 4194304
)

// These cursor constants can be used in LoadCursor, e.g.
//
//	w32.LoadCursor(0, w32.MakeIntResource(w32.IDC_ARROW))
//
// https://learn.microsoft.com/en-us/windows/win32/menurc/about-cursors
const (
	IDC_ARROW       = 32512
	IDC_IBEAM       = 32513
	IDC_WAIT        = 32514
	IDC_CROSS       = 32515
	IDC_UPARROW     = 32516
	IDC_SIZE        = 32640
	IDC_ICON        = 32641
	IDC_SIZENWSE    = 32642
	IDC_SIZENESW    = 32643
	IDC_SIZEWE      = 32644
	IDC_SIZENS      = 32645
	IDC_SIZEALL     = 32646
	IDC_NO          = 32648
	IDC_HAND        = 32649
	IDC_APPSTARTING = 32650
	IDC_HELP        = 32651
	IDC_PIN         = 32671
	IDC_PERSON      = 32672
)

// These icon constants can be used in LoadIcon, e.g.:
//
//	w32.LoadIcon(0, w32.MakeIntResource(w32.IDI_SHIELD))
//
// https://learn.microsoft.com/en-us/windows/win32/menurc/about-icons
const (
	IDI_APPLICATION = 32512
	IDI_ERROR       = 32513
	IDI_QUESTION    = 32514
	IDI_WARNING     = 32515
	IDI_INFORMATION = 32516
	IDI_WINLOGO     = 32517
	IDI_SHIELD      = 32518
)

// These constants are used for the last parameter to LoadImage. They determine
// the way an image is loaded.
//
// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadimagew
const (
	LR_DEFAULTCOLOR     = 0x0000
	LR_MONOCHROME       = 0x0001
	LR_COLOR            = 0x0002
	LR_COPYRETURNORG    = 0x0004
	LR_COPYDELETEORG    = 0x0008
	LR_LOADFROMFILE     = 0x0010
	LR_LOADTRANSPARENT  = 0x0020
	LR_DEFAULTSIZE      = 0x0040
	LR_VGACOLOR         = 0x0080
	LR_LOADMAP3DCOLORS  = 0x1000
	LR_CREATEDIBSECTION = 0x2000
	LR_COPYFROMRESOURCE = 0x4000
	LR_SHARED           = 0x8000
)

// These constants determine the type of image to load in LoadImage.
//
// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-loadimagew
const (
	IMAGE_BITMAP = 0
	IMAGE_ICON   = 1
	IMAGE_CURSOR = 2
)

// https://learn.microsoft.com/en-us/windows/win32/api/wingdi/nf-wingdi-selectobject
const (
	NULLREGION    = 1
	SIMPLEREGION  = 2
	COMPLEXREGION = 3
)

const HGDI_ERROR = 65535

// https://learn.microsoft.com/en-us/windows/win32/controls/button-styles
const (
	BS_3STATE          = 5
	BS_AUTO3STATE      = 6
	BS_AUTOCHECKBOX    = 3
	BS_AUTORADIOBUTTON = 9
	BS_BITMAP          = 128
	BS_BOTTOM          = 2048
	BS_CENTER          = 768
	BS_CHECKBOX        = 2
	BS_DEFPUSHBUTTON   = 1
	BS_GROUPBOX        = 7
	BS_ICON            = 64
	BS_LEFT            = 256
	BS_LEFTTEXT        = 32
	BS_MULTILINE       = 8192
	BS_NOTIFY          = 16384
	BS_OWNERDRAW       = 11
	BS_PUSHBUTTON      = 0
	BS_PUSHLIKE        = 4096
	BS_RADIOBUTTON     = 4
	BS_RIGHT           = 512
	BS_RIGHTBUTTON     = 32
	BS_TEXT            = 0
	BS_TOP             = 1024
	BS_USERBUTTON      = 8
	BS_VCENTER         = 3072
	BS_FLAT            = 32768
)

var (
	WC_TREEVIEW    = String("SysTreeView32")
	MSFTEDIT_CLASS = String("RICHEDIT50W")
)

const (
	ICC_ANIMATE_CLASS      = 128
	ICC_BAR_CLASSES        = 4
	ICC_COOL_CLASSES       = 1024
	ICC_DATE_CLASSES       = 256
	ICC_HOTKEY_CLASS       = 64
	ICC_INTERNET_CLASSES   = 2048
	ICC_LINK_CLASS         = 32768
	ICC_LISTVIEW_CLASSES   = 1
	ICC_NATIVEFNTCTL_CLASS = 8192
	ICC_PAGESCROLLER_CLASS = 4096
	ICC_PROGRESS_CLASS     = 32
	ICC_STANDARD_CLASSES   = 16384
	ICC_TAB_CLASSES        = 8
	ICC_TREEVIEW_CLASSES   = 2
	ICC_UPDOWN_CLASS       = 16
	ICC_USEREX_CLASSES     = 512
	ICC_WIN95_CLASSES      = 255
)

const (
	I_ZERO             = 0
	I_ONE_OR_MORE      = 1
	I_CHILDRENCALLBACK = -1
	I_CHILDRENAUTO     = -2
)

const (
	I_IMAGECALLBACK = -1
)

const INVALID_HANDLE_VALUE = ^HANDLE(0)

var LPSTR_TEXTCALLBACK = UTF16String(unsafe.Pointer(^uintptr(0)))

// https://learn.microsoft.com/en-us/windows/win32/controls/ilc-constants
const (
	ILC_MASK             = 1
	ILC_COLOR            = 0
	ILC_COLORDDB         = 254
	ILC_COLOR4           = 4
	ILC_COLOR8           = 8
	ILC_COLOR16          = 16
	ILC_COLOR24          = 24
	ILC_COLOR32          = 32
	ILC_PALETTE          = 2048
	ILC_MIRROR           = 8192
	ILC_PERITEMMIRROR    = 32768
	ILC_ORIGINALSIZE     = 65536
	ILC_HIGHQUALITYSCALE = 131072
)

// https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_draw
const (
	ILD_NORMAL        = 0
	ILD_TRANSPARENT   = 1
	ILD_BLEND25       = 2
	ILD_FOCUS         = 2
	ILD_BLEND50       = 4
	ILD_SELECTED      = 4
	ILD_BLEND         = 4
	ILD_MASK          = 16
	ILD_IMAGE         = 32
	ILD_ROP           = 64
	ILD_OVERLAYMASK   = 3840
	ILD_PRESERVEALPHA = 4096
	ILD_SCALE         = 8192
	ILD_DPISCALE      = 16384
	ILD_ASYNC         = 32768
)

// https://learn.microsoft.com/en-us/windows/win32/api/commctrl/nf-commctrl-imagelist_drawex
const (
	CLR_NONE    = 0xFFFFFFFF
	CLR_INVALID = 0xFFFFFFFF
	CLR_DEFAULT = 0xFF000000
	CLR_HILIGHT = 0xFF000000
)

const (
	TVSIL_NORMAL = 0
	TVSIL_STATE  = 2
)

// https://learn.microsoft.com/windows/win32/inputdev/virtual-key-codes
const (
	VK_0                               = '0'
	VK_1                               = '1'
	VK_2                               = '2'
	VK_3                               = '3'
	VK_4                               = '4'
	VK_5                               = '5'
	VK_6                               = '6'
	VK_7                               = '7'
	VK_8                               = '8'
	VK_9                               = '9'
	VK_A                               = 'A'
	VK_B                               = 'B'
	VK_C                               = 'C'
	VK_D                               = 'D'
	VK_E                               = 'E'
	VK_F                               = 'F'
	VK_G                               = 'G'
	VK_H                               = 'H'
	VK_I                               = 'I'
	VK_J                               = 'J'
	VK_K                               = 'K'
	VK_L                               = 'L'
	VK_M                               = 'M'
	VK_N                               = 'N'
	VK_O                               = 'O'
	VK_P                               = 'P'
	VK_Q                               = 'Q'
	VK_R                               = 'R'
	VK_S                               = 'S'
	VK_T                               = 'T'
	VK_U                               = 'U'
	VK_V                               = 'V'
	VK_W                               = 'W'
	VK_X                               = 'X'
	VK_Y                               = 'Y'
	VK_Z                               = 'Z'
	VK_ABNT_C1                         = 193
	VK_ABNT_C2                         = 194
	VK_DBE_ALPHANUMERIC                = 240
	VK_DBE_CODEINPUT                   = 250
	VK_DBE_DBCSCHAR                    = 244
	VK_DBE_DETERMINESTRING             = 252
	VK_DBE_ENTERDLGCONVERSIONMODE      = 253
	VK_DBE_ENTERIMECONFIGMODE          = 248
	VK_DBE_ENTERWORDREGISTERMODE       = 247
	VK_DBE_FLUSHSTRING                 = 249
	VK_DBE_HIRAGANA                    = 242
	VK_DBE_KATAKANA                    = 241
	VK_DBE_NOCODEINPUT                 = 251
	VK_DBE_NOROMAN                     = 246
	VK_DBE_ROMAN                       = 245
	VK_DBE_SBCSCHAR                    = 243
	VK__none_                          = 255
	VK_LBUTTON                         = 1
	VK_RBUTTON                         = 2
	VK_CANCEL                          = 3
	VK_MBUTTON                         = 4
	VK_XBUTTON1                        = 5
	VK_XBUTTON2                        = 6
	VK_BACK                            = 8
	VK_TAB                             = 9
	VK_CLEAR                           = 12
	VK_RETURN                          = 13
	VK_SHIFT                           = 16
	VK_CONTROL                         = 17
	VK_MENU                            = 18
	VK_PAUSE                           = 19
	VK_CAPITAL                         = 20
	VK_KANA                            = 21
	VK_HANGEUL                         = 21
	VK_HANGUL                          = 21
	VK_IME_ON                          = 22
	VK_JUNJA                           = 23
	VK_FINAL                           = 24
	VK_HANJA                           = 25
	VK_KANJI                           = 25
	VK_IME_OFF                         = 26
	VK_ESCAPE                          = 27
	VK_CONVERT                         = 28
	VK_NONCONVERT                      = 29
	VK_ACCEPT                          = 30
	VK_MODECHANGE                      = 31
	VK_SPACE                           = 32
	VK_PRIOR                           = 33
	VK_NEXT                            = 34
	VK_END                             = 35
	VK_HOME                            = 36
	VK_LEFT                            = 37
	VK_UP                              = 38
	VK_RIGHT                           = 39
	VK_DOWN                            = 40
	VK_SELECT                          = 41
	VK_PRINT                           = 42
	VK_EXECUTE                         = 43
	VK_SNAPSHOT                        = 44
	VK_INSERT                          = 45
	VK_DELETE                          = 46
	VK_HELP                            = 47
	VK_LWIN                            = 91
	VK_RWIN                            = 92
	VK_APPS                            = 93
	VK_SLEEP                           = 95
	VK_NUMPAD0                         = 96
	VK_NUMPAD1                         = 97
	VK_NUMPAD2                         = 98
	VK_NUMPAD3                         = 99
	VK_NUMPAD4                         = 100
	VK_NUMPAD5                         = 101
	VK_NUMPAD6                         = 102
	VK_NUMPAD7                         = 103
	VK_NUMPAD8                         = 104
	VK_NUMPAD9                         = 105
	VK_MULTIPLY                        = 106
	VK_ADD                             = 107
	VK_SEPARATOR                       = 108
	VK_SUBTRACT                        = 109
	VK_DECIMAL                         = 110
	VK_DIVIDE                          = 111
	VK_F1                              = 112
	VK_F2                              = 113
	VK_F3                              = 114
	VK_F4                              = 115
	VK_F5                              = 116
	VK_F6                              = 117
	VK_F7                              = 118
	VK_F8                              = 119
	VK_F9                              = 120
	VK_F10                             = 121
	VK_F11                             = 122
	VK_F12                             = 123
	VK_F13                             = 124
	VK_F14                             = 125
	VK_F15                             = 126
	VK_F16                             = 127
	VK_F17                             = 128
	VK_F18                             = 129
	VK_F19                             = 130
	VK_F20                             = 131
	VK_F21                             = 132
	VK_F22                             = 133
	VK_F23                             = 134
	VK_F24                             = 135
	VK_NAVIGATION_VIEW                 = 136
	VK_NAVIGATION_MENU                 = 137
	VK_NAVIGATION_UP                   = 138
	VK_NAVIGATION_DOWN                 = 139
	VK_NAVIGATION_LEFT                 = 140
	VK_NAVIGATION_RIGHT                = 141
	VK_NAVIGATION_ACCEPT               = 142
	VK_NAVIGATION_CANCEL               = 143
	VK_NUMLOCK                         = 144
	VK_SCROLL                          = 145
	VK_OEM_NEC_EQUAL                   = 146
	VK_OEM_FJ_JISHO                    = 146
	VK_OEM_FJ_MASSHOU                  = 147
	VK_OEM_FJ_TOUROKU                  = 148
	VK_OEM_FJ_LOYA                     = 149
	VK_OEM_FJ_ROYA                     = 150
	VK_LSHIFT                          = 160
	VK_RSHIFT                          = 161
	VK_LCONTROL                        = 162
	VK_RCONTROL                        = 163
	VK_LMENU                           = 164
	VK_RMENU                           = 165
	VK_BROWSER_BACK                    = 166
	VK_BROWSER_FORWARD                 = 167
	VK_BROWSER_REFRESH                 = 168
	VK_BROWSER_STOP                    = 169
	VK_BROWSER_SEARCH                  = 170
	VK_BROWSER_FAVORITES               = 171
	VK_BROWSER_HOME                    = 172
	VK_VOLUME_MUTE                     = 173
	VK_VOLUME_DOWN                     = 174
	VK_VOLUME_UP                       = 175
	VK_MEDIA_NEXT_TRACK                = 176
	VK_MEDIA_PREV_TRACK                = 177
	VK_MEDIA_STOP                      = 178
	VK_MEDIA_PLAY_PAUSE                = 179
	VK_LAUNCH_MAIL                     = 180
	VK_LAUNCH_MEDIA_SELECT             = 181
	VK_LAUNCH_APP1                     = 182
	VK_LAUNCH_APP2                     = 183
	VK_OEM_1                           = 186
	VK_OEM_PLUS                        = 187
	VK_OEM_COMMA                       = 188
	VK_OEM_MINUS                       = 189
	VK_OEM_PERIOD                      = 190
	VK_OEM_2                           = 191
	VK_OEM_3                           = 192
	VK_OEM_4                           = 219
	VK_OEM_5                           = 220
	VK_OEM_6                           = 221
	VK_OEM_7                           = 222
	VK_OEM_8                           = 223
	VK_OEM_AX                          = 225
	VK_OEM_102                         = 226
	VK_ICO_HELP                        = 227
	VK_ICO_00                          = 228
	VK_PROCESSKEY                      = 229
	VK_ICO_CLEAR                       = 230
	VK_PACKET                          = 231
	VK_OEM_RESET                       = 233
	VK_OEM_JUMP                        = 234
	VK_OEM_PA1                         = 235
	VK_OEM_PA2                         = 236
	VK_OEM_PA3                         = 237
	VK_OEM_WSCTRL                      = 238
	VK_OEM_CUSEL                       = 239
	VK_OEM_ATTN                        = 240
	VK_OEM_FINISH                      = 241
	VK_OEM_COPY                        = 242
	VK_OEM_AUTO                        = 243
	VK_OEM_ENLW                        = 244
	VK_OEM_BACKTAB                     = 245
	VK_GAMEPAD_A                       = 195
	VK_GAMEPAD_B                       = 196
	VK_GAMEPAD_X                       = 197
	VK_GAMEPAD_Y                       = 198
	VK_GAMEPAD_RIGHT_SHOULDER          = 199
	VK_GAMEPAD_LEFT_SHOULDER           = 200
	VK_GAMEPAD_LEFT_TRIGGER            = 201
	VK_GAMEPAD_RIGHT_TRIGGER           = 202
	VK_GAMEPAD_DPAD_UP                 = 203
	VK_GAMEPAD_DPAD_DOWN               = 204
	VK_GAMEPAD_DPAD_LEFT               = 205
	VK_GAMEPAD_DPAD_RIGHT              = 206
	VK_GAMEPAD_MENU                    = 207
	VK_GAMEPAD_VIEW                    = 208
	VK_GAMEPAD_LEFT_THUMBSTICK_BUTTON  = 209
	VK_GAMEPAD_RIGHT_THUMBSTICK_BUTTON = 210
	VK_GAMEPAD_LEFT_THUMBSTICK_UP      = 211
	VK_GAMEPAD_LEFT_THUMBSTICK_DOWN    = 212
	VK_GAMEPAD_LEFT_THUMBSTICK_RIGHT   = 213
	VK_GAMEPAD_LEFT_THUMBSTICK_LEFT    = 214
	VK_GAMEPAD_RIGHT_THUMBSTICK_UP     = 215
	VK_GAMEPAD_RIGHT_THUMBSTICK_DOWN   = 216
	VK_GAMEPAD_RIGHT_THUMBSTICK_RIGHT  = 217
	VK_GAMEPAD_RIGHT_THUMBSTICK_LEFT   = 218
	VK_ATTN                            = 246
	VK_CRSEL                           = 247
	VK_EXSEL                           = 248
	VK_EREOF                           = 249
	VK_PLAY                            = 250
	VK_ZOOM                            = 251
	VK_NONAME                          = 252
	VK_PA1                             = 253
	VK_OEM_CLEAR                       = 254
	VK_PAD_A                           = 22528
	VK_PAD_B                           = 22529
	VK_PAD_X                           = 22530
	VK_PAD_Y                           = 22531
	VK_PAD_RSHOULDER                   = 22532
	VK_PAD_LSHOULDER                   = 22533
	VK_PAD_LTRIGGER                    = 22534
	VK_PAD_RTRIGGER                    = 22535
	VK_PAD_DPAD_UP                     = 22544
	VK_PAD_DPAD_DOWN                   = 22545
	VK_PAD_DPAD_LEFT                   = 22546
	VK_PAD_DPAD_RIGHT                  = 22547
	VK_PAD_START                       = 22548
	VK_PAD_BACK                        = 22549
	VK_PAD_LTHUMB_PRESS                = 22550
	VK_PAD_RTHUMB_PRESS                = 22551
	VK_PAD_LTHUMB_UP                   = 22560
	VK_PAD_LTHUMB_DOWN                 = 22561
	VK_PAD_LTHUMB_RIGHT                = 22562
	VK_PAD_LTHUMB_LEFT                 = 22563
	VK_PAD_LTHUMB_UPLEFT               = 22564
	VK_PAD_LTHUMB_UPRIGHT              = 22565
	VK_PAD_LTHUMB_DOWNRIGHT            = 22566
	VK_PAD_LTHUMB_DOWNLEFT             = 22567
	VK_PAD_RTHUMB_UP                   = 22576
	VK_PAD_RTHUMB_DOWN                 = 22577
	VK_PAD_RTHUMB_RIGHT                = 22578
	VK_PAD_RTHUMB_LEFT                 = 22579
	VK_PAD_RTHUMB_UPLEFT               = 22580
	VK_PAD_RTHUMB_UPRIGHT              = 22581
	VK_PAD_RTHUMB_DOWNRIGHT            = 22582
	VK_PAD_RTHUMB_DOWNLEFT             = 22583
)

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-messageboxw
const (
	IDOK       = 1
	IDCANCEL   = 2
	IDABORT    = 3
	IDRETRY    = 4
	IDIGNORE   = 5
	IDYES      = 6
	IDNO       = 7
	IDCLOSE    = 8
	IDHELP     = 9
	IDTRYAGAIN = 10
	IDCONTINUE = 11
)

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-mouse_event
const (
	MOUSEEVENTF_ABSOLUTE        = 32768
	MOUSEEVENTF_LEFTDOWN        = 2
	MOUSEEVENTF_LEFTUP          = 4
	MOUSEEVENTF_MIDDLEDOWN      = 32
	MOUSEEVENTF_MIDDLEUP        = 64
	MOUSEEVENTF_MOVE            = 1
	MOUSEEVENTF_RIGHTDOWN       = 8
	MOUSEEVENTF_RIGHTUP         = 16
	MOUSEEVENTF_WHEEL           = 2048
	MOUSEEVENTF_XDOWN           = 128
	MOUSEEVENTF_XUP             = 256
	MOUSEEVENTF_HWHEEL          = 4096
	MOUSEEVENTF_MOVE_NOCOALESCE = 8192
	MOUSEEVENTF_VIRTUALDESK     = 16384
)

// https://learn.microsoft.com/openspecs/windows_protocols/ms-emf/a5e722e3-891a-4a67-be1a-ed5a48a7fda1
const (
	DIB_RGB_COLORS  = 0
	DIB_PAL_COLORS  = 1
	DIB_PAL_INDICES = 2
)

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-bitblt
const (
	BLACKNESS      = 0x00000042
	CAPTUREBLT     = 0x40000000
	DSTINVERT      = 0x00550009
	MERGECOPY      = 0x00C000CA
	MERGEPAINT     = 0x00BB0226
	NOMIRRORBITMAP = 0x80000000
	NOTSRCCOPY     = 0x00330008
	NOTSRCERASE    = 0x001100A6
	PATCOPY        = 0x00F00021
	PATINVERT      = 0x005A0049
	PATPAINT       = 0x00FB0A09
	SRCAND         = 0x008800C6
	SRCCOPY        = 0x00CC0020
	SRCERASE       = 0x00440328
	SRCINVERT      = 0x00660046
	SRCPAINT       = 0x00EE0086
	WHITENESS      = 0x00FF0062
)

// https://learn.microsoft.com/openspecs/windows_protocols/ms-wmf/4e588f70-bd92-4a6f-b77f-35d0feaf7a57
const (
	BI_RGB       = 0
	BI_RLE8      = 1
	BI_RLE4      = 2
	BI_BITFIELDS = 3
	BI_JPEG      = 4
	BI_PNG       = 5
	BI_CMYK      = 11
	BI_CMYKRLE8  = 12
	BI_CMYKRLE4  = 13
	BI_1632      = 842217009
)

// https://learn.microsoft.com/windows/win32/dataxchg/standard-clipboard-formats
const (
	CF_BITMAP          = 2
	CF_DIB             = 8
	CF_DIBV5           = 17
	CF_DIF             = 5
	CF_DSPBITMAP       = 130
	CF_DSPENHMETAFILE  = 142
	CF_DSPMETAFILEPICT = 131
	CF_DSPTEXT         = 129
	CF_ENHMETAFILE     = 14
	CF_GDIOBJFIRST     = 768
	CF_GDIOBJLAST      = 1023
	CF_HDROP           = 15
	CF_LOCALE          = 16
	CF_METAFILEPICT    = 3
	CF_OEMTEXT         = 7
	CF_OWNERDISPLAY    = 128
	CF_PALETTE         = 9
	CF_PENDATA         = 10
	CF_PRIVATEFIRST    = 512
	CF_PRIVATELAST     = 767
	CF_RIFF            = 11
	CF_SYLK            = 4
	CF_TEXT            = 1
	CF_TIFF            = 6
	CF_UNICODETEXT     = 13
	CF_WAVE            = 12
	CF_MAX             = 18
)

// https://learn.microsoft.com/windows/win32/api/winbase/nf-winbase-globalalloc
const (
	GMEM_FIXED          = 0
	GMEM_MOVEABLE       = 2
	GMEM_ZEROINIT       = 64
	GMEM_NOCOMPACT      = 16
	GMEM_NODISCARD      = 32
	GMEM_MODIFY         = 128
	GMEM_DISCARDABLE    = 256
	GMEM_NOT_BANKED     = 4096
	GMEM_SHARE          = 8192
	GMEM_DDESHARE       = 8192
	GMEM_NOTIFY         = 16384
	GMEM_LOWER          = 4096
	GMEM_VALID_FLAGS    = 32626
	GMEM_INVALID_HANDLE = 32768
	GMEM_DISCARDED      = 16384
	GMEM_LOCKCOUNT      = 255
	GHND                = GMEM_MOVEABLE | GMEM_ZEROINIT
	GPTR                = GMEM_FIXED | GMEM_ZEROINIT
)

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-setwindowshookexw
const (
	WH_CALLWNDPROC     = 4
	WH_CALLWNDPROCRET  = 12
	WH_CBT             = 5
	WH_DEBUG           = 9
	WH_FOREGROUNDIDLE  = 11
	WH_GETMESSAGE      = 3
	WH_JOURNALPLAYBACK = 1
	WH_JOURNALRECORD   = 0
	WH_KEYBOARD        = 2
	WH_KEYBOARD_LL     = 13
	WH_MOUSE           = 7
	WH_MOUSE_LL        = 14
	WH_MSGFILTER       = -1
	WH_SHELL           = 10
	WH_SYSMSGFILTER    = 6
	WH_MIN             = -1
	WH_HARDWARE        = 8
	WH_MAX             = 14
	WH_MINHOOK         = -1
	WH_MAXHOOK         = 14
)

// https://learn.microsoft.com/windows/win32/api/winuser/ns-winuser-input
const (
	INPUT_MOUSE    = 0
	INPUT_KEYBOARD = 1
	INPUT_HARDWARE = 2
)

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-mapvirtualkeyw
const (
	MAPVK_VK_TO_VSC    = 0
	MAPVK_VSC_TO_VK    = 1
	MAPVK_VK_TO_CHAR   = 2
	MAPVK_VSC_TO_VK_EX = 3
	MAPVK_VK_TO_VSC_EX = 4
)

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-loadkeyboardlayoutw
const (
	KLF_REORDER       = 8
	KLF_RESET         = 1073741824
	KLF_SETFORPROCESS = 256
	KLF_SHIFTLOCK     = 65536
	KLF_ACTIVATE      = 1
	KLF_NOTELLSHELL   = 128
	KLF_REPLACELANG   = 16
	KLF_SUBSTITUTE_OK = 2
)

const (
	KEYEVENTF_EXTENDEDKEY = 1
	KEYEVENTF_KEYUP       = 2
	KEYEVENTF_UNICODE     = 4
	KEYEVENTF_SCANCODE    = 8
)

const (
	LLKHF_EXTENDED          = 1
	LLKHF_ALTDOWN           = 32
	LLKHF_UP                = 128
	LLKHF_INJECTED          = 16
	LLKHF_LOWER_IL_INJECTED = 2
)

const (
	LLMHF_INJECTED          = 1
	LLMHF_LOWER_IL_INJECTED = 2
)

const (
	XBUTTON1 = 1
	XBUTTON2 = 2
)

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-showwindow
const (
	SW_HIDE               = 0
	SW_SHOWNORMAL         = 1
	SW_NORMAL             = 1
	SW_SHOWMINIMIZED      = 2
	SW_SHOWMAXIMIZED      = 3
	SW_MAXIMIZE           = 3
	SW_SHOWNOACTIVATE     = 4
	SW_SHOW               = 5
	SW_MINIMIZE           = 6
	SW_SHOWMINNOACTIVE    = 7
	SW_SHOWNA             = 8
	SW_RESTORE            = 9
	SW_SHOWDEFAULT        = 10
	SW_FORCEMINIMIZE      = 11
	SW_MAX                = 11
	SW_PARENTCLOSING      = 1
	SW_OTHERZOOM          = 2
	SW_PARENTOPENING      = 3
	SW_OTHERUNZOOM        = 4
	SW_SCROLLCHILDREN     = 1
	SW_INVALIDATE         = 2
	SW_ERASE              = 4
	SW_SMOOTHSCROLL       = 16
	SW_AUTOPROF_LOAD_MASK = 1
	SW_AUTOPROF_SAVE_MASK = 2
)

const (
	GWL_EXSTYLE    = -20
	GWL_STYLE      = -16
	GWL_HINSTANCE  = -6
	GWL_ID         = -12
	GWL_USERDATA   = -21
	GWL_WNDPROC    = -4
	GWL_HWNDPARENT = -8
)

const (
	DWL_MSGRESULT = 0
	DWL_DLGPROC   = 4
	DWL_USER      = 8
)

const (
	HWND_BROADCAST = 65535

	// HWND_MESSAGE can be set as the parent window in CreateWindowEx to create a
	// message-only window, i.e. a non-visible window that only exists to receive
	// window messages.
	//
	// https://learn.microsoft.com/en-us/windows/win32/winmsg/window-features
	HWND_MESSAGE = ^HWND(3) + 1

	HWND_DESKTOP   = 0
	HWND_TOP       = 0
	HWND_BOTTOM    = 1
	HWND_TOPMOST   = ^HWND(1) + 1
	HWND_NOTOPMOST = ^HWND(2) + 1
)

const (
	SWP_ASYNCWINDOWPOS = 16384
	SWP_DEFERERASE     = 8192
	SWP_DRAWFRAME      = 32
	SWP_FRAMECHANGED   = 32
	SWP_HIDEWINDOW     = 128
	SWP_NOACTIVATE     = 16
	SWP_NOCOPYBITS     = 256
	SWP_NOMOVE         = 2
	SWP_NOOWNERZORDER  = 512
	SWP_NOREDRAW       = 8
	SWP_NOREPOSITION   = 512
	SWP_NOSENDCHANGING = 1024
	SWP_NOSIZE         = 1
	SWP_NOZORDER       = 4
	SWP_SHOWWINDOW     = 64
)

// https://learn.microsoft.com/windows/win32/api/winuser/ns-winuser-windowplacement
const (
	WPF_ASYNCWINDOWPLACEMENT = 4
	WPF_RESTORETOMAXIMIZED   = 2
	WPF_SETMINPOSITION       = 1
)

const (
	MONITOR_DEFAULTTONULL    = 0
	MONITOR_DEFAULTTOPRIMARY = 1
	MONITOR_DEFAULTTONEAREST = 2
)

const (
	DWMWA_NCRENDERING_ENABLED         = 1
	DWMWA_NCRENDERING_POLICY          = 2
	DWMWA_TRANSITIONS_FORCEDISABLED   = 3
	DWMWA_ALLOW_NCPAINT               = 4
	DWMWA_CAPTION_BUTTON_BOUNDS       = 5
	DWMWA_NONCLIENT_RTL_LAYOUT        = 6
	DWMWA_FORCE_ICONIC_REPRESENTATION = 7
	DWMWA_FLIP3D_POLICY               = 8
	DWMWA_EXTENDED_FRAME_BOUNDS       = 9
	DWMWA_HAS_ICONIC_BITMAP           = 10
	DWMWA_DISALLOW_PEEK               = 11
	DWMWA_EXCLUDED_FROM_PEEK          = 12

	// For cloaking see:
	// https://devblogs.microsoft.com/oldnewthing/20200302-00/?p=103507
	DWMWA_CLOAK   = 13
	DWMWA_CLOAKED = 14

	DWMWA_FREEZE_REPRESENTATION          = 15
	DWMWA_PASSIVE_UPDATE_MODE            = 16
	DWMWA_USE_HOSTBACKDROPBRUSH          = 17
	DWMWA_USE_IMMERSIVE_DARK_MODE        = 20
	DWMWA_WINDOW_CORNER_PREFERENCE       = 33
	DWMWA_BORDER_COLOR                   = 34
	DWMWA_CAPTION_COLOR                  = 35
	DWMWA_TEXT_COLOR                     = 36
	DWMWA_VISIBLE_FRAME_BORDER_THICKNESS = 37
	DWMWA_SYSTEMBACKDROP_TYPE            = 38
	DWMWA_LAST                           = 39
	DWMWA_COLOR_DEFAULT                  = 4294967295
	DWMWA_COLOR_NONE                     = 4294967294
)

const (
	FORMAT_MESSAGE_ALLOCATE_BUFFER = 256
	FORMAT_MESSAGE_ARGUMENT_ARRAY  = 8192
	FORMAT_MESSAGE_FROM_HMODULE    = 2048
	FORMAT_MESSAGE_FROM_STRING     = 1024
	FORMAT_MESSAGE_FROM_SYSTEM     = 4096
	FORMAT_MESSAGE_IGNORE_INSERTS  = 512
	FORMAT_MESSAGE_MAX_WIDTH_MASK  = 255
)

const S_OK = 0

const (
	LANG_NEUTRAL             = 0
	LANG_INVARIANT           = 127
	LANG_AFRIKAANS           = 54
	LANG_ALBANIAN            = 28
	LANG_ALSATIAN            = 132
	LANG_AMHARIC             = 94
	LANG_ARABIC              = 1
	LANG_ARMENIAN            = 43
	LANG_ASSAMESE            = 77
	LANG_AZERI               = 44
	LANG_AZERBAIJANI         = 44
	LANG_BANGLA              = 69
	LANG_BASHKIR             = 109
	LANG_BASQUE              = 45
	LANG_BELARUSIAN          = 35
	LANG_BENGALI             = 69
	LANG_BRETON              = 126
	LANG_BOSNIAN             = 26
	LANG_BOSNIAN_NEUTRAL     = 30746
	LANG_BULGARIAN           = 2
	LANG_CATALAN             = 3
	LANG_CENTRAL_KURDISH     = 146
	LANG_CHEROKEE            = 92
	LANG_CHINESE             = 4
	LANG_CHINESE_SIMPLIFIED  = 4
	LANG_CHINESE_TRADITIONAL = 31748
	LANG_CORSICAN            = 131
	LANG_CROATIAN            = 26
	LANG_CZECH               = 5
	LANG_DANISH              = 6
	LANG_DARI                = 140
	LANG_DIVEHI              = 101
	LANG_DUTCH               = 19
	LANG_ENGLISH             = 9
	LANG_ESTONIAN            = 37
	LANG_FAEROESE            = 56
	LANG_FARSI               = 41
	LANG_FILIPINO            = 100
	LANG_FINNISH             = 11
	LANG_FRENCH              = 12
	LANG_FRISIAN             = 98
	LANG_FULAH               = 103
	LANG_GALICIAN            = 86
	LANG_GEORGIAN            = 55
	LANG_GERMAN              = 7
	LANG_GREEK               = 8
	LANG_GREENLANDIC         = 111
	LANG_GUJARATI            = 71
	LANG_HAUSA               = 104
	LANG_HAWAIIAN            = 117
	LANG_HEBREW              = 13
	LANG_HINDI               = 57
	LANG_HUNGARIAN           = 14
	LANG_ICELANDIC           = 15
	LANG_IGBO                = 112
	LANG_INDONESIAN          = 33
	LANG_INUKTITUT           = 93
	LANG_IRISH               = 60
	LANG_ITALIAN             = 16
	LANG_JAPANESE            = 17
	LANG_KANNADA             = 75
	LANG_KASHMIRI            = 96
	LANG_KAZAK               = 63
	LANG_KHMER               = 83
	LANG_KICHE               = 134
	LANG_KINYARWANDA         = 135
	LANG_KONKANI             = 87
	LANG_KOREAN              = 18
	LANG_KYRGYZ              = 64
	LANG_LAO                 = 84
	LANG_LATVIAN             = 38
	LANG_LITHUANIAN          = 39
	LANG_LOWER_SORBIAN       = 46
	LANG_LUXEMBOURGISH       = 110
	LANG_MACEDONIAN          = 47
	LANG_MALAY               = 62
	LANG_MALAYALAM           = 76
	LANG_MALTESE             = 58
	LANG_MANIPURI            = 88
	LANG_MAORI               = 129
	LANG_MAPUDUNGUN          = 122
	LANG_MARATHI             = 78
	LANG_MOHAWK              = 124
	LANG_MONGOLIAN           = 80
	LANG_NEPALI              = 97
	LANG_NORWEGIAN           = 20
	LANG_OCCITAN             = 130
	LANG_ODIA                = 72
	LANG_ORIYA               = 72
	LANG_PASHTO              = 99
	LANG_PERSIAN             = 41
	LANG_POLISH              = 21
	LANG_PORTUGUESE          = 22
	LANG_PULAR               = 103
	LANG_PUNJABI             = 70
	LANG_QUECHUA             = 107
	LANG_ROMANIAN            = 24
	LANG_ROMANSH             = 23
	LANG_RUSSIAN             = 25
	LANG_SAKHA               = 133
	LANG_SAMI                = 59
	LANG_SANSKRIT            = 79
	LANG_SCOTTISH_GAELIC     = 145
	LANG_SERBIAN             = 26
	LANG_SERBIAN_NEUTRAL     = 31770
	LANG_SINDHI              = 89
	LANG_SINHALESE           = 91
	LANG_SLOVAK              = 27
	LANG_SLOVENIAN           = 36
	LANG_SOTHO               = 108
	LANG_SPANISH             = 10
	LANG_SWAHILI             = 65
	LANG_SWEDISH             = 29
	LANG_SYRIAC              = 90
	LANG_TAJIK               = 40
	LANG_TAMAZIGHT           = 95
	LANG_TAMIL               = 73
	LANG_TATAR               = 68
	LANG_TELUGU              = 74
	LANG_THAI                = 30
	LANG_TIBETAN             = 81
	LANG_TIGRIGNA            = 115
	LANG_TIGRINYA            = 115
	LANG_TSWANA              = 50
	LANG_TURKISH             = 31
	LANG_TURKMEN             = 66
	LANG_UIGHUR              = 128
	LANG_UKRAINIAN           = 34
	LANG_UPPER_SORBIAN       = 46
	LANG_URDU                = 32
	LANG_UZBEK               = 67
	LANG_VALENCIAN           = 3
	LANG_VIETNAMESE          = 42
	LANG_WELSH               = 82
	LANG_WOLOF               = 136
	LANG_XHOSA               = 52
	LANG_YAKUT               = 133
	LANG_YI                  = 120
	LANG_YORUBA              = 106
	LANG_ZULU                = 53
	LANG_BASE_LANGUAGE_INDEX = 0
	LANG_BASE_ENCODING_INDEX = 1
	LANG_BASE_OFFSET_INDEX   = 2
	LANG_DEFAULT_ID          = 256
)

const (
	SUBLANG_NEUTRAL                             = 0
	SUBLANG_DEFAULT                             = 1
	SUBLANG_SYS_DEFAULT                         = 2
	SUBLANG_CUSTOM_DEFAULT                      = 3
	SUBLANG_CUSTOM_UNSPECIFIED                  = 4
	SUBLANG_UI_CUSTOM_DEFAULT                   = 5
	SUBLANG_AFRIKAANS_SOUTH_AFRICA              = 1
	SUBLANG_ALBANIAN_ALBANIA                    = 1
	SUBLANG_ALSATIAN_FRANCE                     = 1
	SUBLANG_AMHARIC_ETHIOPIA                    = 1
	SUBLANG_ARABIC_SAUDI_ARABIA                 = 1
	SUBLANG_ARABIC_IRAQ                         = 2
	SUBLANG_ARABIC_EGYPT                        = 3
	SUBLANG_ARABIC_LIBYA                        = 4
	SUBLANG_ARABIC_ALGERIA                      = 5
	SUBLANG_ARABIC_MOROCCO                      = 6
	SUBLANG_ARABIC_TUNISIA                      = 7
	SUBLANG_ARABIC_OMAN                         = 8
	SUBLANG_ARABIC_YEMEN                        = 9
	SUBLANG_ARABIC_SYRIA                        = 10
	SUBLANG_ARABIC_JORDAN                       = 11
	SUBLANG_ARABIC_LEBANON                      = 12
	SUBLANG_ARABIC_KUWAIT                       = 13
	SUBLANG_ARABIC_UAE                          = 14
	SUBLANG_ARABIC_BAHRAIN                      = 15
	SUBLANG_ARABIC_QATAR                        = 16
	SUBLANG_ARMENIAN_ARMENIA                    = 1
	SUBLANG_ASSAMESE_INDIA                      = 1
	SUBLANG_AZERI_LATIN                         = 1
	SUBLANG_AZERI_CYRILLIC                      = 2
	SUBLANG_AZERBAIJANI_AZERBAIJAN_LATIN        = 1
	SUBLANG_AZERBAIJANI_AZERBAIJAN_CYRILLIC     = 2
	SUBLANG_BANGLA_INDIA                        = 1
	SUBLANG_BANGLA_BANGLADESH                   = 2
	SUBLANG_BASHKIR_RUSSIA                      = 1
	SUBLANG_BASQUE_BASQUE                       = 1
	SUBLANG_BELARUSIAN_BELARUS                  = 1
	SUBLANG_BENGALI_INDIA                       = 1
	SUBLANG_BENGALI_BANGLADESH                  = 2
	SUBLANG_BOSNIAN_BOSNIA_HERZEGOVINA_LATIN    = 5
	SUBLANG_BOSNIAN_BOSNIA_HERZEGOVINA_CYRILLIC = 8
	SUBLANG_BRETON_FRANCE                       = 1
	SUBLANG_BULGARIAN_BULGARIA                  = 1
	SUBLANG_CATALAN_CATALAN                     = 1
	SUBLANG_CENTRAL_KURDISH_IRAQ                = 1
	SUBLANG_CHEROKEE_CHEROKEE                   = 1
	SUBLANG_CHINESE_TRADITIONAL                 = 1
	SUBLANG_CHINESE_SIMPLIFIED                  = 2
	SUBLANG_CHINESE_HONGKONG                    = 3
	SUBLANG_CHINESE_SINGAPORE                   = 4
	SUBLANG_CHINESE_MACAU                       = 5
	SUBLANG_CORSICAN_FRANCE                     = 1
	SUBLANG_CZECH_CZECH_REPUBLIC                = 1
	SUBLANG_CROATIAN_CROATIA                    = 1
	SUBLANG_CROATIAN_BOSNIA_HERZEGOVINA_LATIN   = 4
	SUBLANG_DANISH_DENMARK                      = 1
	SUBLANG_DARI_AFGHANISTAN                    = 1
	SUBLANG_DIVEHI_MALDIVES                     = 1
	SUBLANG_DUTCH                               = 1
	SUBLANG_DUTCH_BELGIAN                       = 2
	SUBLANG_ENGLISH_US                          = 1
	SUBLANG_ENGLISH_UK                          = 2
	SUBLANG_ENGLISH_AUS                         = 3
	SUBLANG_ENGLISH_CAN                         = 4
	SUBLANG_ENGLISH_NZ                          = 5
	SUBLANG_ENGLISH_EIRE                        = 6
	SUBLANG_ENGLISH_SOUTH_AFRICA                = 7
	SUBLANG_ENGLISH_JAMAICA                     = 8
	SUBLANG_ENGLISH_CARIBBEAN                   = 9
	SUBLANG_ENGLISH_BELIZE                      = 10
	SUBLANG_ENGLISH_TRINIDAD                    = 11
	SUBLANG_ENGLISH_ZIMBABWE                    = 12
	SUBLANG_ENGLISH_PHILIPPINES                 = 13
	SUBLANG_ENGLISH_INDIA                       = 16
	SUBLANG_ENGLISH_MALAYSIA                    = 17
	SUBLANG_ENGLISH_SINGAPORE                   = 18
	SUBLANG_ESTONIAN_ESTONIA                    = 1
	SUBLANG_FAEROESE_FAROE_ISLANDS              = 1
	SUBLANG_FILIPINO_PHILIPPINES                = 1
	SUBLANG_FINNISH_FINLAND                     = 1
	SUBLANG_FRENCH                              = 1
	SUBLANG_FRENCH_BELGIAN                      = 2
	SUBLANG_FRENCH_CANADIAN                     = 3
	SUBLANG_FRENCH_SWISS                        = 4
	SUBLANG_FRENCH_LUXEMBOURG                   = 5
	SUBLANG_FRENCH_MONACO                       = 6
	SUBLANG_FRISIAN_NETHERLANDS                 = 1
	SUBLANG_FULAH_SENEGAL                       = 2
	SUBLANG_GALICIAN_GALICIAN                   = 1
	SUBLANG_GEORGIAN_GEORGIA                    = 1
	SUBLANG_GERMAN                              = 1
	SUBLANG_GERMAN_SWISS                        = 2
	SUBLANG_GERMAN_AUSTRIAN                     = 3
	SUBLANG_GERMAN_LUXEMBOURG                   = 4
	SUBLANG_GERMAN_LIECHTENSTEIN                = 5
	SUBLANG_GREEK_GREECE                        = 1
	SUBLANG_GREENLANDIC_GREENLAND               = 1
	SUBLANG_GUJARATI_INDIA                      = 1
	SUBLANG_HAUSA_NIGERIA_LATIN                 = 1
	SUBLANG_HAWAIIAN_US                         = 1
	SUBLANG_HEBREW_ISRAEL                       = 1
	SUBLANG_HINDI_INDIA                         = 1
	SUBLANG_HUNGARIAN_HUNGARY                   = 1
	SUBLANG_ICELANDIC_ICELAND                   = 1
	SUBLANG_IGBO_NIGERIA                        = 1
	SUBLANG_INDONESIAN_INDONESIA                = 1
	SUBLANG_INUKTITUT_CANADA                    = 1
	SUBLANG_INUKTITUT_CANADA_LATIN              = 2
	SUBLANG_IRISH_IRELAND                       = 2
	SUBLANG_ITALIAN                             = 1
	SUBLANG_ITALIAN_SWISS                       = 2
	SUBLANG_JAPANESE_JAPAN                      = 1
	SUBLANG_KANNADA_INDIA                       = 1
	SUBLANG_KASHMIRI_SASIA                      = 2
	SUBLANG_KASHMIRI_INDIA                      = 2
	SUBLANG_KAZAK_KAZAKHSTAN                    = 1
	SUBLANG_KHMER_CAMBODIA                      = 1
	SUBLANG_KICHE_GUATEMALA                     = 1
	SUBLANG_KINYARWANDA_RWANDA                  = 1
	SUBLANG_KONKANI_INDIA                       = 1
	SUBLANG_KOREAN                              = 1
	SUBLANG_KYRGYZ_KYRGYZSTAN                   = 1
	SUBLANG_LAO_LAO                             = 1
	SUBLANG_LATVIAN_LATVIA                      = 1
	SUBLANG_LITHUANIAN                          = 1
	SUBLANG_LOWER_SORBIAN_GERMANY               = 2
	SUBLANG_LUXEMBOURGISH_LUXEMBOURG            = 1
	SUBLANG_MACEDONIAN_MACEDONIA                = 1
	SUBLANG_MALAY_MALAYSIA                      = 1
	SUBLANG_MALAY_BRUNEI_DARUSSALAM             = 2
	SUBLANG_MALAYALAM_INDIA                     = 1
	SUBLANG_MALTESE_MALTA                       = 1
	SUBLANG_MAORI_NEW_ZEALAND                   = 1
	SUBLANG_MAPUDUNGUN_CHILE                    = 1
	SUBLANG_MARATHI_INDIA                       = 1
	SUBLANG_MOHAWK_MOHAWK                       = 1
	SUBLANG_MONGOLIAN_CYRILLIC_MONGOLIA         = 1
	SUBLANG_MONGOLIAN_PRC                       = 2
	SUBLANG_NEPALI_INDIA                        = 2
	SUBLANG_NEPALI_NEPAL                        = 1
	SUBLANG_NORWEGIAN_BOKMAL                    = 1
	SUBLANG_NORWEGIAN_NYNORSK                   = 2
	SUBLANG_OCCITAN_FRANCE                      = 1
	SUBLANG_ODIA_INDIA                          = 1
	SUBLANG_ORIYA_INDIA                         = 1
	SUBLANG_PASHTO_AFGHANISTAN                  = 1
	SUBLANG_PERSIAN_IRAN                        = 1
	SUBLANG_POLISH_POLAND                       = 1
	SUBLANG_PORTUGUESE                          = 2
	SUBLANG_PORTUGUESE_BRAZILIAN                = 1
	SUBLANG_PULAR_SENEGAL                       = 2
	SUBLANG_PUNJABI_INDIA                       = 1
	SUBLANG_PUNJABI_PAKISTAN                    = 2
	SUBLANG_QUECHUA_BOLIVIA                     = 1
	SUBLANG_QUECHUA_ECUADOR                     = 2
	SUBLANG_QUECHUA_PERU                        = 3
	SUBLANG_ROMANIAN_ROMANIA                    = 1
	SUBLANG_ROMANSH_SWITZERLAND                 = 1
	SUBLANG_RUSSIAN_RUSSIA                      = 1
	SUBLANG_SAKHA_RUSSIA                        = 1
	SUBLANG_SAMI_NORTHERN_NORWAY                = 1
	SUBLANG_SAMI_NORTHERN_SWEDEN                = 2
	SUBLANG_SAMI_NORTHERN_FINLAND               = 3
	SUBLANG_SAMI_LULE_NORWAY                    = 4
	SUBLANG_SAMI_LULE_SWEDEN                    = 5
	SUBLANG_SAMI_SOUTHERN_NORWAY                = 6
	SUBLANG_SAMI_SOUTHERN_SWEDEN                = 7
	SUBLANG_SAMI_SKOLT_FINLAND                  = 8
	SUBLANG_SAMI_INARI_FINLAND                  = 9
	SUBLANG_SANSKRIT_INDIA                      = 1
	SUBLANG_SCOTTISH_GAELIC                     = 1
	SUBLANG_SERBIAN_BOSNIA_HERZEGOVINA_LATIN    = 6
	SUBLANG_SERBIAN_BOSNIA_HERZEGOVINA_CYRILLIC = 7
	SUBLANG_SERBIAN_MONTENEGRO_LATIN            = 11
	SUBLANG_SERBIAN_MONTENEGRO_CYRILLIC         = 12
	SUBLANG_SERBIAN_SERBIA_LATIN                = 9
	SUBLANG_SERBIAN_SERBIA_CYRILLIC             = 10
	SUBLANG_SERBIAN_CROATIA                     = 1
	SUBLANG_SERBIAN_LATIN                       = 2
	SUBLANG_SERBIAN_CYRILLIC                    = 3
	SUBLANG_SINDHI_INDIA                        = 1
	SUBLANG_SINDHI_PAKISTAN                     = 2
	SUBLANG_SINDHI_AFGHANISTAN                  = 2
	SUBLANG_SINHALESE_SRI_LANKA                 = 1
	SUBLANG_SOTHO_NORTHERN_SOUTH_AFRICA         = 1
	SUBLANG_SLOVAK_SLOVAKIA                     = 1
	SUBLANG_SLOVENIAN_SLOVENIA                  = 1
	SUBLANG_SPANISH                             = 1
	SUBLANG_SPANISH_MEXICAN                     = 2
	SUBLANG_SPANISH_MODERN                      = 3
	SUBLANG_SPANISH_GUATEMALA                   = 4
	SUBLANG_SPANISH_COSTA_RICA                  = 5
	SUBLANG_SPANISH_PANAMA                      = 6
	SUBLANG_SPANISH_DOMINICAN_REPUBLIC          = 7
	SUBLANG_SPANISH_VENEZUELA                   = 8
	SUBLANG_SPANISH_COLOMBIA                    = 9
	SUBLANG_SPANISH_PERU                        = 10
	SUBLANG_SPANISH_ARGENTINA                   = 11
	SUBLANG_SPANISH_ECUADOR                     = 12
	SUBLANG_SPANISH_CHILE                       = 13
	SUBLANG_SPANISH_URUGUAY                     = 14
	SUBLANG_SPANISH_PARAGUAY                    = 15
	SUBLANG_SPANISH_BOLIVIA                     = 16
	SUBLANG_SPANISH_EL_SALVADOR                 = 17
	SUBLANG_SPANISH_HONDURAS                    = 18
	SUBLANG_SPANISH_NICARAGUA                   = 19
	SUBLANG_SPANISH_PUERTO_RICO                 = 20
	SUBLANG_SPANISH_US                          = 21
	SUBLANG_SWAHILI_KENYA                       = 1
	SUBLANG_SWEDISH                             = 1
	SUBLANG_SWEDISH_FINLAND                     = 2
	SUBLANG_SYRIAC_SYRIA                        = 1
	SUBLANG_TAJIK_TAJIKISTAN                    = 1
	SUBLANG_TAMAZIGHT_ALGERIA_LATIN             = 2
	SUBLANG_TAMAZIGHT_MOROCCO_TIFINAGH          = 4
	SUBLANG_TAMIL_INDIA                         = 1
	SUBLANG_TAMIL_SRI_LANKA                     = 2
	SUBLANG_TATAR_RUSSIA                        = 1
	SUBLANG_TELUGU_INDIA                        = 1
	SUBLANG_THAI_THAILAND                       = 1
	SUBLANG_TIBETAN_PRC                         = 1
	SUBLANG_TIGRIGNA_ERITREA                    = 2
	SUBLANG_TIGRINYA_ERITREA                    = 2
	SUBLANG_TIGRINYA_ETHIOPIA                   = 1
	SUBLANG_TSWANA_BOTSWANA                     = 2
	SUBLANG_TSWANA_SOUTH_AFRICA                 = 1
	SUBLANG_TURKISH_TURKEY                      = 1
	SUBLANG_TURKMEN_TURKMENISTAN                = 1
	SUBLANG_UIGHUR_PRC                          = 1
	SUBLANG_UKRAINIAN_UKRAINE                   = 1
	SUBLANG_UPPER_SORBIAN_GERMANY               = 1
	SUBLANG_URDU_PAKISTAN                       = 1
	SUBLANG_URDU_INDIA                          = 2
	SUBLANG_UZBEK_LATIN                         = 1
	SUBLANG_UZBEK_CYRILLIC                      = 2
	SUBLANG_VALENCIAN_VALENCIA                  = 2
	SUBLANG_VIETNAMESE_VIETNAM                  = 1
	SUBLANG_WELSH_UNITED_KINGDOM                = 1
	SUBLANG_WOLOF_SENEGAL                       = 1
	SUBLANG_XHOSA_SOUTH_AFRICA                  = 1
	SUBLANG_YAKUT_RUSSIA                        = 1
	SUBLANG_YI_PRC                              = 1
	SUBLANG_YORUBA_NIGERIA                      = 1
	SUBLANG_ZULU_SOUTH_AFRICA                   = 1
)

// https://devblogs.microsoft.com/oldnewthing/20200302-00/?p=103507
const (
	DWM_CLOAKED_APP       = 1
	DWM_CLOAKED_SHELL     = 2
	DWM_CLOAKED_INHERITED = 4
)

// https://learn.microsoft.com/windows/win32/api/dwmapi/ne-dwmapi-dwm_systembackdrop_type
const (
	DWMSBT_AUTO            = 0
	DWMSBT_NONE            = 1
	DWMSBT_MAINWINDOW      = 2
	DWMSBT_TRANSIENTWINDOW = 3
	DWMSBT_TABBEDWINDOW    = 4
)

// https://learn.microsoft.com/windows/win32/api/dwmapi/ne-dwmapi-dwmncrenderingpolicy
const (
	DWMNCRP_USEWINDOWSTYLE = 0
	DWMNCRP_DISABLED       = 1
	DWMNCRP_ENABLED        = 2
	DWMNCRP_LAST           = 3
)

const (
	DWMWCP_DEFAULT    = 0
	DWMWCP_DONOTROUND = 1
	DWMWCP_ROUND      = 2
	DWMWCP_ROUNDSMALL = 3
)

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-printwindow?redirectedfrom=MSDN
const (
	PW_CLIENTONLY        = 1
	PW_RENDERFULLCONTENT = 2
)

// https://learn.microsoft.com/windows/win32/gdi/wm-print?redirectedfrom=MSDN
const (
	PRF_CHECKVISIBLE = 1
	PRF_CHILDREN     = 16
	PRF_CLIENT       = 4
	PRF_ERASEBKGND   = 8
	PRF_NONCLIENT    = 2
	PRF_OWNED        = 32
)

const (
	AW_ACTIVATE     = 131072
	AW_BLEND        = 524288
	AW_CENTER       = 16
	AW_HIDE         = 65536
	AW_HOR_POSITIVE = 1
	AW_HOR_NEGATIVE = 2
	AW_SLIDE        = 262144
	AW_VER_POSITIVE = 4
	AW_VER_NEGATIVE = 8
)

const (
	SM_ARRANGE                     = 56
	SM_CLEANBOOT                   = 67
	SM_CMONITORS                   = 80
	SM_CMOUSEBUTTONS               = 43
	SM_CONVERTIBLESLATEMODE        = 8195
	SM_CXBORDER                    = 5
	SM_CXCURSOR                    = 13
	SM_CXDLGFRAME                  = 7
	SM_CXDOUBLECLK                 = 36
	SM_CXDRAG                      = 68
	SM_CXEDGE                      = 45
	SM_CXFIXEDFRAME                = 7
	SM_CXFOCUSBORDER               = 83
	SM_CXFRAME                     = 32
	SM_CXFULLSCREEN                = 16
	SM_CXHSCROLL                   = 21
	SM_CXHTHUMB                    = 10
	SM_CXICON                      = 11
	SM_CXICONSPACING               = 38
	SM_CXMAXIMIZED                 = 61
	SM_CXMAXTRACK                  = 59
	SM_CXMENUCHECK                 = 71
	SM_CXMENUSIZE                  = 54
	SM_CXMIN                       = 28
	SM_CXMINIMIZED                 = 57
	SM_CXMINSPACING                = 47
	SM_CXMINTRACK                  = 34
	SM_CXPADDEDBORDER              = 92
	SM_CXSCREEN                    = 0
	SM_CXSIZE                      = 30
	SM_CXSIZEFRAME                 = 32
	SM_CXSMICON                    = 49
	SM_CXSMSIZE                    = 52
	SM_CXVIRTUALSCREEN             = 78
	SM_CXVSCROLL                   = 2
	SM_CYBORDER                    = 6
	SM_CYCAPTION                   = 4
	SM_CYCURSOR                    = 14
	SM_CYDLGFRAME                  = 8
	SM_CYDOUBLECLK                 = 37
	SM_CYDRAG                      = 69
	SM_CYEDGE                      = 46
	SM_CYFIXEDFRAME                = 8
	SM_CYFOCUSBORDER               = 84
	SM_CYFRAME                     = 33
	SM_CYFULLSCREEN                = 17
	SM_CYHSCROLL                   = 3
	SM_CYICON                      = 12
	SM_CYICONSPACING               = 39
	SM_CYKANJIWINDOW               = 18
	SM_CYMAXIMIZED                 = 62
	SM_CYMAXTRACK                  = 60
	SM_CYMENU                      = 15
	SM_CYMENUCHECK                 = 72
	SM_CYMENUSIZE                  = 55
	SM_CYMIN                       = 29
	SM_CYMINIMIZED                 = 58
	SM_CYMINSPACING                = 48
	SM_CYMINTRACK                  = 35
	SM_CYSCREEN                    = 1
	SM_CYSIZE                      = 31
	SM_CYSIZEFRAME                 = 33
	SM_CYSMCAPTION                 = 51
	SM_CYSMICON                    = 50
	SM_CYSMSIZE                    = 53
	SM_CYVIRTUALSCREEN             = 79
	SM_CYVSCROLL                   = 20
	SM_CYVTHUMB                    = 9
	SM_DBCSENABLED                 = 42
	SM_DEBUG                       = 22
	SM_DIGITIZER                   = 94
	SM_IMMENABLED                  = 82
	SM_MAXIMUMTOUCHES              = 95
	SM_MEDIACENTER                 = 87
	SM_MENUDROPALIGNMENT           = 40
	SM_MIDEASTENABLED              = 74
	SM_MOUSEPRESENT                = 19
	SM_MOUSEHORIZONTALWHEELPRESENT = 91
	SM_MOUSEWHEELPRESENT           = 75
	SM_NETWORK                     = 63
	SM_PENWINDOWS                  = 41
	SM_REMOTECONTROL               = 8193
	SM_REMOTESESSION               = 4096
	SM_SAMEDISPLAYFORMAT           = 81
	SM_SECURE                      = 44
	SM_SERVERR2                    = 89
	SM_SHOWSOUNDS                  = 70
	SM_SHUTTINGDOWN                = 8192
	SM_SLOWMACHINE                 = 73
	SM_STARTER                     = 88
	SM_SWAPBUTTON                  = 23
	SM_SYSTEMDOCKED                = 8196
	SM_TABLETPC                    = 86
	SM_XVIRTUALSCREEN              = 76
	SM_YVIRTUALSCREEN              = 77
	SM_RESERVED1                   = 24
	SM_RESERVED2                   = 25
	SM_RESERVED3                   = 26
	SM_RESERVED4                   = 27
	SM_CMETRICS                    = 76
	SM_CARETBLINKINGENABLED        = 8194
)

const CCHDEVICENAME = 32

const MONITORINFOF_PRIMARY = 1

const (
	ARW_BOTTOMLEFT  = 0
	ARW_BOTTOMRIGHT = 1
	ARW_TOPLEFT     = 2
	ARW_TOPRIGHT    = 3
	ARW_STARTMASK   = 3
	ARW_STARTRIGHT  = 1
	ARW_STARTTOP    = 2
	ARW_LEFT        = 0
	ARW_RIGHT       = 0
	ARW_UP          = 4
	ARW_DOWN        = 4
	ARW_HIDE        = 8
)

const (
	NID_INTEGRATED_TOUCH = 1
	NID_EXTERNAL_TOUCH   = 2
	NID_INTEGRATED_PEN   = 4
	NID_EXTERNAL_PEN     = 8
	NID_MULTI_INPUT      = 64
	NID_READY            = 128
)

// TODO These were all double-checked for correctness and can be used as is:
// Once the above section is empty, this file is good to be checked in.

// https://learn.microsoft.com/windows/win32/controls/bumper-tree-view-control-reference-messages
const (
	TVM_INSERTITEM          = 4402
	TVM_DELETEITEM          = 4353
	TVM_EXPAND              = 4354
	TVM_GETITEMRECT         = 4356
	TVM_GETCOUNT            = 4357
	TVM_GETINDENT           = 4358
	TVM_SETINDENT           = 4359
	TVM_GETIMAGELIST        = 4360
	TVM_SETIMAGELIST        = 4361
	TVM_GETNEXTITEM         = 4362
	TVM_SELECTITEM          = 4363
	TVM_GETITEM             = 4414
	TVM_SETITEM             = 4415
	TVM_EDITLABEL           = 4417
	TVM_GETEDITCONTROL      = 4367
	TVM_GETVISIBLECOUNT     = 4368
	TVM_HITTEST             = 4369
	TVM_CREATEDRAGIMAGE     = 4370
	TVM_SORTCHILDREN        = 4371
	TVM_ENSUREVISIBLE       = 4372
	TVM_SORTCHILDRENCB      = 4373
	TVM_ENDEDITLABELNOW     = 4374
	TVM_GETISEARCHSTRING    = 4416
	TVM_SETTOOLTIPS         = 4376
	TVM_GETTOOLTIPS         = 4377
	TVM_SETINSERTMARK       = 4378
	TVM_SETUNICODEFORMAT    = 8197
	TVM_GETUNICODEFORMAT    = 8198
	TVM_SETITEMHEIGHT       = 4379
	TVM_GETITEMHEIGHT       = 4380
	TVM_SETBKCOLOR          = 4381
	TVM_SETTEXTCOLOR        = 4382
	TVM_GETBKCOLOR          = 4383
	TVM_GETTEXTCOLOR        = 4384
	TVM_SETSCROLLTIME       = 4385
	TVM_GETSCROLLTIME       = 4386
	TVM_SETINSERTMARKCOLOR  = 4389
	TVM_GETINSERTMARKCOLOR  = 4390
	TVM_SETBORDER           = 4387
	TVM_GETITEMSTATE        = 4391
	TVM_SETLINECOLOR        = 4392
	TVM_GETLINECOLOR        = 4393
	TVM_MAPACCIDTOHTREEITEM = 4394
	TVM_MAPHTREEITEMTOACCID = 4395
	TVM_SETEXTENDEDSTYLE    = 4396
	TVM_GETEXTENDEDSTYLE    = 4397
	TVM_SETAUTOSCROLLINFO   = 4411
	TVM_SETHOT              = 4410
	TVM_GETSELECTEDCOUNT    = 4422
	TVM_SHOWINFOTIP         = 4423
	TVM_GETITEMPARTRECT     = 4424
)

// https://learn.microsoft.com/windows/win32/controls/tree-view-control-item-states
const (
	TVIS_SELECTED       = 2
	TVIS_CUT            = 4
	TVIS_DROPHILITED    = 8
	TVIS_BOLD           = 16
	TVIS_EXPANDED       = 32
	TVIS_EXPANDEDONCE   = 64
	TVIS_EXPANDPARTIAL  = 128
	TVIS_OVERLAYMASK    = 3840
	TVIS_STATEIMAGEMASK = 61440
	TVIS_USERMASK       = 61440
)

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-tvitemexw
//
// NOTE At the time of this writing (Jan 2024) this link describes
// TVIS_EX_HWND, which is not defined in C. It does not mention TVIS_EX_ALL
// however.
const (
	TVIS_EX_FLAT     = 1
	TVIS_EX_DISABLED = 2
	TVIS_EX_ALL      = 2
)

// https://learn.microsoft.com/windows/win32/controls/tree-view-control-window-extended-styles
const (
	TVS_HASBUTTONS      = 1
	TVS_HASLINES        = 2
	TVS_LINESATROOT     = 4
	TVS_EDITLABELS      = 8
	TVS_DISABLEDRAGDROP = 16
	TVS_SHOWSELALWAYS   = 32
	TVS_RTLREADING      = 64
	TVS_NOTOOLTIPS      = 128
	TVS_CHECKBOXES      = 256
	TVS_TRACKSELECT     = 512
	TVS_SINGLEEXPAND    = 1024
	TVS_INFOTIP         = 2048
	TVS_FULLROWSELECT   = 4096
	TVS_NOSCROLL        = 8192
	TVS_NONEVENHEIGHT   = 16384
	TVS_NOHSCROLL       = 32768
)

// https://learn.microsoft.com/windows/win32/controls/tree-view-control-window-extended-styles
const (
	TVS_EX_NOSINGLECOLLAPSE    = 1
	TVS_EX_MULTISELECT         = 2
	TVS_EX_DOUBLEBUFFER        = 4
	TVS_EX_NOINDENTSTATE       = 8
	TVS_EX_RICHTOOLTIP         = 16
	TVS_EX_AUTOHSCROLL         = 32
	TVS_EX_FADEINOUTEXPANDOS   = 64
	TVS_EX_PARTIALCHECKBOXES   = 128
	TVS_EX_EXCLUSIONCHECKBOXES = 256
	TVS_EX_DIMMEDCHECKBOXES    = 512
	TVS_EX_DRAWIMAGEASYNC      = 1024
)

// https://learn.microsoft.com/windows/win32/controls/bumper-tree-view-control-reference-notifications
const (
	TVN_FIRST          = 0xFFFFFE70
	TVN_GETDISPINFO    = TVN_FIRST - 52
	TVN_SETDISPINFO    = TVN_FIRST - 53
	TVN_SELCHANGING    = TVN_FIRST - 50
	TVN_SELCHANGED     = TVN_FIRST - 51
	TVN_ITEMEXPANDING  = TVN_FIRST - 54
	TVN_ITEMEXPANDED   = TVN_FIRST - 55
	TVN_BEGINDRAG      = TVN_FIRST - 56
	TVN_BEGINRDRAG     = TVN_FIRST - 57
	TVN_DELETEITEM     = TVN_FIRST - 58
	TVN_BEGINLABELEDIT = TVN_FIRST - 59
	TVN_ENDLABELEDIT   = TVN_FIRST - 60
	TVN_KEYDOWN        = TVN_FIRST - 12
	TVN_GETINFOTIP     = TVN_FIRST - 14
	TVN_SINGLEEXPAND   = TVN_FIRST - 15
	TVN_ITEMCHANGING   = TVN_FIRST - 17
	TVN_ITEMCHANGED    = TVN_FIRST - 19
	TVN_ASYNCDRAW      = TVN_FIRST - 20
	TVN_LAST           = TVN_FIRST - 99
)

// Notifications are used for multiple GUI controls, e.g.:
//
// https://learn.microsoft.com/windows/win32/controls/bumper-tree-view-control-reference-notifications
// https://learn.microsoft.com/windows/win32/controls/bumper-list-view-control-reference-notifications
const (
	NM_FIRST                = 0
	NM_OUTOFMEMORY          = 4294967295
	NM_CLICK                = 4294967294
	NM_DBLCLK               = 4294967293
	NM_RETURN               = 4294967292
	NM_RCLICK               = 4294967291
	NM_RDBLCLK              = 4294967290
	NM_SETFOCUS             = 4294967289
	NM_KILLFOCUS            = 4294967288
	NM_CUSTOMDRAW           = 4294967284
	NM_HOVER                = 4294967283
	NM_NCHITTEST            = 4294967282
	NM_KEYDOWN              = 4294967281
	NM_RELEASEDCAPTURE      = 4294967280
	NM_SETCURSOR            = 4294967279
	NM_CHAR                 = 4294967278
	NM_TOOLTIPSCREATED      = 4294967277
	NM_LDOWN                = 4294967276
	NM_RDOWN                = 4294967275
	NM_THEMECHANGED         = 4294967274
	NM_FONTCHANGED          = 4294967273
	NM_CUSTOMTEXT           = 4294967272
	NM_TVSTATEIMAGECHANGING = 4294967272
	NM_LAST                 = 4294967197
	NM_GETCUSTOMSPLITRECT   = 4294966049
)

// https://learn.microsoft.com/windows/win32/controls/tvm-expand
const (
	TVE_COLLAPSE      = 1
	TVE_EXPAND        = 2
	TVE_TOGGLE        = 3
	TVE_EXPANDPARTIAL = 0x4000
	TVE_COLLAPSERESET = 0x8000
)

// https://learn.microsoft.com/windows/win32/controls/tvn-selchanged
const (
	TVC_UNKNOWN    = 0
	TVC_BYMOUSE    = 1
	TVC_BYKEYBOARD = 2
)

// https://learn.microsoft.com/windows/win32/controls/about-custom-draw
const (
	CDDS_POSTPAINT     = 2
	CDDS_PREERASE      = 3
	CDDS_PREPAINT      = 1
	CDDS_ITEMPOSTERASE = 65540
	CDDS_ITEMPOSTPAINT = 65538
	CDDS_ITEMPREERASE  = 65539
	CDDS_ITEMPREPAINT  = 65537
	CDDS_SUBITEM       = 131072
)

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-nmcustomdraw
const (
	CDIS_SELECTED         = 1
	CDIS_GRAYED           = 2
	CDIS_DISABLED         = 4
	CDIS_CHECKED          = 8
	CDIS_FOCUS            = 0x10
	CDIS_DEFAULT          = 0x20
	CDIS_HOT              = 0x40
	CDIS_MARKED           = 0x80
	CDIS_INDETERMINATE    = 0x100
	CDIS_SHOWKEYBOARDCUES = 0x200
	CDIS_NEARHOT          = 0x400
	CDIS_OTHERSIDEHOT     = 0x800
	CDIS_DROPHILITED      = 0x1000
)

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-nmtvdispinfow
// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-tvitemw
const (
	TVIF_CHILDREN      = 64
	TVIF_DI_SETITEM    = 4096
	TVIF_HANDLE        = 16
	TVIF_IMAGE         = 2
	TVIF_PARAM         = 4
	TVIF_SELECTEDIMAGE = 32
	TVIF_STATE         = 8
	TVIF_TEXT          = 1
	TVIF_EXPANDEDIMAGE = 512
	TVIF_INTEGRAL      = 128
	TVIF_STATEEX       = 256
)

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-tvgetitempartrectinfo
const TVGIPR_BUTTON = 1

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-tvinsertstructw
const (
	TVI_ROOT  = ^HTREEITEM(65536) + 1
	TVI_FIRST = ^HTREEITEM(65535) + 1
	TVI_LAST  = ^HTREEITEM(65534) + 1
	TVI_SORT  = ^HTREEITEM(65533) + 1
)

// https://learn.microsoft.com/windows/win32/controls/tvm-selectitem
const (
	TVGN_ROOT            = 0
	TVGN_NEXT            = 1
	TVGN_PREVIOUS        = 2
	TVGN_PARENT          = 3
	TVGN_CHILD           = 4
	TVGN_FIRSTVISIBLE    = 5
	TVGN_NEXTVISIBLE     = 6
	TVGN_PREVIOUSVISIBLE = 7
	TVGN_DROPHILITE      = 8
	TVGN_CARET           = 9
	TVGN_LASTVISIBLE     = 10
	TVGN_NEXTSELECTED    = 11
)

// https://learn.microsoft.com/windows/win32/controls/tvm-selectitem
const TVSI_NOSINGLEEXPAND = 32768

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-setfocus
const ERROR_INVALID_PARAMETER = 87

// https://learn.microsoft.com/windows/win32/api/commctrl/ns-commctrl-tvhittestinfo
const (
	TVHT_NOWHERE         = 1
	TVHT_ONITEMICON      = 2
	TVHT_ONITEMLABEL     = 4
	TVHT_ONITEMINDENT    = 8
	TVHT_ONITEMBUTTON    = 0x10
	TVHT_ONITEMRIGHT     = 0x20
	TVHT_ONITEMSTATEICON = 0x40
	TVHT_ONITEM          = 0x46
	TVHT_ABOVE           = 0x100
	TVHT_BELOW           = 0x200
	TVHT_TORIGHT         = 0x400
	TVHT_TOLEFT          = 0x800
)

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-messagebox
const (
	MB_COMPOSITE                 = 2
	MB_ERR_INVALID_CHARS         = 8
	MB_PRECOMPOSED               = 1
	MB_USEGLYPHCHARS             = 4
	MB_ABORTRETRYIGNORE          = 2
	MB_CANCELTRYCONTINUE         = 6
	MB_HELP                      = 16384
	MB_OK                        = 0
	MB_OKCANCEL                  = 1
	MB_RETRYCANCEL               = 5
	MB_YESNO                     = 4
	MB_YESNOCANCEL               = 3
	MB_ICONHAND                  = 16
	MB_ICONQUESTION              = 32
	MB_ICONEXCLAMATION           = 48
	MB_ICONASTERISK              = 64
	MB_USERICON                  = 128
	MB_ICONWARNING               = 48
	MB_ICONERROR                 = 16
	MB_ICONINFORMATION           = 64
	MB_ICONSTOP                  = 16
	MB_DEFBUTTON1                = 0
	MB_DEFBUTTON2                = 256
	MB_DEFBUTTON3                = 512
	MB_DEFBUTTON4                = 768
	MB_APPLMODAL                 = 0
	MB_SYSTEMMODAL               = 4096
	MB_TASKMODAL                 = 8192
	MB_NOFOCUS                   = 32768
	MB_SETFOREGROUND             = 65536
	MB_DEFAULT_DESKTOP_ONLY      = 131072
	MB_TOPMOST                   = 262144
	MB_RIGHT                     = 524288
	MB_RTLREADING                = 1048576
	MB_SERVICE_NOTIFICATION      = 2097152
	MB_SERVICE_NOTIFICATION_NT3X = 262144
	MB_TYPEMASK                  = 15
	MB_ICONMASK                  = 240
	MB_DEFMASK                   = 3840
	MB_MODEMASK                  = 12288
	MB_MISCMASK                  = 49152
)

// https://learn.microsoft.com/previous-versions/dd743680(v=vs.85)
const (
	SND_RING        = 1048576
	SND_ALIAS_START = 0
	SND_APPLICATION = 128
	SND_ALIAS       = 65536
	SND_ALIAS_ID    = 1114112
	SND_FILENAME    = 131072
	SND_RESOURCE    = 262148
	SND_ASYNC       = 1
	SND_NODEFAULT   = 2
	SND_LOOP        = 8
	SND_MEMORY      = 4
	SND_NOSTOP      = 16
	SND_NOWAIT      = 8192
	SND_PURGE       = 64
	SND_SENTRY      = 524288
	SND_SYNC        = 0
	SND_SYSTEM      = 2097152
)

// https://learn.microsoft.com/windows/win32/controls/edit-control-window-extended-styles
// https://learn.microsoft.com/windows/win32/controls/rich-edit-control-styles
const (
	ES_EX_ALLOWEOL_CR          = 1
	ES_EX_ALLOWEOL_LF          = 2
	ES_EX_CONVERT_EOL_ON_PASTE = 4
	ES_EX_ZOOMABLE             = 16
	ES_AWAYMODE_REQUIRED       = 64
	ES_CONTINUOUS              = 2147483648
	ES_DISPLAY_REQUIRED        = 2
	ES_SYSTEM_REQUIRED         = 1
	ES_USER_PRESENT            = 4
	ES_SAVESEL                 = 32768
	ES_SUNKEN                  = 16384
	ES_DISABLENOSCROLL         = 8192
	ES_SELECTIONBAR            = 16777216
	ES_NOOLEDRAGDROP           = 8
	ES_EX_NOCALLOLEINIT        = 0
	ES_VERTICAL                = 4194304
	ES_NOIME                   = 524288
	ES_SELFIME                 = 262144
	ES_LEFT                    = 0
	ES_CENTER                  = 1
	ES_RIGHT                   = 2
	ES_MULTILINE               = 4
	ES_UPPERCASE               = 8
	ES_LOWERCASE               = 16
	ES_PASSWORD                = 32
	ES_AUTOVSCROLL             = 64
	ES_AUTOHSCROLL             = 128
	ES_NOHIDESEL               = 256
	ES_OEMCONVERT              = 1024
	ES_READONLY                = 2048
	ES_WANTRETURN              = 4096
	ES_NUMBER                  = 8192
)

// https://learn.microsoft.com/windows/win32/controls/static-control-styles
const (
	SS_LEFT            = 0
	SS_CENTER          = 1
	SS_RIGHT           = 2
	SS_ICON            = 3
	SS_BLACKRECT       = 4
	SS_GRAYRECT        = 5
	SS_WHITERECT       = 6
	SS_BLACKFRAME      = 7
	SS_GRAYFRAME       = 8
	SS_WHITEFRAME      = 9
	SS_USERITEM        = 10
	SS_SIMPLE          = 11
	SS_LEFTNOWORDWRAP  = 12
	SS_OWNERDRAW       = 13
	SS_BITMAP          = 14
	SS_ENHMETAFILE     = 15
	SS_ETCHEDHORZ      = 16
	SS_ETCHEDVERT      = 17
	SS_ETCHEDFRAME     = 18
	SS_TYPEMASK        = 31
	SS_REALSIZECONTROL = 64
	SS_NOPREFIX        = 128
	SS_NOTIFY          = 256
	SS_CENTERIMAGE     = 512
	SS_RIGHTJUST       = 1024
	SS_REALSIZEIMAGE   = 2048
	SS_SUNKEN          = 4096
	SS_EDITCONTROL     = 8192
	SS_ENDELLIPSIS     = 16384
	SS_PATHELLIPSIS    = 32768
	SS_WORDELLIPSIS    = 49152
	SS_ELLIPSISMASK    = 49152
)

// https://learn.microsoft.com/windows/win32/api/winbase/ns-winbase-actctxw
const (
	ACTCTX_FLAG_PROCESSOR_ARCHITECTURE_VALID           = 1
	ACTCTX_FLAG_LANGID_VALID                           = 2
	ACTCTX_FLAG_ASSEMBLY_DIRECTORY_VALID               = 4
	ACTCTX_FLAG_RESOURCE_NAME_VALID                    = 8
	ACTCTX_FLAG_SET_PROCESS_DEFAULT                    = 16
	ACTCTX_FLAG_APPLICATION_NAME_VALID                 = 32
	ACTCTX_FLAG_SOURCE_IS_ASSEMBLYREF                  = 64
	ACTCTX_FLAG_HMODULE_VALID                          = 128
	ACTCTX_RUN_LEVEL_UNSPECIFIED                       = 0
	ACTCTX_RUN_LEVEL_AS_INVOKER                        = 1
	ACTCTX_RUN_LEVEL_HIGHEST_AVAILABLE                 = 2
	ACTCTX_RUN_LEVEL_REQUIRE_ADMIN                     = 3
	ACTCTX_RUN_LEVEL_NUMBERS                           = 4
	ACTCTX_COMPATIBILITY_ELEMENT_TYPE_UNKNOWN          = 0
	ACTCTX_COMPATIBILITY_ELEMENT_TYPE_OS               = 1
	ACTCTX_COMPATIBILITY_ELEMENT_TYPE_MITIGATION       = 2
	ACTCTX_COMPATIBILITY_ELEMENT_TYPE_MAXVERSIONTESTED = 3
)

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-createfontw
const (
	FW_DONTCARE   = 0
	FW_THIN       = 100
	FW_EXTRALIGHT = 200
	FW_ULTRALIGHT = 200
	FW_LIGHT      = 300
	FW_NORMAL     = 400
	FW_REGULAR    = 400
	FW_MEDIUM     = 500
	FW_SEMIBOLD   = 600
	FW_DEMIBOLD   = 600
	FW_BOLD       = 700
	FW_EXTRABOLD  = 800
	FW_ULTRABOLD  = 800
	FW_HEAVY      = 900
	FW_BLACK      = 900
)

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-createfontw
const (
	ANSI_CHARSET        = 0
	DEFAULT_CHARSET     = 1
	SYMBOL_CHARSET      = 2
	CHINESEBIG5_CHARSET = 136
	EASTEUROPE_CHARSET  = 238
	GREEK_CHARSET       = 161
	MAC_CHARSET         = 77
	RUSSIAN_CHARSET     = 204
	SHIFTJIS_CHARSET    = 128
	HANGEUL_CHARSET     = 129
	HANGUL_CHARSET      = 129
	GB2312_CHARSET      = 134
	OEM_CHARSET         = 255
	JOHAB_CHARSET       = 130
	HEBREW_CHARSET      = 177
	ARABIC_CHARSET      = 178
	TURKISH_CHARSET     = 162
	VIETNAMESE_CHARSET  = 163
	THAI_CHARSET        = 222
	BALTIC_CHARSET      = 186
	CFM_CHARSET         = 134217728
)

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-createfontw
const (
	CLIP_DEFAULT_PRECIS   = 0
	CLIP_CHARACTER_PRECIS = 1
	CLIP_STROKE_PRECIS    = 2
)

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-createfontw
const (
	OUT_DEFAULT_PRECIS        = 0
	OUT_STRING_PRECIS         = 1
	OUT_CHARACTER_PRECIS      = 2
	OUT_STROKE_PRECIS         = 3
	OUT_TT_PRECIS             = 4
	OUT_DEVICE_PRECIS         = 5
	OUT_RASTER_PRECIS         = 6
	OUT_TT_ONLY_PRECIS        = 7
	OUT_OUTLINE_PRECIS        = 8
	OUT_SCREEN_OUTLINE_PRECIS = 9
	OUT_PS_ONLY_PRECIS        = 10
)

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-createfontw
const (
	DEFAULT_QUALITY           = 0
	DRAFT_QUALITY             = 1
	PROOF_QUALITY             = 2
	ANTIALIASED_QUALITY       = 4
	NONANTIALIASED_QUALITY    = 3
	CLEARTYPE_NATURAL_QUALITY = 6
	CLEARTYPE_QUALITY         = 5
)

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-createfontw
const (
	DEFAULT_PITCH    = 0
	FIXED_PITCH      = 1
	VARIABLE_PITCH   = 2
	WAVECAPS_PITCH   = 1
	TMPF_FIXED_PITCH = 1
)

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-createfontw
const (
	OBJ_FONT            = 6
	MONO_FONT           = 8
	ANSI_FIXED_FONT     = 11
	ANSI_VAR_FONT       = 12
	DEVICE_DEFAULT_FONT = 14
	DEFAULT_GUI_FONT    = 17
	OEM_FIXED_FONT      = 10
	SYSTEM_FONT         = 13
	SYSTEM_FIXED_FONT   = 16
	VFT_FONT            = 4
	RT_FONT             = 8
)

// https://learn.microsoft.com/windows/win32/api/wingdi/nf-wingdi-createfontw
const (
	FF_DONTCARE   = 0
	FF_ROMAN      = 16
	FF_SWISS      = 32
	FF_MODERN     = 48
	FF_SCRIPT     = 64
	FF_DECORATIVE = 80
)

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-axeslistw
const STAMP_DESIGNVECTOR = 134248036

// https://learn.microsoft.com/windows/win32/api/shtypes/ns-shtypes-logfontw
const LF_FACESIZE = 32

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-enumlogfontexw
const LF_FULLFACESIZE = 64

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-designvector
// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-axeslistw
const MM_MAX_NUMAXES = 16

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-axisinfow
const MM_MAX_AXES_NAMELEN = 16

// https://learn.microsoft.com/windows/win32/api/wingdi/ns-wingdi-newtextmetricw
const (
	NTM_REGULAR        = 64
	NTM_BOLD           = 32
	NTM_ITALIC         = 1
	NTM_NONNEGATIVE_AC = 65536
	NTM_PS_OPENTYPE    = 131072
	NTM_TT_OPENTYPE    = 262144
	NTM_MULTIPLEMASTER = 524288
	NTM_TYPE1          = 1048576
	NTM_DSIG           = 2097152
)

const (
	EM_SCROLLCARET      = 183
	EM_SETLIMITTEXT     = 197
	EM_GETLIMITTEXT     = 213
	EM_POSFROMCHAR      = 214
	EM_CHARFROMPOS      = 215
	EM_SETCUEBANNER     = 5377
	EM_GETCUEBANNER     = 5378
	EM_SHOWBALLOONTIP   = 5379
	EM_HIDEBALLOONTIP   = 5380
	EM_SETHILITE        = 5381
	EM_GETHILITE        = 5382
	EM_NOSETFOCUS       = 5383
	EM_TAKEFOCUS        = 5384
	EM_SETEXTENDEDSTYLE = 5386
	EM_GETEXTENDEDSTYLE = 5387
	EM_SETENDOFLINE     = 5388
	EM_GETENDOFLINE     = 5389
	EM_ENABLESEARCHWEB  = 5390
	EM_SEARCHWEB        = 5391
	EM_SETCARETINDEX    = 5393
	EM_GETCARETINDEX    = 5394
	EM_FILELINEFROMCHAR = 5395
	EM_FILELINEINDEX    = 5396
	EM_FILELINELENGTH   = 5397
	EM_GETFILELINE      = 5398
	EM_GETFILELINECOUNT = 5399
	EM_GETSEL           = 176
	EM_SETSEL           = 177
	EM_GETRECT          = 178
	EM_SETRECT          = 179
	EM_SETRECTNP        = 180
	EM_SCROLL           = 181
	EM_LINESCROLL       = 182
	EM_GETMODIFY        = 184
	EM_SETMODIFY        = 185
	EM_GETLINECOUNT     = 186
	EM_LINEINDEX        = 187
	EM_SETHANDLE        = 188
	EM_GETHANDLE        = 189
	EM_GETTHUMB         = 190
	EM_LINELENGTH       = 193
	EM_REPLACESEL       = 194
	EM_GETLINE          = 196
	EM_LIMITTEXT        = 197
	EM_CANUNDO          = 198
	EM_UNDO             = 199
	EM_FMTLINES         = 200
	EM_LINEFROMCHAR     = 201
	//EM_SETTABSTOPS  uses "dialog units" where 1 dialog unit is one fourth of
	//the average font width.
	//
	// https://learn.microsoft.com/previous-versions/windows/desktop/bb226789(v=vs.85)
	EM_SETTABSTOPS          = 203
	EM_SETPASSWORDCHAR      = 204
	EM_EMPTYUNDOBUFFER      = 205
	EM_GETFIRSTVISIBLELINE  = 206
	EM_SETREADONLY          = 207
	EM_SETWORDBREAKPROC     = 208
	EM_GETWORDBREAKPROC     = 209
	EM_GETPASSWORDCHAR      = 210
	EM_SETMARGINS           = 211
	EM_GETMARGINS           = 212
	EM_SETIMESTATUS         = 216
	EM_GETIMESTATUS         = 217
	EM_ENABLEFEATURE        = 218
	EM_CANPASTE             = 1074
	EM_DISPLAYBAND          = 1075
	EM_EXGETSEL             = 1076
	EM_EXLIMITTEXT          = 1077
	EM_EXLINEFROMCHAR       = 1078
	EM_EXSETSEL             = 1079
	EM_FORMATRANGE          = 1081
	EM_GETCHARFORMAT        = 1082
	EM_GETEVENTMASK         = 1083
	EM_GETOLEINTERFACE      = 1084
	EM_GETPARAFORMAT        = 1085
	EM_GETSELTEXT           = 1086
	EM_HIDESELECTION        = 1087
	EM_PASTESPECIAL         = 1088
	EM_REQUESTRESIZE        = 1089
	EM_SELECTIONTYPE        = 1090
	EM_SETBKGNDCOLOR        = 1091
	EM_SETCHARFORMAT        = 1092
	EM_SETEVENTMASK         = 1093
	EM_SETOLECALLBACK       = 1094
	EM_SETPARAFORMAT        = 1095
	EM_SETTARGETDEVICE      = 1096
	EM_STREAMIN             = 1097
	EM_STREAMOUT            = 1098
	EM_GETTEXTRANGE         = 1099
	EM_FINDWORDBREAK        = 1100
	EM_SETOPTIONS           = 1101
	EM_GETOPTIONS           = 1102
	EM_GETWORDBREAKPROCEX   = 1104
	EM_SETWORDBREAKPROCEX   = 1105
	EM_SETUNDOLIMIT         = 1106
	EM_REDO                 = 1108
	EM_CANREDO              = 1109
	EM_GETUNDONAME          = 1110
	EM_GETREDONAME          = 1111
	EM_STOPGROUPTYPING      = 1112
	EM_SETTEXTMODE          = 1113
	EM_GETTEXTMODE          = 1114
	EM_AUTOURLDETECT        = 1115
	EM_GETAUTOURLDETECT     = 1116
	EM_SETPALETTE           = 1117
	EM_GETTEXTEX            = 1118
	EM_GETTEXTLENGTHEX      = 1119
	EM_SHOWSCROLLBAR        = 1120
	EM_SETTEXTEX            = 1121
	EM_SETPUNCTUATION       = 1124
	EM_GETPUNCTUATION       = 1125
	EM_SETWORDWRAPMODE      = 1126
	EM_GETWORDWRAPMODE      = 1127
	EM_SETIMECOLOR          = 1128
	EM_GETIMECOLOR          = 1129
	EM_SETIMEOPTIONS        = 1130
	EM_GETIMEOPTIONS        = 1131
	EM_CONVPOSITION         = 1132
	EM_SETLANGOPTIONS       = 1144
	EM_GETLANGOPTIONS       = 1145
	EM_GETIMECOMPMODE       = 1146
	EM_FINDTEXTW            = 1147
	EM_FINDTEXTEXW          = 1148
	EM_RECONVERSION         = 1149
	EM_SETIMEMODEBIAS       = 1150
	EM_GETIMEMODEBIAS       = 1151
	EM_SETBIDIOPTIONS       = 1224
	EM_GETBIDIOPTIONS       = 1225
	EM_SETTYPOGRAPHYOPTIONS = 1226
	EM_GETTYPOGRAPHYOPTIONS = 1227
	EM_SETEDITSTYLE         = 1228
	EM_GETEDITSTYLE         = 1229
	EM_OUTLINE              = 1244
	EM_GETSCROLLPOS         = 1245
	EM_SETSCROLLPOS         = 1246
	EM_SETFONTSIZE          = 1247
	EM_GETZOOM              = 1248
	EM_SETZOOM              = 1249
	EM_GETVIEWKIND          = 1250
	EM_SETVIEWKIND          = 1251
	EM_GETPAGE              = 1252
	EM_SETPAGE              = 1253
	EM_GETHYPHENATEINFO     = 1254
	EM_SETHYPHENATEINFO     = 1255
	EM_GETPAGEROTATE        = 1259
	EM_SETPAGEROTATE        = 1260
	EM_GETCTFMODEBIAS       = 1261
	EM_SETCTFMODEBIAS       = 1262
	EM_GETCTFOPENSTATUS     = 1264
	EM_SETCTFOPENSTATUS     = 1265
	EM_GETIMECOMPTEXT       = 1266
	EM_ISIME                = 1267
	EM_GETIMEPROPERTY       = 1268
	EM_GETQUERYRTFOBJ       = 1293
	EM_SETQUERYRTFOBJ       = 1294
	EM_INSERTTABLE          = 1256
	EM_GETAUTOCORRECTPROC   = 1257
	EM_SETAUTOCORRECTPROC   = 1258
	EM_CALLAUTOCORRECTPROC  = 1279
	EM_GETTABLEPARMS        = 1289
	EM_SETEDITSTYLEEX       = 1299
	EM_GETEDITSTYLEEX       = 1300
	EM_GETSTORYTYPE         = 1314
	EM_SETSTORYTYPE         = 1315
	EM_GETELLIPSISMODE      = 1329
	EM_SETELLIPSISMODE      = 1330
	EM_SETTABLEPARMS        = 1331
	EM_GETTOUCHOPTIONS      = 1334
	EM_SETTOUCHOPTIONS      = 1335
	EM_INSERTIMAGE          = 1338
	EM_SETUIANAME           = 1344
	EM_GETELLIPSISSTATE     = 1346
)

// https://learn.microsoft.com/windows/win32/api/winuser/nf-winuser-settimer
const (
	USER_TIMER_MINIMUM = 10
	USER_TIMER_MAXIMUM = 2147483647
)

// https://learn.microsoft.com/windows/win32/api/winuser/ns-winuser-accel
const (
	FVIRTKEY  = 1
	FNOINVERT = 2
	FSHIFT    = 4
	FCONTROL  = 8
	FALT      = 16
)

// https://learn.microsoft.com/windows/win32/inputdev/wm-mousewheel
const (
	MK_ALT      = 32
	MK_LBUTTON  = 1
	MK_RBUTTON  = 2
	MK_SHIFT    = 4
	MK_CONTROL  = 8
	MK_MBUTTON  = 16
	MK_XBUTTON1 = 32
	MK_XBUTTON2 = 64
)

// https://learn.microsoft.com/windows/win32/controls/bumper-edit-control-reference-notifications
const (
	EN_SEARCHWEB         = 4294965776
	EN_MSGFILTER         = 1792
	EN_REQUESTRESIZE     = 1793
	EN_SELCHANGE         = 1794
	EN_DROPFILES         = 1795
	EN_PROTECTED         = 1796
	EN_CORRECTTEXT       = 1797
	EN_STOPNOUNDO        = 1798
	EN_IMECHANGE         = 1799
	EN_SAVECLIPBOARD     = 1800
	EN_OLEOPFAILED       = 1801
	EN_OBJECTPOSITIONS   = 1802
	EN_LINK              = 1803
	EN_DRAGDROPDONE      = 1804
	EN_PARAGRAPHEXPANDED = 1805
	EN_PAGECHANGE        = 1806
	EN_LOWFIRTF          = 1807
	EN_ALIGNLTR          = 1808
	EN_ALIGNRTL          = 1809
	EN_CLIPFORMAT        = 1810
	EN_STARTCOMPOSITION  = 1811
	EN_ENDCOMPOSITION    = 1812
	EN_SETFOCUS          = 256
	EN_KILLFOCUS         = 512
	EN_CHANGE            = 768
	EN_UPDATE            = 1024
	EN_ERRSPACE          = 1280
	EN_MAXTEXT           = 1281
	EN_HSCROLL           = 1537
	EN_VSCROLL           = 1538
	EN_ALIGN_LTR_EC      = 1792
	EN_ALIGN_RTL_EC      = 1793
	EN_BEFORE_PASTE      = 2048
	EN_AFTER_PASTE       = 2049
)

// TODO
const (
	SB_SETTEXTA         = 1025
	SB_SETTEXTW         = 1035
	SB_GETTEXTA         = 1026
	SB_GETTEXTW         = 1037
	SB_GETTEXTLENGTHA   = 1027
	SB_GETTEXTLENGTHW   = 1036
	SB_GETTEXT          = 1037
	SB_SETTEXT          = 1035
	SB_GETTEXTLENGTH    = 1036
	SB_SETPARTS         = 1028
	SB_GETPARTS         = 1030
	SB_GETBORDERS       = 1031
	SB_SETMINHEIGHT     = 1032
	SB_SIMPLE           = 1033
	SB_GETRECT          = 1034
	SB_ISSIMPLE         = 1038
	SB_SETICON          = 1039
	SB_SETTIPTEXTA      = 1040
	SB_SETTIPTEXTW      = 1041
	SB_GETTIPTEXTA      = 1042
	SB_GETTIPTEXTW      = 1043
	SB_GETICON          = 1044
	SB_SETUNICODEFORMAT = 8197
	SB_GETUNICODEFORMAT = 8198
	SB_SETBKCOLOR       = 8193
	SB_SIMPLEID         = 255
	SB_NONE             = 0
	SB_CONST_ALPHA      = 1
	SB_PIXEL_ALPHA      = 2
	SB_PREMULT_ALPHA    = 4
	SB_GRAD_RECT        = 16
	SB_GRAD_TRI         = 32
	SB_CTL              = 2
	SB_HORZ             = 0
	SB_VERT             = 1
	SB_BOTH             = 3
	SB_LINEUP           = 0
	SB_LINELEFT         = 0
	SB_LINEDOWN         = 1
	SB_LINERIGHT        = 1
	SB_PAGEUP           = 2
	SB_PAGELEFT         = 2
	SB_PAGEDOWN         = 3
	SB_PAGERIGHT        = 3
	SB_THUMBPOSITION    = 4
	SB_THUMBTRACK       = 5
	SB_TOP              = 6
	SB_LEFT             = 6
	SB_RIGHT            = 7
	SB_BOTTOM           = 7
	SB_ENDSCROLL        = 8
)
