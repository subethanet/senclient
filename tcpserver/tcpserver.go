package tcpserver

import (
	"net"
	"fmt"
	"senclient/con"
	"strconv"
)

/*
	Responsible for the Server object, which is meant to be 1 one-per-app object.
	The Server accepts network connections and passes them along to a Connection object,
	which handles the actual subethanet protocol.
*/


type Server struct {
	listenPort int
}


func Create(port int) Server {
	s := Server{
		port,
	}
	return s
}


func (client *Server) Start() {
	fmt.Println("Listening on tcpserver port " + strconv.Itoa(client.listenPort))
	// Listen for incoming connections.
	l, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(client.listenPort))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		panic(err)
	}
	// Close the tcpserver when the application closes.
	defer l.Close()
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			panic(err)
		}
		// Handle connections in a new goroutine.
		go con.FromIncomingCon(conn)
	}
	fmt.Println("Shouldn't get here...")
}
