package con

import (
	"net"
)

/*
	Functions for creating Connection objects (which can vary by context).
*/


func FromIncomingCon(conn net.Conn) Connection  {
	return Connection{
		tcpCon: conn,
		isMyTurn: true,
	}
}


func FromOutgoingCon(conn net.Conn) Connection  {
	return Connection{
		tcpCon: conn,
		isMyTurn: false,
	}
}