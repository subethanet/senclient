package senclient

import (
	"senclient/tcpserver"
	"crypto/tls"
)


type client struct {
	listeningPort int
	server tcpserver.Server
	cert tls.Certificate
}