package core

import (
	"fmt"
	"github.com/elarity/gNet/iface"
	"net"
)

type TcpConn struct {
	UniqueID     uint64
	RawTcpConnFd *net.TCPConn
	Status       int
	handler      iface.TcpConnHandler
}

func (conn *TcpConn) Extinguish() {

}

func (conn *TcpConn) Fire() {
	// process bussiness
	go func() {
		defer fmt.Println("tcp connection fire end")
		for {
			clientDataBuffer := make([]byte, 512)
			clientDataLength, err := conn.RawTcpConnFd.Read(clientDataBuffer)
			if err != nil {

			}
			fmt.Println("client data length=", clientDataLength)
			conn.handler(conn.RawTcpConnFd, clientDataBuffer)
		}
	}()
}

func InitTcpConn() iface.ItcpConn {
	tcpConn := &TcpConn{}
	return tcpConn
}

var _ iface.ItcpConn = (*TcpConn)(nil)
