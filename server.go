package senclient

import (
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
)

// The server accepts network connections and passes them along to a tcpConnection object,
// which handles the actual subethanet protocol.
type server struct {
	listenPort        int
	certificate       tls.Certificate
	connectionHandler func(conn net.Conn)
	shutdownFlag      bool
}

func newServer(port int, cert tls.Certificate, connectionHandler func(conn net.Conn)) server {
	s := server{
		port,
		cert,
		connectionHandler,
		false,
	}
	return s
}

// start the server using a goroutine.
func (s *server) start() {
	go s.startSynchronous()
}

func (s *server) startSynchronous() {
	fmt.Println("Starting tcp server.")
	service := "0.0.0.0:" + strconv.Itoa(s.listenPort)
	config := tls.Config{Certificates: []tls.Certificate{s.certificate}}

	l, err := tls.Listen("tcp", service, &config)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		panic(err)
	}
	defer l.Close() // Close the tcpserver when the application closes.

	// Hang and listen for an incoming connection.
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			panic(err)
		}
		go s.connectionHandler(conn) // Handle connections in a new goroutine.

		if s.shutdownFlag == true {
			break
		}
	}

	fmt.Println("TCP server shutdown.")
}

// Trigger a server shutdown.
func (s *server) stop() {
	fmt.Println("Triggering server shutdown.")
	s.shutdownFlag = true
	config := tls.Config{InsecureSkipVerify: true}
	conn, _ := tls.Dial("tcp", "127.0.0.1:"+strconv.Itoa(s.listenPort), &config) // Throwaway connection.
	conn.Close()
}
