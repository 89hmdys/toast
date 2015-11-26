package cipher

import "bytes"

func PKCS7Padding(size int, src []byte) []byte {
	padding := size - len(src)%size
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func PKCS7UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}
