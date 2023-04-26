package iface

type Request interface {
	SetConn(TcpConn)
	GetConn() TcpConn
	SetMessage(Message)
	SetClientData([]byte)
	GetClientData() []byte
	GetMessageId() uint64
}
