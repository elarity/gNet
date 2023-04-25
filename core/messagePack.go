package core

import (
	"bytes"
	"encoding/binary"
	"github.com/elarity/gNet/iface"
)

/*
一个pack的组成是下面这样的：
| 4字节，放整个pack的字节长度 | 8字节，放msgid | 这部分是客户端请求参数 |
*/
type MessagePack struct{}

func (mp *MessagePack) GetMessagePackHeaderLength() uint64 {
	// 4字节放长度数字，8字节mid（uint64）
	return 12
}

func (mp *MessagePack) Pack(msg iface.Message) ([]byte, error) {
	packData := bytes.NewBuffer([]byte{})
	err := binary.Write(packData, binary.LittleEndian, msg.GetMessageLength())
	if err != nil {
		return nil, err
	}
	err = binary.Write(packData, binary.LittleEndian, msg.GetMid())
	if err != nil {
		return nil, err
	}
	err = binary.Write(packData, binary.LittleEndian, msg.GetMessageData())
	if err != nil {
		return nil, err
	}
	return packData.Bytes(), nil
}

func (mp *MessagePack) UnPack(binaryData []byte) (iface.Message, error) {
	var messageDataLength uint32
	var mid uint64
	packData := bytes.NewReader(binaryData)
	message := &Message{}
	err := binary.Read(packData, binary.LittleEndian, &messageDataLength)
	if err != nil {
		return nil, err
	}
	message.SetMessageLength(messageDataLength)
	err = binary.Read(packData, binary.LittleEndian, &mid)
	if err != nil {
		return nil, err
	}
	message.SetMid(mid)
	return message, nil
}

func InitMessagePack() iface.MessagePack {
	msgPack := &MessagePack{}
	return msgPack
}

var _ iface.MessagePack = (*MessagePack)(nil)
