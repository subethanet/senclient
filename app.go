package senclient

import (
	"senclient/tcpserver"
	"senclient/con"
)

/*
	Go library for the subethanet protocol.
*/


type Client struct {
	server tcpserver.Server
}


/*
	Create and start the senclient app.
*/
func Create(port int) Client {
	server := tcpserver.Create(port)
	app := Client{
		server: server,
	}
	go server.Start()
	return app
}


func (*Client) ConnectAndListen(ip string, port int) {
	conn := con.Connect(ip, port)
	conn.Listen()
}