package senclient

import (
	"fmt"
	"net"
)

/*
	Functions for creating Connection objects (which can vary by context).
*/

func FromIncomingCon(tcpCon net.Conn) {
	fmt.Println("Created incomming connection")
	conn := Connection{
		tcpCon: tcpCon,
	}
	conn.Send("This is a temporary test auto-response.")
}

func FromOutgoingCon(tcpCon net.Conn) Connection {
	return Connection{
		tcpCon: tcpCon,
	}
}
