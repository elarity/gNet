package core

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

func TestPack(t *testing.T) {

	msgPackFactory := InitMessagePack()

	msgOne := InitMessage(1, []byte{'p', 'i', 'n', 'g', 'o'})
	msgOneByte, _ := msgPackFactory.Pack(msgOne)
	msgTwo := InitMessage(2, []byte{'p', 'o', 'n', 'g'})
	msgTwoByte, _ := msgPackFactory.Pack(msgTwo)
	totalMsgByte := append(msgOneByte, msgTwoByte...)

	rawMsg, _ := msgPackFactory.UnPack(totalMsgByte)
	fmt.Println("mid = ", rawMsg.GetMid())
	fmt.Println("data = ", string(rawMsg.GetMessageData()))
	fmt.Println("length=", rawMsg.GetMessageLength())

	rawMsg2, _ := msgPackFactory.UnPack(totalMsgByte)
	fmt.Println("mid = ", rawMsg2.GetMid())
	fmt.Println("data = ", string(rawMsg2.GetMessageData()))
	fmt.Println("length=", rawMsg.GetMessageLength())
}

func TestB(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.BigEndian, 1.004); err != nil {
		panic(err)
	}
	if err := binary.Write(buf, binary.BigEndian,
		[]byte("Hello")); err != nil {
		panic(err)
	}

	var num float64
	if err := binary.Read(buf, binary.BigEndian, &num); err != nil {
		panic(err)
	}
	fmt.Printf("float64: %.3f\n", num)

	if err := binary.Read(buf, binary.BigEndian, &num); err != nil {
		panic(err)
	}
	fmt.Printf("float64: %.3f\n", num)

	greeting := make([]byte, 5)
	if err := binary.Read(buf, binary.BigEndian, &greeting); err != nil {
		panic(err)
	}
	fmt.Printf("string: %s\n", string(greeting))
}
