package senclient

import "crypto/x509"

type CertificateChain struct {
	clientCert   x509.Certificate
	identityCert x509.Certificate
}

/*func encryptToCert(unencryptedSecret []byte, targetPubkey *rsa.PublicKey) ([]byte, error) {
	encryptedSecret, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		targetPubkey,
		unencryptedSecret,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return encryptedSecret, err
}*/
