package senclient

import (
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	client := Create(4546)

	client.LoadCertAndKey(
		"tcpserver/test/testCert.pem",
		"tcpserver/test/testKey.key",
	)

	client.Start()
	time.Sleep(time.Millisecond)
	client.Stop()
}
