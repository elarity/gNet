package iface

import (
	"net"
)

type TcpConn interface {
	Fire()
	Extinguish()
	GetRouter() Router
	GetRawTcpConnFd() *net.TCPConn
}

type TcpConnHandler func(tcpRawConn *net.TCPConn, clientData []byte) error
