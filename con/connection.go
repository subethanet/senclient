package con

import "net"


type Connection struct {
	tcpCon net.Conn
}