package main

import (
	"fmt"
	"github.com/elarity/gNet/core"
	"github.com/elarity/gNet/iface"
)

func main() {
	core.InitConf()
	svr := core.InitServer()
	svr.AddRouter(&TestRouter{})
	svr.Serv()

}

type TestRouter struct {
	core.MainRouter
}

func (r *TestRouter) HandlerBefore(request iface.Request) {
	fmt.Println("HandlerBefore")
	_, err := request.GetConn().GetRawTcpConnFd().Write([]byte("HandlerBefore fire\n"))
	if err != nil {
		fmt.Println(err)
	}
}

func (r *TestRouter) HandlerFire(request iface.Request) {
	fmt.Println("HandlerFire")
	_, err := request.GetConn().GetRawTcpConnFd().Write([]byte("HandlerFire fire\n"))
	if err != nil {
		fmt.Println(err)
	}
}

func (r *TestRouter) HandlerAfter(request iface.Request) {
	fmt.Println("HandlerAfter")
	_, err := request.GetConn().GetRawTcpConnFd().Write([]byte("HandlerAfter fire\n"))
	if err != nil {
		fmt.Println(err)
	}
}
