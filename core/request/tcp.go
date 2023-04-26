package request

import (
	"github.com/elarity/gNet/iface"
)

type Tcp struct {
	// 这里不用指针的原因是为了，避免误操作修改内容...
	// 直接复制一份出来即可
	conn    iface.TcpConn
	data    []byte
	message iface.Message
}

func (t *Tcp) SetConn(conn iface.TcpConn) {
	t.conn = conn
}

func (t *Tcp) GetConn() iface.TcpConn {
	return t.conn
}

func (t *Tcp) SetClientData(data []byte) {
	t.data = data
}

func (t *Tcp) GetClientData() []byte {
	return t.data
}

func (t *Tcp) SetMessage(message iface.Message) {
	t.message = message
}

func (t *Tcp) GetMessageId() uint64 {
	return t.message.GetMid()
}

var _ iface.Request = (*Tcp)(nil)
