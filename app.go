package senclient

import (
	"crypto/tls"
	"senclient/tcpserver"
)

type Client struct {
	listeningPort int
	server        tcpserver.Server
	cert          tls.Certificate
}
