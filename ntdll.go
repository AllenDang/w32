// Copyright 2010-2012 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package w32

import (
	"errors"
	"syscall"
	"unsafe"
)

var (
	modntdll = syscall.NewLazyDLL("ntdll.dll")

	procNtAlpcCreatePort = modntdll.NewProc("NtAlpcCreatePort")
)

func newUnicodeString(s string) (us UNICODE_STRING, e error) {
	ustr, err := syscall.UTF16FromString(s)
	if err != nil {
		e = err
		return
	}
	us.Length = len(ustr)
	us.MaximumLength = len(ustr)
	us.Buffer = unsafe.Pointer(&ustr[0])
	return
}

// (this is a macro)
// VOID InitializeObjectAttributes(
//   [out]           POBJECT_ATTRIBUTES InitializedAttributes,
//   [in]            PUNICODE_STRING ObjectName,
//   [in]            ULONG Attributes,
//   [in]            HANDLE RootDirectory,
//   [in, optional]  PSECURITY_DESCRIPTOR SecurityDescriptor
// )
func NewObjectAttributes(
	name string,
	attributes uint32,
	rootDir HANDLE,
	pSecurityDescriptor *SECURITY_DESCRIPTOR,
) (objectAttributes OBJECT_ATTRIBUTES, e error) {

	unicodeString, err := newUnicodeString(name)
	if err != nil {
		e = err
		return
	}

	objectAttributes = OBJECT_ATTRIBUTES{
		RootDirectory:      rootDir,
		ObjectName:         &unicodeString,
		Attributes:         attributes,
		SecurityDescriptor: pSecurityDescriptor,
	}
	return
}

// # NTSTATUS
// # NtAlpcCreatePort(
// #   __out PHANDLE PortHandle,
// #   __in POBJECT_ATTRIBUTES ObjectAttributes,
// #   __in_opt PALPC_PORT_ATTRIBUTES PortAttributes
// #   );
func NtAlpcCreatePort(pObjectAttributes *OBJECT_ATTRIBUTES, pPortAttributes *ALPC_PORT_ATTRIBUTES) (hPort HANDLE, e error) {

	pHandle := &hPort
	ret, _, _ := procNtAlpcCreatePort.Call(
		uintptr(unsafe.Pointer(pHandle)),
		uintptr(unsafe.Pointer(pObjectAttributes)),
		uintptr(unsafe.Pointer(pPortAttributes)),
	)
	if ret != ERROR_SUCCESS {
		return hPort, errors.New(ret)
	}

	return
}
