package rsa

func pkcs1Padding(src []byte, keySize int) [][]byte {

	srcSize := len(src)

	blockSize := keySize - 11

	var v [][]byte

	if srcSize <= blockSize {
		v = append(v, src)
	} else {
		groups := len(src) / blockSize
		for i := 0; i < groups; i++ {
			block := src[:blockSize]

			v = append(v, block)
			src = src[blockSize:]

			if len(src) < blockSize {
				v = append(v, src)
			}
		}
	}
	return v
}

func unPadding(src []byte, keySize int) [][]byte {

	srcSize := len(src)

	blockSize := keySize

	var v [][]byte

	if srcSize == blockSize {
		v = append(v, src)
	} else {
		groups := len(src) / blockSize
		for i := 0; i < groups; i++ {
			block := src[:blockSize]

			v = append(v, block)
			src = src[blockSize:]
		}
	}
	return v
}
