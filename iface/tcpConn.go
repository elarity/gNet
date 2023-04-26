package iface

import (
	"net"
)

type TcpConn interface {
	Fire()
	Extinguish()
	GetRouter() Router
	GetRawTcpConnFd() *net.TCPConn
	// 给客户端回写信息...
	Write(uint64, []byte) error
}

//type TcpConnHandler func(tcpRawConn *net.TCPConn, clientData []byte) error
