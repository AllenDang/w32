package com

import (
    . "github.com/AllenDang/w32"
)

type pIStreamVtbl struct {
    pQueryInterface uintptr
    pAddRef         uintptr
    pRelease        uintptr
}

type IStream struct {
    lpVtbl *pIStreamVtbl
}

func (this *IStream) QueryInterface(id *GUID) *IDispatch {
    return ComQueryInterface((*IUnknown)(unsafe.Pointer(this)), id)
}

func (this *IStream) AddRef() int32 {
    return ComAddRef((*IUnknown)(unsafe.Pointer(this)))
}

func (this *IStream) Release() int32 {
    return ComRelease((*IUnknown)(unsafe.Pointer(this)))
}
