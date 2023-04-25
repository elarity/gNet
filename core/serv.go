package core

import (
	"fmt"
	"github.com/elarity/gNet/iface"
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
				tcpConn := InitTcpConn(tcpConnection, svr.Router)
				tcpConn.Fire()
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
