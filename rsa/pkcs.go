package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
)

type pkcsClient struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func (this *pkcsClient) Encrypt(plaintext []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, this.publicKey, plaintext)
}
func (this *pkcsClient) Decrypt(ciphertext []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, this.privateKey, ciphertext)
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
