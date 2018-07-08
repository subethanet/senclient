package senclient

import (
	"testing"
	"time"
)

func TestStart(t *testing.T) {
	client := CreateApp(4546)

	client.LoadCertAndKey(
		"test/testCert.pem",
		"test/testKey.key",
	)

	client.Start()
	time.Sleep(time.Millisecond)
	client.Stop()
}
