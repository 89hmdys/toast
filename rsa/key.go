package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"strings"
	"io/ioutil"
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

func ParsePKCS1Key(publicKey, privateKey []byte) (Key, error) {
	puk, err := x509.ParsePKIXPublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	prk, err := x509.ParsePKCS1PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return &key{publicKey:puk.(*rsa.PublicKey), privateKey:prk.(*rsa.PrivateKey)}, nil
}

func LoadPKCS8KeyFromPEMFile(publicKeyFilePath, privateKeyFilePath string) (Key, error) {

	//TODO 断言如果入参为"" ，则直接报错

	publicKeyFilePath = strings.TrimSpace(publicKeyFilePath)

	pukBytes, err := ioutil.ReadFile(publicKeyFilePath)
	if err != nil {
		return nil, err
	}

	puk, _ := pem.Decode(pukBytes)
	if puk == nil {
		return nil, errors.New("publicKey is not pem formate")
	}

	privateKeyFilePath = strings.TrimSpace(privateKeyFilePath)

	prkBytes, err := ioutil.ReadFile(privateKeyFilePath)
	if err != nil {
		return nil, err
	}

	prk, _ := pem.Decode(prkBytes)
	if prk == nil {
		return nil, errors.New("privateKey is not pem formate")
	}

	return ParsePKCS8Key(puk.Bytes, prk.Bytes)
}

func LoadPKCS1KeyFromPEMFile(publicKeyFilePath, privateKeyFilePath string) (Key, error) {

	//TODO 断言如果入参为"" ，则直接报错

	publicKeyFilePath = strings.TrimSpace(publicKeyFilePath)

	pukBytes, err := ioutil.ReadFile(publicKeyFilePath)
	if err != nil {
		return nil, err
	}

	puk, _ := pem.Decode(pukBytes)
	if puk == nil {
		return nil, errors.New("publicKey is not pem formate")
	}

	privateKeyFilePath = strings.TrimSpace(privateKeyFilePath)

	prkBytes, err := ioutil.ReadFile(privateKeyFilePath)
	if err != nil {
		return nil, err
	}

	prk, _ := pem.Decode(prkBytes)
	if prk == nil {
		return nil, errors.New("privateKey is not pem formate")
	}

	return ParsePKCS1Key(puk.Bytes, prk.Bytes)
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
