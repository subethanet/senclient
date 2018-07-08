package senclient

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
)

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
