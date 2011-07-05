package com

import (
    "unsafe"
    . "w32"
)

var (
    IID_NULL                      = &GUID{0x00000000, 0x0000, 0x0000, [8]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}}
    IID_IUnknown                  = &GUID{0x00000000, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
    IID_IDispatch                 = &GUID{0x00020400, 0x0000, 0x0000, [8]byte{0xC0, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x46}}
    IID_IConnectionPointContainer = &GUID{0xB196B284, 0xBAB4, 0x101A, [8]byte{0xB6, 0x9C, 0x00, 0xAA, 0x00, 0x34, 0x1D, 0x07}}
    IID_IConnectionPoint          = &GUID{0xB196B286, 0xBAB4, 0x101A, [8]byte{0xB6, 0x9C, 0x00, 0xAA, 0x00, 0x34, 0x1D, 0x07}}
)

type pIDispatchVtbl struct {
    pQueryInterface   uintptr
    pAddRef           uintptr
    pRelease          uintptr
    pGetTypeInfoCount uintptr
    pGetTypeInfo      uintptr
    pGetIDsOfNames    uintptr
    pInvoke           uintptr
}

type IDispatch struct {
    lpVtbl *pIDispatchVtbl
}

func (this *IDispatch) QueryInterface(id *GUID) (disp *IDispatch) {
    disp = ComQueryInterface((*IUnknown)(unsafe.Pointer(this)), id)
    return
}

func (this *IDispatch) AddRef() int32 {
    return ComAddRef((*IUnknown)(unsafe.Pointer(this)))
}

func (this *IDispatch) Release() int32 {
    return ComRelease((*IUnknown)(unsafe.Pointer(this)))
}

func (this *IDispatch) GetIDsOfName(names []string) []int32 {
    return ComGetIDsOfName(this, names)
}

func (this *IDispatch) Invoke(dispid int32, dispatch int16, params ...interface{}) *VARIANT {
	//TODO: Implement Invoke here.
    return nil
}

type pIUnknownVtbl struct {
    pQueryInterface uintptr
    pAddRef         uintptr
    pRelease        uintptr
}

type IUnknown struct {
    lpVtbl *pIUnknownVtbl
}

func (this *IUnknown) QueryInterface(id *GUID) (disp *IDispatch) {
    disp = ComQueryInterface(this, id)
    return
}

func (this *IUnknown) AddRef() int32 {
    return ComAddRef(this)
}

func (this *IUnknown) Release() int32 {
    return ComRelease(this)
}
