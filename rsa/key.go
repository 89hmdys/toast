package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"encoding/base64"
	"errors"
)

type Key interface {
	PublicKey() *rsa.PublicKey
	PrivateKey() *rsa.PrivateKey
	Modulus() int
}

func ParsePKCS8Key(publicKey, privateKey []byte) (Key, error) {
	puk, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	prk, err := x509.ParsePKCS8PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return &key{publicKey:puk.(*rsa.PublicKey), privateKey:prk.(*rsa.PrivateKey)}, nil
}

func ParsePKCS8KeyWithBase64(publicKey, privateKey string) (Key, error) {
	puk, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}
	prk, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	return ParsePKCS8Key(puk, prk)
}

func ParsePKCS8KeyWithPEM(publicKey, privateKey string) (Key, error) {

	puk, _ := pem.Decode([]byte(publicKey))
	prk, _ := pem.Decode([]byte(privateKey))

	if puk == nil || prk == nil {
		return nil, errors.New("is not pem formate")
	}
	return ParsePKCS8Key(puk.Bytes, prk.Bytes)
}

type key struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
}

func (key *key) Modulus() int {
	return len(key.publicKey.N.Bytes())
}

func (key *key) PublicKey() *rsa.PublicKey {
	return key.publicKey
}

func (key *key) PrivateKey() *rsa.PrivateKey {
	return key.privateKey
}
