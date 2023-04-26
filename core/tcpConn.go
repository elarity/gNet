package core

import (
	"fmt"
	"github.com/elarity/gNet/core/request"
	"github.com/elarity/gNet/iface"
	"io"
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

func (conn *TcpConn) GetRawTcpConnFd() *net.TCPConn {
	return conn.RawTcpConnFd
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
			var bodyMessage []byte
			// 使用pack与unpack...
			/*
			   一个pack的组成是下面这样的：
			   | 4字节，放整个pack的字节长度 | 8字节，放msgid | 这部分是客户端请求参数 |
			*/
			rawTcpConnection := conn.GetRawTcpConnFd()
			// 服务器这里要拆包，所以创建一个messagePack后，准备开始拆包
			messagePack := InitMessagePack()
			// 读取header的bytes二进制流
			// header是有 length + msgid组成
			headDataBytes := make([]byte, messagePack.GetMessagePackHeaderLength())
			// actuallyReadLength
			_, err := io.ReadFull(rawTcpConnection, headDataBytes)
			if err != nil {
				break
			}
			/*
				headerMessage struct
				type Message struct {
					mid    uint64
					data   []byte
					length uint32 // data的length
				}
			*/
			headerMessage, err := messagePack.UnPack(headDataBytes)
			if err != nil {
				break
			}
			if headerMessage.GetMessageLength() > 0 {
				bodyMessage = make([]byte, headerMessage.GetMessageLength())
				_, err := io.ReadFull(rawTcpConnection, bodyMessage)
				if err != nil {
					fmt.Println("server unpack data err:", err)
					return
				}
				headerMessage.SetMessageData(bodyMessage)
				/*
					tcpConn := InitTcpConn(tcpConnection, svr.Router)
					tcpConn.Fire()
				*/
				fmt.Println("message id=", headerMessage.GetMid(), " | message length=", headerMessage.GetMessageLength(), " | message body=", string(headerMessage.GetMessageData()))
			}
			/*
				fmt.Println("Fire for begin...")
				clientDataBuffer := make([]byte, 512)
				clientDataLength, err := conn.RawTcpConnFd.Read(clientDataBuffer)
				if err != nil {

				}
				fmt.Println("client data length=", clientDataLength)
			*/
			// 构造request对象
			tcpRequest := request.Tcp{}
			tcpRequest.SetConn(conn)
			//tcpRequest.SetClientData(bodyMessage)
			tcpRequest.SetMessage(headerMessage)
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
