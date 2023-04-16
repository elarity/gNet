package iface

type MessagePack interface {
	GetMessagePackLength() uint64
	Pack(Message) ([]byte, error)
	UnPack([]byte) (Message, error)
}
