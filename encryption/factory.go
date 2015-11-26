package encryption

import (
	"crypto/aes"
	"crypto/des"
)

func DES(key []byte) Builder {
	block, err := des.NewCipher(key) //秘钥长度64位
	if err != nil {
		panic(err)
	}
	return &builder{block: block, key: key, cipher: &cipher{}}
}

func AES(key []byte) Builder {
	block, err := aes.NewCipher(key) //秘钥长度128 192 256 位
	if err != nil {
		panic(err)
	}
	return &builder{block: block, key: key, cipher: &cipher{}}
}

func RSA(privateKey []byte, publicKey []byte) {

}
