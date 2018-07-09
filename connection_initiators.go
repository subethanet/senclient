package senclient

import (
	"fmt"
	"net"
)

/*
	Functions for creating tcpConnection objects (which can vary by context).
*/

func FromIncomingCon(tcpCon net.Conn) {
	fmt.Println("Created incomming connection")
	conn := tcpConnection{
		tcpCon: tcpCon,
	}
	conn.send("This is a temporary test auto-response.")
}

func FromOutgoingCon(tcpCon net.Conn) tcpConnection {
	return tcpConnection{
		tcpCon: tcpCon,
	}
}
