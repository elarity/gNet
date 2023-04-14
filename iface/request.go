package iface

type Request interface {
	GetConn() TcpConn
	GetClientData() []byte
}
