package crypto

import (
	"crypto/aes"
//"crypto/des"
	. "github.com/89hmdys/toast/cipher"
	"crypto/des"
)

//秘钥长度128 192 256 位 , 使用秘钥作为初始向量
func NewAES(key []byte) (Cipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return NewECBMode().Cipher(block, key[:block.BlockSize()]), err
}

//秘钥长度128 192 256 位 , 使用秘钥作为初始向量
func NewAESWith(key []byte, mode CipherMode) (Cipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return mode.Cipher(block, key[:block.BlockSize()]), nil
}

//秘钥长度64位 , 使用秘钥作为初始向量
func NewDES(key []byte) (Cipher, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return NewECBMode().Cipher(block, key[:block.BlockSize()]), nil
}

func NewDESWith(key []byte, mode CipherMode) (Cipher, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return mode.Cipher(block, key[:block.BlockSize()]), nil
}