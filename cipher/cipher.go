package cipher

import (
	. "crypto/cipher"
)

type Cipher interface {
	Encrypt(src []byte) []byte
	Decrypt(src []byte) []byte
}

func NewBlockCipher(padding paddingFunc, unPadding unPaddingFunc, encrypt BlockMode, decrypt BlockMode) Cipher {
	return &blockCipher{
		blockSize: encrypt.BlockSize(),
		encrypt:   encrypt.CryptBlocks,
		decrypt:   decrypt.CryptBlocks,
		padding:   padding,
		unPadding: unPadding}
}

func NewStreamCipher(encrypt Stream, decrypt Stream) Cipher {
	return &streamCipher{
		encrypt: encrypt.XORKeyStream,
		decrypt: decrypt.XORKeyStream}
}

type blockCipher struct {
	blockSize int
	padding   paddingFunc
	unPadding unPaddingFunc
	encrypt   func(ciphertext, plaintext []byte)
	decrypt   func(plaintext, ciphertext []byte)
}

func (this *blockCipher) Encrypt(plaintext []byte) []byte {
	plaintext = this.padding(plaintext, this.blockSize)
	ciphertext := make([]byte, len(plaintext))
	this.encrypt(ciphertext, plaintext)
	return ciphertext
}
func (this *blockCipher) Decrypt(ciphertext []byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	this.decrypt(plaintext, ciphertext)
	plaintext = this.unPadding(plaintext)
	return plaintext
}

type streamCipher struct {
	encrypt func(dst, src []byte)
	decrypt func(dst, src []byte)
}

func (this *streamCipher) Encrypt(plaintext []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	this.encrypt(ciphertext, plaintext)
	return ciphertext
}
func (this *streamCipher) Decrypt(ciphertext []byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	this.decrypt(plaintext, ciphertext)
	return plaintext
}
