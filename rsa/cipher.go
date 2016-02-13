package rsa

import (
	"crypto"
	"crypto/rsa"
	"bytes"
	"crypto/rand"

	"github.com/Sirupsen/logrus"
)

type Cipher interface {
	Encrypt(plainText []byte) ([]byte, error)
	Decrypt(cipherText []byte) ([]byte, error)
	Sign(src []byte, hash crypto.Hash) ([]byte, error)
	Verify(src []byte, sign []byte, hash crypto.Hash) error
}

func NewCipher(key Key, padding Padding) Cipher {
	return &cipher{key:key, padding:padding}
}

type cipher struct {
	key     Key
	padding Padding
}

func (cipher *cipher) Encrypt(plainText []byte) ([]byte, error) {
	groups := cipher.padding.Padding(plainText)
	buffer := bytes.Buffer{}
	for _, plainTextBlock := range groups {
		cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, cipher.key.PublicKey(), plainTextBlock)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		buffer.Write(cipherText)
	}
	return buffer.Bytes(), nil
}

func (cipher *cipher) Decrypt(cipherText []byte) ([]byte, error) {
	groups := grouping(cipherText, cipher.key.Modulus())
	buffer := bytes.Buffer{}
	for _, cipherTextBlock := range groups {
		plainText, err := rsa.DecryptPKCS1v15(rand.Reader, cipher.key.PrivateKey(), cipherTextBlock)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		buffer.Write(plainText)
	}
	return buffer.Bytes(), nil
}

func (cipher *cipher) Sign(src []byte, hash crypto.Hash) ([]byte, error) {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, cipher.key.PrivateKey(), hash, hashed)
}

func (cipher *cipher) Verify(src []byte, sign []byte, hash crypto.Hash) error {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(cipher.key.PublicKey(), hash, hashed, sign)
}
