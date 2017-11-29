package tcpserver

import (
	"testing"
	"senclient/crypt"
	"crypto/tls"
	"time"
	"net"
	"fmt"
	"bufio"
)


var testPort = 9994


func loadTestCert() tls.Certificate {
	return crypt.LoadCert("./test/testCert.pem", "./test/testKey.key")
}


func createAndStartServer() *Server {
	time.Sleep(100*time.Millisecond)
	server := Create(
		testPort,
		loadTestCert(),
		demoConnectionHandler,
	)
	server.Start()
	time.Sleep(50*time.Millisecond)
	return &server
}


func demoConnectionHandler(conn net.Conn) {
	fmt.Fprintf(conn, "This is a test string.\n")
}


func TestServerStart(t *testing.T) {
	s := createAndStartServer()
	s.Stop()
}


func TestServerStartAgain(t *testing.T) {
	s := createAndStartServer()
	s.Stop()
}


func TestServerAcceptConnection(t *testing.T) {
	s := createAndStartServer()

	connConfig := tls.Config{InsecureSkipVerify: true}
	_, err := tls.Dial("tcp", "127.0.0.1:9994", &connConfig)
	if err != nil {
		panic(err)
	}

	s.Stop()
}


func TestServerConnectionResponse(t *testing.T) {
	s := createAndStartServer()

	connConfig := tls.Config{InsecureSkipVerify: true}
	conn, err := tls.Dial("tcp", "127.0.0.1:9994", &connConfig)
	if err != nil {
		panic(err)
	}

	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		panic(err)
	}
	if msg != "This is a test string.\n" {
		t.Fatal("Didn't get the correct test response. Got: " + msg)
	}

	s.Stop()
}