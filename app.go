package senclient

import (
	"senclient/tcpserver"
	"senclient/con"
	"crypto/tls"
	sencrypto "senclient/con/crypt"
)

/*
	Go library for the subethanet protocol.
*/


type Client struct {
	listeningPort int
	server tcpserver.Server
	cert tls.Certificate
}


/*
	Create the senclient app.
*/
func Create(port int, cert tls.Certificate) Client {
	server := tcpserver.Create(port, cert, con.FromIncomingCon)
	app := Client{
		server: server,
	}
	return app
}


/*
	Load the client's certificate and key.
*/
func (app *Client) LoadCertAndKey(certPath string, keyPath string) {
	app.cert = sencrypto.LoadCert(certPath, keyPath)
}


/*
	Run the senclient app (once everything has been set up).
*/
func (app *Client) Start() {
	app.server.Start()
}


/*
	Test function.
*/
func (app *Client) ConnectAndListen(ip string, port int) {
	conn := con.Connect(ip, port)
	conn.Listen()
}