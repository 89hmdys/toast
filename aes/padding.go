package aes

import (
	"bytes"
)

type pkcs7Padding struct {
	size int
}

func (this *pkcs7Padding) Padding(src []byte) []byte {
	padding := this.size - len(src)%this.size
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func (this *pkcs7Padding) UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
