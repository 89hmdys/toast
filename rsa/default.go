package rsa

//import (
//	"crypto"
//	"crypto/rand"
//	"crypto/rsa"
//)

//type defaultClient struct {
//	client
//}
//
//func (this *defaultClient) Encrypt(plaintext []byte) ([]byte, error) {
////	return rsa.EncryptOAEP(rand.Reader, this.publicKey, plaintext)
//}
//func (this *defaultClient) Decrypt(ciphertext []byte) ([]byte, error) {
//	return rsa.EncryptPKCS1v15(rand.Reader, this.publicKey, ciphertext)
//}
//
//func (this *defaultClient) Sign(src []byte, hash crypto.Hash) ([]byte, error) {
//	h := hash.New()
//	h.Write(src)
//	digest := h.Sum(nil)
//	return rsa.SignPKCS1v15(rand.Reader, this.privateKey, hash, digest)
//}
//
//func (this *defaultClient) Verify(src []byte, sign []byte, hash crypto.Hash) ([]byte, error) {
//	h := hash.New()
//	h.Write(src)
//	digest := h.Sum(nil)
//	return rsa.VerifyPKCS1v15(rand.Reader, this.privateKey, hash, digest)
//}
