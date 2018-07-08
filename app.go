package senclient

import (
	"crypto/tls"
)

type Client struct {
	listeningPort int
	server        Server
	cert          tls.Certificate
}

// Create the senclient app.
func CreateApp(port int) Client {
	app := Client{
		listeningPort: port,
	}
	return app
}

// Load the Client's certificate and key.
func (app *Client) LoadCertAndKey(certPath string, keyPath string) {
	app.cert = LoadCert(certPath, keyPath)
}

// Run the senclient app (once everything has been set up).
func (app *Client) Start() {
	app.server = Create(app.listeningPort, app.cert, FromIncomingCon)
	app.server.Start()
}

func (app *Client) Stop() {
	app.server.Stop()
}
