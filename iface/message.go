package iface

type Message interface {
	GetMid() uint64
	SetMid(uint64)
	GetMessageData() []byte
	SetMessageData([]byte)
	GetMessageLength() uint32
	SetMessageLength(uint32)
}
