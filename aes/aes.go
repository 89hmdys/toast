package aes

import (
	. "crypto/aes"
	"errors"
	"toast/aes/cipher"
	pt "toast/aes/padding"
)

type client struct {
	cipher  Cipher
	padding padding
}

func (this *client) Encrypt(plaintext []byte) []byte {
	plaintext = this.padding.Padding(plaintext)
	ciphertext := this.cipher.Encrypt(plaintext)
	return ciphertext
}

func (this *client) Decrypt(ciphertext []byte) []byte {
	plaintext := this.cipher.Decrypt(ciphertext)
	plaintext = this.padding.UnPadding(plaintext)
	return plaintext
}

//创建默认加密客户端，使用CBC工作模式，PKCS7方式填充
func NewDefault(key []byte) (Cipher, error) {
	block, err := NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}
	return &client{padding: &pkcs7{size: block.BlockSize()}, cipher: &cbc{key: key, block: block}}, nil
}

func New(key []byte, cipherType, paddingType int64) (Cipher, error) {

	block, err := NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}

	var p padding
	var c Cipher

	switch paddingType {
	case pt.PKCS7:
		{
			p = &pkcs7{size: block.BlockSize()}
		}
	default:
		{
			return nil, errors.New("illegal padding,only support PKCS7 for now")
		}
	}

	switch cipherType {
	case cipher.CBC:
		{
			c = &cbc{key: key, block: block}
		}
	case cipher.CFB:
		{
			c = &cfb{key: key, block: block}
		}
	case cipher.OFB:
		{
			c = &ofb{key: key, block: block}
		}
	default:
		{
			return nil, errors.New("illegal padding,only support PKCS7 for now")
		}
	}

	return &client{padding: p, cipher: c}, nil
}
