package senclient

import (
	"senclient/tcpserver"
	"senclient/con"
	"senclient/crypt"
)

/*
	Go library for the subethanet protocol.
*/


/*
	Create the senclient app.
*/
func Create(port int) client {
	app := client{
		listeningPort: port,
	}
	return app
}


/*
	Load the client's certificate and key.
*/
func (app *client) LoadCertAndKey(certPath string, keyPath string) {
	app.cert = crypt.LoadCert(certPath, keyPath)
}


/*
	Run the senclient app (once everything has been set up).
*/
func (app *client) Start() {
	app.server = tcpserver.Create(app.listeningPort, app.cert, con.FromIncomingCon)
	app.server.Start()
}


func (app *client) Stop() {
	app.server.Stop()
}