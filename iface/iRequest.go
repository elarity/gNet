package iface

type IRequest interface {
	GetConn() ItcpConn
	GetClientData() []byte
}