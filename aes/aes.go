package aes

import (
	. "crypto/aes"
	"errors"
	ct "github.com/89hmdys/toast/aes/cipher"
	pt "github.com/89hmdys/toast/aes/padding"
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

//aes加密的秘钥长度必须是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])
func vaildKey(length int) error {
	switch length {
	case 16, 24, 32:
		break
	default:
		return errors.New("illegal")
	}
	return nil
}

//创建默认加密客户端，使用CBC工作模式，PKCS7方式填充
func NewDefault(key []byte) (Cipher, error) {

	errKey := vaildKey(len(key))
	if errKey != nil {
		return nil, errKey
	}
	block, err := NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}
	return &client{padding: &pkcs7Padding{size: block.BlockSize()}, cipher: &cbc{key: key, block: block}}, nil
}

func New(key []byte, cipherType ct.Type, paddingType pt.Type) (Cipher, error) {

	block, err := NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}

	var p padding
	var c Cipher

	switch paddingType {
	case pt.PKCS7:
		{
			p = &pkcs7Padding{size: block.BlockSize()}
		}
	default:
		{
			return nil, errors.New("illegal padding,only support PKCS7 for now")
		}
	}

	switch cipherType {
	case ct.CBC:
		{
			c = &cbc{key: key, block: block}
		}
	case ct.CFB:
		{
			c = &cfb{key: key, block: block}
		}
	case ct.OFB:
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
