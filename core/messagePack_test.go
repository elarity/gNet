package core

import (
	"fmt"
	"testing"
)

func TestPack(t *testing.T) {

	msgPackFactory := InitMessagePack()

	msgOne := InitMessage(1, []byte{'p', 'i', 'n', 'g'})
	msgOneByte, _ := msgPackFactory.Pack(msgOne)
	msgTwo := InitMessage(2, []byte{'p', 'o', 'n', 'g'})
	msgTwoByte, _ := msgPackFactory.Pack(msgTwo)
	totalMsgByte := append(msgOneByte, msgTwoByte...)

	rawMsg, _ := msgPackFactory.UnPack(totalMsgByte)
	fmt.Println("mid = ", rawMsg.GetMid())
	fmt.Println("data = ", rawMsg.GetMessageData())

	rawMsg, _ = msgPackFactory.UnPack(totalMsgByte)
	fmt.Println("mid = ", rawMsg.GetMid())
	fmt.Println("data = ", rawMsg.GetMessageData())
}
