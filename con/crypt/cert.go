package crypt

import (
	"crypto/tls"
	"log"
)


func LoadCert(certPath string, keyPath string) tls.Certificate {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	return cert
}
