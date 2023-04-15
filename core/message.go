package core

import "github.com/elarity/gNet/iface"

type Message struct {
	mid    uint64
	data   []byte
	length uint32
}

func (m *Message) GetMid() uint64 {
	return m.mid
}

func (m *Message) SetMid(mid uint64) {
	m.mid = mid
}

func (m *Message) GetMessageData() []byte {
	return m.data
}

func (m *Message) SetMessageData(data []byte) {
	m.data = data
}

func (m *Message) GetMessageLength() uint32 {
	return m.length
}

func (m *Message) SetMessageLength(length uint32) {
	m.length = length
}

func InitMessage(mid uint64, data []byte) iface.Message {
	message := &Message{
		mid:    mid,
		data:   data,
		length: uint32(len(data)),
	}
	return message
}

var _ iface.Message = (*Message)(nil)
