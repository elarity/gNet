package iface

type Request interface {
	SetConn(TcpConn)
	GetConn() TcpConn
	SetClientData([]byte)
	GetClientData() []byte
}
