package iface

type Router interface {
	HandlerBefore(request Request)
	HandlerFire(request Request)
	HandlerAfter(request Request)
}
