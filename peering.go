package senclient

import (
	"crypto/tls"
	"crypto/x509"
)

func PeerFromDeviceCert(clientCert tls.Certificate) *peer {
	p := peer{
		clientCert: clientCert,
	}
	return &p
}

// openssl genrsa -out rootCA.key 2048
// openssl req -x509 -new -nodes -key rootCA.key -sha256 -days 90 -out rootCA.pem
func validateClientCert(clientCert tls.Certificate, idCert tls.Certificate) bool {
	caPool := x509.CertPool{}
	caPool.AddCert(idCert.Leaf)
	verifyOpts := x509.VerifyOptions{ // Only allow the idCert as a valid CA.
		Roots: &caPool,
	}

	_, err := clientCert.Leaf.Verify(verifyOpts)
	if err == nil {
		return true
	}
	return false
}
