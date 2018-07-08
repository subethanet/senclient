package senclient

import (
	"bufio"
	"fmt"
	"net"
)

type Connection struct {
	tcpCon net.Conn
}

func (conn *Connection) Send(msg string) {
	fmt.Fprintf(conn.tcpCon, msg+"\n")
	fmt.Println("Sent msg:" + msg)
}

func (conn *Connection) Listen() string {
	msg, err := bufio.NewReader(conn.tcpCon).ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println("Got msg:" + msg)
	return msg
}
