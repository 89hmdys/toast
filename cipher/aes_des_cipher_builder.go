package cipher

import (
	. "crypto/cipher"
	. "github.com/89hmdys/toast/padding"
)

type Builder interface {
	CBC() Cipher
	CFB() Cipher
	OFB() Cipher
	ECB() Cipher
	CTR() Cipher
}

func NewBuilder(block Block, iv []byte) Builder {
	return &builder{block: block, iv: iv}
}

type builder struct {
	block Block
	iv    []byte
}

//TODO 下版本支持自定义PADDING
func (this *builder) CBC() Cipher {

	encrypter := NewCBCEncrypter(this.block, this.iv)
	decrypter := NewCBCDecrypter(this.block, this.iv)

	return &blockCipher{blockSize: this.block.BlockSize(),
		encrypt:   encrypter.CryptBlocks,
		decrypt:   decrypter.CryptBlocks,
		padding:   PKCS7Padding,
		unPadding: PKCS7UnPadding}
}

func (this *builder) ECB() Cipher {

	encrypter := NewECBEncrypter(this.block)
	decrypter := NewECBDecrypter(this.block)

	return &blockCipher{blockSize: this.block.BlockSize(),
		encrypt:   encrypter.CryptBlocks,
		decrypt:   decrypter.CryptBlocks,
		padding:   PKCS7Padding,
		unPadding: PKCS7UnPadding}
}

func (this *builder) CFB() Cipher {

	encrypter := NewCFBEncrypter(this.block, this.iv)
	decrypter := NewCFBDecrypter(this.block, this.iv)

	return &streamCipher{encrypt: encrypter.XORKeyStream, decrypt: decrypter.XORKeyStream}
}

func (this *builder) OFB() Cipher {
	encrypter := NewOFB(this.block, this.iv)
	decrypter := NewOFB(this.block, this.iv)
	return &streamCipher{encrypt: encrypter.XORKeyStream, decrypt: decrypter.XORKeyStream}
}

func (this *builder) CTR() Cipher {
	encrypter := NewCTR(this.block, this.iv)
	decrypter := NewCTR(this.block, this.iv)
	return &streamCipher{encrypt: encrypter.XORKeyStream, decrypt: decrypter.XORKeyStream}
}
