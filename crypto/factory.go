package crypto

import (
	"crypto/aes"
	"crypto/des"
	. "github.com/89hmdys/toast/cipher"
)

//秘钥长度64位 , 使用秘钥作为初始向量
func DES(key []byte) Builder {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	return NewBuilder(block, key[:block.BlockSize()])
}

//秘钥长度128 192 256 位 , 使用秘钥作为初始向量
func AES(key []byte) Builder {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	return NewBuilder(block, key[:block.BlockSize()])
}
