package rsa

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
)

type pkcsClient struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func (this *pkcsClient) Encrypt(plaintext []byte) ([]byte, error) {

	blocks := pkcs1Padding(plaintext, this.publicKey.N.BitLen()/8)

	buffer := bytes.Buffer{}
	for _, block := range blocks {
		ciphertextPart, err := rsa.EncryptPKCS1v15(rand.Reader, this.publicKey, block)
		if err != nil {
			return nil, err
		}
		buffer.Write(ciphertextPart)
	}

	return buffer.Bytes(), nil
}

func (this *pkcsClient) Decrypt(ciphertext []byte) ([]byte, error) {

	ciphertextBlocks := unPadding(ciphertext, this.privateKey.N.BitLen()/8)

	buffer := bytes.Buffer{}
	for _, ciphertextBlock := range ciphertextBlocks {
		plaintextBlock, err := rsa.DecryptPKCS1v15(rand.Reader, this.privateKey, ciphertextBlock)
		if err != nil {
			return nil, err
		}
		buffer.Write(plaintextBlock)
	}

	return buffer.Bytes(), nil
}

func (this *pkcsClient) Sign(src []byte, hash crypto.Hash) ([]byte, error) {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, this.privateKey, hash, hashed)
}

func (this *pkcsClient) Verify(src []byte, sign []byte, hash crypto.Hash) error {
	h := hash.New()
	h.Write(src)
	hashed := h.Sum(nil)
	return rsa.VerifyPKCS1v15(this.publicKey, hash, hashed, sign)
}
