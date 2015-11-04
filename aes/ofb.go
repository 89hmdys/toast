package aes

import (
	. "crypto/cipher"
)

type ofb struct {
	block Block
	key   []byte
}

func (this *ofb) Encrypt(plaintext []byte) []byte {

	blockModel := NewOFB(this.block, this.key[:this.block.BlockSize()])

	ciphertext := make([]byte, len(plaintext))

	blockModel.XORKeyStream(ciphertext, plaintext)

	return ciphertext
}

func (this *ofb) Decrypt(ciphertext []byte) []byte {

	blockModel := NewOFB(this.block, this.key[:this.block.BlockSize()])

	plaintext := make([]byte, len(ciphertext))

	blockModel.XORKeyStream(plaintext, ciphertext)

	return plaintext
}
