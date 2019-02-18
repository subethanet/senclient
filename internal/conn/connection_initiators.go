package conn

import (
	"fmt"
	"net"
)

/*
	Functions for creating TcpConnection objects (which can vary by context).
*/

func FromIncomingCon(tcpCon net.Conn) {
	fmt.Println("Created incomming connection")
	conn := TcpConnection{
		tcpCon: tcpCon,
	}
	conn.send("This is a temporary test auto-response.")
}

func FromOutgoingCon(tcpCon net.Conn) TcpConnection {
	return TcpConnection{
		tcpCon: tcpCon,
	}
}
