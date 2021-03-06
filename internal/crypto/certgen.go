package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"
)

func LoadCert(certPath string, keyPath string) tls.Certificate {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatalf("server: loadkeys: %s", err)
	}
	return cert
}

func generateKeyPair(byteCount int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privKeyMinSize := 2048
	if byteCount < privKeyMinSize {
		return nil, nil, errors.New(fmt.Sprintf("private key size must be at least %d", privKeyMinSize))
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, byteCount)
	if err != nil {
		return nil, nil, err
	}

	pubKey := &privateKey.PublicKey
	if err != nil {
		return nil, nil, err
	}

	return privateKey, pubKey, err
}

func newCertificateSigningRequest(keySize int) (*x509.CertificateRequest, *rsa.PrivateKey, error) {
	privateKey, pubKey, err := generateKeyPair(keySize)
	if err != nil {
		return nil, nil, err
	}

	template := &x509.CertificateRequest{
		SignatureAlgorithm: x509.SHA256WithRSA,
		PublicKeyAlgorithm: x509.RSA,
		PublicKey:          pubKey,
		Subject:            pkix.Name{CommonName: "example.org"},
		DNSNames:           []string{"example.org"},
	}

	csrDER, err := x509.CreateCertificateRequest(
		rand.Reader,
		template,
		privateKey,
	)
	if err != nil {
		return nil, nil, err
	}

	csr, err := x509.ParseCertificateRequest(csrDER)
	if err != nil {
		return nil, nil, err
	}

	return csr, privateKey, err
}

func newCaCertificate(keySize int) (*x509.Certificate, *rsa.PrivateKey, error) {
	template := &x509.Certificate{
		IsCA: true,
		BasicConstraintsValid: true,
		SubjectKeyId:          []byte{1, 2, 3},
		SerialNumber:          big.NewInt(42),
		Subject: pkix.Name{
			Country:      []string{"Earth"},
			Organization: []string{"Individual"},
		},
		NotBefore:   time.Now(),
		NotAfter:    time.Now().AddDate(0, 3, 0),
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
	}

	privateKey, pubKey, err := generateKeyPair(keySize)
	if err != nil {
		return nil, nil, err
	}

	cert, err := x509.CreateCertificate(rand.Reader, template, template, pubKey, privateKey)
	if err != nil {
		return nil, nil, err
	}

	certObject, err := x509.ParseCertificate(cert)
	if err != nil {
		return nil, nil, err
	}

	return certObject, privateKey, err
}

func newCertificate(caCert *x509.Certificate, caKey *rsa.PrivateKey, csr *x509.CertificateRequest) (*x509.Certificate, error) {
	if csr == nil || csr.Raw == nil || csr.CheckSignature() != nil {
		return nil, errors.New("invalid certificate request passed")
	}

	template := x509.Certificate{
		SerialNumber: new(big.Int).Lsh(big.NewInt(1), 128),
		Subject: pkix.Name{
			Organization: []string{"Individual"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(time.Hour * 24 * 30),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA: true,
	}
	certBytes, err := x509.CreateCertificate(
		rand.Reader,
		&template,
		caCert,
		csr.PublicKey,
		caKey,
	)
	if err != nil {
		return nil, err
	}

	certObject, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, err
	}

	return certObject, err
}

func verifyCertChain(cert *x509.Certificate, caCert *x509.Certificate) bool {
	caCertPool := x509.NewCertPool()
	caCertPool.AddCert(caCert)

	verifyOpts := x509.VerifyOptions{
		Roots: caCertPool,
	}
	// WARNING: (c *Certificate) Verify does not check for revocation.
	chain, err := cert.Verify(verifyOpts)

	if chain != nil && err == nil {
		return true
	}
	return false
}
