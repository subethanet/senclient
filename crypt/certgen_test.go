package crypt

import (
	"testing"
	"crypto/x509"
	"crypto/rsa"
	"fmt"
)


func TestNewCertificateSigningRequest(t *testing.T) {
	csr, key, err := newCertificateSigningRequest(2048)
	if err != nil {
		panic(err)
	}
	if csr == nil {
		panic("Got null CSR")
	}
	if key == nil {
		panic("Got null key")
	}
}


func TestNewCaCertificate(t *testing.T) {
	cert, privateKey, err := newCaCertificate(2048)
	if err != nil {
		panic(err)
	}
	if cert == nil {
		panic("Got null cert")
	}
	if privateKey == nil {
		panic("Got null privateKey")
	}
}


func TestNewCertificate(t *testing.T) {
	caCert, caKey, err := newCaCertificate(2048)
	if err != nil {
		panic(err)
	}

	csr, _, err := newCertificateSigningRequest(2048)
	if err != nil {
		panic(err)
	}

	cert, err := newCertificate(caCert, caKey, csr)
	if err != nil {
		panic(err)
	}
	if cert == nil {
		panic("cert was nil")
	}
}


func TestNewCertificateWithBadInputs(t *testing.T) {
	testKeySizes := []int{2048, 2048, 4096}

	var tmpCas []*x509.Certificate
	tmpCas = make([]*x509.Certificate, 3)
	var tmpCaKeys []*rsa.PrivateKey
	tmpCaKeys = make([]*rsa.PrivateKey, 3)
	var tmpCsrs []*x509.CertificateRequest
	tmpCsrs = make([]*x509.CertificateRequest, 3)

	for index, keySize := range testKeySizes {
		ca, caKey, _ := newCaCertificate(keySize)
		tmpCas[index] = ca
		tmpCaKeys[index] = caKey

		csr, _, _ := newCertificateSigningRequest(keySize)
		tmpCsrs[index] = csr
	}

	panics := 0

	// Runs 3^3 (27) times.
	for caNo := 0; caNo < 3; caNo++ {
		for caKeyNo := 0; caKeyNo < 3; caKeyNo++ {
			for csrNo := 0; csrNo < 3; csrNo++ {

				cert, err := newCertificate(tmpCas[caNo], tmpCaKeys[caKeyNo], tmpCsrs[csrNo])
				if caNo == caKeyNo {
					if cert == nil || err != nil {
						fmt.Printf("newCertificate failed with caNo=%v, caKeyNo=%v, csrNo=%v\n", caNo, caKeyNo, csrNo)
						panics += 1
					} else if verifyCertChain(cert, tmpCas[caNo]) == false {
						fmt.Printf("verifyCertChain failed with caNo=%v, caKeyNo=%v, csrNo=%v\n", caNo, caKeyNo, csrNo)
						panics += 1
					}
				} else {
					if err == nil {
						if verifyCertChain(cert, tmpCas[caNo]) == true {
							fmt.Printf("newCertificate SHOULD HAVE failed to validate with caNo=%v, caKeyNo=%v, csrNo=%v\n", caNo, caKeyNo, csrNo)
							panics += 1
						}
					}
				}
			}

		}
	}

	if panics != 0 {
		panic(fmt.Sprintf("%v problems out of 27 cases", panics))
	}

}