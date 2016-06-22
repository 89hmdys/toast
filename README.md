# toast(土司)

## 起因
   * 项目从JAVA迁移到Go后，由于之前JAVA项目AES加密时使用的是默认ECB在Go中不被支持，只能自己动手实现一个，最后越做越多，就把Go AES/DES加密解密都做了个简单的封装，方便日后使用。目前支持RSA/AES/DES加密解密

## 例子

### AES

    mode := cipher.NewCBCMode()                         //加密工作模式，支持 CBC ECB CFB CTR 四种工作模式
    cipher, err := crypto.NewAESWith([]byte(key), mode) //创建一个AES 加密的builder
    if err != nil {
    	t.Error(err)
    	return
    }

    planttext := `故经之以五事，校之以计而索其情：一曰道，二曰天，三曰地，四曰将，五曰法。道者，令民与上同意也，故可与之死，可与之生，而不畏危。天者，阴阳、寒暑、时制也。地者，高下、远近、险易、广狭、死生也。将者，智、信、仁、勇、严也。法者，曲制、官道、主用也。凡此五者，将莫不闻，知之者胜，不知者不胜。故校之以计而索其情，曰：主孰有道？将孰有能？天地孰得？法令孰行？兵众孰强？士卒孰练？赏罚孰明？吾以此知胜负矣。`

    ciphertext := cipher.Encrypt([]byte(planttext))

    ciphertextWithBase64 := base64.URLEncoding.EncodeToString(ciphertext)

    fmt.Println(ciphertextWithBase64)

    ciphertext, err = base64.URLEncoding.DecodeString(ciphertextWithBase64)
    if err != nil {
    	t.Error(err)
    }
    planttextBytes := cipher.Decrypt(ciphertext)

    fmt.Println(string(planttextBytes))

### DES


    cipher, err := crypto.NewDES([]byte("Z'{ru/^e"))
	if err != nil {
		t.Error(err)
		return
	}

	plant := `故经之以五事，校之以计而索其情：一曰道，二曰天，三曰地，四曰将，五曰法。道者，令民与上同意也，故可与之死，可与之生，而不畏危。天者，阴阳、寒暑、时制也。地者，高下、远近、险易、广狭、死生也。将者，智、信、仁、勇、严也。法者，曲制、官道、主用也。凡此五者，将莫不闻，知之者胜，不知者不胜。故校之以计而索其情，曰：主孰有道？将孰有能？天地孰得？法令孰行？兵众孰强？士卒孰练？赏罚孰明？吾以此知胜负矣。`

	cp := cipher.Encrypt([]byte(plant))

	cpStr := base64.URLEncoding.EncodeToString(cp)

	fmt.Println(cpStr)

	ppBy, err := base64.URLEncoding.DecodeString(cpStr)
	if err != nil {
		t.Error(err)
	}
	pp := cipher.Decrypt(ppBy)
	fmt.Println(string(pp))

	fmt.Println("Test_DES_CBC ok")

### RSA

    plant := `故经之以五事，校之以计而索其情：一曰道，二曰天，三曰地，四曰将，五曰法。道者，令民与上同意也，故可与之死，可与之生，而不畏危。天者，阴阳、寒暑、时制也。地者，高下、远近、险易、广狭、死生也。将者，智、信、仁、勇、严也。法者，曲制、官道、主用也。凡此五者，将莫不闻，知之者胜，不知者不胜。故校之以计而索其情，曰：主孰有道？将孰有能？天地孰得？法令孰行？兵众孰强？士卒孰练？赏罚孰明？吾以此知胜负矣。`

	key, err := rsa.LoadKeyFromPEMFile(
		`/Users/alex/Documents/go/src/toast/crypto/rsa_public_key.pem`,
		`/Users/alex/Documents/go/src/toast/crypto/rsa_private_key.pem`,
		rsa.ParsePKCS8Key)
	if err != nil {
		t.Error(err)
		return
	}

	cipher, err := crypto.NewRSA(key)
	if err != nil {
		t.Error(err)
		return
	}

	enT, err := cipher.Encrypt([]byte(plant))
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(base64.StdEncoding.EncodeToString(enT))

	deT, err := cipher.Decrypt(enT)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(string(deT))

	signBytes, err := cipher.Sign([]byte(plant), SHA1)
	if err != nil {
		t.Error(err)
		return
	}

	sign := base64.StdEncoding.EncodeToString(signBytes)

	fmt.Println(sign)

	errV := cipher.Verify([]byte(plant), signBytes, SHA1)
	if errV != nil {
		t.Error(errV)
		return
	}