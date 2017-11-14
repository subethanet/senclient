package con

import (
	"net"
	"fmt"
)

/*
	Functions for creating Connection objects (which can vary by context).
*/


func FromIncomingCon(tcpCon net.Conn) Connection {
	fmt.Println("Created incomming connection")
	conn := Connection{
		tcpCon:   tcpCon,
	}
	conn.Send("This is a temporary test auto-response.")
	return Connection{}
}


func FromOutgoingCon(tcpCon net.Conn) Connection  {
	return Connection{
		tcpCon:   tcpCon,
	}
}
