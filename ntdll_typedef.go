package w32

import (
	"errors"
)

// nt!_ALPC_MESSAGE_ATTRIBUTES
//  +0x000 AllocatedAttributes : Uint4B
//  +0x004 ValidAttributes  : Uint4B
type ALPC_MESSAGE_ATTRIBUTES struct {
	AllocatedAttributes uint32
	ValidAttributes     uint32
}

type ALPC_CONTEXT_ATTR struct {
	PortContext    *AlpcPortContext
	MessageContext uintptr
	Sequence       uint32
	MessageId      uint32
	CallbackId     uint32
}

type ALPC_HANDLE_ATTR struct {
	Flags         uint32
	Handle        HANDLE
	ObjectType    uint32
	DesiredAccess uint32
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
	_             [4]byte // align to 0x08
	Buffer        *uint16
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
	_                        [4]byte // align to 0x08
	RootDirectory            HANDLE
	ObjectName               *UNICODE_STRING
	Attributes               uint32
	_                        [4]byte // align to 0x20
	SecurityDescriptor       *SECURITY_DESCRIPTOR
	SecurityQualityOfService *SECURITY_QUALITY_OF_SERVICE
}

// cf: http://j00ru.vexillium.org/?p=502 for legacy RPC
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
	_              [4]byte // align up to 0x20
	ClientViewSize uint64
}

func (pm PORT_MESSAGE) CallbackId() uint32 {
	return uint32(pm.ClientViewSize >> 32)
}

func (pm PORT_MESSAGE) DoNotUseThisField() float64 {
	panic("WE TOLD YOU NOT TO USE THIS FIELD")
}

const PORT_MESSAGE_SIZE = 0x28

// http://www.nirsoft.net/kernel_struct/vista/SECURITY_QUALITY_OF_SERVICE.html
type SECURITY_QUALITY_OF_SERVICE struct {
	Length              uint32
	ImpersonationLevel  uint32
	ContextTrackingMode byte
	EffectiveOnly       byte
	_                   [2]byte // align to 12 bytes
}

const SECURITY_QOS_SIZE = 12

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
	MaxMessageLength    uint64 // must be filled out
	MemoryBandwidth     uint64
	MaxPoolUsage        uint64
	MaxSectionSize      uint64
	MaxViewSize         uint64
	MaxTotalSectionSize uint64
	DupObjectTypes      uint32
	Reserved            uint32
}

const SHORT_MESSAGE_MAX_SIZE = 65535 // MAX_USHORT

type AlpcShortMessage struct {
	PORT_MESSAGE
	Command uint32
	Data    [SHORT_MESSAGE_MAX_SIZE - PORT_MESSAGE_SIZE - 4]byte
}

func NewAlpcShortMessage() AlpcShortMessage {
	sm := AlpcShortMessage{}
	sm.TotalLength = uint16(SHORT_MESSAGE_MAX_SIZE)
	return sm
}

func (sm *AlpcShortMessage) SetData(d []byte) (e error) {

	copy(sm.Data[:], d)
	if len(d) > len(sm.Data) {
		e = errors.New("data too big - truncated")
		// the 4 everywhere is the Command struct member
		sm.DataLength = uint16(4 + len(sm.Data))
		sm.TotalLength = uint16(PORT_MESSAGE_SIZE + 4 + len(sm.Data))
		return
	}
	sm.TotalLength = uint16(PORT_MESSAGE_SIZE + 4 + len(d))
	sm.DataLength = uint16(4 + len(d))
	return

}

func (sm *AlpcShortMessage) GetData() []byte {
	if len(sm.Data) < int(sm.DataLength-4) {
		return []byte{}
	}
	return sm.Data[:sm.DataLength-4]
}

func (sm *AlpcShortMessage) Reset() {
	// zero the PORT_MESSAGE header
	sm.PORT_MESSAGE = PORT_MESSAGE{}
	sm.TotalLength = uint16(SHORT_MESSAGE_MAX_SIZE)
	sm.DataLength = 0
}

type AlpcPortContext struct {
	Handle HANDLE
}
