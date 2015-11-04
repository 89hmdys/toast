package aes

type Cipher interface {
	Encrypt(plaintext []byte) []byte
	Decrypt(ciphertext []byte) []byte
}

type padding interface {
	Padding(src []byte) []byte
	UnPadding(src []byte) []byte
}
