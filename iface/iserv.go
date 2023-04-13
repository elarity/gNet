package iface

/*
 * @desc : 定义服务器核心interface～
 */
type IServ interface {
	Start()
	Stop()
	Serv()
}