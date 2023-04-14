package main

import (
	"github.com/elarity/gNet/core"
)

func main() {

	svr := core.InitServer()
	svr.Serv()

}