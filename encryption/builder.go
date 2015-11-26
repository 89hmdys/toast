package encryption

import (
	cc "crypto/cipher"
	. "github.com/89hmdys/toast/cipher"
	tc "toast/cipher"
)

type Builder interface {
	CBC() Builder
	CFB() Builder
	OFB() Builder
	ECB() Builder
	Cipher() Cipher
}

type Cipher interface {
	Encrypt(src []byte) []byte
	Decrypt(src []byte) []byte
}

type builder struct {
	block  cc.Block
	key    []byte
	cipher *cipher
}

func (this *builder) padding() {
	this.cipher.BlockSize = this.block.BlockSize()
	//	this.cipher.Padding = padding.PKCS7Padding
	//	this.cipher.UnPadding = padding.PKCS7UnPadding
}

func (this *builder) CBC() Builder {

	encrypter := cc.NewCBCEncrypter(this.block, this.key[:this.block.BlockSize()])
	this.cipher.Encrypter = Encrypter{
		BlockSize: encrypter.BlockSize(),
		Invoke: func(src []byte) []byte {
			src = PKCS7Padding(encrypter.BlockSize(), src)
			dst := make([]byte, len(src))
			encrypter.CryptBlocks(dst, src)
			return dst
		}}

	decrypter := cc.NewCBCDecrypter(this.block, this.key[:this.block.BlockSize()])
	this.cipher.Decrypter = Decrypter{
		Invoke: func(src []byte) []byte {
			dst := make([]byte, len(src))
			decrypter.CryptBlocks(dst, src)
			PKCS7UnPadding(dst)
			return dst
		}}

	return this
}

func (this *builder) ECB() Builder {

	encrypter := tc.NewECBEncrypter(this.block)

	encrypt := func(src []byte) []byte {
		src = PKCS7Padding(encrypter.BlockSize(), src)
		dst := make([]byte, len(src))
		encrypter.CryptBlocks(dst, src)
		return dst
	}

	this.cipher.Encrypter = Encrypter{
		BlockSize: encrypter.BlockSize(),
		Invoke:    encrypt}

	decrypter := tc.NewECBDecrypter(this.block)

	decrypt := func(src []byte) []byte {
		dst := make([]byte, len(src))
		decrypter.CryptBlocks(dst, src)
		PKCS7UnPadding(dst)
		return dst
	}

	this.cipher.Decrypter = Decrypter{
		Invoke: decrypt}

	return this
}

func (this *builder) CFB() Builder {

	encrypter := cc.NewCFBEncrypter(this.block, this.key[:this.block.BlockSize()])

	encrypt := func(src []byte) []byte {
		dst := make([]byte, len(src))
		encrypter.XORKeyStream(dst, src)
		return dst
	}

	this.cipher.Encrypter = Encrypter{
		Invoke: encrypt}

	decrypter := cc.NewCFBDecrypter(this.block, this.key[:this.block.BlockSize()])

	decrypt := func(src []byte) []byte {
		dst := make([]byte, len(src))
		decrypter.XORKeyStream(dst, src)
		return dst
	}

	this.cipher.Decrypter = Decrypter{
		Invoke: decrypt}

	return this
}
func (this *builder) OFB() Builder {
	ofb := cc.NewOFB(this.block, this.key[:this.block.BlockSize()])

	encrypt := func(src []byte) []byte {
		dst := make([]byte, len(src))
		ofb.XORKeyStream(dst, src)
		return dst
	}

	this.cipher.Encrypter = Encrypter{
		Invoke: encrypt}

	decrypt := func(src []byte) []byte {
		dst := make([]byte, len(src))
		ofb.XORKeyStream(dst, src)
		return dst
	}

	this.cipher.Decrypter = Decrypter{
		Invoke: decrypt}

	return this
}

func (this *builder) Cipher() Cipher {
	this.padding()
	return this.cipher
}
