package senclient

import (
	"senclient/con"
	"senclient/crypt"
	"senclient/tcpserver"
)

// Create the senclient app.
func Create(port int) Client {
	app := Client{
		listeningPort: port,
	}
	return app
}

// Load the Client's certificate and key.
func (app *Client) LoadCertAndKey(certPath string, keyPath string) {
	app.cert = crypt.LoadCert(certPath, keyPath)
}

// Run the senclient app (once everything has been set up).
func (app *Client) Start() {
	app.server = tcpserver.Create(app.listeningPort, app.cert, con.FromIncomingCon)
	app.server.Start()
}

func (app *Client) Stop() {
	app.server.Stop()
}
