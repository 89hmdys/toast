package cipher

import "bytes"

type Padding interface {
	Padding(src []byte, blockSize int) []byte
	UnPadding(src []byte) []byte
}

type padding struct {

}

type pkcs57Padding  padding

func NewPKCS57Padding() Padding {
	return &pkcs57Padding{}
}

func (p *pkcs57Padding) Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func (p *pkcs57Padding) UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length - 1])
	return src[:(length - unpadding)]
}
