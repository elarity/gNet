package iface

import (
	"net"
)

type ItcpConn interface {
	Fire()
	Extinguish()
	GetUniqueID() uint64
}

type TcpConnHandler func(tcpRawConn *net.TCPConn, clientData []byte) error