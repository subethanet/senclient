package conn

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
)

type TcpConnection struct {
	tcpCon net.Conn
}

func newTcpConnection(ip string, port int) TcpConnection {
	fmt.Println("Outgoing connection...")
	tcpCon, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	conn := FromOutgoingCon(tcpCon)
	return conn
}

func (conn *TcpConnection) listen() string {
	msg, err := bufio.NewReader(conn.tcpCon).ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println("Got msg:" + msg)
	return msg
}

func (conn *TcpConnection) send(msg string) {
	fmt.Fprintf(conn.tcpCon, msg+"\n")
	fmt.Println("Sent msg:" + msg)
}
