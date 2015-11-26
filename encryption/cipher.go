package encryption

type Encrypter struct {
	BlockSize int
	Invoke    func(src []byte) []byte
}

type Decrypter struct {
	Invoke func(src []byte) []byte
}

func (this *Encrypter) Execute(src []byte) []byte {
	dst := this.Invoke(src)
	return dst
}

func (this *Decrypter) Execute(src []byte) []byte {
	dst := this.Invoke(src)
	return dst
}

type cipher struct {
	Encrypter Encrypter
	Decrypter Decrypter
	BlockSize int
}

func (this *cipher) Encrypt(src []byte) []byte {
	return this.Encrypter.Execute(src)
}

func (this *cipher) Decrypt(src []byte) []byte {
	return this.Decrypter.Execute(src)
}
