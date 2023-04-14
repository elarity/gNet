package core

import (
	"fmt"
	"github.com/elarity/gNet/iface"
	"net"
)

// 打开和关闭状态...
const TcpConnOpenStatus = 1
const TcpConnCloseStatus = 0

type TcpConn struct {
	UniqueID     uint64
	RawTcpConnFd *net.TCPConn
	Status       int
	handler      iface.TcpConnHandler
}

func (conn *TcpConn) Extinguish() {
	fmt.Println("tcp connction Extinguish()")
	if TcpConnCloseStatus == conn.Status {
		return
	}
	// 设置status状态
	conn.Status = TcpConnCloseStatus

	// 关闭fd
	err := conn.RawTcpConnFd.Close()
	if err != nil {

	}
}

func (conn *TcpConn) Fire() {
	// process bussiness
	go func() {
		defer fmt.Println("tcp connection fire end")
		defer conn.Extinguish()
		for {
			fmt.Println("Fire for begin...")
			clientDataBuffer := make([]byte, 512)
			clientDataLength, err := conn.RawTcpConnFd.Read(clientDataBuffer)
			if err != nil {

			}
			fmt.Println("client data length=", clientDataLength)
			conn.handler(conn.RawTcpConnFd, clientDataBuffer)
		}
	}()

	fmt.Println("over over...")
}

func InitTcpConn(rawTcpConnFd *net.TCPConn, handler iface.TcpConnHandler) iface.ItcpConn {
	tcpConn := &TcpConn{
		RawTcpConnFd: rawTcpConnFd,
		handler: handler,
	}
	return tcpConn
}

var _ iface.ItcpConn = (*TcpConn)(nil)
