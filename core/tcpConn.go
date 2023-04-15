package core

import (
	"fmt"
	"github.com/elarity/gNet/core/request"
	"github.com/elarity/gNet/iface"
	"net"
)

// 打开和关闭状态...
const TcpConnOpenStatus = 1
const TcpConnCloseStatus = 0

type TcpConn struct {
	UniqueID     uint64
	RawTcpConnFd *net.TCPConn
	Status       int
	Router       iface.Router
}

func (conn *TcpConn) GetRouter() iface.Router {
	return conn.Router
}

func (conn *TcpConn) Extinguish() {
	fmt.Println("tcp connction Extinguish()")
	if TcpConnCloseStatus == conn.Status {
		return
	}
	// 设置status状态
	conn.Status = TcpConnCloseStatus

	// 关闭fd
	err := conn.RawTcpConnFd.Close()
	if err != nil {

	}
}

func (conn *TcpConn) Fire() {
	// process bussiness
	go func() {
		defer fmt.Println("tcp connection fire end")
		defer conn.Extinguish()
		for {
			fmt.Println("Fire for begin...")
			clientDataBuffer := make([]byte, 512)
			clientDataLength, err := conn.RawTcpConnFd.Read(clientDataBuffer)
			if err != nil {

			}
			fmt.Println("client data length=", clientDataLength)

			// 构造request对象
			tcpRequest := request.Tcp{}
			tcpRequest.SetConn(conn)
			tcpRequest.SetClientData(clientDataBuffer)
			go func(req iface.Request) {
				//conn := req.GetConn()
				//router := conn.GetRouter()
				conn.Router.HandlerBefore(req)
				conn.Router.HandlerFire(req)
				conn.Router.HandlerAfter(req)
			}(&tcpRequest)
		}
	}()

	fmt.Println("over over...")
}

func InitTcpConn(rawTcpConnFd *net.TCPConn, router iface.Router) iface.TcpConn {
	tcpConn := &TcpConn{
		RawTcpConnFd: rawTcpConnFd,
		Status:       TcpConnOpenStatus,
		Router:       router,
	}
	return tcpConn
}

var _ iface.TcpConn = (*TcpConn)(nil)
