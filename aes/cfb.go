package aes

import (
	. "crypto/cipher"
)

type cfb struct {
	key   []byte
	block Block
}

func (this *cfb) Encrypt(plaintext []byte) []byte {

	blockModel := NewCFBEncrypter(this.block, this.key[:this.block.BlockSize()])

	ciphertext := make([]byte, len(plaintext))

	blockModel.XORKeyStream(ciphertext, plaintext)

	return ciphertext
}

func (this *cfb) Decrypt(ciphertext []byte) []byte {

	blockModel := NewCFBDecrypter(this.block, this.key[:this.block.BlockSize()])

	plaintext := make([]byte, len(ciphertext))

	blockModel.XORKeyStream(plaintext, ciphertext)

	return plaintext
}
