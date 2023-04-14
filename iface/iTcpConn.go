package iface

import (
	"net"
)

type ItcpConn interface {
	Fire()
	Extinguish()
}

type TcpConnHandler func(tcpRawConn *net.TCPConn, clientData []byte) error
