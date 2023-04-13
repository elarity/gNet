package core

import (
	"github.com/elarity/gNet/base"
	"fmt"
	"time"
	"net"
	"github.com/elarity/gNet/iface"
)

type Server struct {
	IpAddress  string
	Port 	   int
	ServerName string
	NetFamily  string
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
			for {
				buffer := make([]byte, 500)
				readContentLength, err := tcpConnection.Read(buffer)
				if err != nil {
					fmt.Printf("tcpConnection.Read err=%+v", err)
					continue
				}
				_, err = tcpConnection.Write(buffer[:readContentLength])
				if err != nil {
					fmt.Printf("tcpConnection.Write err=%+v", err)
					continue
				}
			}
		}
	}()
}

func (svr *Server) Stop() {
	fmt.Println("svr *Server.Stop")
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
func InitServer() iface.IServ {
	svr := Server{
		IpAddress: "0.0.0.0",
		Port: 9191,
		NetFamily: "tcp4",
		ServerName: "gNet-core",
	}
	return &svr
}