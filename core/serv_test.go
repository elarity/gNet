package core

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestSvr(t *testing.T) {
	svr := InitServer()

	go func() {
		// 保证svr一定要在client运行前ready
		time.Sleep(3 * time.Second)
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
		buffer := make([]byte, 4)
		readContentLenth, err := clientTcpConn.Read(buffer)
		if err != nil {
			fmt.Printf("clientTcpConn.Read error=%+v", err)
			return
		}
		fmt.Printf("Message From Svr:%s, readContentLenth=%d, writeContentLength=%d\n", buffer, readContentLenth, writeContentLength)
	}()

	svr.Serv()
}
