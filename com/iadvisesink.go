package com

import (
    . "github.com/AllenDang/w32"
    "unsafe"
)

type pIAdviseSinkVtbl struct {
    pQueryInterface uintptr
    pAddRef         uintptr
    pRelease        uintptr
    pOnClose        uintptr
    pOnDataChange   uintptr
    pOnRename       uintptr
    pOnSave         uintptr
    pOnViewChange   uintptr
}

type IAdviseSink struct {
    lpVtbl *pIAdviseSinkVtbl
}

func (this *IAdviseSink) QueryInterface(id *GUID) *IDispatch {
    return ComQueryInterface((*IUnknown)(unsafe.Pointer(this)), id)
}

func (this *IAdviseSink) AddRef() int32 {
    return ComAddRef((*IUnknown)(unsafe.Pointer(this)))
}

func (this *IAdviseSink) Release() int32 {
    return ComRelease((*IUnknown)(unsafe.Pointer(this)))
}

func (this *IAdviseSink) OnClose() {
    ComOnClose(this)
}

func (this *IAdviseSink) OnDataChange(pFormatetc *FORMATETC, pStgmed *STGMEDIUM) {
    ComOnDataChange(this, pFormatetc, pStgmed)
}
