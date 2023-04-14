package iface

import (
	"net"
)

type TcpConn interface {
	Fire()
	Extinguish()
}

type TcpConnHandler func(tcpRawConn *net.TCPConn, clientData []byte) error
