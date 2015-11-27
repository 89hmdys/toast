package cipher

type Cipher interface {
	Encrypt(src []byte) []byte
	Decrypt(src []byte) []byte
}

type blockCipher struct {
	blockSize int
	padding   func(size int, src []byte) []byte
	unPadding func(src []byte) []byte
	encrypt   func(ciphertext, plaintext []byte)
	decrypt   func(plaintext, ciphertext []byte)
}

func (this *blockCipher) Encrypt(plaintext []byte) []byte {
	plaintext = this.padding(this.blockSize, plaintext)
	ciphertext := make([]byte, len(plaintext))
	this.encrypt(ciphertext, plaintext)
	return ciphertext
}
func (this *blockCipher) Decrypt(ciphertext []byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	this.decrypt(plaintext, ciphertext)
	plaintext = this.unPadding(plaintext)
	return plaintext
}

type streamCipher struct {
	blockSize int
	encrypt   func(dst, src []byte)
	decrypt   func(dst, src []byte)
}

func (this *streamCipher) Encrypt(plaintext []byte) []byte {
	ciphertext := make([]byte, len(plaintext))
	this.encrypt(ciphertext, plaintext)
	return ciphertext
}
func (this *streamCipher) Decrypt(ciphertext []byte) []byte {
	plaintext := make([]byte, len(ciphertext))
	this.decrypt(plaintext, ciphertext)
	return plaintext
}

//
//type Encrypter struct {
//	BlockSize int
//	Invoke    func(src []byte) []byte
//}
//
//type Decrypter struct {
//	Invoke func(src []byte) []byte
//}
//
//func (this *Encrypter) Execute(src []byte) []byte {
//	dst := this.Invoke(src)
//	return dst
//}
//
//func (this *Decrypter) Execute(src []byte) []byte {
//	dst := this.Invoke(src)
//	return dst
//}
//
//type cipher struct {
//	Encrypter Encrypter
//	Decrypter Decrypter
//	BlockSize int
//}
//
//func (this *cipher) Encrypt(src []byte) []byte {
//	return this.Encrypter.Execute(src)
//}
//
//func (this *cipher) Decrypt(src []byte) []byte {
//	return this.Decrypter.Execute(src)
//}
