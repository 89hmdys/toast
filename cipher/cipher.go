package cipher

import (
	. "crypto/cipher"
)

type Cipher interface {
	Encrypt(src []byte) []byte
	Decrypt(src []byte) []byte
}

func NewBlockCipher(padding Padding, encrypt, decrypt BlockMode) Cipher {
	return &blockCipher{
		encrypt:   encrypt,
		decrypt:   decrypt,
		padding:   padding}
}

type blockCipher struct {
	padding Padding
	encrypt BlockMode
	decrypt BlockMode
}

func (this *blockCipher) Encrypt(plaintext []byte) []byte {
	plaintext = this.padding.Padding(plaintext, this.encrypt.BlockSize())
	ciphertext := make([]byte, len(plaintext))
	this.encrypt.CryptBlocks(ciphertext, plaintext)
	return ciphertext
}

func (this *blockCipher) Decrypt(ciphertext []byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	this.decrypt.CryptBlocks(plaintext, ciphertext)
	plaintext = this.padding.UnPadding(plaintext)
	return plaintext
}

func NewStreamCipher(encrypt Stream, decrypt Stream) Cipher {
	return &streamCipher{
		encrypt: encrypt,
		decrypt: decrypt}
}

type streamCipher struct {
	encrypt Stream
	decrypt Stream
}

func (this *streamCipher) Encrypt(plaintext []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	this.encrypt.XORKeyStream(ciphertext, plaintext)
	return ciphertext
}
func (this *streamCipher) Decrypt(ciphertext []byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	this.decrypt.XORKeyStream(plaintext, ciphertext)
	return plaintext
}
