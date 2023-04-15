package iface

import (
	"net"
)

type TcpConn interface {
	Fire()
	Extinguish()
	GetRouter() Router
}

type TcpConnHandler func(tcpRawConn *net.TCPConn, clientData []byte) error
