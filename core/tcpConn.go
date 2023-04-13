package core

import (
	"net"
	"github.com/elarity/gNet/iface"
)

type TcpConn struct {
	UniqueID uint64
	RawTcpConnFd net.TCPConn
	Status int
}

func (conn *TcpConn) GetUniqueID() uint64 {
	return conn.UniqueID
}

func (conn *TcpConn) Extinguish() {

}

func (conn *TcpConn) Fire() {

}

func InitTcpConn() *iface.ItcpConn {
	tcpConn := &TcpConn{}
	return tcpConn
}