package crypto

import (
	"crypto/aes"
//"crypto/des"
	. "github.com/89hmdys/toast/cipher"
	"crypto/des"
)

/*
介绍:创建默认的AES Cipher,使用ECB工作模式、pkcs57填充,算法秘钥长度128 192 256 位 , 使用秘钥作为初始向量

作者:Alex
版本:release-1.1
*/
func NewAES(key []byte) (Cipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	return NewECBMode().Cipher(block, key[:block.BlockSize()]), err
}

/*
介绍:根据指定的工作模式，创建AESCipher,算法秘钥长度128 192 256 位 , 使用秘钥作为初始向量

作者:Alex
版本:release-1.1
*/
func NewAESWith(key []byte, mode CipherMode) (Cipher, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return mode.Cipher(block, key[:block.BlockSize()]), nil
}

/*
介绍:创建默认DESCipher,使用ECB工作模式、pkcs57填充,算法秘钥长度64位 , 使用秘钥作为初始向量

作者:Alex
版本:release-1.1
*/
func NewDES(key []byte) (Cipher, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return NewECBMode().Cipher(block, key[:block.BlockSize()]), nil
}

/*
介绍:根据指定的工作模式，创建DESCipher,算法秘钥长度64位,使用秘钥作为初始向量

作者:Alex
版本:release-1.1
*/
func NewDESWith(key []byte, mode CipherMode) (Cipher, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return mode.Cipher(block, key[:block.BlockSize()]), nil
}