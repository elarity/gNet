package core

/*
	IpAddress:  "0.0.0.0",
	Port:       9191,
	NetFamily:  "tcp4",
	ServerName: "gNet-core",
*/

var ConfInstance *Conf

type Conf struct {
	IpAddress  string
	Port       int
	NetFamily  string
	ServerName string
}

func InitConf() *Conf {
	ConfInstance = &Conf{
		IpAddress:  "0.0.0.0",
		Port:       9191,
		NetFamily:  "tcp4",
		ServerName: "gNet-core",
	}
	return ConfInstance
}
