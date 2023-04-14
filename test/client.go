package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("tcp-client net.Dial")
	clientTcpConn, err := net.Dial("tcp", "127.0.0.1:9191")
	if err != nil {
		fmt.Printf("net.Dial err=%+v", err)
		return
	}
	writeContentLength, err := clientTcpConn.Write([]byte("hello-world-hello-world"))
	if err != nil {
		fmt.Printf("clientTcpConn.Write err=%+v", err)
		return
	}
	buffer := make([]byte, 512)
	fmt.Println("tcp-client net.Dial Read")
	readContentLenth, err := clientTcpConn.Read(buffer)
	fmt.Println("tcp-client net.Dial Read Over")
	if err != nil {
		fmt.Printf("clientTcpConn.Read error=%+v", err)
		return
	}
	fmt.Printf("Message From Svr:%s, readContentLenth=%d, writeContentLength=%d\n", buffer, readContentLenth, writeContentLength)

}