package senclient

import (
	"crypto/tls"
	"senclient/tcpserver"
)

type client struct {
	listeningPort int
	server        tcpserver.Server
	cert          tls.Certificate
}
