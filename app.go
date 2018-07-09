package senclient

import (
	"crypto/tls"
)

type Client struct {
	listeningPort int
	server        server
	cert          tls.Certificate
}

func CreateApp(port int) Client {
	app := Client{
		listeningPort: port,
	}
	return app
}

// Load the Client's certificate and key.
func (app *Client) LoadCertAndKey(certPath string, keyPath string) {
	app.cert = loadCert(certPath, keyPath)
}

// Run the senclient app (once everything has been set up).
func (app *Client) Start() {
	app.server = newServer(app.listeningPort, app.cert, FromIncomingCon)
	app.server.start()
}

func (app *Client) Stop() {
	app.server.stop()
}
