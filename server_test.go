package senclient

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"senclient/crypt"
	"testing"
	"time"
)

var testPort = 9994

func loadTestCert() tls.Certificate {
	return crypt.LoadCert("./test/testCert.pem", "./test/testKey.key")
}

func createAndStartServer() *server {
	time.Sleep(100 * time.Millisecond)
	server := newServer(
		testPort,
		loadTestCert(),
		demoConnectionHandler,
	)
	server.start()
	time.Sleep(50 * time.Millisecond)
	return &server
}

func demoConnectionHandler(conn net.Conn) {
	fmt.Fprintf(conn, "This is a test string.\n")
}

func TestServerStart(t *testing.T) {
	s := createAndStartServer()
	s.stop()
}

func TestServerStartAgain(t *testing.T) {
	s := createAndStartServer()
	s.stop()
}

func TestServerAcceptConnection(t *testing.T) {
	s := createAndStartServer()

	connConfig := tls.Config{InsecureSkipVerify: true}
	_, err := tls.Dial("tcp", "127.0.0.1:9994", &connConfig)
	if err != nil {
		panic(err)
	}

	s.stop()
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

	s.stop()
}
