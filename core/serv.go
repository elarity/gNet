package core

import (
	"fmt"
	"github.com/elarity/gNet/iface"
	"io"
	"net"
	"time"
)

type Server struct {
	IpAddress  string
	Port       int
	ServerName string
	NetFamily  string
	Router     iface.Router
}

func tcpConnHandler(tcpRawConn *net.TCPConn, clientData []byte) error {
	fmt.Println("tcp connection callback, content is=", string(clientData), ", length=", len(clientData))
	return nil
}

/*
Start()
Stop()
Serv()
*/
func (svr *Server) Start() {
	fmt.Println("svr *Server.Start")
	// 单独开一个新的goroutine去处理...
	go func() {
		/*
			golang标准库下创建tcp socket的流程也是固定的...
		*/
		tcpAddr, err := net.ResolveTCPAddr(svr.NetFamily, fmt.Sprintf("%s:%d", svr.IpAddress, svr.Port))
		if err != nil {
			panic(any(err))
		}
		// 本质上
		listenerSocket, err := net.ListenTCP(svr.NetFamily, tcpAddr)
		if err != nil {
			panic(any(err))
		}

		for {
			tcpConnection, err := listenerSocket.AcceptTCP()
			// 这里遇到错误就不要panic了，打印一下错误就赶紧continue
			if err != nil {
				fmt.Printf("Accept.Tcp err=%+v", err)
				continue
			}

			// accept后，使用goroutine去做这些事，考虑下不用goroutine会发生什么事
			go func(tcpConnection *net.TCPConn) {
				/*
				   一个pack的组成是下面这样的：
				   | 4字节，放整个pack的字节长度 | 8字节，放msgid | 这部分是客户端请求参数 |
				*/
				// 服务器这里要拆包，所以创建一个messagePack后，准备开始拆包
				messagePack := InitMessagePack()
				for {
					// 读取header的bytes二进制流
					// header是有 length + msgid组成
					headDataBytes := make([]byte, messagePack.GetMessagePackHeaderLength())
					// actuallyReadLength
					_, err := io.ReadFull(tcpConnection, headDataBytes)
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
						bodyMessage := make([]byte, headerMessage.GetMessageLength())
						_, err := io.ReadFull(tcpConnection, bodyMessage)
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
				}
			}(tcpConnection)

			//time.Sleep(1000 * time.Second)
			/*
				for {
					buffer := make([]byte, 8)
					readContentLength, err := tcpConnection.Read(buffer)
					if err != nil {
						fmt.Printf("tcpConnection.Read err=%+v", err)
						continue
					}
					fmt.Println("Message from client:", string(buffer), " length:", readContentLength)
					_, err = tcpConnection.Write(buffer[:readContentLength])
					if err != nil {
						fmt.Printf("tcpConnection.Write err=%+v", err)
						continue
					}
				}
			*/
		}
	}()
}

func (svr *Server) Stop() {
	fmt.Println("svr *Server.Stop")
}

func (svr *Server) AddRouter(router iface.Router) {
	svr.Router = router
}

func (svr *Server) Serv() {
	svr.Start()

	// 阻止退出...
	for {
		time.Sleep(1 * time.Second)
	}
}

/*
 * @desc : 初始化一个服务器实例
 */
func InitServer() iface.Serv {
	svr := Server{
		IpAddress:  ConfInstance.IpAddress,
		Port:       ConfInstance.Port,
		NetFamily:  ConfInstance.NetFamily,
		ServerName: ConfInstance.ServerName,
		Router:     nil,
	}
	return &svr
}

var _ iface.Serv = (*Server)(nil)
