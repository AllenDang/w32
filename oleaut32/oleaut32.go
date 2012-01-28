package oleaut32

import (
    "syscall"
    "unsafe"
    . "w32"
)

var (
    lib uintptr

    procVariantInit        uintptr
    procSysAllocString     uintptr
    procSysFreeString      uintptr
    procSysStringLen       uintptr
    procCreateDispTypeInfo uintptr
    procCreateStdDispatch  uintptr
)

func init() {
    lib = LoadLib("oleaut32")

    procVariantInit = GetProcAddr(lib, "VariantInit")
    procSysAllocString = GetProcAddr(lib, "SysAllocString")
    procSysFreeString = GetProcAddr(lib, "SysFreeString")
    procSysStringLen = GetProcAddr(lib, "SysStringLen")
    procCreateDispTypeInfo = GetProcAddr(lib, "CreateDispTypeInfo")
    procCreateStdDispatch = GetProcAddr(lib, "CreateStdDispatch")
}

func VariantInit(v *VARIANT) {
    hr, _, _ := syscall.Syscall(procVariantInit, 1,
        uintptr(unsafe.Pointer(v)),
        0,
        0)
    if hr != 0 {
        panic("Invoke VariantInit error.")
    }
    return
}

func SysAllocString(v string) (ss *int16) {
    pss, _, _ := syscall.Syscall(procSysAllocString, 1,
        uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(v))),
        0,
        0)
    ss = (*int16)(unsafe.Pointer(pss))
    return
}

func SysFreeString(v *int16){
    hr, _, _ := syscall.Syscall(procSysFreeString, 1,
        uintptr(unsafe.Pointer(v)),
        0,
        0)
    if hr != 0 {
        panic("Invoke SysFreeString error.")
    }
    return
}

func SysStringLen(v *int16) uint {
    l, _, _ := syscall.Syscall(procSysStringLen, 1,
        uintptr(unsafe.Pointer(v)),
        0,
        0)
    return uint(l)
}
