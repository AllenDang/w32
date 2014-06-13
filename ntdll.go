// Copyright 2010-2012 The W32 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package w32

import (
	"errors"
	"fmt"
	// "github.com/davecgh/go-spew/spew"
	"syscall"
	"unsafe"
)

var (
	modntdll = syscall.NewLazyDLL("ntdll.dll")

	procNtAlpcCreatePort          = modntdll.NewProc("NtAlpcCreatePort")
	procNtAlpcConnectPort         = modntdll.NewProc("NtAlpcConnectPort")
	procNtAlpcSendWaitReceivePort = modntdll.NewProc("NtAlpcSendWaitReceivePort")
	procNtAlpcAcceptConnectPort   = modntdll.NewProc("NtAlpcAcceptConnectPort")
	procAlpcGetMessageAttribute   = modntdll.NewProc("AlpcGetMessageAttribute")
	procNtAlpcCancelMessage       = modntdll.NewProc("NtAlpcCancelMessage")
)

func newUnicodeString(s string) (us UNICODE_STRING, e error) {
	ustr, err := syscall.UTF16FromString(s)
	if err != nil {
		e = err
		return
	}
	us.Length = uint16(len(ustr) * 2) // in bytes
	us.MaximumLength = uint16(len(ustr) * 2)
	us.Buffer = &ustr[0]
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
func newObjectAttributes(
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
	objectAttributes.Length = uint32(unsafe.Sizeof(objectAttributes))
	return
}

func basicObjectAttributes(name string) (objectAttributes OBJECT_ATTRIBUTES, e error) {

	sd, e := InitializeSecurityDescriptor(1)
	if e != nil {
		return
	}
	e = SetSecurityDescriptorDacl(sd, nil)
	if e != nil {
		return
	}

	objectAttributes, e = newObjectAttributes(name, 0, 0, sd)
	return
}

// NTSTATUS
// NtAlpcCreatePort(
//   __out PHANDLE PortHandle,
//   __in POBJECT_ATTRIBUTES ObjectAttributes,
//   __in_opt PALPC_PORT_ATTRIBUTES PortAttributes
//   );
func NtAlpcCreatePort(pObjectAttributes *OBJECT_ATTRIBUTES, pPortAttributes *ALPC_PORT_ATTRIBUTES) (hPort HANDLE, e error) {

	pHandle := &hPort
	ret, _, _ := procNtAlpcCreatePort.Call(
		uintptr(unsafe.Pointer(pHandle)),
		uintptr(unsafe.Pointer(pObjectAttributes)),
		uintptr(unsafe.Pointer(pPortAttributes)),
	)
	if ret != ERROR_SUCCESS {
		return hPort, errors.New(fmt.Sprintf("0x%x", ret))
	}

	return
}

// NTSTATUS
// NtAlpcConnectPort(
//     __out PHANDLE PortHandle,
//     __in PUNICODE_STRING PortName,
//     __in POBJECT_ATTRIBUTES ObjectAttributes,
//     __in_opt PALPC_PORT_ATTRIBUTES PortAttributes,
//     __in ULONG Flags,
//     __in_opt PSID RequiredServerSid,
//     __inout PPORT_MESSAGE ConnectionMessage,
//     __inout_opt PULONG BufferLength,
//     __inout_opt PALPC_MESSAGE_ATTRIBUTES OutMessageAttributes,
//     __inout_opt PALPC_MESSAGE_ATTRIBUTES InMessageAttributes,
//     __in_opt PLARGE_INTEGER Timeout
//     );
func NtAlpcConnectPort(
	destPort string,
	pClientObjAttrs *OBJECT_ATTRIBUTES,
	pClientAlpcPortAttrs *ALPC_PORT_ATTRIBUTES,
	flags uint32,
	pRequiredServerSid *SID,
	pConnMsg *AlpcShortMessage,
	pBufLen *uint32,
	pOutMsgAttrs *ALPC_MESSAGE_ATTRIBUTES,
	pInMsgAttrs *ALPC_MESSAGE_ATTRIBUTES,
	timeout *int64,
) (hPort HANDLE, e error) {

	destPortU, e := newUnicodeString(destPort)
	if e != nil {
		return
	}

	ret, _, _ := procNtAlpcConnectPort.Call(
		uintptr(unsafe.Pointer(&hPort)),
		uintptr(unsafe.Pointer(&destPortU)),
		uintptr(unsafe.Pointer(pClientObjAttrs)),
		uintptr(unsafe.Pointer(pClientAlpcPortAttrs)),
		uintptr(flags),
		uintptr(unsafe.Pointer(pRequiredServerSid)),
		uintptr(unsafe.Pointer(pConnMsg)),
		uintptr(unsafe.Pointer(pBufLen)),
		uintptr(unsafe.Pointer(pOutMsgAttrs)),
		uintptr(unsafe.Pointer(pInMsgAttrs)),
		uintptr(unsafe.Pointer(timeout)),
	)

	if ret != ERROR_SUCCESS {
		e = errors.New(fmt.Sprintf("0x%x", ret))
	}
	return
}

// NTSTATUS
// NtAlpcAcceptConnectPort(
//     __out PHANDLE PortHandle,
//     __in HANDLE ConnectionPortHandle,
//     __in ULONG Flags,
//     __in POBJECT_ATTRIBUTES ObjectAttributes,
//     __in PALPC_PORT_ATTRIBUTES PortAttributes,
//     __in_opt PVOID PortContext,
//     __in PPORT_MESSAGE ConnectionRequest,
//     __inout_opt PALPC_MESSAGE_ATTRIBUTES ConnectionMessageAttributes,
//     __in BOOLEAN AcceptConnection
//     );

func NtAlpcAcceptConnectPort(
	hSrvConnPort HANDLE,
	flags uint32,
	pObjAttr *OBJECT_ATTRIBUTES,
	pPortAttr *ALPC_PORT_ATTRIBUTES,
	pContext *AlpcPortContext,
	pConnReq *AlpcShortMessage,
	pConnMsgAttrs *ALPC_MESSAGE_ATTRIBUTES,
	accept uintptr,
) (hPort HANDLE, e error) {

	ret, _, _ := procNtAlpcAcceptConnectPort.Call(
		uintptr(unsafe.Pointer(&hPort)),
		uintptr(hSrvConnPort),
		uintptr(flags),
		uintptr(unsafe.Pointer(pObjAttr)),
		uintptr(unsafe.Pointer(pPortAttr)),
		uintptr(unsafe.Pointer(pContext)),
		uintptr(unsafe.Pointer(pConnReq)),
		uintptr(unsafe.Pointer(pConnMsgAttrs)),
		accept,
	)

	if ret != ERROR_SUCCESS {
		e = errors.New(fmt.Sprintf("0x%x", ret))
	}
	return
}

// NTSTATUS
// NtAlpcSendWaitReceivePort(
//     __in HANDLE PortHandle,
//     __in ULONG Flags,
//     __in_opt PPORT_MESSAGE SendMessage,
//     __in_opt PALPC_MESSAGE_ATTRIBUTES SendMessageAttributes,
//     __inout_opt PPORT_MESSAGE ReceiveMessage,
//     __inout_opt PULONG BufferLength,
//     __inout_opt PALPC_MESSAGE_ATTRIBUTES ReceiveMessageAttributes,
//     __in_opt PLARGE_INTEGER Timeout
//     );
func NtAlpcSendWaitReceivePort(
	hPort HANDLE,
	flags uint32,
	sendMsg *AlpcShortMessage, // Should actually point to PORT_MESSAGE + payload
	sendMsgAttrs *ALPC_MESSAGE_ATTRIBUTES,
	recvMsg *AlpcShortMessage,
	recvBufLen *uint32,
	recvMsgAttrs *ALPC_MESSAGE_ATTRIBUTES,
	timeout *int64, // use native int64
) (e error) {

	ret, _, _ := procNtAlpcSendWaitReceivePort.Call(
		uintptr(hPort),
		uintptr(flags),
		uintptr(unsafe.Pointer(sendMsg)),
		uintptr(unsafe.Pointer(sendMsgAttrs)),
		uintptr(unsafe.Pointer(recvMsg)),
		uintptr(unsafe.Pointer(recvBufLen)),
		uintptr(unsafe.Pointer(recvMsgAttrs)),
		uintptr(unsafe.Pointer(timeout)),
	)

	if ret != ERROR_SUCCESS {
		e = errors.New(fmt.Sprintf("0x%x", ret))
	}
	return
}

// NTSYSAPI
// PVOID
// NTAPI
// AlpcGetMessageAttribute(
//     __in PALPC_MESSAGE_ATTRIBUTES Buffer,
//     __in ULONG AttributeFlag
//     );

// This basically returns a pointer to the correct struct for whichever
// message attribute you asked for. In Go terms, it returns unsafe.Pointer
// which you should then cast. Example:

// ptr := AlpcGetMessageAttribute(&recvMsgAttrs, ALPC_MESSAGE_CONTEXT_ATTRIBUTE)
// if ptr != nil {
//     context := (*ALPC_CONTEXT_ATTR)(ptr)
// }
func AlpcGetMessageAttribute(buf *ALPC_MESSAGE_ATTRIBUTES, attr uint32) uintptr {
	ret, _, _ := procAlpcGetMessageAttribute.Call(
		uintptr(unsafe.Pointer(buf)),
		uintptr(attr),
	)
	return ret
}

// NTSYSCALLAPI
// NTSTATUS
// NTAPI
// NtAlpcCancelMessage(
//     __in HANDLE PortHandle,
//     __in ULONG Flags,
//     __in PALPC_CONTEXT_ATTR MessageContext
//     );
func NtAlpcCancelMessage(hPort HANDLE, flags uint32, pMsgContext *ALPC_CONTEXT_ATTR) (e error) {
	ret, _, _ := procNtAlpcCancelMessage.Call(
		uintptr(hPort),
		uintptr(flags),
		uintptr(unsafe.Pointer(pMsgContext)),
	)
	if ret != ERROR_SUCCESS {
		e = errors.New(fmt.Sprintf("0x%x", ret))
	}
	return
}

var basicPortAttr = ALPC_PORT_ATTRIBUTES{
	MaxMessageLength: uint64(SHORT_MESSAGE_MAX_SIZE),
	SecurityQos: SECURITY_QUALITY_OF_SERVICE{
		Length:              SECURITY_QOS_SIZE,
		ContextTrackingMode: SECURITY_DYNAMIC_TRACKING,
		EffectiveOnly:       1,
		ImpersonationLevel:  SecurityAnonymous,
	},
	Flags:          ALPC_PORFLG_ALLOW_LPC_REQUESTS,
	DupObjectTypes: ALPC_SYNC_OBJECT_TYPE,
}

func BasicAlpcSend(
	hPort HANDLE,
	msg *AlpcShortMessage,
	flags uint32,
	pMsgAttrs *ALPC_MESSAGE_ATTRIBUTES,
	timeout *int64,
) (e error) {

	e = NtAlpcSendWaitReceivePort(hPort, flags, msg, pMsgAttrs, nil, nil, nil, timeout)
	return

}

func BasicAlpcRecv(
	hPort HANDLE,
	pMsg *AlpcShortMessage,
	pMsgAttrs *ALPC_MESSAGE_ATTRIBUTES,
	timeout *int64,
) (bufLen uint32, e error) {

	bufLen = uint32(pMsg.TotalLength)
	e = NtAlpcSendWaitReceivePort(hPort, 0, nil, nil, pMsg, &bufLen, pMsgAttrs, timeout)
	return

}

// Convenience method to create an ALPC port with a NULL DACL. Requires an
// absolute port name ( where / is the root of the kernel object directory )
func BasicAlpcCreatePort(name string) (hPort HANDLE, e error) {

	objAttr, e := basicObjectAttributes(name)
	if e != nil {
		return
	}

	hPort, e = NtAlpcCreatePort(&objAttr, &basicPortAttr)

	return
}

func BasicAlpcConnectPort(clientName, serverName string, pConnMsg *AlpcShortMessage) (hPort HANDLE, e error) {

	objAttr, e := basicObjectAttributes(clientName)
	if e != nil {
		return
	}

	hPort, e = NtAlpcConnectPort(
		serverName,
		&objAttr,
		&basicPortAttr,
		0,
		nil,
		pConnMsg,
		nil,
		nil,
		nil,
		nil,
	)

	return
}

func BasicAlpcAccept(
	hSrv HANDLE,
	context *AlpcPortContext,
	pConnReq *AlpcShortMessage,
	portName string,
	accept bool,
) (hPort HANDLE, e error) {

	objAttr, e := basicObjectAttributes(portName)
	if e != nil {
		return
	}

	var accepted uintptr
	if accept {
		accepted++
	}

	hPort, e = NtAlpcAcceptConnectPort(
		hSrv,
		0,
		&objAttr,
		&basicPortAttr,
		context,
		pConnReq,
		nil,
		accepted,
	)

	return
}
