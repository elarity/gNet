package request

import (
	"github.com/elarity/gNet/iface"
)

type Tcp struct {
	// 这里不用指针的原因是为了，避免误操作修改内容...
	// 直接复制一份出来即可
	conn iface.TcpConn
	data []byte
}

func (t *Tcp) GetConn() iface.TcpConn {
	return t.conn
}

func (t *Tcp) GetClientData() []byte {
	return t.data
}

var _ iface.Request = (*Tcp)(nil)
