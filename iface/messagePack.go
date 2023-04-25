package iface

type MessagePack interface {
	GetMessagePackHeaderLength() uint64
	Pack(Message) ([]byte, error)
	UnPack([]byte) (Message, error)
}
