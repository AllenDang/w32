package com

import (
    . "github.com/AllenDang/w32"
)

type pIUnknownVtbl struct {
    pQueryInterface uintptr
    pAddRef         uintptr
    pRelease        uintptr
}

type IUnknown struct {
    lpVtbl *pIUnknownVtbl
}

func (this *IUnknown) QueryInterface(id *GUID) *IDispatch {
    return ComQueryInterface(this, id)
}

func (this *IUnknown) AddRef() int32 {
    return ComAddRef(this)
}

func (this *IUnknown) Release() int32 {
    return ComRelease(this)
}
