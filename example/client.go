package main

import (
	"fmt"
	"github.com/elarity/gNet/core"
	"net"
	"time"
)

func main() {

	for {

		fmt.Println("tcp-client net.Dial")
		clientTcpConn, err := net.Dial("tcp", "127.0.0.1:9191")
		if err != nil {
			fmt.Printf("net.Dial err=%+v", err)
			return
		}

		msgPack := core.InitMessagePack()

		msg1 := core.InitMessage(1, []byte{'h', 'e', 'l', 'l', '9'})
		msg1Byte, _ := msgPack.Pack(msg1)
		//writeContentLength, err := clientTcpConn.Write(msg1Byte)
		//fmt.Println("send to svr length=", writeContentLength)

		msg2 := core.InitMessage(2, []byte{'w', 'o', 'r', 'l', '0'})
		msg2Byte, _ := msgPack.Pack(msg2)

		msg := append(msg1Byte, msg2Byte...)

		writeContentLength, err := clientTcpConn.Write(msg)
		fmt.Println("send to svr length=", writeContentLength)
		/*
			writeContentLength, err := clientTcpConn.Write([]byte("hello-world-hello-world"))
			if err != nil {
				fmt.Printf("clientTcpConn.Write err=%+v", err)
				return
			}
			buffer := make([]byte, 5120)
			fmt.Println("tcp-client net.Dial Read")
			readContentLenth, err := clientTcpConn.Read(buffer)
			fmt.Println("tcp-client net.Dial Read Over")
			if err != nil {
				fmt.Printf("clientTcpConn.Read error=%+v", err)
				return
			}
			fmt.Printf("Message From Svr:%s, readContentLenth=%d, writeContentLength=%d\n", buffer, readContentLenth, writeContentLength)
		*/

		time.Sleep(1 * time.Second)
	}
}
