package cipher

import (
	cc "crypto/cipher"
)

type Builder interface {
	CBC() Cipher
	CFB() Cipher
	OFB() Cipher
	CTR() Cipher
	ECB() Cipher
}

func NewBuilder(block cc.Block, iv []byte) Builder {
	return &builder{block: block, iv: iv}
}

type builder struct {
	block cc.Block
	iv    []byte
}

//TODO 下版本支持自定义PADDING
func (this *builder) CBC() Cipher {
	encrypter := cc.NewCBCEncrypter(this.block, this.iv)
	decrypter := cc.NewCBCDecrypter(this.block, this.iv)
	return NewBlockCipher(pkcs7Padding, pkcs7UnPadding, encrypter, decrypter)
}

func (this *builder) ECB() Cipher {
	encrypter := NewECBEncrypter(this.block)
	decrypter := NewECBDecrypter(this.block)
	return NewBlockCipher(pkcs7Padding, pkcs7UnPadding, encrypter, decrypter)
}

func (this *builder) CFB() Cipher {
	encrypter := cc.NewCFBEncrypter(this.block, this.iv)
	decrypter := cc.NewCFBDecrypter(this.block, this.iv)
	return NewStreamCipher(encrypter, decrypter)
}

func (this *builder) OFB() Cipher {
	encrypter := cc.NewOFB(this.block, this.iv)
	decrypter := cc.NewOFB(this.block, this.iv)
	return NewStreamCipher(encrypter, decrypter)
}

func (this *builder) CTR() Cipher {
	encrypter := cc.NewCTR(this.block, this.iv)
	decrypter := cc.NewCTR(this.block, this.iv)
	return NewStreamCipher(encrypter, decrypter)
}
