package con

import (
	"net"
	"strconv"
	"fmt"
)


func Connect(ip string, port int) Connection {
	fmt.Println("Outgoing connection...")
	tcpCon, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	conn := FromOutgoingCon(tcpCon)
	return conn
}