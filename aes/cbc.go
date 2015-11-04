package aes

import (
	. "crypto/cipher"
)

type cbc struct {
	block Block
	key   []byte
}

func (this *cbc) Encrypt(plaintext []byte) []byte {

	blockModel := NewCBCEncrypter(this.block, this.key[:this.block.BlockSize()])

	ciphertext := make([]byte, len(plaintext))

	blockModel.CryptBlocks(ciphertext, plaintext)

	return ciphertext
}

func (this *cbc) Decrypt(ciphertext []byte) []byte {

	blockModel := NewCBCDecrypter(this.block, this.key[:this.block.BlockSize()])

	plaintext := make([]byte, len(ciphertext))

	blockModel.CryptBlocks(plaintext, ciphertext)

	return plaintext
}
