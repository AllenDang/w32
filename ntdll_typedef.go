package w32

import (
	"unsafe"
)

// nt!_ALPC_MESSAGE_ATTRIBUTES
//  +0x000 AllocatedAttributes : Uint4B
//  +0x004 ValidAttributes  : Uint4B
type ALPC_MESSAGE_ATTRIBUTES struct {
	AllocatedAttributes uint32
	ValidAttributes     uint32
}

// nt!_CLIENT_ID
//  +0x000 UniqueProcess    : Ptr64 Void
//  +0x008 UniqueThread     : Ptr64 Void
type CLIENT_ID struct {
	UniqueProcess uintptr
	UniqueThread  uintptr
}

// nt!_UNICODE_STRING
//  +0x000 Length           : Uint2B
//  +0x002 MaximumLength    : Uint2B
//  +0x008 Buffer           : Ptr64 Uint2B
type UNICODE_STRING struct {
	Length        uint16
	MaximumLength uint16
	Buffer        unsafe.Pointer
}

// nt!_OBJECT_ATTRIBUTES
//  +0x000 Length           : Uint4B
//  +0x008 RootDirectory    : Ptr64 Void
//  +0x010 ObjectName       : Ptr64 _UNICODE_STRING
//  +0x018 Attributes       : Uint4B
//  +0x020 SecurityDescriptor : Ptr64 Void
//  +0x028 SecurityQualityOfService : Ptr64 Void
type OBJECT_ATTRIBUTES struct {
	Length                   uint32
	padding1                 [4]byte
	RootDirectory            HANDLE
	ObjectName               *UNICODE_STRING
	Attributes               uint32
	padding2                 [4]byte
	SecurityDescriptor       *SECURITY_DESCRIPTOR
	SecurityQualityOfService *SECURITY_QUALITY_OF_SERVICE
}

// cf: http://j00ru.vexillium.org/?p=502
// nt!_PORT_MESSAGE
//    +0x000 u1               : <unnamed-tag>
//    +0x004 u2               : <unnamed-tag>
//    +0x008 ClientId         : _CLIENT_ID
//    +0x008 DoNotUseThisField : Float
//    +0x018 MessageId        : Uint4B
//    +0x020 ClientViewSize   : Uint8B
//    +0x020 CallbackId       : Uint4B
type PORT_MESSAGE struct {
	DataLength     uint16 // These are the two unnamed unions
	TotalLength    uint16 // without Length and ZeroInit
	Type           uint16
	DataInfoOffset uint16
	ClientId       CLIENT_ID
	MessageId      uint32
	padding        [4]byte
	ClientViewSize uint64
}

func (pm PORT_MESSAGE) CallbackId() uint32 {
	return uint32(pm.ClientViewSize >> 32)
}

func (pm PORT_MESSAGE) DoNotUseThisField() float {
	panic("WE TOLD YOU NOT TO USE THIS FIELD")
}

type SECURITY_IMPERSONATION_LEVEL int

const (
	SecurityAnonymous SECURITY_IMPERSONATION_LEVEL = iota
	SecurityIdentification
	SecurityImpersonation
	SecurityDelegation
)

// http://www.nirsoft.net/kernel_struct/vista/SECURITY_QUALITY_OF_SERVICE.html
// Added internal padding to make it 0xC bytes, as per the dt output below
type SECURITY_QUALITY_OF_SERVICE struct {
	Length              uint32
	ImpersonationLevel  SECURITY_IMPERSONATION_LEVEL
	ContextTrackingMode byte
	EffectiveOnly       byte
	padding             [2]byte
}

// nt!_ALPC_PORT_ATTRIBUTES
//  +0x000 Flags            : Uint4B
//  +0x004 SecurityQos      : _SECURITY_QUALITY_OF_SERVICE
//  +0x010 MaxMessageLength : Uint8B
//  +0x018 MemoryBandwidth  : Uint8B
//  +0x020 MaxPoolUsage     : Uint8B
//  +0x028 MaxSectionSize   : Uint8B
//  +0x030 MaxViewSize      : Uint8B
//  +0x038 MaxTotalSectionSize : Uint8B
//  +0x040 DupObjectTypes   : Uint4B
//  +0x044 Reserved         : Uint4B
type ALPC_PORT_ATTRIBUTES struct {
	Flags               uint32
	SecurityQos         SECURITY_QUALITY_OF_SERVICE
	MaxMessageLength    uint64
	MemoryBandwidth     uint64
	MaxPoolUsage        uint64
	MaxSectionSize      uint64
	MaxViewSize         uint64
	MaxTotalSectionSize uint64
	DupObjectTypes      uint32
	Reserved            uint32
}

//   typedef struct _TRANSFERRED_MESSAGE
// {
//     PORT_MESSAGE Header;
//     ULONG   Command;
//     WCHAR   MessageText[48];
// }
type TRANSFERRED_MESSAGE struct {
	Header      PORT_MESSAGE
	Command     uint32
	MessageText [48]uint16
}
