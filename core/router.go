package core

import (
	"fmt"
	"github.com/elarity/gNet/iface"
)

type MainRouter struct{}

/*
HandlerBefore()
HandlerFire()
HandlerAfter()
*/

func (r *MainRouter) HandlerBefore(req iface.Request) {
	fmt.Println("okokkkkk")
}

func (r *MainRouter) HandlerFire(req iface.Request) {

}

func (r *MainRouter) HandlerAfter(req iface.Request) {

}

var _ iface.Router = (*MainRouter)(nil)
