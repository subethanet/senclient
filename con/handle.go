package con

import (
	"net"
)


func HandleNewTcpConnection(conn net.Conn) Connection  {
	return Connection{
		tcpCon: conn,
	}
}